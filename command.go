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
	"net/textproto"
	"regexp"
	"strings"
)

// A Command represents an NNTP command received by a server or issued
// by a client.
type Command struct {
	// Keyword specifies the type of command.
	Keyword string

	// Args are the arguments to the command.
	Args []string
}

var spacesOrTabs = regexp.MustCompile("[ \t]+")

// ReadCommand reads and parses an incoming command from b.
func ReadCommand(b *bufio.Reader) (com *Command, err error) {
	tp := textproto.NewReader(b)
	com = new(Command)

	// Get first line: Keyword Arg1 Arg2...
	var s string
	if s, err = tp.ReadLine(); err != nil {
		return nil, err
	}

	parts := spacesOrTabs.Split(s, -1)
	com.Keyword = strings.ToUpper(parts[0])
	com.Args = parts[1:]

	return com, nil
}
