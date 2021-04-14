package main

import (
	_ "flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid parameter count")
		// TODO: print help text.
	}
	section := os.Args[1]
	switch section {
	case "airspace":
		processAirspace(os.Args)
	}
}

func processAirspace(args []string) {
	result, err := RESTGet("/airspaces")
	if err != nil {
		fmt.Printf("error occured in request: %s\n", err.Error())
		return
	}
	fmt.Println(result)
}

func RESTGet(path string) (string, error) {
	uri := fmt.Sprintf("%s%s", os.Getenv("CAPTAIN_URL"), path)
	resp, err := http.Get(uri)
	if err != nil {
		return "", fmt.Errorf("unable to perform REST action against %s: %w", path, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read server response: %w", err)
	}
	return string(body), nil
}
