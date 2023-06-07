package api

import (
	"fmt"
	"net/http"
	httperrors "weavestore/httpErrors"
	jsontypes "weavestore/jsonTypes"
	"weavestore/kvs"
	"weavestore/messages"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	kvStore *kvs.KVS
}

func NewAPIInstance() *Handler {
	return &Handler{
		kvStore: kvs.NewStore(
			kvs.WithMaxRAMSize(1000),
		),
	}
}

func (h *Handler) InsertObject(c echo.Context) error {
	object := &jsontypes.InsertObject{}

	if err := c.Bind(object); err != nil {
		return c.JSON(http.StatusInternalServerError, httperrors.ErrInsertionFailed)
	}

	respCode := h.kvStore.InsertItem(object.Key, object.Value)
	return c.JSON(http.StatusOK, messages.Resp(fmt.Sprintf("key %s", object.Key), respCode))
}

func (h *Handler) Read(c echo.Context) error {
	object := &jsontypes.GetObject{}

	if err := c.Bind(object); err != nil {
		return c.JSON(http.StatusInternalServerError, httperrors.ErrReadFailed)
	}

	value := h.kvStore.GetItem(object.Key)
	return c.JSON(http.StatusOK, value)
}

func (h *Handler) DeleteObject(c echo.Context) error {
	object := &jsontypes.DelObject{}

	if err := c.Bind(object); err != nil {
		return c.JSON(http.StatusInternalServerError, httperrors.ErrDelFailed)
	}

	respCode := h.kvStore.DeleteItem(object.Key)
	return c.JSON(http.StatusOK, messages.Resp(fmt.Sprintf("key %s", object.Key), respCode))
}

func (h *Handler) UpdateObject(c echo.Context) error {
	object := &jsontypes.InsertObject{}

	if err := c.Bind(object); err != nil {
		return c.JSON(http.StatusInternalServerError, httperrors.ErrUpdateFailed)
	}

	respCode := h.kvStore.UpdateItem(object.Key, object.Value)
	return c.JSON(http.StatusOK, messages.Resp(fmt.Sprintf("key %s", object.Key), respCode))
}

func (h *Handler) UpdateBulk(c echo.Context) error {
	objects := &[]jsontypes.InsertObject{}

	if err := c.Bind(objects); err != nil {
		return c.JSON(http.StatusInternalServerError, httperrors.ErrUpdateFailed)
	}

	var respCode int
	for _, obj := range *objects {
		respCode = h.kvStore.UpdateItem(obj.Key, obj.Value)
	}

	return c.JSON(http.StatusOK, messages.Resp("bulk operation completed", respCode))
}
