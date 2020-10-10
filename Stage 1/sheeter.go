package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getSheet() (string, error) {
	// Download remote sheet, notice that no api usage is made here
	// This is due to the possability that if your key is outdated or expired
	// The program will stop working so instead we get the spread sheet in a raw json
	// Format, parse it, and either way continue to run commands
	// NOTE: YOUR SHEEET MUST BE PUBLISHED TO THE INTERNET!!!
	requestURL := fmt.Sprintf(spreadsheetJsonTmplate, spreadsheetID)
	resp, err := http.Get(requestURL)
	if err != nil {
		return "Unable to retrieve data from sheet: %v", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	if err != nil {

		return "Error getting sheet body: %v", err
	}
	return bodyString, nil
}

func checkSeed(seed []string) bool {
	// Check if seed has changed and if its in the correct row
	if seed[1] == "1" && seed[2] == "1" && seed[3] != flagSID {
		flagSID = seed[3]
		return true
	}

	return false
}
func checkToken(Token []string) {
	// Check if correct Row
	if Token[1] == "1" && Token[2] == "2" {
		if Token[3] != "none" && strings.Contains(Token[3], "access_token") && Token[3] != googleToken {
			// Checking if token is not empty, that it contains the access_token just for sanity, and that it is different
			// From what we already have
			fixedToken := strings.Replace(Token[3], "\\\"", "\"", -1)
			googleToken = fixedToken
		}

	}
}

func checkCommand(commandType []string, commandLine []string, index string) []string {

	// Parsing commands and their method
	if commandType[1] != index || commandType[2] != "1" {
		return nil
	}
	if commandLine[1] != index || commandLine[2] != "2" {
		return nil
	}
	return []string{commandType[3], commandLine[3]}
}

func getCommandsFromSheet() [][]string {
	bodyString, err := getSheet()
	if err != nil {
		log.Println(bodyString, err)
		return nil
	}

	// Check to see if format is somehow wrong, try to stop as many crashes as possible....
	matches := regexCompiled.FindAllStringSubmatch(bodyString, -1)
	if matches == nil {
		return nil
	}
	lastMatch := matches[len(matches)-1]
	highestRow, err := strconv.Atoi(lastMatch[1])
	if err != nil {
		log.Println("Error Parsing max row to int: ", err)
		return nil
	}

	if highestRow < 2 {
		log.Println("Not enough rows: ", err)
		return nil
	}
	if len(matches)%2 != 0 {
		log.Println("Wrong Format: ", err)
		return nil
	}
	if len(matches)/2 != highestRow {
		log.Println("Somebody is trying to buffer me....:", err)
		return nil
	}

	// Checking if seed changed and only then running.... no need to run commands twice
	seed := matches[0]
	if !checkSeed(seed) {
		return nil
	}

	// On the way check if you recived a new token from the spredsheet
	checkToken(matches[1])

	// Parsing command to send back a nice array other functions can work with
	counter := 2
	var commandArray [][]string
	for i := 2; i < len(matches); i += 2 {
		index := strconv.Itoa(counter)
		command := checkCommand(matches[i], matches[i+1], index)
		if command == nil {
			log.Println("Command array is broken, check sheet ", err)
			break
		}
		commandArray = append(commandArray, command)
		counter++
	}
	return commandArray
}

func writeOutToDisk(output []string) {

	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	for _, value := range output {
		_, err = fmt.Fprintln(f, value)
		if err != nil {
			log.Println("Failed to write to file: ", err)
			f.Close()
			break
		}
	}
	f.Close()
}
