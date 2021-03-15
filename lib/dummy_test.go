package lib

import (
	"bytes"
	"io/ioutil"
	"runtime"
	"testing"
)

func TestDummy(t *testing.T) {
	var wants []byte
	if runtime.GOOS == "windows" {
		wants = []byte{0x31, 0x32, 0x33, 0x0a, 0x0d, 0x61, 0x62, 0x63, 0x0a, 0x0d}
	} else {
		wants = []byte{0x31, 0x32, 0x33, 0x0a, 0x61, 0x62, 0x63, 0x0a}
	}

	data, err := ioutil.ReadFile("../testdata/some.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, wants) {
		t.Fatalf("\nread : %x\nwants: %x\n", data, wants)
	}
}
