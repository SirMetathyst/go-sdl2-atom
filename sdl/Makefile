prog = atom component add -p sdl -n

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

testreport: test
	go tool cover -html=coverage.txt -o ./coverage.html

bench:
	go test -bench . -benchmem -benchtime 15s -cpu 1,4 -count 2

clean:
	rm -f ./coverage.html
	rm -f ./coverage.txt

generate:
	$(prog) SDLCreateWindow    -d title:string
	$(prog) SDLWindow          -d github.com/veandco/go-sdl2/sdl:window:*sdl.Window 