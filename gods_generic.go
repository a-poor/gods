package gods

import (
	"github.com/cheekybits/genny/generic"
)

type Numeric generic.Type

type ArrNumeric struct {
	Data []Numeric
}

func (a ArrNumeric) String() string {
	return "I'm a Numeric array"
}
