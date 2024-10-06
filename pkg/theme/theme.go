package theme

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const ColorNamePanelBackground fyne.ThemeColorName = "fynedeskPanelBackground"

//go:embed assets/argent.svg
var argentSvg []byte

var ResourceBatterySvg = &fyne.StaticResource{
	StaticName:    "argent.svg",
	StaticContent: argentSvg,
}

var (

	// ArgentDarkYellowIcon is the material design icon for logo  in light and dark theme
	ArgentDarkYellowIcon = theme.NewThemedResource(ResourceBatterySvg)

	BorderWidth      = float32(4)
	ButtonWidth      = float32(32)
	NarrowBarWidth   = float32(36)
	TitleHeight      = float32(28)
	WidgetPanelWidth = float32(196)
)
