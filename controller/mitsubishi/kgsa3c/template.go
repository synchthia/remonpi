package kgsa3c

import (
	"github.com/synchthia/remonpi/template"
)

// ModeTemplate - Template individual Modes
type ModeTemplate struct {
	Temp           *template.Value `json:"temp"`
	Fan            *template.Value `json:"fan"`
	HorizontalVane *template.Value `json:"horizontal_vane"`
	VerticalVane   *template.Value `json:"vertical_vane"`
}

var (
	// TemplateData - Entries of Template
	TemplateData = map[string]*ModeTemplate{
		"cool": &ModeTemplate{
			Temp: &template.Value{
				Type:    "range",
				Default: float32(21.0),
				Range: &template.Range{
					Step: 1,
					From: 16.0,
					To:   31.0,
				},
			},
			Fan: &template.Value{
				Type:    "step",
				Default: "auto",
				Step: []string{
					"auto",
					"low",
					"mid",
					"high",
				},
			},
			HorizontalVane: &template.Value{
				Type:    "step",
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
			VerticalVane: &template.Value{
				Type:    "shot",
				Default: "keep",
				Shot: &template.Shot{
					Value:    "toggle",
					SendOnly: true,
				},
			},
		},
		"dry": &ModeTemplate{
			Fan: &template.Value{
				Type:    "step",
				Default: "auto",
				Step: []string{
					"auto",
				},
			},
			HorizontalVane: &template.Value{
				Type:    "step",
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
			VerticalVane: &template.Value{
				Type:    "shot",
				Default: "keep",
				Shot: &template.Shot{
					Value:    "toggle",
					SendOnly: true,
				},
			},
		},
		"heat": &ModeTemplate{
			Temp: &template.Value{
				Type:    "range",
				Default: float32(27.0),
				Range: &template.Range{
					Step: 1,
					From: 16.0,
					To:   31.0,
				},
			},
			Fan: &template.Value{
				Type:    "step",
				Default: "auto",
				Step: []string{
					"auto",
					"low",
					"mid",
					"high",
				},
			},
			HorizontalVane: &template.Value{
				Type:    "step",
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
			VerticalVane: &template.Value{
				Type:    "shot",
				Default: "keep",
				Shot: &template.Shot{
					Value:    "toggle",
					SendOnly: true,
				},
			},
		},
	}
)
