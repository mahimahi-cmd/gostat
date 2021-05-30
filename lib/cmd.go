package lib

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"runtime"
	"strings"
)

var (
	newline = "<br>"
)

const (
	//deray time for log write
	wait = 60 * time.Second
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
		
		CmdLog(b3)
		return m4
	}

	return string(a)
}

func Run() {
	ticker := time.NewTicker(wait)
	defer ticker.Stop()

	for{
		select{
		case <-ticker.C:
			HttpsOs()
		}
	}

}