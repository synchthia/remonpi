package kgsa3c

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/synchthia/remonpi/models"
)

// Database - Controller Database
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

// Load - load from database
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

// Save - Save to Database
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

// GetState - Get current State
func (d *Database) GetState() *models.State {
	return d.State
}

// UpdateState - Update state from RemoteData
func (d *Database) UpdateState(r *models.RemoteData) error {
	state := d.State
	state.Operation = r.Operation
	state.Mode = r.Mode
	if r.Temp != 0 {
		state.ModeData[state.Mode].Temp = r.Temp
	}
	state.ModeData[state.Mode].Fan = r.Fan
	state.ModeData[state.Mode].HorizontalVane = r.HorizontalVane

	// Disabled: this option only use for send. KEEP DEFAULT VALUE.
	//state.ModeData[state.Mode].VerticalVane = r.VerticalVane
	return nil
}

func generateState() *models.State {
	template := TemplateData
	s := &models.State{
		Operation: false,
		Mode:      "cool",
		ModeData:  make(map[string]*models.ModeData),
	}
	for mode, data := range template {
		modeData := &models.ModeData{}
		if data.Temp != nil {
			modeData.Temp = data.Temp.Default.(float32)
		}
		if data.Fan != nil {
			modeData.Fan = data.Fan.Default.(string)
		}
		if data.HorizontalVane != nil {
			modeData.HorizontalVane = data.HorizontalVane.Default.(string)
		}
		if data.VerticalVane != nil {
			modeData.VerticalVane = data.VerticalVane.Default.(string)
		}

		s.ModeData[mode] = modeData
	}

	return s
}
