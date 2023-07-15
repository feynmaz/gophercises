package parser

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParseLinks(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		links []Link
	}{
		{
			name:  "empty",
			input: "",
			links: []Link{},
		},
		{
			name:  "simple link",
			input: `<a href="/dog">the name of the dog is Max</a>`,
			links: []Link{
				{
					Href: "/dog",
					Text: "the name of the dog is Max",
				},
			},
		},
		{
			name: "single link with other attributes",
			input: `<div> 
			<span>some text<span>
			<a href="/dog">the name of the dog is Max</a>
			</div>`,
			links: []Link{
				{
					Href: "/dog",
					Text: "the name of the dog is Max",
				},
			},
		},
		{
			name: "multiple links with other attributes",
			input: `<div> 
			<span>some text<span>
			<a href="/dog">the name of the dog is Max</a>
			</div>
			<a href="/cat">the name of the cat is Polly</a>
			`,
			links: []Link{
				{
					Href: "/dog",
					Text: "the name of the dog is Max",
				},
				{
					Href: "/cat",
					Text: "the name of the cat is Polly",
				},
			},
		},
	}
	

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			links := ParseLinks(tc.input)

			assert.Equal(t, links, tc.links)
		})
	}
}
