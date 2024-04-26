package ssllabs

import (
	"net/http"
	"encoding/json"
	"errors"
	"io/ioutil"
)


var baseApiURL = "https://api.ssllabs.com/api/v2"

// GetAPIInfo will return current API info including currentAssessments and maxAssessments
func GetAPIInfo() (APIInfo, error) {
	client := &http.Client{}

	var url = baseApiURL + "/info"
	var apiInfo APIInfo

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return apiInfo, errors.New("Error doing creating HTTP request: " + err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return apiInfo, errors.New("Error doing HTTP request: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return apiInfo, errors.New("Error while reading body from response: " + err.Error())
	}


	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal([]byte(body), &apiInfo)
		if err != nil {
			return apiInfo, errors.New("Error while Unmarshaling data: " + err.Error())
		}
		return apiInfo, nil
	}

	return apiInfo, errors.New("Error response code is not expected, body is: " + string(body))
}

// AnalyzeHostname will start the scan for a hostname
// only use startNew for the first api call, and then for all other calls (ie when waiting for results), set it to off
func AnalyzeHostname(hostname string, startNew string) (Report, error) {
	client := &http.Client{}

	var url = baseApiURL + "/analyze?host=" + hostname + "&startNew=" + startNew + "&all=on"
	var report Report

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return report, errors.New("Error doing creating HTTP request: " + err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return report, errors.New("Error doing HTTP request: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return report, errors.New("Error while reading body from response: " + err.Error())
	}


	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal([]byte(body), &report)
		if err != nil {
			return report, errors.New("Error while Unmarshaling data: " + err.Error())
		}
		return report, nil
	}

	return report, errors.New("Error response code is not expected, body is: " + string(body))
}