package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

type Site_04250200 struct {
	title        string
	renderedData string
}

func (s Site_04250200) Title() string        { return s.title }
func (s Site_04250200) RenderedData() string { return s.renderedData }

type reservoirResponse struct {
	Status int32
	Data   []reservoirData
}
type reservoirData struct {
	Start string
	End   string
	Flow  string
}

type templateData struct {
	Start string
	End   string
	Flow  string
}

func Render_Site_04250200() *Site_04250200 {
	b := loadReservoirData()
	if len(b) == 0 {
		return nil
	}

	templateDatas := parseResponse(b)
	if templateDatas == nil {
		return nil
	}

	tmpl, err := template.ParseFiles("templates/04250200.tmpl.html")
	if err != nil {
		fmt.Println("Error parsing 04250200 template:", err.Error())
		return nil
	}
	w := bytes.NewBufferString("")
	err = tmpl.Execute(w, templateDatas)
	if err != nil {
		fmt.Println("Template error, (042502000)", err.Error())
		return nil
	}
	return &Site_04250200{
		title:        "Lighthouse Hill Reservoir - Water Release Schedule",
		renderedData: w.String(),
	}
}

func parseResponse(b []byte) *[]templateData {
	var data reservoirResponse
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("Error parsing extra response (042502000):", err, string(b))
		return nil
	}
	renderedDatas := []templateData{}

	for _, d := range data.Data {
		layout := "01/02/2006 03:04 PM"

		s, err := time.Parse(layout, d.Start)
		if err != nil {
			fmt.Println("Error parsing date (lighthouse):", err)
		}
		e, err := time.Parse(layout, d.End)
		if err != nil {
			fmt.Println("Error parsing end date (lighthouse):", err)
		}
		renderedDatas = append(renderedDatas, templateData{
			Start: s.Format("Jan 2 3:04 pm"),
			End:   e.Format("Jan 2 3:04 pm"),
			Flow:  fmt.Sprintf("%s cfs", d.Flow),
		})
	}
	return &renderedDatas
}

func loadReservoirData() []byte {
	url := "https://api.safewaters.com/api/schedule/3524dbf0-00c8-11ec-9351-dd66b05aaa5c"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request to Site_04250200:", err)
		return []byte{}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading Site_04250200 body:", err)
		return []byte{}
	}
	return body
}
