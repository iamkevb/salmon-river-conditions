package data

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type OttawaRiverFlowData map[string]string

type KML struct {
	XMLName  xml.Name `xml:"kml"`
	Document Document `xml:"Document"`
}

type Document struct {
	Placemarks []Placemark `xml:"Placemark"`
}

type Placemark struct {
	Name         string       `xml:"name"`
	ExtendedData ExtendedData `xml:"ExtendedData"`
}

type ExtendedData struct {
	Data []DataEntry `xml:"Data"`
}

type DataEntry struct {
	Name  string `xml:"name,attr"` // attribute from <Data name="...">
	Value string `xml:"value"`     // inner <value> element text
}

func fetchWaterOfficeData() (string, error) {
	url := "https://wateroffice.ec.gc.ca/services/current_conditions/xml/inline?stations[]=02KF005&stations[]=02LA004&stations[]=02KF001&lang=en"

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

	return string(body), nil
}

func OttawaWaterData() []byte {
	data, err := fetchWaterOfficeData()
	if err != nil {
		return []byte{}
	}

	var kmlData KML
	err = xml.Unmarshal([]byte(data), &kmlData)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}

	re := regexp.MustCompile(`^([\d.]+)`)

	var flowData OttawaRiverFlowData = map[string]string{}
	doc := kmlData.Document
	for _, place := range doc.Placemarks {
		extendedData := place.ExtendedData
		var n, v string
		switch place.Name {
		case "02KF005":
			n = "Ottawa River (Britannia)"
		case "02LA004":
			n = "Rideau River (Ottawa)"
		case "02KF001":
			n = "Mississippi River (Ferguson Falls)"
		}
		for _, item := range extendedData.Data {
			switch item.Name {
			case "Latest Discharge Value":
				match := re.FindStringSubmatch(item.Value)
				v = match[0]
			}
		}
		flowData[n] = v
	}

	json, err := json.Marshal(flowData)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}

	return json
}
