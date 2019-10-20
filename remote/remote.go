package remote

import (
	"fmt"

	"github.com/synchthia/remonpi/controller"
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
	}
	return r
}

func (r *Remote) GetState() *models.State {
	d := r.Database.GetState()
	fmt.Println("getstate...")
	return d
}

func (r *Remote) Send(d *models.RemoteData) error {
	if err := r.Controller.Set(d); err != nil {
		return err
	}

	//if err := r.Controller.Send(d); err != nil {
	//	return err
	//}

	//if err := r.Database.UpdateState(d); err != nil {
	//	return err
	//}

	//if err := r.Database.Save(); err != nil {
	//	return err
	//}
	//return errors.New("send request was not matched any vendor/model")
	return nil
}
