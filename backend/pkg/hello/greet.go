package hello

import (
	"fmt"
)

func Greet(name *string) string {
	if len(*name) == 0 {
		return fmt.Sprintf("Hello, Anonymous")
	} else {
		return fmt.Sprintf("Hello, %s", *name)
	}
}