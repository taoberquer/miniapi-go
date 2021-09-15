package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"bufio"
)

func hourHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(res, time.Now().Format("15:04"))
	default:
		res.WriteHeader(404)
	}
}

func addHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(res, "Something went bad")
			return
		}
		data := req.PostForm["author"][0] + ":" + req.PostForm["entry"][0]
		saveEntry(data)
		fmt.Fprintln(res, data)
	default:
		res.WriteHeader(404)
	}
}

func entriesHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, readEntries())
}

func saveEntry(data string) {
	saveFile, err := os.OpenFile("./save.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	defer saveFile.Close()

	w := bufio.NewWriter(saveFile)
	if err == nil {
		fmt.Fprintf(w, "%s\n", data)
	}
	w.Flush()
}

func readEntries() string {
	saveData, err := os.ReadFile("./save.data")
	if err == nil {
		return string(saveData)
	}
	return ""
}

func main() {
	http.HandleFunc("/", hourHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/entries", entriesHandler)
	http.ListenAndServe(":4567", nil)
}
