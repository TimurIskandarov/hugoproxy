package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func saveUser(user User, users map[string]string) error {
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
	err = saveUser(user, users)
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

	// 	если пользователь не зарегистрирован
	if password, ok := users[user.Name]; !ok && password != user.Password {
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
