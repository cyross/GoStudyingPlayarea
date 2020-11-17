package trans

import "testing"

func TestTransportSlow(t *testing.T) {
	expect := "Slow"
	actual := TransportSlow()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
