package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func getFirstAuth() (*oauth2.Config, error) {
	config, err := google.ConfigFromJSON(configFile, googleSpreadSheetPrivURL)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
	}
	return config, err
}

func getClient() *http.Client {
	config, err := getFirstAuth()
	if err != nil {
		log.Println("Failed to make basic auth: ", err)
	}
	tok := tokenFromSheet()
	// Check if token expired
	now := time.Now()
	if now.Before(tok.Expiry) {
		return config.Client(context.Background(), tok)
	}

	newTok, err := config.TokenSource(context.TODO(), tok).Token()
	if err != nil {
		log.Println("Token expired and unable to update: ", err)
		return nil
	}
	rawJsonNewToken, err := json.Marshal(newTok)
	if err != nil {
		log.Println("Unable to marshal token: ", err)
		return nil
	}
	googleToken = string(rawJsonNewToken)
	return config.Client(context.Background(), newTok)
}

func tokenFromSheet() *oauth2.Token {
	r := strings.NewReader(googleToken)
	tok := &oauth2.Token{}
	json.NewDecoder(r).Decode(tok)

	return tok
}

func checkConnection() error {
	client := getClient()
	if client == nil {
		return fmt.Errorf("Unable to retrieve Sheets client")
	}
	srv, err := sheets.New(client)
	if err != nil {
		log.Println("Unable to retrieve Sheets client: ", err)
		return err
	}

	readRange := "Sheet1!A1"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Println("Unable to retrieve data from sheet: ", err, resp)
		return err
	}
	return nil
}

func sendOutputToSheet(output []string) error {
	client := getClient()
	if client == nil {
		return fmt.Errorf("Unable to retrieve Sheets client")
	}
	srv, err := sheets.New(client)
	if err != nil {
		log.Println("Unable to retrieve Sheets client: ", err)
		return err
	}

	writeRange := "Sheet1!C1"

	var vr sheets.ValueRange

	vr.Values = append(vr.Values, []interface{}{"Output"})
	for _, value := range output {
		vr.Values = append(vr.Values, []interface{}{value})
	}

	_, err = srv.Spreadsheets.Values.Update(spreadsheetID, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Println("Unable to retrieve data from sheet ", err)
		return err
	}

	return nil
}
