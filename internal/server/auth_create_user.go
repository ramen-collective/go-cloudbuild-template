package server

import (
	"encoding/json"
	"net/http"

	"github.com/ramen-collective/go-cloudbuild-template/internal/repository"
)

func CreateUserRest(w http.ResponseWriter, r *http.Request, repositories *repository.Container) {
	r.ParseForm()
	username := r.Form.Get("username")
	if len(username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := repositories.User.GetByName(username); err == nil {
		w.WriteHeader(http.StatusConflict)
		jsonResponse, _ := json.Marshal("Username already taken")
		w.Write(jsonResponse)
		return
	}
	user, err := repositories.User.Create(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(err)
		w.Write(jsonResponse)
	}

	jsonResponse, _ := json.Marshal(
		map[string]interface{}{
			"userInfo": map[string]interface{}{
				"uuid": user.UUID,
				"name": user.Name,
			},
		})
	w.Write(jsonResponse)
}
