package utils

import (
	"errors"
)

func warpError(err error, message string) error {
	//ind := strings.Index(message, "hint")
	//if ind == -1 {
	//	return err
	//}
	//message = message[ind:]
	return errors.Join(err, errors.New(message))
}
