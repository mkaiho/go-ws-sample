package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mkaiho/go-ws-sample/controller/web/handlers"
	"github.com/mkaiho/go-ws-sample/usecase"
)

func NoMatchPathHandler() handlers.Handler {
	return func(gc *gin.Context) {
		gc.Error(usecase.ErrNotFoundEntity).SetType(gin.ErrorTypePublic)
	}
}
