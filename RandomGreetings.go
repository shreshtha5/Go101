package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func RandomHello( name string) (string, error){
	if name == "" {
		return "", errors.New("Empty name, Please provide name")
    }

	if name == "Glady" {
		return "", errors.New("I only greet People with Indian Names :P ")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string{
	formats := [] string{
		"Hi %v. Welcome",
		"Learn Go asap!, %v",
		"Nice to meet you %v",
	}

	return formats[rand.Intn(len(formats))]

}
