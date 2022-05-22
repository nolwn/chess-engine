package board

import "errors"

type coord uint8

func (c coord) typeError() error {
	if c >= boardSize {
		return errors.New("rank and file must be numbers between 0 and 7")
	}

	return nil
}
