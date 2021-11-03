default: test
	@echo ""

gods_arr.go: gods_generic.go
	cat gods_generic.go | genny gen "Numeric=int,int32,int64,float32,float64" > gods_arr.go

test: gods_arr.go
	go test .
