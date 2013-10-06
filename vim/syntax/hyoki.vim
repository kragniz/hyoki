if exists("b:current_syntax")
  finish
endif

syntax match section "^\(\w\|\-\)*"
highlight link section todo

syntax match hashtag "#\w*"
highlight link hashtag constant

syntax match item "^\s*-"
highlight link item comment

syntax match subitem "^\s*\*"
highlight link subitem comment

syntax match done "\[done\]"
highlight link done todo

let b:current_syntax = "hyoki"
