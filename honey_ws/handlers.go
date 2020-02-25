package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func generateRepoName(url string) string {
	gh := strings.Split(url, "/")
	ghName := gh[len(gh)-1]
	timeNow := time.Now().Unix()
	repoName := fmt.Sprintf("%s_%d", ghName, timeNow)
	return repoName
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// HoneyProcess ...
func HoneyProcess(w http.ResponseWriter, r *http.Request) {
	var repo RepoCreate

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("net/http: (error) while reading body: %v\n", err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panicf("net/http: (error) while closing the reader: %v\n", err)
	}
	if err := json.Unmarshal(body, &repo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panicf("http/json: (error) while encoding the error: %v\n", err)
		}
	}

	if repo.Key != HoneyKey {
		log.Printf("net/http: (error) wrong key")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	repoName := generateRepoName(repo.URL)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := GitClone(repo.URL, repoName); err != nil {
		log.Panicf("cloning: (error) %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panicf("http/json: (error) while encoding the error: %v\n", err)
		}

	}
	log.Printf("cloning: (success) repo has just been cloned in /tmp/%s.", repoName)
	w.WriteHeader(http.StatusOK)
}
