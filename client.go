package minecomm

type Client struct {
	playerName string
}

func NewClient(playerName string) (*Client, error) {
	return &Client{
		playerName: playerName,
	}, nil
}

func (client *Client) initialize() {

}

type ConnectedMsg struct{}
type ErrorMsg error

// Add more methods for handling Minecraft protocol here
