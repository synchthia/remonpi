package crw

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/synchthia/remonpi/controller"
	"github.com/synchthia/remonpi/models"
	"github.com/synchthia/remonpi/sender/hexpi"
	"github.com/synchthia/remonpi/util"
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
	if template.GetByMode(d.Mode) == nil {
		return errors.New("invalid mode provided")
	}
	data.Mode = d.Mode

	// Temp
	if template.GetByMode(d.Mode).Temp != nil {
		if err := template.GetByMode(d.Mode).Temp.Validate(d.Temp); err != nil {
			return err
		}
		data.Temp = d.Temp
	}

	// Fan
	if template.GetByMode(d.Mode).Fan != nil {
		if err := template.GetByMode(d.Mode).Fan.Validate(d.Fan); err != nil {
			return err
		}
		data.Fan = d.Fan
	}

	signal, err := c.Generate(d, &models.GenerateOption{
		MutateOperation: false,
	})
	if err != nil {
		return err
	}

	if err := c.Send(signal); err != nil {
		return err
	}

	signal, err = c.Generate(d, &models.GenerateOption{
		MutateOperation: true,
	})
	if err != nil {
		return err
	}

	if err := c.Send(signal); err != nil {
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

// Send - Generate & Send IR Signal
func (c *remoteController) Send(signal [][]int) error {
	for _, s := range signal {
		for _, c := range s {
			fmt.Printf("%x ", c)
		}
		fmt.Println("")
	}

	code := util.SignalToCode(430, signal, 13300)
	logrus.Debugln("Emitted")
	if err := hexpi.SendIR(code); err != nil {
		return err
	}

	return nil
}

// Generate - Generate Hex Code
func (c *remoteController) Generate(d *models.RemoteData, opt *models.GenerateOption) ([][]int, error) {
	template := TemplateData
	templateByMode := template.GetByMode(d.Mode)

	baseSignal := [][]int{
		{0x28, 0x61, 0x3D, 0x10, 0xEF, 0x94, 0x6B},
		{0x28, 0x61, 0x6D, 0xFF, 0x00, 0xFF, 0x00},
		{0x28, 0x61, 0xCD, 0xFF, 0x00, 0xFF, 0x00},
	}

	signal := baseSignal

	// Operation
	if d.Operation {
		if opt.MutateOperation {
			signal[0][5] = 0x20 // Only Action
		}
	} else {
		signal[0][5] = 0x10
	}

	// Mode
	if d.Mode == "cool" {
		signal[0][5] += 0x90
	} else if d.Mode == "dry" {
		signal[0][5] += 0x02
	} else if d.Mode == "fan" {
		signal[0][5] += 0x01
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	if templateByMode.Temp != nil {
		if d.Temp >= templateByMode.Temp.Range.From && d.Temp <= templateByMode.Temp.Range.To {
			signal[0][5] += int(d.Temp - templateByMode.Temp.Range.From)
		} else {
			return nil, errors.New("invalid temp provided")
		}
	}

	// Fan
	if d.Fan == "auto" {
		signal[0][3] = 0x10
	} else if d.Fan == "low" {
		signal[0][3] = 0x11
	} else if d.Fan == "mid" {
		signal[0][3] = 0x12
	} else if d.Fan == "high" {
		signal[0][3] = 0x13
	} else {
		return nil, errors.New("invalid fan parameters")
	}

	// Sum (Parity)
	sum := 0
	for _, c := range signal[0] {
		sum += c
	}

	baseSignalSum := 0
	for _, c := range baseSignal[0] {
		baseSignalSum += c
	}

	signal[0][len(signal[0])-1] = sum + (sum - baseSignalSum)

	return signal, nil
}
