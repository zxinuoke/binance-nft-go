package bapi

import "testing"

func TestNftList(t *testing.T) {
	what, err := NFTMysteryBoxList()
	if err != nil {
		t.Fatalf("error=%v\n", err)
	}
	t.Logf("data len=%d\n", len(what.Data))
}
