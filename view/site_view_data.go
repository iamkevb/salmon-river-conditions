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

func (s SiteViewData) TemperatureChartData() string {
	labels := []string{}
	bgColors := []string{}
	hoverColors := []string{}
	borderColors := []string{}
	data := []any{}

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

func (s SiteViewData) PrecipitationChartData() string {
	labels := []string{}

	rainBgColors := []string{}
	rainHoverColors := []string{}
	rainBorderColors := []string{}
	rainData := []any{}

	snowBgColors := []string{}
	snowHoverColors := []string{}
	snowBorderColors := []string{}
	snowData := []any{}

	for i, v := range s.model.WeatherData.Dates {
		labels = append(labels, v.Format("Mon Jan 2"))
		rainBgColors = append(rainBgColors, colorForDate(v, RainColor).String())
		rainHoverColors = append(rainHoverColors, RainColor.String())
		rainBorderColors = append(rainBorderColors, RainColor.String())
		rainData = append(rainData, s.model.WeatherData.Rain[i])

		snowBgColors = append(snowBgColors, colorForDate(v, SnowColor).String())
		snowHoverColors = append(snowHoverColors, SnowColor.String())
		snowBorderColors = append(snowBorderColors, SnowColor.String())
		snowData = append(snowData, s.model.WeatherData.Snow[i])
	}
	rainDataset := ChartDataset{
		BackgroundColor:      rainBgColors,
		HoverBackgroundColor: rainHoverColors,
		BorderColor:          rainBorderColors,
		BorderWidth:          2,
		BorderSkipped:        false,
		Data:                 rainData,
	}
	snowDataset := ChartDataset{
		BackgroundColor:      snowBgColors,
		HoverBackgroundColor: snowHoverColors,
		BorderColor:          snowBorderColors,
		BorderWidth:          2,
		BorderSkipped:        false,
		Data:                 snowData,
	}
	chartData := ChartData{
		Labels:   labels,
		Datasets: []ChartDataset{rainDataset, snowDataset},
	}

	jsonData, err := json.Marshal(chartData)
	if err != nil {
		fmt.Println("Error serializing temperature chart data", err.Error())
	}
	return string(jsonData)
}

func (s SiteViewData) AtmosphericPressureChartData() string {
	labels := []string{}
	data := []any{}
	for i, v := range s.model.WeatherData.Pressure {
		time := s.model.WeatherData.Times[i]
		labels = append(labels, time.Format("Mon Jan 2, 3:04pm"))
		data = append(data, v)
	}
	dataset := ChartDataset{
		BorderColor: []string{PrimaryColor.String()},
		Data:        data,
		BorderWidth: 2,
		PointRadius: 0,
		Tension:     0.4,
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

func colorForDate(d time.Time, c ChartColor) ChartColor {
	today := time.Now().UTC().Truncate(24 * time.Hour)
	compareDay := d.Truncate(24 * time.Hour)

	if compareDay.Before(today) {
		return c
	} else if compareDay.Equal(today) {
		return c.WithAlpha(0.3)
	} else {
		return c.WithAlpha(0.0)
	}
}

func (s SiteViewData) ExtraData() data.ExtraData {
	return s.model.ExtraData
}

func NewSiteViewData(data *data.SiteData) *SiteViewData {
	return &SiteViewData{
		model: data,
	}
}
