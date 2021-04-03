package main

import (
	"github.com/calini/crabbie/pkg/strings"
	"net/http"

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

	c.String(http.StatusAccepted, "Here's your new room: "+code)
}

func GetRoom(c *gin.Context) {

	room, found := activeRooms[c.Param("code")]

	if found {
		c.String(http.StatusOK, "camera numarul: "+room.Code)
	} else {
		c.String(http.StatusNotFound, "Nu exista camera")
	}
}

type Room struct {
	Code    string
	Players []string
}
