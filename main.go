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

}

func RESTGet() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/airspaces", os.Getenv("CAPTAIN_URL")))
	if err != nil {
		fmt.Printf("Unable to connect to Captain cluster at %s", os.Getenv("CAPTAIN_URL"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Unable to read server response.")
		return
	}
	fmt.Println(string(body))
}
