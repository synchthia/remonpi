package models

import "strings"

// RemoteData - Scheme for Remote Controller
type RemoteData struct {
	Operation      bool    `json:"operation"`
	Mode           string  `json:"mode"`
	Temp           float32 `json:"temp,omitempty"`
	Fan            string  `json:"fan"`
	HorizontalVane string  `json:"horizontal_vane,omitempty"`
	VerticalVane   string  `json:"vertical_vane,omitempty"`
}

// State - Scheme for State management (Database etc.)
type State struct {
	Operation bool                 `json:"operation"`
	Mode      string               `json:"mode"`
	ModeData  map[string]*ModeData `json:"mode_data"`
}

// ModeData - RemoteData by individual Modes.
type ModeData struct {
	Temp           float32 `json:"temp,omitempty"`
	Fan            string  `json:"fan"`
	HorizontalVane string  `json:"horizontal_vane,omitempty"`
	VerticalVane   string  `json:"vertical_vane,omitempty"`
}

// ToRemoteData - Convert to RemoteData for current mode.
func (s *State) ToRemoteData() *RemoteData {
	return s.ToRemoteDataByMode(s.Mode)
}

// ToRemoteDataByMode - Convert State to RemoteData for provided mode.
func (s *State) ToRemoteDataByMode(mode string) *RemoteData {
	mode = strings.ToLower(mode)
	return &RemoteData{
		Operation:      s.Operation,
		Mode:           mode,
		Temp:           s.ModeData[mode].Temp,
		Fan:            s.ModeData[mode].Fan,
		HorizontalVane: s.ModeData[mode].HorizontalVane,
		VerticalVane:   s.ModeData[mode].VerticalVane,
	}
}
