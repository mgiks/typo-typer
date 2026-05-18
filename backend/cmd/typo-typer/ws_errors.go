package main

import (
	"github.com/coder/websocket"
)

const StatusInavalidData = 4000

func (app application) unsupportedDataClose(c Client) {
	c.conn.Close(websocket.StatusUnsupportedData, "")
}

func (app application) inavalidDataClose(c Client) {
	c.conn.Close(StatusInavalidData, "invalid data")
}

func (app application) abnormalClosureClose(c Client) {
	c.conn.Close(websocket.StatusAbnormalClosure, "")
}

func (app application) internalErrorClose(c Client, err error) {
	app.logger.Error("internal error close", "path", c.urlPath, "err", err)
	c.conn.Close(websocket.StatusInternalError, "internal error")
}

func (app application) normalClosureClose(c Client) {
	c.conn.Close(websocket.StatusNormalClosure, "normal closure")
}
