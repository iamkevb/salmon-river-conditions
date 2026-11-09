package data

type Site_04250200 struct {
	title        string
	renderedData string
}

func (s Site_04250200) Title() string        { return s.title }
func (s Site_04250200) RenderedData() string { return s.renderedData }

func Render_Site_04250200() *Site_04250200 {
	return &Site_04250200{
		title:        "Lighthouse Hill Reservoir",
		renderedData: "<a href=\"https://www.safewaters.com/facility/lighthouse-hill/\">Lighthouse Hill Reservoir</a>",
	}
}
