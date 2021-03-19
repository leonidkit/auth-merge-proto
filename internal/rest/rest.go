package rest

import "github.com/gin-gonic/gin"

type API struct {
	router *gin.Engine
}

func New(engine *gin.Engine) *API {
	return &API{
		router: engine,
	}
}

func (a *API) endponts() {
	// TODO: endpoints init
}

func (a *API) ZamIOHandler(c *gin.Context) {
}
