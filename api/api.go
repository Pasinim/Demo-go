package api

import (
	"demo/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EcommerceApi struct {
	repository repo.Repository
}

// ----------- LIVELLO 1: API ----------- //
func (e *EcommerceApi) GETCollezione(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, e.repository.GetAllCollezioniREPO())
	} else {
		idInt, _ := strconv.Atoi(id)
		c.JSON(http.StatusOK, e.repository.GetArticoliCollezioneREPO(idInt))
	}
}

func (e *EcommerceApi) GETArticoli(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, e.repository.GetArticoliREPO())
	} else {
		idInt, _ := strconv.Atoi(id)
		c.JSON(http.StatusOK, e.repository.GetArticoloREPO(idInt))
	}
}

/*Funzione che inizializza il server*/
func New(r repo.Repository) *gin.Engine {
	api := EcommerceApi{repository: r}
	router := gin.Default()
	router.GET("/collezione/", api.GETCollezione) //creo gli handler
	router.GET("/collezione/:id", api.GETCollezione)
	router.GET("/articolo/", api.GETArticoli)
	router.GET("/articolo/:id/", api.GETArticoli)
	router.GET("/collezione/:id/articoli", api.GETCollezione)
	return router
}
