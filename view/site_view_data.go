package view

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"com.iamkevb.fishing/data"
)

type SiteViewData struct {
	model *data.SiteData
}

func (s SiteViewData) Title() string {
	return s.model.WaterData.Title
}

func (s SiteViewData) Latitude() string {
	return fmt.Sprintf("%f", s.model.WaterData.Latitude)
}

func (s SiteViewData) Longitude() string {
	return fmt.Sprintf("%f", s.model.WaterData.Longitude)
}

var (
	PrimaryColor = ChartColor{R: 54, G: 162, B: 235, A: 1}
	SnowColor    = ChartColor{R: 255, G: 128, B: 64, A: 1}
	RainColor    = ChartColor{R: 255, G: 128, B: 255, A: 1}
)

func (s SiteViewData) TemperatureChartData() string {
	labels := []string{}
	bgColors := []string{}
	hoverColors := []string{}
	borderColors := []string{}
	data := [][]float32{}

	for i, v := range s.model.WeatherData.Dates {
		labels = append(labels, v.Format("Mon Jan 2"))
		bgColors = append(bgColors, colorForDate(v, PrimaryColor).String())
		hoverColors = append(hoverColors, PrimaryColor.WithAlpha(0.9).String())
		borderColors = append(borderColors, PrimaryColor.String())
		t := s.model.WeatherData.Temps[i]
		data = append(data, []float32{t[0], t[1]})
	}
	dataset := ChartDataset{
		BackgroundColor:      bgColors,
		HoverBackgroundColor: hoverColors,
		BorderColor:          borderColors,
		BorderWidth:          2,
		BorderSkipped:        false,
		Data:                 data,
	}
	chartData := ChartData{
		Labels:   labels,
		Datasets: []ChartDataset{dataset},
	}

	jsonData, err := json.Marshal(chartData)
	if err != nil {
		fmt.Println("Error serializing temperature chart data", err.Error())
	}
	return string(jsonData)
}

func colorForDate(d time.Time, c ChartColor) ChartColor {
	today := time.Now().UTC().Truncate(24 * time.Hour)
	compareDay := d.Truncate(24 * time.Hour)

	if compareDay.Before(today) {
		return c.WithAlpha(0.7)
	} else if compareDay.Equal(today) {
		return c.WithAlpha(0.3)
	} else {
		return c.WithAlpha(0.0)
	}
}

// / Maybe delete below here??!!
func (s SiteViewData) FlowLabels() string {
	formatted := []string{}
	for _, t := range s.model.WaterData.Times {
		ft := t.Format("Mon Jan 2, 3:04pm")
		formatted = append(formatted, ft)
	}
	return jsStringArray(formatted)
}

func (s SiteViewData) FlowReadings() string {
	vals := []string{}
	for _, v := range s.model.WaterData.Readings {
		vals = append(vals, fmt.Sprintf("%f", v))
	}
	return fmt.Sprintf("[%s]", strings.Join(vals, ","))
}

func (s SiteViewData) WeatherDateLabels() string {
	formatted := []string{}
	for _, d := range s.model.WeatherData.Dates {
		formatted = append(formatted, d.Format("Mon Jan 2"))

	}
	return jsStringArray(formatted)
}

func (s SiteViewData) WeatherTimeLabels() string {
	formatted := []string{}
	for _, t := range s.model.WeatherData.Times {
		formatted = append(formatted, t.Format("Mon Jan 2, 3:04pm"))

	}
	return jsStringArray(formatted)
}

func (s SiteViewData) WeatherTemps() string {
	var dailyTemps = []string{}
	for _, t := range s.model.WeatherData.Temps {
		dt := fmt.Sprintf("[%f,%f]", t[0], t[1])
		dailyTemps = append(dailyTemps, dt)
	}
	return fmt.Sprintf("[%s]", strings.Join(dailyTemps, ","))
}

func (s SiteViewData) Rains() string {
	var daily = []string{}
	for _, v := range s.model.WeatherData.Rain {
		daily = append(daily, fmt.Sprintf("%f", v))
	}
	return fmt.Sprintf("[%s]", strings.Join(daily, ","))
}

func (s SiteViewData) Snows() string {
	var daily = []string{}
	for _, v := range s.model.WeatherData.Snow {
		daily = append(daily, fmt.Sprintf("%f", v))
	}
	return fmt.Sprintf("[%s]", strings.Join(daily, ","))
}

func (s SiteViewData) AirPressures() string {
	var daily = []string{}
	for _, v := range s.model.WeatherData.Pressure {
		daily = append(daily, fmt.Sprintf("%f", v))
	}
	return fmt.Sprintf("[%s]", strings.Join(daily, ","))
}

func (s SiteViewData) ExtraData() data.ExtraData {
	return s.model.ExtraData
}

func jsStringArray(s []string) string {
	quoted := []string{}
	for _, v := range s {
		quoted = append(quoted, fmt.Sprintf("'%s'", v))
	}
	return fmt.Sprintf("[%s]", strings.Join(quoted, ","))
}

func NewSiteViewData(data *data.SiteData) *SiteViewData {
	return &SiteViewData{
		model: data,
	}
}
