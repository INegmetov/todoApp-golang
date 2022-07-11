package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inegmetov/todoApp-golang"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	var input todoApp.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
