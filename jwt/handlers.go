package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key = []byte("The true crimefighter always carries everything he needs in his utility belt, Robin.")

var users = map[string]string{
	"bruce": "batman",
	"peter": "spider-man",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Index(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	page.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	credentials := &Credentials{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		},
	)
}
