all: hyoki

hyoki: hyoki.go
	go build hyoki.go

install:
	cp hyoki /usr/local/bin/
	cp bash/hyoki /etc/bash_completion.d/

viminstall:
	cp -Rv vim/* ~/.vim
