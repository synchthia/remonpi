package remote

import (
	"github.com/synchthia/remonpi/controller"
	"github.com/synchthia/remonpi/controller/corona/crw"
	"github.com/synchthia/remonpi/controller/mitsubishi/kgsa3c"
	"github.com/synchthia/remonpi/models"
)

// Remote - Defined Remote Data
type Remote struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`

	Controller controller.Controller
	Database   controller.Database
}

// NewRemote - Initialize Remote
func NewRemote(vendor string, model string, dbPath string) *Remote {
	r := &Remote{
		Vendor: vendor,
		Model:  model,
	}

	if vendor == "mitsubishi" && model == "kgsa3-c" {
		r.Database = kgsa3c.NewDatabase(vendor, model, dbPath)
		r.Controller = kgsa3c.EnsureController(r.Database)
	} else if vendor == "corona" && model == "cr-w" {
		r.Database = crw.NewDatabase(vendor, model, dbPath)
		r.Controller = crw.EnsureController(r.Database)
	}
	return r
}

// GetState - Get Current State
func (r *Remote) GetState() *models.State {
	d := r.Database.GetState()
	return d
}

// Send - Send IR Signal
func (r *Remote) Send(d *models.RemoteData) error {
	if err := r.Controller.Set(d); err != nil {
		return err
	}

	return nil
}
