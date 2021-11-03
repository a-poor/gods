package gods

import (
	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ndarr=Arr numeric=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"

type ndarr generic.Type
type numeric generic.Type

type ndarrnumeric struct {
	Data  []numeric
	Shape []int
}
