all: hyoki

hyoki: hyoki.go
	go build hyoki.go

install: hyoki
	cp hyoki /usr/local/bin/
	cp bash/hyoki /etc/bash_completion.d/

plugin: ~/.vim/syntax/hyoki.vim ~/.vim/ftdetect/hyoki.vim

~/.vim/syntax/hyoki.vim: vim/syntax/hyoki.vim
	cp -Rv vim/syntax/* ~/.vim/syntax/

~/.vim/ftdetect/hyoki.vim: vim/ftdetect/hyoki.vim
	cp -Rv vim/ftdetect/* ~/.vim/ftdetect/
