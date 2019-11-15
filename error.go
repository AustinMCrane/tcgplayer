package tcgplayer

import (
	"errors"
	"strings"
)

func wrapResponseErrors(errStrs []string) error {
	return errors.New(strings.Join(errStrs, ", "))
}
