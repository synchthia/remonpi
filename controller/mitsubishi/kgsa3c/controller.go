package kgsa3c

import (
	"errors"

	"github.com/synchthia/remonpi/controller"
	"github.com/synchthia/remonpi/models"
	"github.com/synchthia/remonpi/sender"
)

type remoteController struct {
	database controller.Database
}

// EnsureController - Return implemented interface
func EnsureController(database controller.Database) controller.Controller {
	return &remoteController{
		database: database,
	}
}

// Set - Set data and send signal
func (c *remoteController) Set(d *models.RemoteData) error {
	template := TemplateData

	data := &models.RemoteData{}

	// Operation
	data.Operation = d.Operation

	// Mode
	if template[d.Mode] == nil {
		return errors.New("invalid mode provided")
	}
	data.Mode = d.Mode

	// Temp
	if template[d.Mode].Temp != nil {
		if err := template[d.Mode].Temp.Validate(d.Temp); err != nil {
			return err
		}
		data.Temp = d.Temp
	}

	// Fan
	if template[d.Mode].Fan != nil {
		if err := template[d.Mode].Fan.Validate(d.Fan); err != nil {
			return err
		}
		data.Fan = d.Fan
	}

	// HorizontalVane
	if template[d.Mode].HorizontalVane != nil {
		if err := template[d.Mode].HorizontalVane.Validate(d.HorizontalVane); err != nil {
			return err
		}
		data.HorizontalVane = d.HorizontalVane
	}

	// VerticalVane
	if template[d.Mode].VerticalVane != nil {
		if err := template[d.Mode].VerticalVane.Validate(d.VerticalVane); err != nil {
			return err
		}
		data.VerticalVane = d.VerticalVane
	}

	signal, err := c.Generate(d, &models.GenerateOption{})
	if err != nil {
		return err
	}

	if err := sender.Send(signal); err != nil {
		return err
	}

	if err := c.database.UpdateState(data); err != nil {
		return err
	}

	if err := c.database.Save(); err != nil {
		return err
	}

	return nil
}

// Generate - Generate Hex Code
func (c *remoteController) Generate(d *models.RemoteData, opt *models.GenerateOption) ([][]int, error) {
	template := TemplateData
	templateByMode := template[d.Mode]

	signal := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x24, 0x03, 0x0B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x24, 0x03, 0x0B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	}

	// Operation
	if d.Operation {
		signal[0][5] = 0x24
	} else {
		signal[0][5] = 0x20
	}

	// Mode
	if d.Mode == "cool" {
		signal[0][6] = 0x03
	} else if d.Mode == "dry" {
		signal[0][6] = 0x02
	} else if d.Mode == "heat" {
		signal[0][6] = 0x01
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	if templateByMode.Temp != nil {
		if d.Temp >= templateByMode.Temp.Range.From && d.Temp <= templateByMode.Temp.Range.To {
			signal[0][7] = int(templateByMode.Temp.Range.To - d.Temp)
		} else {
			return nil, errors.New("invalid temp provided")
		}
	}
	signal[0][8] = 0x00

	// Fan
	if d.Fan == "auto" {
		signal[0][8] += 0x00
	} else if d.Fan == "low" {
		signal[0][8] += 0x02
	} else if d.Fan == "mid" {
		signal[0][8] += 0x03
	} else if d.Fan == "high" {
		signal[0][8] += 0x05
	} else {
		return nil, errors.New("invalid fan parameters")
	}

	// Horizontal Vane
	if d.HorizontalVane == "auto" {
		signal[0][8] += 0x00
	} else if d.HorizontalVane == "1" {
		signal[0][8] += 0x08
	} else if d.HorizontalVane == "2" {
		signal[0][8] += 0x10
	} else if d.HorizontalVane == "3" {
		signal[0][8] += 0x18
	} else if d.HorizontalVane == "4" {
		signal[0][8] += 0x20
	} else if d.HorizontalVane == "5" {
		signal[0][8] += 0x28
	} else {
		return nil, errors.New("invalid horizontal_vane parameters")
	}

	// Vertical Vane
	if d.VerticalVane == "keep" {
		signal[0][11] = 0x00
	} else if d.VerticalVane == "swing" {
		signal[0][11] = 0x04
	} else {
		return nil, errors.New("invalid vertical_Vane parameters")
	}

	// Sum (Parity)
	sum := 0
	for _, c := range signal[0] {
		sum += c
	}
	signal[0][len(signal[0])-1] = sum

	return signal, nil
}
