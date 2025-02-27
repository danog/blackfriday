//
// Blackfriday Markdown Processor
// Available at http://github.com/danog/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
// Unit tests for full document parsing and rendering
//

package blackfriday

import "testing"

func TestDocument(t *testing.T) {
	t.Parallel()
	var tests = []string{
		// Empty document.
		"",
		"",

		" ",
		"",

		// This shouldn't panic.
		// https://github.com/danog/blackfriday/issues/172
		"[]:<",
		"<p>[]:&lt;</p>\n",

		// This shouldn't panic.
		// https://github.com/danog/blackfriday/issues/173
		"   [",
		"<p>[</p>\n",
	}
	doTests(t, tests)
}
