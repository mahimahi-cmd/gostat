package lib

import (
	"testing"
	"fmt"
)

func TestAPI(t *testing.T) {
	buff := GetLog()
	fmt.Println(buff)
}