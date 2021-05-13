package lib

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	newline = "<br>"
)

func HttpsOs() string {
	var a []byte
	var err error

	switch os := runtime.GOOS; os{
	case "linux":
		a, err = exec.Command("vmstat").Output()
	case "darwin":
		a, err = exec.Command("vm_stat").Output()
	case "windows":
		a, err = exec.Command("tasklist").Output()
	default:
		a, err = exec.Command("echo", "Command Not Found").Output()
	}

	if err != nil {
		fmt.Println("Command error")
		os.Exit(1)
	}

	if runtime.GOOS == "linux" {
		b := strings.Fields(string(a))
		b1, b2, b3 := b[:6], b[6:23],  b[23:]
		m1 := strings.Join(b1, " ") + newline
		m2 := strings.Join(b2, " ") + newline
		m3 := strings.Join(b3, " ") + newline
		m4 := m1 + m2 + m3
		
		return m4
	}

	return string(a)
}