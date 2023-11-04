package main

import (
	"log"
	"net/http"
	"os"
	"io"
)

func read_file(fname string) (error,[]byte) {
	f, err__open_file := os.Open(fname)

	// check if the file does not exist
	if os.IsNotExist(err__open_file) {
		return err__open_file, []byte{}
	} 

	// check for any other errors
	if err__open_file != nil{
		log.Fatalf("[Error]: %s\n", err__open_file)
	}

	// read all data from file 
	data, err__read_data := io.ReadAll(f)
	if err__read_data != nil {
		return err__read_data, []byte{}
	}

	return nil, data
} 

func handle__FileServer(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Event] got a request\n")
	// get the requested file
	req_path := r.URL.Path

	if req_path == "/" {
		req_path = "/index.html"
	}
	
	base_dir := "./web"
	req_path = base_dir + req_path

	// read_file should return content-type to
	err__read_file, fdata := read_file(req_path)
	if err__read_file != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
	} else {
		w.WriteHeader(http.StatusOK)		
		// TODO: use content-type from read_file()
		w.Header().Set("Content-Type", "text/html")
		w.Write(fdata)
	}
}

func main() {
	// handle all incoming connection
	http.HandleFunc("/", handle__FileServer)

	log.Printf("[Event]: Server listening on port :5656\n")

	// start listening on port :5656
	err := http.ListenAndServe(":5656", nil)
	if err != nil {
		log.Fatalf("[Error]: %s\n", err)
	}
}
