package data

type ExtraData interface {
	Title() string
	RenderedData() string
}

func RenderExtraData(code string) ExtraData {
	switch code {
	case "04250200":
		return Render_Site_04250200()
	default:
		return nil
	}
}
