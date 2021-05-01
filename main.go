package main

import (
	"bytes"
	"encoding/json"
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
		break
	case "commit":
		processCommitResource()
		break
	default:
		PrintArray(os.Args)
	}
}

func processCommitResource() {
	// Each resource type is going to need it's own logic.
	url := fmt.Sprintf("/%s", os.Args[2])
	switch os.Args[2] {
	case "formation":
		resp, err := RESTPost(fmt.Sprintf("%s/%s", url, os.Args[3]), map[string]string{
			"TargetCount": os.Args[4],
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
		break
	}
}

func processStatusResource() {
	if len(os.Args) == 3 {
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

func RESTPost(path string, payload map[string]string) (string, error) {
	uri := fmt.Sprintf("%s%s", os.Getenv("CAPTAIN_URL"), path)
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(uri, "application/json", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(uri)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
