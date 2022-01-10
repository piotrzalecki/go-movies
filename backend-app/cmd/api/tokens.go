package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"movies-backend/models"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"

	"golang.org/x/crypto/bcrypt"
)

var validUser = models.User{
	ID:       10,
	Email:    "me@here.com",
	Password: "$2a$12$CVsG9rZvjUOJ29omTe0sb.49vt5xrNLJXHZbpFmZcYtAC3QbKT9tW",
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *application) Signin(w http.ResponseWriter, r *http.Request) {

	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJson(w, errors.New("unauthorised"))
		return
	}

	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errorJson(w, errors.New("unauthorised"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errorJson(w, errors.New("error signing in"))
		return
	}

	app.writeJson(w, http.StatusOK, string(jwtBytes), "response")

}
