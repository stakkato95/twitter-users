package domain

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-users/config"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type UserRepo interface {
	Create(*User) (*User, error)
	Authenticate(*User) (string, error)
	GetUserIdByUsername(string) (int, error)
}

type defaultUserRepo struct {
	db *sql.DB
}

func NewUserRepo() UserRepo {
	db, err := sql.Open(config.AppConfig.DbDriver, config.GetConnectionString())

	if err != nil {
		logger.Panic("can not open database connection: " + err.Error())
	}
	if err = db.Ping(); err != nil {
		logger.Panic("can not ping database: " + err.Error())
	}

	repo := &defaultUserRepo{db: db}
	repo.migrate()
	return repo
}

func (r *defaultUserRepo) Create(user *User) (*User, error) {
	statement, err := r.db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	if err != nil {
		logger.Fatal("can not create insert statement: " + err.Error())
		return nil, err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		logger.Fatal("can not hash password: " + err.Error())
	}

	var result sql.Result
	result, err = statement.Exec(user.Username, hashedPassword)
	if err != nil {
		logger.Fatal("can not insert user: " + err.Error())
	}

	var id int64
	id, err = result.LastInsertId()
	user.Id = id
	user.Password = ""

	return user, nil
}

func (r *defaultUserRepo) Authenticate(user *User) (string, error) {
	statement, err := r.db.Prepare("SELECT password FROM user WHERE username = ?")
	if err != nil {
		logger.Fatal("can not create select statement: " + err.Error())
	}

	row := statement.QueryRow(user.Username)
	var hashedPassword string
	if err = row.Scan(&hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("wrong password")
		} else {
			return "", errors.New("user does not exist")
		}
	}

	return hashedPassword, nil
}

func (r *defaultUserRepo) GetUserIdByUsername(username string) (int, error) {
	statement, err := r.db.Prepare("select ID from user WHERE username = ?")
	if err != nil {
		logger.Fatal("can not create select statement: " + err.Error())
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Fatal(err.Error())
		}
		logger.Fatal("unknown error when retrieving user id")
		return 0, err
	}

	return Id, nil
}

func (r *defaultUserRepo) migrate() {
	driver, _ := mysql.WithInstance(r.db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		config.AppConfig.DbName,
		driver,
	)
	if err != nil {
		logger.Fatal("can not migrate up: " + err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Fatal("can not migrate up: " + err.Error())
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
