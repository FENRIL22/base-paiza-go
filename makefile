all:
	make test

judge:
	#please set test file to out.dat
	go build p.go
	./p < judge.dat > out.dat
	@echo "-- Output ---------"
	cat out.dat
	@echo "-------------------"
	@echo ""

test:
	go test .

debug:
	go test . -v

clean:
	rm out.dat || true
	rm p || true
