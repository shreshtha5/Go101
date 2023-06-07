package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func MultipleHello( name string) (string, error){
	if name == "" {
		return "", errors.New("Empty name, Please provide name")
    }

	if name == "Glady" {
		return "", errors.New("I only greet People with Indian Names :P ")
	}

	message := fmt.Sprintf(randomFormatt(), name)
	return message, nil
}

func randomFormatt() string{
	formats := [] string{
		"Hi %v. Welcome",
		"Learn Go asap!, %v",
		"Nice to meet you %v",
	}

	return formats[rand.Intn(len(formats))]

}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
        message, err := MultipleHello(name)
        if err != nil {
            return nil, err
        }
	
		messages[name] = message
    }
    return messages, nil
}	
