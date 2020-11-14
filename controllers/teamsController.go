package controllers

import (
	"encoding/json"
	"go-team/models"
	u "go-team/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var GetTeams = func(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	limit, _ := strconv.Atoi(v.Get("limit"))
	page, _ := strconv.Atoi(v.Get("page"))

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	offset := (limit * page) - limit

	params := make(map[string]int)
	params["limit"] = limit
	params["offset"] = offset

	data := models.GetTeams(params)
	resp := u.Message(true, "success")
	resp["limit"] = limit
	resp["page"] = page
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateTeam = func(w http.ResponseWriter, r *http.Request) {

	team := &models.Team{}

	err := json.NewDecoder(r.Body).Decode(team)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := team.Create()
	if resp == nil {
		u.Respond(w, u.Message(false, "Create team failed"))
		return
	}

	u.Respond(w, resp)
}

var FindTeam = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var params_id string

	params_id = vars["id"]

	id, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Team ID not valid"))
		return
	}

	data := models.GetTeam(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateTeam = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var params_id string

	params_id = vars["id"]

	id, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Team ID not valid"))
		return
	}

	team := &models.Team{}

	error := json.NewDecoder(r.Body).Decode(team)
	if error != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := team.Update(uint(id))
	u.Respond(w, resp)
}

var DeleteTeam = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var params_id string

	params_id = vars["id"]

	id, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Team ID not valid"))
		return
	}

	models.DeleteTeam(uint(id))
	resp := u.Message(true, "Delete team success")
	u.Respond(w, resp)
}
