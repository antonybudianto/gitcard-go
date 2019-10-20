package main

import (
	"encoding/json"
	"fmt"
	"gogithub/github"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing argument for username. Example: ./githubcli antonybudianto")
		os.Exit(1)
	}
	username := os.Args[1]

	data, err := github.FetchAllRepos(username)

	if err != nil {
		fmt.Println("Error fetch all repo", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", username)
	fmt.Printf("%d stars\n", data.StarCount)
	fmt.Printf("%d repos\n", data.RepoCount)
	fmt.Printf("%d forks\n", data.ForkCount)
	b, _ := json.MarshalIndent(data.LanguageMap, "", "  ")
	str := string(b)
	fmt.Printf("LangMap: %v\n", str)
}
