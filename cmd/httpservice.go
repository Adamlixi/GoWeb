package main

import "net/http"

func main()  {
	mux := http.NewServeMux()
	file := http.FileServer(http.Dir("/public"))
	
}
