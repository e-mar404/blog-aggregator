package main 

import "fmt"

type command struct {
	name string
	arguments []string
}

type commands struct {
	available map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.available[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.available[cmd.name]
	if !ok {
		return fmt.Errorf("there is no available handler for this command")
	}
	if err := handler(s, cmd); err != nil {
		return fmt.Errorf("error running command %s: %v", cmd.name, err)
	}
	return nil
}

func newHandlerList() map[string]func(*state, command) error {
	return make(map[string]func(*state, command) error)
}
