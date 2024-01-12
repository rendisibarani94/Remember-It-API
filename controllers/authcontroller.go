package controllers

import (
	"encoding/json"
	"first-jwt/configs"
	"first-jwt/helpers"
	"first-jwt/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register
	if err := json.NewDecoder(r.Body).Decode(&register); err != nil{
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 500, "Confirm Password not Match", nil)
		return
	}

PasswordHash, err := helpers.HashPassword(register.Password)
if err != nil {
	helpers.Response(w, 400, err.Error(), nil)
	return
}

user := models.User {
	Name: register.Name,
	Email: register.Email,
	Password: PasswordHash,
}

if err := configs.DB.Create(&user).Error; err != nil{
	helpers.Response(w, 500, err.Error(), nil)
	return
}
helpers.Response(w, 201, "Register Successfully", nil)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	if err := configs.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong Email or Password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong Email or Password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "SuccessFuly Login", token)
}