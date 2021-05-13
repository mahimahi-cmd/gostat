package lib

import (
	"fmt"
	"testing"
)

func TestHttpOs(t *testing.T) {
	var i interface{} = HttpsOs()

	switch i.(type) {
	case string:
		fmt.Println("OK")
	default:
		t.Fatal("fatal")
	}
}