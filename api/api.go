package api

import (
	"demo/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Repository interface {
	GetAllCollezioniREPO() []core.Collection
	GetArticoliREPO() []core.Item
	GetArticoloREPO(int) core.Item
	GetArticoliCollezioneREPO(int) []core.Item
}

type EcommerceApi struct {
	rep Repository //accetto tutto quello che e compatibile con l'interfaccia
}

// ----------- LIVELLO 1: API ----------- //
func (e *EcommerceApi) GETCollezione(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, e.rep.GetAllCollezioniREPO())
	} else {
		idInt, _ := strconv.Atoi(id)
		c.JSON(http.StatusOK, e.rep.GetArticoliCollezioneREPO(idInt))
	}
}

func (e *EcommerceApi) GETArticoli(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, e.rep.GetArticoliREPO())
	} else {
		idInt, _ := strconv.Atoi(id)
		c.JSON(http.StatusOK, e.rep.GetArticoloREPO(idInt))
	}
}

/*Funzione che inizializza il server*/
func New(r Repository) *gin.Engine {
	api := EcommerceApi{rep: r}
	router := gin.Default()
	router.GET("/collezione/", api.GETCollezione) //creo gli handler
	router.GET("/collezione/:id", api.GETCollezione)
	router.GET("/articolo/", api.GETArticoli)
	router.GET("/articolo/:id/", api.GETArticoli)
	router.GET("/collezione/:id/articoli", api.GETCollezione)
	return router
}
