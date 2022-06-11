package app

import (
	"encoding/json"
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/handlers"
	"github.com/stakkato95/twitter-service-users/dto"
	"github.com/stakkato95/twitter-service-users/service"
)

type userHandlers struct {
	service service.UserService
}

func (h *userHandlers) hello(w http.ResponseWriter, r *http.Request) {
	handlers.WriteResponse(w, http.StatusOK, "hello")
}

func (h *userHandlers) create(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		handlers.WriteResponse(w, http.StatusBadRequest, "invalid body")
		return
	}
	defer r.Body.Close()

	entity := userDto.ToEntity()
	token, createdUser, err := h.service.Create(&entity)
	if err != nil {
		handlers.WriteResponse(w, err.Code, err)
		return
	}

	handlers.WriteResponse(w, http.StatusCreated, dto.NewUserDto{
		User:  dto.ToDto(*createdUser),
		Token: dto.TokenDto{Token: token},
	})
}

func (h *userHandlers) auth(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		handlers.WriteResponse(w, http.StatusBadRequest, "invalid body")
		return
	}
	defer r.Body.Close()

	entity := userDto.ToEntity()
	token, err := h.service.Authenticate(&entity)
	if err != nil {
		handlers.WriteResponse(w, err.Code, err)
		return
	}

	handlers.WriteResponse(w, http.StatusOK, dto.TokenDto{Token: token})
}
