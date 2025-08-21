package e

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %q", msg, err)
}
