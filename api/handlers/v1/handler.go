package v1

import (
	log "product/pkg/logger"
	"product/storage"

	"github.com/gin-gonic/gin"
)

type handlers struct{
	storage  storage.StorageI
	log 	 log.Log
}

type Handlers struct{
	Storage  storage.StorageI
	Log 	 log.Log
}

func NewHandler(h Handlers) handlers{
	return handlers{h.Storage, h.Log}
}


func (h *handlers) Ping(ctx *gin.Context) {
	ctx.JSON(200,map[string]string{"meesage":"pong"})
}

