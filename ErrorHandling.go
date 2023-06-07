package greetings

import "fmt" 
import "errors"

func HelloNew (name string) (string, error){

	if name == "" {
		return "", errors.New("Empty name, Please provide name")
    }

	if name == "Glady" {
		return "", errors.New("I only greet People with Indian Names :P ")
	}

	message := fmt.Sprintf("Hi , %v. welcome", name)
	return message, nil

}

