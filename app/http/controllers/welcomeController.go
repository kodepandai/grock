package controllers

type WelcomeController struct{}

func (c *WelcomeController) Init() {
	// TODO: register controller middleweares here
}

func (c *WelcomeController) Index() string {
	return "HELLO GROCK"
}
