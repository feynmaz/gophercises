package main

import (
	"fmt"

	"github.com/feynmaz/gophercises/exercise_4_html_link_parser/parser"
)

func main() {
	input :=
		`
	<div>
		body
	</div>
	<a href="/cat">cat</a>
	<a href="/dog">
		<span>Something in a span</span>
		Text not in a span
		<b>Bold text!</b>
	</a>
	<a href="/nested">
		some text
		<a href="/internal">more text</a>
	</a>
	`
	links := parser.ParseLinks(input)
	fmt.Printf("%#v \n", links)
}
