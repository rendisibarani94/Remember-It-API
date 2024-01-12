package controllers

import (
	"first-jwt/helpers"
	"first-jwt/models"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.CustomClaims) // convert into a struct
	userResponse := &models.Profile{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

helpers.Response(w, 200, "My Profile", userResponse)
}