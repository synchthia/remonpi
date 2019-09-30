package template

// Type:
// DISABLE 	- disabled (will be hidden)
// RANGE 	- numeric range
// TOGGLE 	- toggle values (on/off)
// STEP 	- step values (low, mid, high)
// BUTTON 	- multiple values (actionA, actionB)

// Value - Value of Template
type Value struct {
	Type    string      `json:"type"`
	Default interface{} `json:"default,omitempty"`
	Range   *Range      `json:"range,omitempty"`
	Toggle  []string    `json:"toggle,omitempty"`
	Step    []string    `json:"step,omitempty"`
	Button  []string    `json:"button,omitempty"`
}

type Range struct {
	Step float32 `json:"step"`
	From float32 `json:"from"`
	To   float32 `json:"to"`
}
