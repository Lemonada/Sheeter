package main

import (
	"regexp"
	"time"
)

var configFile = []byte("{\"installed\":{\"client_id\":\"1083258992139-51ntmqqalss55vde1hrf7t1h29hot4fb.apps.googleusercontent.com\",\"project_id\":\"sheeter-1602291888463\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_secret\":\"DN6jeohgy_C_5-aIh2udXelA\",\"redirect_uris\":[\"urn:ietf:wg:oauth:2.0:oob\",\"http://localhost\"]}}")

var googleToken = "{\"access_token\":\"ya29.a0AfH6SMA3gJ0i39RmR-prdosCk-0A5mCqe3wNIUafkiuN-BXfkjiTGtYe-FN5EzjpEt3d5_3-zlSpjruMSI9qV_3PFMRP68mO2wc4LZhLM004kgVlLUWm697jQcpNNONOTpiv82UsGpUy7kpg1_bfdUr_W4rY31VKya8\",\"token_type\":\"Bearer\",\"refresh_token\":\"1//03ZJbZ8XKiA2XCgYIARAAGAMSNwF-L9IrLAOowRmAcPhDrcX9VPFB6B7kePaYbyYQ_MhKwhf_uZkVlCLSFajczTlVukckSQKQkSA\",\"expiry\":\"2020-10-10T05:16:56.4063201+03:00\"}"


var spreadsheetID = "1X1ATfmiV_P_ZYyhFSdAKDklPQSYq-9DCBfZtLR5qmes"
var spreadsheetJsonTmplate = "https://spreadsheets.google.com/feeds/cells/%s/1/public/full?alt=json"
var googleSpreadSheetPrivURL = "https://www.googleapis.com/auth/spreadsheets"

// Group 1 should be the row
// Group 2 should be the col
// Group 3 should be the value
var regexCompiled, _ = regexp.Compile("\\\"gs\\$cell\\\":{\\\"row\\\":\\\"(.*?)\\\",\\\"col\\\":\\\"(.*?)\\\",\\\"inputValue\\\":\\\"(.*?)\\\",\\\".*?\\\"}}")

var flagSID = "x"

var loopSleepSeconds = 10 * time.Second
var commandTimeout = 60 * time.Second

var writeToDisk = true
var outputFile = "c:\\temp\\output.txt"
