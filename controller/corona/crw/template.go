package crw

import (
	"github.com/synchthia/remonpi/template"
)

// Template - Remote Template
type Template struct {
	Cool *ModeTemplate `json:"cool"`
	Dry  *ModeTemplate `json:"dry"`
	Fan  *ModeTemplate `json:"fan"`
}

// ModeTemplate - Template individual Modes
type ModeTemplate struct {
	Temp *template.Value `json:"temp"`
	Fan  *template.Value `json:"fan"`
}

// GetByMode - Get Value by Mode
func (t *Template) GetByMode(mode string) *ModeTemplate {
	switch mode {
	case "cool":
		return t.Cool
	case "dry":
		return t.Dry
	case "fan":
		return t.Fan
	default:
		return nil
	}
}

var (
	// ModeList - List of Mode
	ModeList = []string{"cool", "dry", "fan"}

	// TemplateData - Entries of Template
	TemplateData = Template{
		Cool: &ModeTemplate{
			Temp: &template.Value{
				Type:    "RANGE",
				Default: float32(21.0),
				Range: &template.Range{
					Step: 1,
					From: 20.0,
					To:   30.0,
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
		},
		Dry: &ModeTemplate{
			Temp: &template.Value{
				Type:    "RANGE",
				Default: float32(27.0),
				Range: &template.Range{
					Step: 1,
					From: 20.0,
					To:   30.0,
				},
			},
			Fan: &template.Value{
				Type:    "STEP",
				Default: "auto",
				Step: []string{
					"auto",
				},
			},
		},
		Fan: &ModeTemplate{
			Temp: &template.Value{
				Type:    "RANGE",
				Default: float32(27.0),
				Range: &template.Range{
					Step: 1,
					From: 20.0,
					To:   30.0,
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
		},
	}
)
