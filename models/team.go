package models

import (
	"fmt"
	u "go-team/utils"

	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	Name string `json:"name"`
}

func (team *Team) Validate() (map[string]interface{}, bool) {

	if team.Name == "" {
		return u.Message(false, "Team name should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (team *Team) Create() map[string]interface{} {

	if resp, ok := team.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(team).Error
	if err != nil {
		return nil
	}

	resp := u.Message(true, "success")
	resp["data"] = team
	return resp
}

func (team *Team) Update(id uint) map[string]interface{} {

	if resp, ok := team.Validate(); !ok {
		return resp
	}

	err := GetDB().Table("teams").Where("id = ?", id).Updates(team).Error
	if err != nil {
		return u.Message(false, "Update team failed!")
	}

	resp := u.Message(true, "success")
	team.ID = id
	resp["data"] = team
	return resp
}

func GetTeam(id uint) *Team {

	team := &Team{}
	err := GetDB().First(team, id).Error
	if err != nil {
		return nil
	}
	return team
}

func DeleteTeam(id uint) *Team {

	team := &Team{}
	err := GetDB().Delete(team, id).Error

	if err != nil {
		return nil
	}
	return team
}

func GetTeams(params map[string]int) []*Team {
	teams := make([]*Team, 0)
	err := GetDB().Limit(params["limit"]).Offset(params["offset"]).Find(&teams).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return teams
}
