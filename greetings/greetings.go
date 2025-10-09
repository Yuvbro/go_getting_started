package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name is invalid")
	}
	//msg := fmt.Sprintf("Hello, %s from Yuval Brosh!", name)
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = msg
	}

	return messages, nil
}

func randomFormat() string {
	formatsWithDynamicValue := []string{
		"Hi, %v. Welcome from Brosh!",
		"Great to see you, %v!",
		"Hey Mate, %v!",
	}

	return formatsWithDynamicValue[rand.Intn(len(formatsWithDynamicValue))]

}
