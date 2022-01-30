package main

import (
	"fmt"
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
	// put creds provided into a JSON struct
	credentials := &Credentials{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// check credentials
	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// create a JWT
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

	// return the JWT as a cookie
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		},
	)
	http.Redirect(w, r, "/welcome.html", http.StatusSeeOther)
}

func Browse(w http.ResponseWriter, r *http.Request) {
	// check if there is a cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	// parse cookie for signed JWT
	tokenStr := cookie.Value
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if JWT is valid
	if !parsedToken.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
	}
}
