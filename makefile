default:
	go fmt
	go build
install:
	go fmt
	go build
	mkdir -p $HOME/.config/mathlang
	cp syntax_regexp.json $HOME/.config/mathlang/
	cp mathlang /usr/local/bin
