package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type OttawaRiverFlowData map[string]string

func formatOttawaDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func fetchWaterOfficeData() (string, error) {
	d := formatOttawaDate()

	url := fmt.Sprintf("https://wateroffice.ec.gc.ca/services/real_time_data/csv/inline?stations[]=02KF005&stations[]=02LA004&stations[]=02KF001&parameters[]=6&start_date=%s%%2000:00:00&end_date=%s%%2023:59:59", d, d)
	fmt.Println("URL: ", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("BODY: ", string(body))
	return string(body), nil
}

func OttawaWaterData() []byte {
	data, err := fetchWaterOfficeData()
	if err != nil {
		return []byte{}
	}
	var flowData OttawaRiverFlowData = map[string]string{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		switch tokens[0] {
		case "02KF005":
			flowData["Ottawa River (Britannia)"] = tokens[3]
		case "02LA004":
			flowData["Rideau River (Ottawa)"] = tokens[3]
		case "02KF001":
			flowData["Mississippi River (Ferguson Falls)"] = tokens[3]
		}
	}
	json, err := json.Marshal(flowData)
	if err != nil {
		return []byte{}
	}
	return json
}
