package controllers

import (
    "github.com/robfig/revel"
    "gorp/app/models"
)

type Channels struct {
  App
}

func (c Channels) Index() revel.Result {
    return c.Render()
}

func (c Channels) Show(id string) revel.Result {
    conn := Pool.Get()
    defer conn.Close()

    //key := c.Params.Get("id")

    channel := models.GetChannel(conn, id)

    return c.RenderJson(channel.Messages)
}

func (c Channels) Append(id string, data string) revel.Result {
    conn := Pool.Get()
    defer conn.Close()

    models.AppendChannel(conn, id, data)

    return c.RenderJson("OK")
}

//func (c Channels) Create() revel.Result {
    //return c.Render()
//}
