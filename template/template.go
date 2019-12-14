package template

// Value - Value of Template
type Value struct {
	Type string `json:"type"`

	// Default - default values
	Default interface{} `json:"default,omitempty"`

	// NoSave - No save to database
	NoSave bool `json:"nosave,omitempty"`

	// Range - numeric range
	Range *Range `json:"range,omitempty"`

	// Toggle - toggle values (on/off)
	Toggle []string `json:"toggle,omitempty"`

	// Shot - single value
	Shot *Shot `json:"shot,omitempty"`

	// Step - multiple values (auto, low, mid, high)
	Step []string `json:"step,omitempty"`
}

// Range - numeric range
type Range struct {
	Step float32 `json:"step"`
	From float32 `json:"from"`
	To   float32 `json:"to"`
}

// Shot - single value
type Shot struct {
	Value    interface{} `json:"value"`
	SendOnly bool        `json:"sendonly,omitempty"`
}
