package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello world"))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		jsonData, err := json.Marshal(items)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			if len(items) == 0 {
				w.Write([]byte("\"message\":\"belum ada data!!\""))
			} else {
				w.Write([]byte(jsonData))
			}
		}
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			item := Item{}
			json.Unmarshal(body, &item)
			nextId := len(items) + 1
			items = append(items, Item{Id: nextId, Name: item.Name})
			w.Write([]byte("\"message\":\"Item berhasil ditambahkan!!\""))
		}
	}
}

func postItems() {

}
