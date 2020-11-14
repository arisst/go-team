package controllers

import (
	"encoding/json"
	"go-team/models"
	u "go-team/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var GetPlayers = func(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	limit, _ := strconv.Atoi(v.Get("limit"))
	page, _ := strconv.Atoi(v.Get("page"))
	team_id, _ := strconv.Atoi(v.Get("team_id"))

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

	if team_id != 0 {
		params["team_id"] = team_id
	}

	data := models.GetPlayers(params)
	resp := u.Message(true, "success")
	resp["limit"] = limit
	resp["page"] = page
	resp["data"] = data
	u.Respond(w, resp)
}

var CreatePlayer = func(w http.ResponseWriter, r *http.Request) {

	player := &models.Player{}

	err := json.NewDecoder(r.Body).Decode(player)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := player.Create()
	u.Respond(w, resp)
}

var FindPlayer = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var params_id string

	params_id = vars["id"]

	id, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Player ID not valid"))
		return
	}

	data := models.GetPlayer(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdatePlayer = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var params_id string

	params_id = vars["id"]

	id, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Player ID not valid"))
		return
	}

	player := &models.Player{}

	error := json.NewDecoder(r.Body).Decode(player)
	if error != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := player.Update(uint(id))
	u.Respond(w, resp)
}

var DeletePlayer = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var params_id string

	params_id = vars["id"]

	id, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Player ID not valid"))
		return
	}

	models.DeletePlayer(uint(id))
	resp := u.Message(true, "Delete player success")
	u.Respond(w, resp)
}
