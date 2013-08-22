all: hyoki

hyoki: hyoki.go
	go build hyoki.go

install:
	cp hyoki /usr/local/bin/

viminstall:
	cp -Rv vim/* ~/.vim
