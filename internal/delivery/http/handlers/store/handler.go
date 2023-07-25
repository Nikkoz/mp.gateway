package store

import (
	"github.com/Nikkoz/mp.gateway/internal/delivery/http/error"
	"github.com/Nikkoz/mp.gateway/internal/useCase/interfaces"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	ucStore interfaces.Store
}

func New(ucStore interfaces.Store) *Handler {
	return &Handler{
		ucStore: ucStore,
	}
}

func (handler *Handler) Create(c *gin.Context) {
	store, access, err := makeData(c)
	if err != nil {
		error.SetError(c, http.StatusInternalServerError, err)

		return
	}

	ctx := context.New(c)
	response, err := handler.ucStore.Create(ctx, store, access)
	if err != nil {
		error.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusCreated, ToResponse(response))
}
