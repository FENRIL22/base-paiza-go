all:
	make test

test:
	go test .

debug:
	go test . -v
