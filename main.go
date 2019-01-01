package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	port := flag.String("port", "9999", "http port")
	dir := flag.String("dir", "./", "server path")
	flag.Parse()
	absPath, err := filepath.Abs(*dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("start server port:" + *port + " path :" + absPath)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*dir))))
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
