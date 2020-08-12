package main

import (
	"encoding/json"
	"fmt"
	"gogithub/config"
	"gogithub/github"
	"gogithub/model"
	"log"
	"net/http"
	"strings"
	"time"
)

var cacheSummary []byte
var lastCache time.Time
var cacheTopStar []byte
var lastCacheTopStar time.Time
var cacheJobs chan string

const (
	cacheHours        = 24 * 31
	cacheHoursTopStar = 24 * 31

	cacheTypeSummary string = "summary"
	cacheTypeTopStar string = "topstar"
)

// ProfilePayload for profile response payload
type ProfilePayload struct {
	Username      string                     `json:"username"`
	StarCount     int                        `json:"star_count"`
	RepoCount     int                        `json:"repo_count"`
	ForkCount     int                        `json:"fork_count"`
	LanguageCount int                        `json:"language_count"`
	LanguageMap   map[string]int32           `json:"language_map"`
	AvatarURL     string                     `json:"avatar_url"`
	TopRepo       *github.UserRepositoryEdge `json:"top_repo"`
}

func checkCache(lastCache time.Time, cacheHours int, cacheBytes []byte) bool {
	hoursElapsed := time.Since(lastCache).Hours()
	return hoursElapsed < float64(cacheHours) && len(cacheBytes) != 0
}

func processMemoryCache(jobType string) {
	if jobType == cacheTypeSummary {
		if checkCache(lastCache, cacheHours, cacheSummary) {
			return
		}
		data, err := github.FetchTopUserSummary()
		if err != nil {
			fmt.Println("ERR", err)
			return
		}
		// Prevents caching rate limit error data
		if data["message"] != nil {
			fmt.Println("ERR", data["message"])
			return
		}
		b, _ := json.Marshal(data)
		cacheSummary = b
		lastCache = time.Now()
		fmt.Println("summary cache is now updated!")
	} else if jobType == cacheTypeTopStar {
		if checkCache(lastCacheTopStar, cacheHoursTopStar, cacheTopStar) {
			return
		}
		dataTopStar, err := github.FetchAllStars()
		if err != nil {
			fmt.Println("ERR", "FetchAllStars:", err)
			return
		}
		// Prevents caching empty result
		if len(dataTopStar) == 0 {
			fmt.Println("ERR", "FetchAllStars: Empty result")
			return
		}
		data := &model.ResponsePayload{
			Data: dataTopStar,
		}
		b, _ := json.Marshal(data)
		cacheTopStar = b
		lastCacheTopStar = time.Now()
		fmt.Println("topstar cache is now updated!")
	}
}

func checkMemoryCache(cacheJobs <-chan string) {
	for job := range cacheJobs {
		processMemoryCache(job)
	}
}

func handleGithubProfile(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	urlPath = strings.TrimPrefix(urlPath, "/")
	urlSeg := strings.SplitN(urlPath, "/", 3)
	username := urlSeg[2]

	data, err := github.FetchAllRepos(username)

	if err != nil {
		fmt.Println("ERR handleGitHubProfile:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		payload := model.ResponsePayload{
			Error: "Error fetch profile",
		}
		b, _ := json.Marshal(payload)
		w.Write(b)
		return
	}

	profilePayload := ProfilePayload{
		Username:      username,
		AvatarURL:     data.AvatarURL,
		StarCount:     data.StarCount,
		RepoCount:     data.RepoCount,
		ForkCount:     data.ForkCount,
		LanguageCount: len(data.LanguageMap),
		LanguageMap:   data.LanguageMap,
		TopRepo:       data.TopRepo,
	}

	payload := model.ResponsePayload{
		Data: profilePayload,
	}

	b, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func handleGithubSummary(w http.ResponseWriter, r *http.Request) {
	// Asynchronously check cache and update it when it's necessary
	cacheJobs <- cacheTypeSummary
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Gogithub-Cache", "true")
	w.Write(cacheSummary)
	return
}

func handleTopStars(w http.ResponseWriter, r *http.Request) {
	// Asynchronously check cache and update it when it's necessary
	cacheJobs <- cacheTypeTopStar
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Gogithub-Cache", "true")
	w.Write(cacheTopStar)
	return
}

func handleTopCompanies(w http.ResponseWriter, r *http.Request) {
	// Asynchronously check cache and update it when it's necessary
	cacheJobs <- cacheTypeTopStar
	data, _ := github.FetchAllCompanies(cacheTopStar)
	resp := &model.ResponsePayload{
		Data: data,
	}
	b, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Gogithub-Cache", "true")
	w.Write(b)
	return
}

func main() {
	http.HandleFunc("/gh/summary", handleGithubSummary)
	http.HandleFunc("/gh/profile/", handleGithubProfile)
	http.HandleFunc("/gh/topstars", handleTopStars)
	http.HandleFunc("/gh/topcompanies", handleTopCompanies)

	// For testing purpose
	// http.HandleFunc("/gh/test", handleTest)

	// Initialize chan
	cacheJobs = make(chan string, 100) // maximum number of jobs that can be queued
	go checkMemoryCache(cacheJobs)     // start coroutine

	// Initialize cache
	processMemoryCache(cacheTypeSummary)
	processMemoryCache(cacheTypeTopStar)
	if len(cacheSummary) == 0 || len(cacheTopStar) == 0 {
		log.Fatal("Failed to initialize in-memory cache!")
		return
	}

	addr := config.WebAddress()
	log.Println("Web server will be listening at " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
