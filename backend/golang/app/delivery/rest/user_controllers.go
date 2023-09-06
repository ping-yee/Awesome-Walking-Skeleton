package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/taimoor99/three-tier-golang/app/entities"
	"github.com/taimoor99/three-tier-golang/app/models"
	"github.com/taimoor99/three-tier-golang/utills"
)

type user struct {
	Models models.UserModel
}

type UserControllers interface {
	GetUserByIdHandler(w http.ResponseWriter, r *http.Request)
	PostCreateUserHandler(w http.ResponseWriter, r *http.Request)
	GetAllUsersHandler(writer http.ResponseWriter, request *http.Request)
}

func NewUser(userModels models.UserModel) UserControllers {
	return &user{
		Models: userModels,
	}
}

func (m *user) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) < 1 {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, utills.UserIdNotFoundInParam)
		return
	}

	user, err := m.Models.FindUserByID(context.Background(), id)
	if err != nil {
		utills.WriteJsonRes(w, http.StatusNotFound, nil, err.Error())
		return
	}

	utills.WriteJsonRes(w, http.StatusOK, user, "")
	return
}

func (m *user) PostCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userReq entities.UsersCreateInput
	if err := entities.DecodeAndValidate(r, &userReq); err != nil {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	userId, err := m.Models.CreateUser(context.Background(), userReq)
	if err != nil {
		utills.WriteJsonRes(w, http.StatusNotFound, nil, err.Error())
		return
	}

	utills.WriteJsonRes(w, http.StatusOK, userId, utills.UserCreated)
	return
}

func (m *user) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(chi.URLParam(r, "limit"))
	if err != nil {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, utills.LimitNotFoundInParam)
		return
	}
	offset, err := strconv.Atoi(chi.URLParam(r, "offset"))
	if err != nil {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, utills.OffsetNotFoundInParam)
		return
	}

	users, err := m.Models.GetAllUsers(context.Background(), int64(limit), int64(offset))
	if err != nil {
		utills.WriteJsonRes(w, http.StatusNotFound, nil, err.Error())
		return
	}

	utills.WriteJsonRes(w, http.StatusOK, users, "")
	return
}