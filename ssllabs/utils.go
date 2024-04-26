package ssllabs

import (
	"time"
	"strconv"
	"os"
	"errors"
	"encoding/csv"
)

// ValidateLimits will return if we are allowed to make a new request to the api
func ValidateLimits(apiInfo APIInfo) (bool, string) {
	if apiInfo.MaxAssessments <= apiInfo.CurrentAssessments {
		return false, "Reached max assessments. Need to allow assessments to finish before sending more."
	}

	return true, ""
}

// FormatReport will format relevant info from the report into a [][]string which is what the csv library can understand
func FormatReport(report Report) ([][]string) {
	fieldNames := append([]string{"Host", "IP", "Grade", "Has Warnings", "Cert Expired", "Chain Issues", "Forward Secrecy"})
	fieldNames = append(fieldNames, VULNERABLES...)
	fieldNames = append(fieldNames, RC4...)
	fieldNames = append(fieldNames, PROTOCOLS...)
	records := [][]string{fieldNames}

	for _, endpoint := range report.Endpoints {
		details := endpoint.Details
		expired := time.Now().After(time.Unix(0, int64(details.Cert.NotAfter) * int64(time.Millisecond)))

		supportedProtocals := map[int]bool{}
		for _, p := range details.Protocols {
			supportedProtocals[p.Id] = true
		}

		var protocalsToAppend []string
		for _, p := range PROTOCOLS {
			id := PROTOCOL_IDS[p]
			_, ok := supportedProtocals[id]
			if ok {
				protocalsToAppend = append(protocalsToAppend, "Yes")
			} else {
				protocalsToAppend = append(protocalsToAppend, "No")
			}
		}

		record := []string{
			report.Host, 
			endpoint.IpAddress, 
			endpoint.Grade, 
			strconv.FormatBool(endpoint.HasWarnings), 
			strconv.FormatBool(expired), 
			CHAIN_ISSUES[details.Chain.Issues], 
			FORWARD_SECRECY[details.ForwardSecrecy], 
			strconv.FormatBool(details.VulnBeast), 
			strconv.FormatBool(details.DrownVulnerable), 
			strconv.FormatBool(details.Heartbleed), 
			strconv.FormatBool(details.Freak), 
			VUL_TEST_RESULTS[details.OpenSslCcs],
			VUL_TEST_RESULTS[details.OpenSSLLuckyMinus20],
			strconv.FormatBool(details.Poodle),
			VUL_TEST_RESULTS[details.PoodleTls],
			strconv.FormatBool(details.SupportsRc4),
			strconv.FormatBool(details.Rc4WithModern),
			strconv.FormatBool(details.Rc4Only),
		}

		record = append(record, protocalsToAppend...)

		records = append(records, record)
	}

	return records
}

// CsvExport will write a csv file
func CsvExport(data [][]string) error {
    file, err := os.Create(CSV_LOCATION)
    if err != nil {
        return errors.New("Error creating file: " + err.Error())
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range data {
        if err := writer.Write(value); err != nil {
			return errors.New("Error writing to file: " + err.Error())
        }
    }
    return nil
}