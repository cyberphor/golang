package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var key = []byte("The true crimefighter always carries everything he needs in his utility belt, Robin.")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func CookiePage(w http.ResponseWriter, r *http.Request) {
	tokenString, expires, err := CreateToken("victor", "user")
	if err != nil {
		http.Error(w, "Server error.", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "token", Value: tokenString, Expires: expires})
	w.Write([]byte("Hello to the cookie jar!"))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func ScoreboardPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Team 1: 0, Team 2: 0"))
}

func CreateToken(username string, role string) (string, time.Time, error) {
	expires := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	return tokenString, expires, err
}

func VerifyToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Cookie not found.", http.StatusBadRequest)
			return
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(*jwt.Token) (interface{}, error) { return key, nil })
		// claims := token.Claims.(*Claims)
		// if VerifyRole(claims.Role)
		if err != nil {
			http.Error(w, "Bad token.", http.StatusBadRequest)
			return
		}
		if !token.Valid {
			http.Error(w, "Invalid token.", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	IndexPageHandler := http.HandlerFunc(IndexPage)
	CookiePageHandler := http.HandlerFunc(CookiePage)
	ScoreboardPageHandler := http.HandlerFunc(ScoreboardPage)
	mux.Handle("/", IndexPageHandler)
	mux.Handle("/cookiejar.html", CookiePageHandler)
	mux.Handle("/scoreboard.html", VerifyToken(ScoreboardPageHandler))
	http.ListenAndServe(":666", mux)
}
