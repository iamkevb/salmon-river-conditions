package view

import (
	"encoding/json"
	"fmt"
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
	PrimaryColor = ChartColor{R: 102, G: 204, B: 255, A: 1}
	RainColor    = ChartColor{R: 102, G: 204, B: 255, A: 1}
	SnowColor    = ChartColor{R: 176, G: 196, B: 222, A: 1}
)

func (s SiteViewData) ForecastData() []data.ForecastData {
	return s.model.WeatherData.Forecast
}

func (s SiteViewData) FlowChartData() string {
	labels := []string{}
	data := []any{}
	for i, v := range s.model.WaterData.Times {
		labels = append(labels, v.Format("Mon Jan 2, 3:04pm"))
		data = append(data, s.model.WaterData.Readings[i])
	}
	dataset := ChartDataset{
		BorderColor: []string{PrimaryColor.String()},
		Data:        data,
		BorderWidth: 2,
		PointRadius: 0,
		Tension:     0.5,
	}
	chartData := ChartData{
		Labels:   labels,
		Datasets: []ChartDataset{dataset},
	}
	jsonData, err := json.Marshal(chartData)
	if err != nil {
		fmt.Println("Error serializing atmospheric pressure chart data", err.Error())
	}
	return string(jsonData)
}

func (s SiteViewData) AtmosphericPressureChartData() string {
	labels := []string{}
	now := time.Now()
	beforeNow := []any{}
	afterNow := []any{}
	for i, v := range s.model.WeatherData.Pressure.Pressure {
		time := s.model.WeatherData.Pressure.Times[i]
		labels = append(labels, time.Format("Mon Jan 2, 3:04pm"))
		if time.Before(now) {
			beforeNow = append(beforeNow, v)
			afterNow = append(afterNow, nil)
		} else {
			afterNow = append(afterNow, v)
		}
	}
	afterNow[len(beforeNow)-1] = beforeNow[len(beforeNow)-1]
	dataset1 := ChartDataset{
		BorderColor: []string{PrimaryColor.String()},
		Data:        beforeNow,
		BorderWidth: 2,
		PointRadius: 0,
		Tension:     0.4,
	}
	dataset2 := ChartDataset{
		BorderColor: []string{SnowColor.String()},
		Data:        afterNow,
		BorderWidth: 2,
		PointRadius: 0,
		Tension:     0.4,
	}
	chartData := ChartData{
		Labels:   labels,
		Datasets: []ChartDataset{dataset1, dataset2},
	}
	jsonData, err := json.Marshal(chartData)
	if err != nil {
		fmt.Println("Error serializing atmospheric pressure chart data", err.Error())
	}
	return string(jsonData)
}

func (s SiteViewData) ExtraData() data.ExtraData {
	return s.model.ExtraData
}

func NewSiteViewData(data *data.SiteData) *SiteViewData {
	return &SiteViewData{
		model: data,
	}
}
