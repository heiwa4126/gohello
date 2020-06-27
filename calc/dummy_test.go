package calc

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestDummy(t *testing.T) {

	wants := []byte{0x31, 0x32, 0x33, 0x0a, 0x61, 0x62, 0x63, 0x0a}
	data, err := ioutil.ReadFile("../examples/some.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, wants) {
		t.Fatalf("\nread : %x\nwants: %x\n", data, wants)
	}
}
