package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	vdf "github.com/wakeful-cloud/vdf"
)

func addIconToGame(user User) {
	shortcutsVdf := filepath.Join(user.Dir, "config", "shortcuts.vdf")
	if _, err := os.Stat(shortcutsVdf); err != nil {
		return
	}
	shortcutBytes, err := ioutil.ReadFile(shortcutsVdf)
	if err != nil {
		return
	}

	//Read VDF
	vdfMap, err := vdf.ReadVdf(shortcutBytes)

	if err != nil {
		panic(err)
	}

	//Covert to JSON
	rawJSON, err := json.Marshal(vdfMap)

	if err != nil {
		panic(err)
	}

	//Log
	log.Print(string(rawJSON))

	//Write VDF
	rawVdf, err := vdf.WriteVdf(vdfMap)

	if err != nil {
		panic(err)
	}

	//Write the file
	err = os.WriteFile("./out.vdf", rawVdf, 0666)

	if err != nil {
		panic(err)
	}
}
