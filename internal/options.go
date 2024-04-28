package internal

import "fmt"

type Options struct {
	Port       int
	ServerAddr string
	data       string
}

func (o Options) Validate() error {
	if o.Port == 0 {
		return fmt.Errorf("Port is required")
	}
	return nil
}
