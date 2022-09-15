package module

import "testing"

func TestSayTheQuote(t *testing.T) {
	if SayTheQuote(0) != "1+1 = 2" {
		t.Fail()
	}
}
