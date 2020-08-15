package crypto

import (
	"testing"
)

const (
	Encoded = "uggcf://lbhgh.or/bUt5FWLEUN0"
	Decoded = "https://youtu.be/oHg5SJYRHA0"
)

func TestDecode(t *testing.T) {
	expected := Decoded
	got := Decode(Encoded, 13)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	// test -1 step
	got = Decode("gssor://xntst.ad/nGf5RIXQGZ0", -1)
	if got != Decoded {
		t.Errorf("expected %v, got %v", expected, got)
	}

	// test 25 step
	got = Decode("gssor://xntst.ad/nGf5RIXQGZ0", 25)
	if got != Decoded {
		t.Errorf("expected %v, got %v", expected, got)
	}

	// test 0 step
	got = Decode(Decoded, 0)
	if got != Decoded {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestEncode(t *testing.T) {
	expected := Encoded
	got := Encode(Decoded, 13)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	// test -1 step
	expected = "gssor://xntst.ad/nGf5RIXQGZ0"
	got = Encode(Decoded, -1)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	// test 25 step
	expected = "gssor://xntst.ad/nGf5RIXQGZ0"
	got = Encode(Decoded, 25)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	// test 0 step
	got = Encode(Decoded, 0)
	if got != Decoded {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

// разобраться gpg и запушить в новую репу
