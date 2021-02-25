package greetings

import "fmt"

// Hello returns a greetings for a supplied name
func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}