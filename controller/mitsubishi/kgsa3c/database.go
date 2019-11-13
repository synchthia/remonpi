package kgsa3c

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/synchthia/remonpi/models"
)

// Database
type Database struct {
	Vendor   string
	Model    string
	FilePath string

	// Data
	State *models.State
}

// NewDatabase - Ensure new Database
func NewDatabase(vendor, model, path string) *Database {
	// Check files
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logrus.WithError(err).Fatal("[DB] Directory not found.")
		return nil
	}

	d := &Database{
		Vendor:   vendor,
		Model:    model,
		FilePath: fmt.Sprintf("%s/%s_%s.json", path, vendor, model),
	}

	// Check file is not exists
	if _, err := os.Stat(d.FilePath); os.IsNotExist(err) {
		logrus.Infof("[DB] File not found, Creating...")

		// Create DB
		d.State = generateState()

		// Save
		if err := d.Save(); err != nil {
			panic(err)
		}
	}

	// Load
	if err := d.Load(); err != nil {
		panic(err)
	}

	// Check State
	// if state has not exists, fill state from template.
	return d
}

func (d *Database) Load() error {
	b, err := ioutil.ReadFile(d.FilePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &d.State); err != nil {
		return err
	}

	return nil
}

func (d *Database) Save() error {
	b, err := json.Marshal(d.State)
	if err != nil {
		return err
	}

	// Save
	if err := ioutil.WriteFile(d.FilePath, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (d *Database) GetState() *models.State {
	return d.State
}

func (d *Database) UpdateState(r *models.RemoteData) error {
	state := d.State
	state.Operation = r.Operation
	state.Mode = r.Mode
	if r.Temp != 0 {
		state.ModeData[state.Mode].Temp = r.Temp
	}
	state.ModeData[state.Mode].Fan = r.Fan
	state.ModeData[state.Mode].HorizontalVane = r.HorizontalVane
	state.ModeData[state.Mode].VerticalVane = r.VerticalVane
	return nil
}

func generateState() *models.State {
	template := TemplateData
	s := &models.State{
		Operation: false,
		Mode:      "cool",
		ModeData:  make(map[string]*models.ModeData),
	}
	for _, m := range ModeList {
		//d, err := json.Marshal(v)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Printf("%s, %s\n", k, d)
		modeData := &models.ModeData{}
		v := template.GetByMode(m)
		if v.Temp != nil {
			modeData.Temp = v.Temp.Default.(float32)
		}
		if v.Fan != nil {
			modeData.Fan = v.Fan.Default.(string)
		}
		if v.VerticalVane != nil {
			modeData.VerticalVane = v.VerticalVane.Default.(string)
		}
		if v.HorizontalVane != nil {
			modeData.HorizontalVane = v.HorizontalVane.Default.(string)
		}

		s.ModeData[m] = modeData
	}

	return s
}

func (d *Database) Test() {
	fmt.Println("test")
}
