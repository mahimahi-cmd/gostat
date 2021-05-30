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

	go lib.Run()

	//Start HTTP
	fmt.Println("running on http://localhost",port_num,"/gostat")
	http.HandleFunc("/gostat", func(w http.ResponseWriter, r *http.Request) {
		lib.Logwrite(r)
	
	http.ServeFile(w, r, "./view/gostat.html")
	})

	http.HandleFunc("/json", JsHubdleFunc)

	//mux := Router()
	http.ListenAndServe(port_num, nil)

}

func JsHubdleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write(lib.GetLog())
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/gostat",func(w http.ResponseWriter, r *http.Request) {
	//lib.Logwrite(r)
	
	//http.ServeFile(w, r, "./log/cmd.log")
	//fmt.Fprintf(w, lib.GetLog())
	})

	return mux
}