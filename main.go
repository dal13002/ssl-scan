package main

import (
	"time"
	"strconv"
	"log"

	"ssl-scan/ssllabs"
)

func main() {
	// ::TODO:: move to external config
	var hostnameToAnalyze = "www.elliottmgmt.com"

	log.Println("SSL scan script started for host: " + hostnameToAnalyze)

	// Get the current api info
	apiInfo, err := ssllabs.GetAPIInfo()
	if err != nil {
		log.Fatalf("Error from 'GetAPIInfo' func: " + err.Error())
	}

	// Validate we are able to make a new request
	// ::TODO:: We are enable sleep and retry cycle here; Or if using queue, retry the message with a delay
	newRequestAllowed, errMess := ssllabs.ValidateLimits(apiInfo)
	if !newRequestAllowed {
		log.Fatalf("Not allowed to make new requests: " + errMess)
	}

	log.Println("Max assessments are: " + strconv.Itoa(apiInfo.MaxAssessments) + ", and current assessments are: " + strconv.Itoa(apiInfo.CurrentAssessments))

	// We will retry for 10 (sleep time in seconds) * 60 (number of times) =  600 seconds (10 mins)
	var retryCount = 0
	var report ssllabs.Report
	var startNew = "on"
	for retryCount < 60 {
		// try to analyze the hostname
		report, err = ssllabs.AnalyzeHostname(hostnameToAnalyze, startNew)
		if err != nil {
			log.Fatalf("Error from 'AnalyzeHostname' func: " + err.Error())
		}

		startNew = "off" // for rest of the calls, we only want the status

		if (report.Status == "READY") || (report.Status == "ERROR") {
			break // we found what we need
		} else {
			log.Println("Polling for result. Retry count: " + strconv.Itoa(retryCount))
			retryCount = retryCount + 1
			time.Sleep(10 * time.Second)
		}
	}

	if report.Status != "READY" {
		log.Fatalf("Error while getting report / retry time limit hit: " + report.StatusMessage)
	}

	// Create the cvs, which we will be emailing
	formatedReport := ssllabs.FormatReport(report)

	// write out the cvs
	err = ssllabs.CsvExport(formatedReport)
	if err != nil {
		log.Fatalf("Error from 'CsvExport' func: " + err.Error())
	}

	log.Println("Successfully created report at: " + ssllabs.CSV_LOCATION)
}