default:
	go fmt
	go build
install:
	go fmt
	go build
	cp mathlang /usr/local/bin
	mkdir -p /usr/local/share/mathlang
	cp syntax_regexp.json /usr/local/share/mathlang