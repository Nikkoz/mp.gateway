package name

import "github.com/pkg/errors"

const (
	MinLength = 3
	MaxLength = 100
)

var ErrWrongLength = errors.Errorf("store's name must be greater or equal to %d and less or equal to %d count of characters", MinLength, MaxLength)

type Name string

func New(n string) (*Name, error) {
	ln := len([]byte(n))
	if ln < MinLength || ln > MaxLength {
		return nil, ErrWrongLength
	}

	name := Name(n)
	return &name, nil
}

func (sm *Name) String() string {
	return string(*sm)
}
