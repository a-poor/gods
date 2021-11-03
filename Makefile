default: test
	@echo ""

test:
	go generate
	go test .
