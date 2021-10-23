default:
	go fmt
	go build
install:
	go fmt
	go build
	cp mathlang /usr/local/bin
