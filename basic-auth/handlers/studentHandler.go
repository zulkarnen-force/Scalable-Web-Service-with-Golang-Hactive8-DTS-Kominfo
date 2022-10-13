package handlers

import (
	"basic-auth/middlewares"
	"basic-auth/models"
	"encoding/json"
	"net/http"
)


func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}


func StudentHandler(w http.ResponseWriter, r *http.Request) {

	if !middlewares.Auth(w, r) {return}
	if !middlewares.AllowOnlyGET(w, r) {return}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, models.SelectStudent(id))
		return
	}

	OutputJSON(w, models.GetStudents())
}