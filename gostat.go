 /*
 logrus Golang Logger
 https://github.com/sirupsen/logrus
 
 Released under the MIT license
 */

package main

import (
	"strconv"
	"fmt"
	"net/http"
	"os"
	"flag"

	"work/lib"
)

//Option
const defaultport = 8080
var version = "1.0.3"

func main() {

	var port int
	var versionbool bool

	//Flag Parses CLI Oprtion
	flag.IntVar(&port,"p",defaultport,"port number")
	flag.BoolVar(&versionbool,"v",false,"show version")
	flag.Parse()

	if port <= 1024 || port >= 49152 {
		fmt.Println("error : port range 1025 ~ 49151")
		os.Exit(1)
	}
	port_num := ":" + strconv.Itoa(port)

	//return version
	if versionbool {
		fmt.Println("version:",version)
		return
	}

	//Start HTTP
	fmt.Println("running on http://localhost",port_num,"/gostat")
	mux := Router()
	http.ListenAndServe(port_num, mux)
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/gostat",func(w http.ResponseWriter, r *http.Request) {
		lib.Logwrite(r)

		fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <body>
            %s
        </body>
        </html>
    `, lib.HttpsOs())
	})

	return mux
}
