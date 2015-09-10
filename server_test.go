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
