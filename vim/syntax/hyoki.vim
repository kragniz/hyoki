if exists("b:current_syntax")
  finish
endif

syntax match section "^\w*"
highlight link section todo

syntax match hashtag "#\w*"
highlight link hashtag constant

syntax match item "^\s*-"
highlight link item statement

syntax match subitem "^\s*\*"
highlight link subitem comment

let b:current_syntax = "hyoki"
