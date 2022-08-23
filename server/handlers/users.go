package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"userbase-api/server/dal"
	service "userbase-api/server/service/users"
	ue "userbase-api/server/utils/errors"
	"userbase-api/server/utils/helpers"

	"github.com/gorilla/mux"
)

// CreateUser ... Create User
// @Summary Create new user based on parameters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body dal.User true "User Data"
// @Success 201 {object} dal.User
// @Failure 400,500 {object} object
// @Router /users [post]
func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		code := http.StatusBadRequest
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	var user dal.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		code := http.StatusBadRequest
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	ctx := r.Context()
	newUser, err := service.Create(ctx, &user)
	if err != nil {
		code := ue.GetStatusCode(err)
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	code := http.StatusCreated
	helpers.RespondWithJSON(w, code, newUser)
}

// UpdateUser ... Update User
// @Summary Update existing user based on parameters
// @Description update existing user
// @Tags Users
// @Accept json
// @Param user body dal.User true "User Data"
// @Success 200 {object} dal.User
// @Failure 400,403,500 {object} object
// @Router /users/{id} [put]
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		code := http.StatusBadRequest
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	var user dal.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		code := http.StatusBadRequest
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	if user.Id == nil || *user.Id != id {
		code := http.StatusForbidden
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	ctx := r.Context()
	newUser, err := service.Update(ctx, &user)
	if err != nil {
		code := ue.GetStatusCode(err)
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	code := http.StatusOK
	helpers.RespondWithJSON(w, code, newUser)
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {array} dal.User
// @Failure 404,500 {object} object
// @Router /users [get]
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := service.GetAll(ctx)
	if err != nil {
		code := ue.GetStatusCode(err)
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}
	if users == nil {
		code := http.StatusNotFound
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}
	code := http.StatusOK
	helpers.RespondWithJSON(w, code, users)
}

// GetUser ... Get user by ID
// @Summary Get one user
// @Description get user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} dal.User
// @Failure 404,500 {object} object
// @Router /users/{id} [get]
func GetSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := service.Get(ctx, id)
	if err != nil {
		code := ue.GetStatusCode(err)
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}

	if user == nil {
		code := http.StatusNotFound
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}
	code := http.StatusOK
	helpers.RespondWithJSON(w, code, user)
}

// DeleteUser ... Delete user by ID
// @Summary Delete one user
// @Description delete user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 204 {object} object
// @Failure 500 {object} object
// @Router /users/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	if err := service.Delete(ctx, id); err != nil {
		code := ue.GetStatusCode(err)
		helpers.RespondWithError(w, code, http.StatusText(code))
		return
	}
	code := http.StatusNoContent
	http.Error(w, http.StatusText(code), code)
}
