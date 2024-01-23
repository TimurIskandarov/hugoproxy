package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

func saveUser(user User) error {
	if _, ok := users[user.Name]; ok {
		return errors.New("this user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Сохраняем хэшированный пароль в карту
	users[user.Name] = string(hashedPassword)
	return nil
}

// при регистрации храним пользователя в памяти
func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// пароль храним в хешированном виде
	err = saveUser(user)
	if err != nil {
		fmt.Println("error save:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		os.Exit(1)
	}

	_, token, _ := TokenAuth.Encode(map[string]interface{}{
		user.Name: user.Password,
	})

	response := ResponseRegister{
		Token: token,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("register response error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// проверка пароля
	isMatchPassword := func(password string) bool {
		return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
	}
	// если пользователь не зарегистрирован
	if password, ok := users[user.Name]; !ok && !isMatchPassword(password) {
		http.Error(w, "user not found", http.StatusOK)
		return
	}

	_, token, _ := TokenAuth.Encode(map[string]interface{}{
		user.Name: user.Password,
	})

	response := ResponseLogin{
		Token: token,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("login response error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
