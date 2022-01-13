
.PHONY: build install clean rebuild

rebuild: clean build install  

build:
	mkdir -p build && go build -o ./build/loc ./src/main

install: ./build/loc
	cp ./build/loc /usr/local/bin

clean:
	rm -rf ./build

