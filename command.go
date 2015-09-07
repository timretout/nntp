// Copyright (C) 2015 Tim Retout <tim@retout.co.uk>

package nntp

// A Command represents an NNTP command received by a server or issued
// by a client.
type Command struct {
	// Keyword specifies the type of command.
	Keyword string
}
