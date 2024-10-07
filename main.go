package main

import (
	"ascii-art/functions"
	"fmt"
	"html/template"
	"net/http"

	"strings"
)


type Data struct {
	Str string
	Banner string
	Res string
	A	template.HTML
}


func processHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Data
	data.Str = r.FormValue("data")
	if len(data.Str) > 200 {
		http.Error(w, "Input data exceeds 200 characters limit.", http.StatusBadRequest)
		return
	}

	data.Banner = r.FormValue("banner")
	if !function.BannerExists(data.Banner) {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	data.Str = strings.ReplaceAll(data.Str, "\r\n", "\n")
	
	data.Res = function.TraitmentData(data.Banner, data.Str)
	if data.Res == "" { // If TraitmentData failed to generate the result
		http.Error(w, "Internal Server Error: Failed to process data.", http.StatusInternalServerError)
		return
	}
	data.A = template.HTML(strings.ReplaceAll(data.Res, "\n", "<br>"))

	if err := temp.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func main() {

	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/ascii-art", processHandler)
	fmt.Println("Server is running at http://localhost:8088")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}