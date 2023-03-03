package py

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ValueError = errors.New("empty sequence")

type addable interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

type numeric interface {
	constraints.Integer | constraints.Float
}
