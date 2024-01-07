package view

import (
	"fmt"
	"strings"

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
		fmt.Println(d)
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
