package greet

import "fmt"

func Hello(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s", name)
	}

	return "Hello stranger"
}
