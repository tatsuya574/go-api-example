package request

import (
	"api.example.com/pkg/entity"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UserCreate(req *http.Request) (*entity.User, error) {
	defer req.Body.Close()
	body := struct {
		User struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&body)
	if err != nil {
		return nil, fmt.Errorf("HTTP Request UserCreate: %w", err)
	}

	password, err := entity.NewPassword(body.User.Password)
	if err != nil {
		return nil, fmt.Errorf("HTTP Request UserCreate: %w", err)
	}

	return entity.NewUser(
		body.User.Name,
		password,
	), nil
}

func UserRead(req *http.Request) (entity.UserID, error) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		return 0, fmt.Errorf("HTTP Request UserRead: %w", err)
	}

	return entity.UserID(id), nil
}
