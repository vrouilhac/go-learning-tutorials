package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}

	msg := fmt.Sprintf(randomfmt(), name)
	return msg, nil
}

func Greetings(names []string) (map[string]string, error) {
	greetings := make(map[string]string)

	for _, name := range names {
		msg, err := Greeting(name)

		if err != nil {
			return nil, err
		}

		greetings[name] = msg
	}

	return greetings, nil
}

func randomfmt() string {
	formats := []string{
		"Hi, %s",
		"Hey %s!",
		"Hello %s",
		"Welcome back %s!",
	}

	return formats[rand.Intn(len(formats))]
}
