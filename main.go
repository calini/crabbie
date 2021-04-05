package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/calini/crabbie/pkg/strings"
	"github.com/gin-gonic/gin"
)

const CODE_LENGTH = 4

var activeRooms map[string]Room
var activePlayers map[string]Player

func main() {
	activeRooms = make(map[string]Room)

	r := gin.Default()
	r.GET("/room/", GetNewRoom)
	r.GET("/room/:room_type/:code", GetRoom)
	r.POST("/room/:room_type/:code/user/:user_name", CreateNewUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func GetNewRoom(c *gin.Context) {

	code := strings.GenerateCode(CODE_LENGTH)

	activeRooms[code] = Room{code, "1", []Player{}, []string{}}
	response := fmt.Sprintf("Here's your new room: %s, %s", activeRooms[code].Code, activeRooms[code].RoomType)
	c.String(http.StatusAccepted, response)
}

func CreateNewUser(c *gin.Context) {

	room, found := activeRooms[c.Param("code")]

	if found {
		player := Player{c.Param("user_name"), []string{}}
		room.Players = append(room.Players, player)
		activeRooms[room.Code] = room
		c.String(http.StatusOK, player.Name)
	} else {
		c.String(http.StatusNotFound, "Could not add user for room!")
	}
}

func GetRoom(c *gin.Context) {

	room, found := activeRooms[c.Param("code")]

	if found {
		e, err := json.Marshal(room)
		if err != nil {
			fmt.Println(err)
			return
		}
		response := fmt.Sprintf("Room info: \n %s", string(e))
		c.String(http.StatusOK, response)
	} else {
		c.String(http.StatusNotFound, "No room was found!")
	}
}

type Player struct {
	Name    string   `json:"player_name"`
	Answers []string `json:"player_answers"`
}

type Room struct {
	Code      string   `json:"room_code"`
	RoomType  string   `json:"room_type"`
	Players   []Player `json:"players_in_room"`
	Questions []string `json:"sets_of_questions"`
}
