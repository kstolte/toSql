package toSql

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_SimpleList(t *testing.T) {
	in := `a
b
c
d
e
`
	got, err := ToInList(in)
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)

}
