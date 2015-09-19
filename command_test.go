// Copyright (C) 2015 Tim Retout <tim@retout.co.uk>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package nntp

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadCommand(t *testing.T) {
	var readCommandTests = []struct {
		in  string
		com Command
		err error
	}{
		{"HELP\r\n", Command{Keyword: "HELP"}, nil},
	}

	for i, tt := range readCommandTests {
		com, err := ReadCommand(bufio.NewReader(strings.NewReader(tt.in)))
		if err != tt.err {
			t.Errorf("%d. got error %v, expected nil", i, err)
		}
		if *com != tt.com {
			t.Errorf("%d. got command %v, expected %v", i, com, tt.com)
		}
	}
}
