package lib

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"encoding/json"
)

type Data struct {
	label []int
	data []int
}

var (
	lines string
	line []string
	buf []string
	buf2 string
	buff = []string{"", "", "", "", "", "", "", "", "", ""}
	buff2 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fre []string
	fre2 string
	free2 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	cp []string
	cp2 string
	cpu2 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

//This func is retrun encoded json.
//From ../log/cmd.log file.
func GetLog() []byte {
	
	f, err := os.Open("./log/cmd.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = scanner.Text()

		buff[0] = buff[1]
		buff[1] = buff[2]
		buff[2] = buff[3]
		buff[3] = buff[4]
		buff[4] = buff[5]
		buff[5] = buff[6]
		buff[6] = buff[7]
		buff[7] = buff[8]
		buff[8] = buff[9]
		buff[9] = lines
		
		//Check scanner finis on end line
		if scanner.Err() != nil {
			panic(err)
		}

	}

	for n, i := range buff {
		line = strings.Fields(i)

		buf = strings.Split(line[3], "=")
		buf2 = buf[1]

		fre = strings.Split(line[4], "=")
		fre2 = fre[1]

		cp = strings.Split(line[5], "=")
		cp2 = cp[1]

		buff2[n], _ = strconv.Atoi(buf2)
		free2[n], _ = strconv.Atoi(fre2)
		cpu2[n], _ = strconv.Atoi(cp2)
	}

	i := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	
	bytes, _ := json.Marshal(&struct {
        Label []int	`json:"label"`
        Buff []int	`json:"buff"`
		Free []int	`json:"free"`
		Cpu []int	`json:"cpu"`
    }{
        Label:   i,
        Buff: buff2,
		Free: free2,
		Cpu:  cpu2,
    })
	return bytes
}