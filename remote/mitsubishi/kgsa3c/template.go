package kgsa3c

import (
	"errors"

	"github.com/synchthia/remonpi/template"
)

type Template struct {
	Cool *ModeTemplate `json:"cool"`
	Dry  *ModeTemplate `json:"dry"`
	Heat *ModeTemplate `json:"heat"`
}

type ModeTemplate struct {
	Temp           *template.Value `json:"temp"`
	Fan            *template.Value `json:"fan"`
	VerticalVane   *template.Value `json:"vertical_vane"`
	HorizontalVane *template.Value `json:"horizontal_vane"`
}

func (t *Template) GetTemplateByMode(mode string) (*ModeTemplate, error) {
	switch mode {
	case "cool":
		return t.Cool, nil
	case "dry":
		return t.Dry, nil
	case "heat":
		return t.Heat, nil
	default:
		return nil, errors.New("invalid mode provided")
	}
}

var TemplateData = Template{
	Cool: &ModeTemplate{
		Temp: &template.Value{
			Type:    "RANGE",
			Default: 21.0,
			Range: &template.Range{
				Step: 1,
				From: 16.0,
				To:   31.0,
			},
		},
		Fan: &template.Value{
			Type:    "STEP",
			Default: "AUTO",
			Step: []string{
				"AUTO",
				"LOW",
				"MID",
				"HIGH",
			},
		},
		VerticalVane: &template.Value{
			Type:    "STEP",
			Default: "AUTO",
			Step: []string{
				"AUTO",
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		HorizontalVane: &template.Value{
			Type: "BUTTON",
			Button: []string{
				"KEEP",
				"SWING",
			},
		},
	},
	Dry: &ModeTemplate{
		Temp: &template.Value{
			Type: "DISABLED",
		},
		Fan: &template.Value{
			Type:    "STEP",
			Default: "AUTO",
			Step: []string{
				"AUTO",
			},
		},
		VerticalVane: &template.Value{
			Type:    "STEP",
			Default: "AUTO",
			Step: []string{
				"AUTO",
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		HorizontalVane: &template.Value{
			Type: "BUTTON",
			Button: []string{
				"KEEP",
				"SWING",
			},
		},
	},
	Heat: &ModeTemplate{
		Temp: &template.Value{
			Type:    "RANGE",
			Default: 27.0,
			Range: &template.Range{
				Step: 1,
				From: 16.0,
				To:   31.0,
			},
		},
		Fan: &template.Value{
			Type:    "STEP",
			Default: "AUTO",
			Step: []string{
				"AUTO",
				"LOW",
				"MID",
				"HIGH",
			},
		},
		VerticalVane: &template.Value{
			Type:    "STEP",
			Default: "AUTO",
			Step: []string{
				"AUTO",
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		HorizontalVane: &template.Value{
			Type: "BUTTON",
			Button: []string{
				"KEEP",
				"SWING",
			},
		},
	},
}

func GetTemplate() *Template {
	return &TemplateData
}
