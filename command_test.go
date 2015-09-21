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
	"reflect"
	"strings"
	"testing"
)

func TestReadCommand(t *testing.T) {
	var readCommandTests = []struct {
		in  string
		com Command
		err error
	}{
		{"HELP\r\n", Command{Keyword: "HELP", Args: []string{}}, nil},
		{"help\r\n", Command{Keyword: "HELP", Args: []string{}}, nil},
		{ // treat keyword variants as arguments for now
			"LIST OVERVIEW.FMT\r\n",
			Command{Keyword: "LIST", Args: []string{"OVERVIEW.FMT"}},
			nil,
		},
		{ // separated by one or more space or TAB characters
			"NEWNEWS\tnews.*,sci.*  19990624   000000 \tGMT",
			Command{Keyword: "NEWNEWS", Args: []string{"news.*,sci.*", "19990624", "000000", "GMT"}},
			nil,
		},
	}

	for i, tt := range readCommandTests {
		com, err := ReadCommand(bufio.NewReader(strings.NewReader(tt.in)))
		if err != tt.err {
			t.Errorf("%d. got error %v, expected nil", i, err)
		}
		if com.Keyword != tt.com.Keyword {
			t.Errorf("%d. got command %v, expected %v", i, com.Keyword, tt.com.Keyword)
		}
		if !reflect.DeepEqual(com.Args, tt.com.Args) {
			t.Errorf("%d. got args %v, expected %v", i, com.Args, tt.com.Args)
		}
	}
}
