package controllers

import (
	"github.com/robfig/revel"
	"gorp/app/models"
	"log"
)

type Channels struct {
  App
}

func (c Channels) Index() revel.Result {
	return c.Render()
}

func (c Channels) Show(id string) revel.Result {
	log.Print(id)
	channel := models.NewChannel(c.Conn, id)
	channel.Get()

	return c.RenderJson(channel.Messages)
}

func (c Channels) Pop(id string) revel.Result {
	//key := c.Params.Get("id")
	log.Print(id)
	channel := models.NewChannel(c.Conn, id)
	channel.Pop()

	return c.RenderJson(channel.Messages)
}


func (c Channels) Append(id string, message string) revel.Result {

	channel := models.NewChannel(c.Conn, id)
	channel.Append(message)

	return c.RenderJson("OK")
}

//func (c Channels) Create() revel.Result {
	//return c.Render()
//}
