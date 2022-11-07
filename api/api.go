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
	GetArticoliCollezioniREPO() []core.Item
	GetArticoliCollezioneByIdREPO(int) []core.Item
}

type EcommerceApi struct {
	rep Repository //accetto tutto quello che e compatibile con l'interfaccia
}

func (e *EcommerceApi) GETArticoliCollezioneById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil) //quando non metto il numero della collezione di cui voglio gli articoli
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	res := e.rep.GetArticoliCollezioneByIdREPO(idInt)
	c.JSON(http.StatusOK, res)
}

func (e *EcommerceApi) GETAllCollezioni(c *gin.Context) {
	c.JSON(http.StatusOK, e.rep.GetAllCollezioniREPO())
}

func (e *EcommerceApi) GETCollezione(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok { //se non viene specificato l'id della collezione
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Nessun parametro inserito",
		})
	} else {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, e.rep.GetArticoliCollezioneByIdREPO(idInt))
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
	router.GET("/collezione/", api.GETAllCollezioni) //creo gli handler
	router.GET("/collezione/:id", api.GETArticoliCollezioneById)
	router.GET("/articolo/", api.GETArticoli)
	router.GET("/articolo/:id/", api.GETArticoli)
	return router
}
