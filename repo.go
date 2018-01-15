package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Repo struct {
	Path string
}

func NewRepo(path string) *Repo {
	repo := Repo{
		Path: path,
	}

	return &repo
}

func (repo *Repo) Init() error {
	_, err := os.Stat(repo.Path)

	if os.IsNotExist(err) {
		file, err := os.Create(repo.Path)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	_, err = ioutil.ReadFile(repo.Path)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) Save(gif GIF) error {
	gifs, err := repo.All()
	if err != nil {
		return err
	}

	gif.ID = len(gifs) + 1
	// Append new GIF to collection
	gifs = append(gifs, gif)

	// Marshal collection and save it in YAML file
	yamlBlob, err := yaml.Marshal(gifs)
	err = ioutil.WriteFile(repo.Path, yamlBlob, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) ByTag(tag string) ([]GIF, error) {
	var result []GIF

	gifs, err := repo.All()
	if err != nil {
		return result, err
	}

	for _, gif := range gifs {
		for _, t := range gif.Tags {
			if tag == t {
				result = append(result, gif)
			}
		}
	}

	return result, nil
}

func (repo *Repo) Delete(ID int) error {
	var filtered []GIF
	gifs, err := repo.All()
	if err != nil {
		return err
	}

	for _, gif := range gifs {
		if gif.ID != ID {
			filtered = append(filtered, gif)
		}
	}

	yamlBlob, err := yaml.Marshal(filtered)
	err = ioutil.WriteFile(repo.Path, yamlBlob, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) All() ([]GIF, error) {
	var gifs []GIF

	source, err := ioutil.ReadFile(repo.Path)
	if err != nil {
		return gifs, err
	}

	err = yaml.Unmarshal(source, &gifs)
	if err != nil {
		return gifs, err
	}

	return gifs, nil
}
