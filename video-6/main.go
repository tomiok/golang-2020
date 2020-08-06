package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

const signKey = "changeMeInProduction"

func main() {
	server := newWebServer("9999")
	routes()
	server.start()
}

type webServer struct {
	s *http.Server
}

func newWebServer(port string) *webServer {
	s := &http.Server{
		Addr: ":" + port,
	}

	return &webServer{
		s: s,
	}
}

func (w *webServer) start() {
	log.Println("server start at port" + w.s.Addr)
	log.Fatal(w.s.ListenAndServe())
}

func routes() {
	http.HandleFunc("/login", userHandler())
	http.HandleFunc("/greeting", validateTokenMiddleware(secureGreeting()))
}

func secureGreeting() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello from secure endpoint"))
	}
}

func userHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "POST" {
			body := request.Body
			defer body.Close()
			var user User
			err := json.NewDecoder(body).Decode(&user)

			if err != nil {
				http.Error(writer, "cannot decode json", http.StatusBadRequest)
				return
			}

			if user.valid() {
				res := getTokens()
				_ = json.NewEncoder(writer).Encode(&res)
			} else {
				http.Error(writer, "bad credentials", http.StatusUnauthorized)
				return
			}

		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getTokens() *JWTResponse {
	token := signToken()
	refresh := signRefreshToken()

	return &JWTResponse{
		Token:   token,
		Refresh: refresh,
	}
}

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

func (u *User) valid() bool {
	if u.Username == "tomas" && u.Password == "1234" {
		return true
	}
	return false
}

type JWTResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh_token"`
}

func signToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Date(2020, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	mySigningKey := []byte("AllYourBase")
	t, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return t
}

func signRefreshToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "baz",
		"exp":     time.Date(2020, 11, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"refresh": true,
	})
	mySigningKey := []byte("AllYourBase")
	t, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return t
}

// this middleware
func validateTokenMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if validateToken(bearerToken) {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "invalid token", http.StatusUnauthorized)
	}
}

func validateToken(s string) bool {
	if strings.Contains(strings.ToLower(s), "bearer ") {
		// token := strings.Replace(s, "bearer ", "", 1) // handle bearer ignoring casing
		token := s[7:]
		mySigningKey := []byte("AllYourBase") // centralize this secret key
		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return mySigningKey, nil
		})

		if err != nil {
			log.Println(err.Error())
			return false
		}

		return t.Valid
	}
	return false
}

