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
	"html/template"
	"log"

	"work/lib"
)

//Option
const defaultport = 8080
const defaultip = "localhost"
var version = "1.1.0 β"

func main() {

	var port int
	var versionbool bool
	var ip string
	var tpl *template.Template

	//Flag Parses CLI Oprtion
	flag.IntVar(&port,"p",defaultport,"port number")
	flag.BoolVar(&versionbool,"v",false,"show version")
	flag.StringVar(&ip,"i",defaultip,"ip address of this server")
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
	fmt.Println("running on http://",ip ,port_num,"/gostat")
	http.HandleFunc("/gostat", func(w http.ResponseWriter, r *http.Request) {
		lib.Logwrite(r)

		tpl = template.Must(template.ParseFiles("./view/gostat.html"))
		if err := tpl.ExecuteTemplate(w, "gostat.html", ip); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/json", JsHubdleFunc)

	//mux := Router()
	http.ListenAndServe(port_num, nil)

}

func JsHubdleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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