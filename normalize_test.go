package arp

import "testing"

func TestNormalizeMACAddr(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{"00:00:00:00:00", "00:00:00:00:00"},
		{"00-00-00-00-00", "00:00:00:00:00"},
		{"0-0-0-0-0", "00:00:00:00:00"},
		{"0-00-0-00-0", "00:00:00:00:00"},
		{"09:00:26:e:37:61", "09:00:26:0e:37:61"},
		{"09-00-26-e-37-61", "09:00:26:0e:37:61"},
		{"FF-FF-FF-FF-FF", "ff:ff:ff:ff:ff"},
	}

	for _, ex := range examples {
		actual := normalizeMACAddr(ex.in)
		if actual != ex.out {
			t.Errorf("Expected normalizeMACAddr(%#v) to equal %#v, got %#v", ex.in, ex.out, actual)
		}
	}
}
