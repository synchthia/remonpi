package kgsa3c

import (
	"github.com/synchthia/remonpi/template"
)

// Template - Remote Template
type Template struct {
	Cool *ModeTemplate `json:"cool"`
	Dry  *ModeTemplate `json:"dry"`
	Heat *ModeTemplate `json:"heat"`
}

// ModeTemplate - Template individual Modes
type ModeTemplate struct {
	Temp           *template.Value `json:"temp"`
	Fan            *template.Value `json:"fan"`
	VerticalVane   *template.Value `json:"vertical_vane"`
	HorizontalVane *template.Value `json:"horizontal_vane"`
}

// GetByMode - Get Value by Mode
func (t *Template) GetByMode(mode string) *ModeTemplate {
	switch mode {
	case "cool":
		return t.Cool
	case "dry":
		return t.Dry
	case "heat":
		return t.Heat
	default:
		return nil
	}
}

var (
	// ModeList - List of Mode
	ModeList = []string{"cool", "dry", "heat"}

	// TemplateData - Entries of Template
	TemplateData = Template{
		Cool: &ModeTemplate{
			Temp: &template.Value{
				Type:    "RANGE",
				Default: float32(21.0),
				Range: &template.Range{
					Step: 1,
					From: 16.0,
					To:   31.0,
				},
			},
			Fan: &template.Value{
				Type:    "STEP",
				Default: "auto",
				Step: []string{
					"auto",
					"low",
					"mid",
					"high",
				},
			},
			VerticalVane: &template.Value{
				Type:    "STEP",
				Default: "auto",
				Step: []string{
					"auto",
					"1",
					"2",
					"3",
					"4",
					"5",
				},
			},
			HorizontalVane: &template.Value{
				Type:    "BUTTON",
				Default: "keep",
				Button: []string{
					"keep",
					"swing",
				},
			},
		},
		Dry: &ModeTemplate{
			Fan: &template.Value{
				Type:    "STEP",
				Default: "auto",
				Step: []string{
					"auto",
				},
			},
			VerticalVane: &template.Value{
				Type:    "STEP",
				Default: "auto",
				Step: []string{
					"auto",
					"1",
					"2",
					"3",
					"4",
					"5",
				},
			},
			HorizontalVane: &template.Value{
				Type:    "BUTTON",
				Default: "keep",
				Button: []string{
					"keep",
					"swing",
				},
			},
		},
		Heat: &ModeTemplate{
			Temp: &template.Value{
				Type:    "RANGE",
				Default: float32(27.0),
				Range: &template.Range{
					Step: 1,
					From: 16.0,
					To:   31.0,
				},
			},
			Fan: &template.Value{
				Type:    "STEP",
				Default: "auto",
				Step: []string{
					"auto",
					"low",
					"mid",
					"high",
				},
			},
			VerticalVane: &template.Value{
				Type:    "STEP",
				Default: "auto",
				Step: []string{
					"auto",
					"1",
					"2",
					"3",
					"4",
					"5",
				},
			},
			HorizontalVane: &template.Value{
				Type:    "BUTTON",
				Default: "keep",
				Button: []string{
					"keep",
					"swing",
				},
			},
		},
	}
)
