
.PHONY: build install clean rebuild

rebuild: clean build install  

build:
	mkdir -p build && go build -o ./build/loc ./src/main

install: ./build/loc
	cp ./build/loc /usr/local/bin

package: build
	mkdir -p ./build/deb/usr/local/bin && \
	mkdir -p ./build/deb/DEBIAN  && \
	cp ./build/loc ./build/deb/usr/local && \
	cp ./deb/control ./build/deb/DEBIAN && \
  	chown -R root:root ./build/deb && \	
	dpkg -b ./build/deb ./build/loc_linux_amd64_v1.0.deb

clean:
	rm -rf ./build

