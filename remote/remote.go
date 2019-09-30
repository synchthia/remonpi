package remote

// Remote - Defined Remote Data
type Remote struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
}

// NewRemote - Initialize Remote
func NewRemote(vendor string, model string) *Remote {
	return &Remote{
		Vendor: vendor,
		Model:  model,
	}
}
