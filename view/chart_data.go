package view

import "fmt"

type ChartData struct {
	Labels   []string       `json:"labels"`
	Datasets []ChartDataset `json:"datasets"`
}

type ChartDataset struct {
	BackgroundColor      []string    `json:"backgroundColor"`
	HoverBackgroundColor []string    `json:"hoverBackgroundColor"`
	BorderColor          []string    `json:"borderColor"`
	BorderWidth          int         `json:"borderWidth"`
	BorderSkipped        bool        `json:"borderSkipped"`
	Data                 [][]float32 `json:"data"`
}

type ChartColor struct {
	R int
	G int
	B int
	A float32
}

func (c ChartColor) String() string {
	return fmt.Sprintf("rgba(%d, %d, %d, %f)", c.R, c.G, c.B, c.A)
}

func (c ChartColor) WithAlpha(a float32) ChartColor {
	c.A = a
	return c
}
