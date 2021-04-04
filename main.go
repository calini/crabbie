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

func main() {
	activeRooms = make(map[string]Room)

	r := gin.Default()
	r.GET("/room", GetNewRoom)
	r.GET("/room/:code", GetRoom)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func GetNewRoom(c *gin.Context) {
	code := strings.GenerateCode(CODE_LENGTH)
	activeRooms[code] = Room{code, []string{}}
	response := "Here's your new room: " + code
	c.String(http.StatusAccepted, response)
}

func GetRoom(c *gin.Context) {

	room, found := activeRooms[c.Param("code")]

	if found {
		e, err := json.Marshal(room)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(e))
		c.String(http.StatusOK, "Room info: \n"+string(e))
	} else {
		c.String(http.StatusNotFound, "No room was found!")
	}
}

type Room struct {
	Code    string   `json:"room_code"`
	Players []string `json:"players"`
}
