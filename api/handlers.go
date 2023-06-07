package api

import (
	"fmt"
	"net/http"
	jsontypes "weavestore/jsonTypes"
	"weavestore/kvs"
	"weavestore/messages"
	"weavestore/resp"

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
		return c.JSON(http.StatusInternalServerError, messages.Resp(fmt.Sprintf("key %s", object.Key), resp.Insert, resp.Fail))
	}

	respCode := h.kvStore.InsertItem(object.Key, object.Value)
	return c.JSON(http.StatusOK, messages.Resp(fmt.Sprintf("key %s", object.Key), resp.Insert, respCode))
}

func (h *Handler) Read(c echo.Context) error {
	object := &jsontypes.GetObject{}

	if err := c.Bind(object); err != nil {
		return c.JSON(http.StatusInternalServerError, messages.Resp(fmt.Sprintf("key %s", object.Key), resp.Read, resp.Fail))
	}

	value := h.kvStore.GetItem(object.Key)
	return c.JSON(http.StatusOK, messages.Resp(value, resp.Read, resp.NoOp))
}

func (h *Handler) DeleteObject(c echo.Context) error {
	object := &jsontypes.DelObject{}

	if err := c.Bind(object); err != nil {
		return c.JSON(http.StatusInternalServerError, messages.Resp(fmt.Sprintf("key %s", object.Key), resp.Delete, resp.Fail))
	}

	respCode := h.kvStore.DeleteItem(object.Key)
	return c.JSON(http.StatusOK, messages.Resp(fmt.Sprintf("key %s", object.Key), resp.Delete, respCode))
}

func (h *Handler) UpdateObject(c echo.Context) error {
	object := &jsontypes.InsertObject{}

	if err := c.Bind(object); err != nil {
		return c.JSON(http.StatusInternalServerError, messages.Resp(fmt.Sprintf("key %s", object.Key), resp.Update, resp.Fail))
	}

	respCode := h.kvStore.UpdateItem(object.Key, object.Value)
	return c.JSON(http.StatusOK, messages.Resp(fmt.Sprintf("key %s", object.Key), resp.Update, respCode))
}

func (h *Handler) UpdateBulk(c echo.Context) error {
	objects := &[]jsontypes.InsertObject{}

	if err := c.Bind(objects); err != nil {
		return c.JSON(http.StatusInternalServerError, messages.Resp(err, resp.Update, resp.Fail))
	}

	for _, obj := range *objects {
		h.kvStore.UpdateItem(obj.Key, obj.Value)
	}

	return c.JSON(http.StatusOK, messages.Resp("bulk operation", resp.Update, resp.Success))
}
