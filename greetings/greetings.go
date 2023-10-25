// go function in a package

package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hello %v. Welcome !!", name)
	return message
}

func HelloE(name string) (string, error) {
	if name == ""{
		return "", errors.New("error: empty name")
	}
	message := fmt.Sprintf("Hello %v. Welcome !!", name)
	return message, nil
}