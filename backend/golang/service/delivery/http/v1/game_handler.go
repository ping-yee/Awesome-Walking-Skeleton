package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repository "github.com/ping-yee/Awesome-Walking-Skeleton/service/repository"
)

type GameHandler struct {
	GameRepo repository.GameRepository
}

func (g *GameHandler) GetGameById(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	game, err := g.GameRepo.GetGameById(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repository.Game{
		Id:   game.Id,
		Name: game.Name,
	})
}

func (g *GameHandler) CreateGame(c *gin.Context) {
	game := new(repository.Game)
	game.Name = "Game"

	game, err := g.GameRepo.CreateGame(c, game)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repository.Game{
		Id:   game.Id,
		Name: game.Name,
	})
}
