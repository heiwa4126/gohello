package calc

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDummy(t *testing.T) {

	wants := []byte{0x31, 0x32, 0x33, 0x0a, 0x61, 0x62, 0x63, 0x0a}
	data, err := ioutil.ReadFile("../examples/some.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("read : %x\n", data)
	fmt.Printf("wants: %x\n", wants)
	t.Fatal("TestDummy()")
}
