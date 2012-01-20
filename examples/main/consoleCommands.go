package main

import (
	"fmt"
	"github.com/teomat/mater/engine"
	"os"
)

type consoleCommand func(scene *engine.Scene, params []string)

var commands = map[string]func(*engine.Scene, []string){
	"save":    command_save,
	"load":    command_load,
	"exit":    command_quit,
	"quit":    command_quit,
	"pause":   command_pause,
	"unpause": command_unpause,
	"resume":  command_unpause,
	"help":    command_help,
}
var commandNames []string

func init() {
	commandNames = make([]string, 0, len(commands))
	for key, _ := range commands {
		commandNames = append(commandNames, key)
	}
}

var lastSave string

func command_save(scene *engine.Scene, params []string) {
	var path string
	if len(params) < 1 {
		if lastSave != "" {
			path = lastSave
		} else {
			fmt.Printf("Usage: save <filename>\n")
			return
		}
	} else {
		path = params[0]
	}

	fmt.Printf("Saving to %v\n", path)
	err := saveScene(path)

	if err == nil {
		lastSave = path
	}
}

func command_load(scene *engine.Scene, params []string) {
	var path string
	if len(params) < 1 {
		if lastSave != "" {
			path = lastSave
		} else {
			fmt.Printf("Usage: load <filename>\n")
			return
		}
	} else {
		path = params[0]
	}

	fmt.Printf("Loading from %v\n", path)
	Settings.Paused = true
	err := loadScene(path)

	if err == nil {
		lastSave = path
	}
}

func command_quit(scene *engine.Scene, params []string) {
	os.Exit(0)
}

func command_pause(scene *engine.Scene, params []string) {
	Settings.Paused = true
}

func command_unpause(scene *engine.Scene, params []string) {
	Settings.Paused = false
}

func command_help(scene *engine.Scene, params []string) {
	fmt.Printf("Possible commands are:\n")
	for _, cmdName := range commandNames {
		fmt.Printf("%v\n", cmdName)
	}
	fmt.Printf("\n")
}