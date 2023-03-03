package response

import "io"

type Colors struct {
	Status `json:"status"`
	Result ColorsResult `json:"result"`
}

type ColorsResult struct {
	Colors struct {
		BackgroundColors      []BackgroundColor `json:"background_colors"`
		ColorPercentThreshold float64           `json:"color_percent_threshold"`
		ColorVariance         int               `json:"color_variance"`
		ForegroundColors      []ForegroundColor `json:"foreground_colors"`
		ImageColors           []ImageColor      `json:"image_colors"`
		ObjectPercentage      float64           `json:"object_percentage"`
	} `json:"colors"`
}

type BackgroundColor struct {
	B                           int     `json:"b"`
	ClosestPaletteColor         string  `json:"closest_palette_color"`
	ClosestPaletteColorHTMLCode string  `json:"closest_palette_color_html_code"`
	ClosestPaletteColorParent   string  `json:"closest_palette_color_parent"`
	ClosestPaletteDistance      float64 `json:"closest_palette_distance"`
	G                           int     `json:"g"`
	HTMLCode                    string  `json:"html_code"`
	Percent                     float64 `json:"percent"`
	R                           int     `json:"r"`
}

type ForegroundColor struct {
	B                           int     `json:"b"`
	ClosestPaletteColor         string  `json:"closest_palette_color"`
	ClosestPaletteColorHTMLCode string  `json:"closest_palette_color_html_code"`
	ClosestPaletteColorParent   string  `json:"closest_palette_color_parent"`
	ClosestPaletteDistance      float64 `json:"closest_palette_distance"`
	G                           int     `json:"g"`
	HTMLCode                    string  `json:"html_code"`
	Percent                     float64 `json:"percent"`
	R                           int     `json:"r"`
}

type ImageColor struct {
	B                           int     `json:"b"`
	ClosestPaletteColor         string  `json:"closest_palette_color"`
	ClosestPaletteColorHTMLCode string  `json:"closest_palette_color_html_code"`
	ClosestPaletteColorParent   string  `json:"closest_palette_color_parent"`
	ClosestPaletteDistance      float64 `json:"closest_palette_distance"`
	G                           int     `json:"g"`
	HTMLCode                    string  `json:"html_code"`
	Percent                     float64 `json:"percent"`
	R                           int     `json:"r"`
}

func (c *Colors) Decode(body io.ReadCloser) error {
	return decode(body, c)
}

func (c *Colors) SetBody(body io.ReadCloser) {}
