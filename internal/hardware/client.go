package hardware

import (
	"os"
	"strings"
)

const (
	PlatformProfilePath = "/sys/firmware/acpi/platform_profile"
)

type State struct {
	CurrentProfile    string
	AvailableProfiles []string
}

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) read(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func (c *Client) write(path string, val string) error {
	return os.WriteFile(path, []byte(strings.TrimSpace(val)), 0644)
}

func (c *Client) GetState() (*State, error) {
	current, err := c.read(PlatformProfilePath)
	if err != nil {
		return nil, err
	}

	choices, err := c.read(PlatformProfilePath + "_choices")
	if err != nil {
		return &State{CurrentProfile: current, AvailableProfiles: []string{"quiet", "balanced", "performance"}}, nil
	}

	return &State{
		CurrentProfile:    current,
		AvailableProfiles: strings.Fields(choices),
	}, nil
}

func (c *Client) SetProfile(profile string) error {
	return c.write(PlatformProfilePath, profile)
}
