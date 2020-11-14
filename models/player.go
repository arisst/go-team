package models

import (
	"fmt"
	u "go-team/utils"

	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model
	Name     string `json:"name"`
	Position string `json:"position"`
	TeamId   uint   `json:"team_id"`
	Team     Team
}

func (player *Player) Validate() (map[string]interface{}, bool) {

	if player.Name == "" {
		return u.Message(false, "Player name should be on the payload"), false
	}
	if player.Position == "" {
		return u.Message(false, "Player position should be on the payload"), false
	}
	if player.TeamId == 0 {
		return u.Message(false, "TeamId should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (player *Player) Create() map[string]interface{} {

	if resp, ok := player.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(&player).Error
	if err != nil {
		return u.Message(false, "Create player failed!")
	}

	team := Team{}
	GetDB().Table("teams").Where("id = ?", player.TeamId).Find(&team)

	player.Team = team

	resp := u.Message(true, "success")
	resp["data"] = player
	return resp
}

func (player *Player) Update(id uint) map[string]interface{} {

	if resp, ok := player.Validate(); !ok {
		return resp
	}

	err := GetDB().Table("players").Where("id = ?", id).Updates(player).Error
	if err != nil {
		return u.Message(false, "Update player failed!")
	}

	team := Team{}
	GetDB().Table("teams").Where("id = ?", player.TeamId).Find(&team)

	player.Team = team

	resp := u.Message(true, "success")
	player.ID = id
	resp["data"] = player
	return resp
}

func GetPlayer(id uint) *Player {

	player := &Player{}
	player.ID = id
	err := GetDB().Preload("Team").Find(player).Error
	if err != nil {
		return nil
	}
	return player
}

func DeletePlayer(id uint) *Player {

	player := &Player{}
	err := GetDB().Delete(player, id).Error

	if err != nil {
		return nil
	}
	return player
}

func GetPlayers(params map[string]int) []*Player {
	players := make([]*Player, 0)

	where := map[string]interface{}{}
	if team_id, ok := params["team_id"]; ok {
		where["team_id"] = team_id
	}

	err := GetDB().Preload("Team").Limit(params["limit"]).Offset(params["offset"]).Where(where).Find(&players).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return players
}
