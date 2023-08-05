package util

import (
	"fmt"
	"errors"
)


func NotYetImplemented(feat string) error {
	return errors.New(fmt.Sprintf("Sorry, %s is not yet implemented. Please check back for a later release.", feat))
}

func GenericError(msg string) error {
	return errors.New(msg)
}