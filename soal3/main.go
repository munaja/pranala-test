package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func fixData(input *user) {
	// email
	pos := strings.Index(input.Email, "@")
	if pos >= 0 {

		input.Email = strings.Replace(input.Email[0:pos], ".", "", -1) + string(input.Email[pos:])
	}

	// age
	input.Age++
}

func getData(input *user) {
	jsonFile, err := os.Open("file.json")
	if err != nil {
		panic("failed to open file, " + err.Error())
	}

	theBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		panic("failed read file, " + err.Error())
	}

	json.Unmarshal(theBytes, &input)
}

func main() {
	data := user{}
	getData(&data)
	fixData(&data)

	fmt.Println(data)
}
