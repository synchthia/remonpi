package template

import (
	"fmt"
)

// Validate - Validate values in Type
func (v *Value) Validate(target interface{}) error {
	switch v.Type {
	case "range":
		t := target.(float32)
		if t < v.Range.From || t > v.Range.To {
			return fmt.Errorf("out of range: %v", target)
		}
	case "toggle":
		t := target.(string)
		for _, val := range v.Toggle {
			if val == target {
				return nil
			}
		}
		return fmt.Errorf("invalid toggle provided: %v", t)
	case "step":
		t := target.(string)
		for _, val := range v.Step {
			if val == target {
				return nil
			}
		}
		return fmt.Errorf("invalid step provided: %v", t)
	case "shot":
		t := target.(string)
		if v.Shot.Value == target || v.Default == target {
			return nil
		}
		return fmt.Errorf("invalid button provided: %v", t)

	default:
		return fmt.Errorf("unknown type provided: %v", v.Type)
	}

	return nil
}
