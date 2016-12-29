package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Post(postID int) revel.Result {
	return c.Render(postID)
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Projects() revel.Result {
	return c.Render()
}

func (c App) Encryption() revel.Result {
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}
