package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/csxuejin/kodo"
)

type Config struct {
	BookDir string       `json:"book_dir"`
	Kodo    *kodo.Config `json:"kodo"`
}

var (
	config     *Config
	kodoClient *kodo.Kodo
)

func init() {
	data, err := ioutil.ReadFile("kodo.json")
	if err != nil {
		log.Panicf("ioutil.ReadFile(kodo.json): %v", err)
		return
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("json.Unmarshal(): %v", err)
	}
	return
}

func main() {
	kodoClient = kodo.New(config.Kodo)
	dfs(config.BookDir, "")
}

func dfs(filePath, prefix string) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatalf("ioutil.ReadDir(%v): %v", filePath, err)
		return
	}

	for _, v := range files {
		if v.IsDir() {
			dfs(filePath+"/"+v.Name(), prefix+"/"+v.Name())
		} else {
			key := v.Name()
			if prefix != "" {
				key = prefix + "/" + key
			}

			err = kodoClient.PutFile(key, filePath+"/"+v.Name())
			if err != nil {
				log.Fatalf("upload failed: %v\n", key)
			} else {
				log.Printf("upload successfully: %v\n", key)
			}
		}
	}
}
