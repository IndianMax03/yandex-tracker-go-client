package client

import "fmt"

type Client struct{}

func New() (c *Client) {
	return &Client{}
}

func (c *Client) Greet(name string) (greet string) {
	greet = fmt.Sprintf("Greeting, %s!", name)
	return
}
