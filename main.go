package main

import (
	_ "flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid parameter count")
		// TODO: print help text.
	}

	switch os.Args[1] {
	case "status":
		processStatusResource()
	case "commit":
		processCommitResource()
	default:
		PrintArray(os.Args)
	}
}

func processCommitResource() {
	// Each resource type is going to need it's own logic.
	var payload string
	url := fmt.Sprintf("/%s", os.Args[2])
	switch os.Args[2] {
	case "airspace":
		
	}
}

func processStatusResource() {
	if(len(os.Args) == 3) {
		// Get the status of all of resource
		result, err := RESTGet(fmt.Sprintf("/%ss", os.Args[2]))
		if err != nil {
			fmt.Printf("unable to print %s info: %s\n", os.Args[2], err.Error())
			return
		} else {
			fmt.Println(result)
		}
	} else {
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("%s is not a valid ID format", os.Args[3])
			return
		}
		result, err := RESTGet(fmt.Sprintf("/%s/%d", os.Args[2], id))
		if err != nil {
			fmt.Printf("unable to print %s info: %s\n", os.Args[2], err.Error())
		}
		fmt.Println(result)
	}
}

func PrintArray(args []string) {
	for i := 0; i < len(args); i++ {
		fmt.Printf("'%s'\n", args[i])
	}
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
