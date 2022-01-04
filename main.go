package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func determine_config_path() (string, error) {
	f, err := ioutil.ReadDir("~/.config")
	if err != nil {
		return "", err
	}
	for _, entry := range f {
		if entry.IsDir() && (entry.Name() == "wmsel") {
			return "~/.config/wmsel", nil
		}
	}
	err = os.Mkdir("~/.config/wmsel", 0755)
	if err != nil {
		return "", err
	}
	return "~/.config/wmsel", nil
}

func main() {
	CONFIG_PATH, err := determine_config_path()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CONFIG_PATH)
}
