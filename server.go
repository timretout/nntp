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

type Handler interface {
	ServeNNTP(ResponseWriter, *Command)
}

// A ResponseWriter interface is used by an NNTP handler to
// construct an NNTP response.
type ResponseWriter interface {
	// Write writes the data to the connection as part of an NNTP
	// reply.  If WriteHeader has not yet been called, Write calls
	// WriteHeader(nntp.StatusOK) before writing the data.
	Write([]byte) (int, error)

	// WriteHeader sends an NNTP response header with status code.
	// If WriteHeader is not called explicitly, the first call to Write
	// will trigger an implicit WriteHeader(nntp.StatusInternalFault).
	WriteHeader(int)
}

// A Server represents an NNTP server.
type Server struct {
}
