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
	"testing"

	"github.com/diocles/nntp/nntptest"
)

func TestServerHandler(t *testing.T) {
	handler := serverHandler{}

	w := nntptest.NewRecorder()

	c := Command{Keyword: "HELP"}
	handler.ServeNNTP(w, &c)

	if w.Code != 100 {
		t.Errorf("HELP should return status 100")
	}
	if w.Body.String() == "" {
		t.Errorf("HELP body should not be empty")
	}
}
