default:
	go fmt
	go build
install:
	go fmt
<<<<<<< HEAD
	go build --buildvcs=false
=======
	go build
	mkdir -p $HOME/.config/mathlang
	cp syntax_regexp.json $HOME/.config/mathlang/
>>>>>>> parent of aaa7ce7 (Completed project)
	cp mathlang /usr/local/bin
