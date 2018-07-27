package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/csxuejin/kodo"
)

type Config struct {
	BookDir string       `json:"book_dir"`
	Kodo    *kodo.Config `json:"kodo"`
}

var (
	kodoClient *kodo.Kodo
)

func main() {
	config, err := initConfig()
	if err != nil {
		log.Fatalf("initConfig(): %v", err)
		return
	}

	dirname := config.BookDir
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatalf("ioutil.ReadDir(%v): %v", dirname, err)
		return
	}

	kodoClient = kodo.New(config.Kodo)
	for _, v := range files {
		if v.IsDir() {
			helper(dirname+"/"+v.Name(), v.Name(), v)
		} else {
			err = kodoClient.PutFile(v.Name(), dirname+"/"+v.Name())
			if err != nil {
				log.Printf("upload failed: %v\n", dirname+"/"+v.Name())
			}
		}
	}

	return
}

func helper(filePath, prefix string, f os.FileInfo) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatalf("ioutil.ReadDir(%v): %v", filePath, err)
		return
	}

	for _, v := range files {
		if v.IsDir() {
			helper(filePath+"/"+v.Name(), prefix+"/"+v.Name(), v)
		} else {
			var key string
			if prefix == "" {
				key = v.Name()
			} else {
				key = prefix + "/" + v.Name()
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

func initConfig() (config *Config, err error) {
	data, err := ioutil.ReadFile("kodo.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &config)
	return
}
