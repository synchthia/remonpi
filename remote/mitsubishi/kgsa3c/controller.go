package kgsa3c

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/synchthia/remonpi/util"
)

// Controller - Controller Data
type Controller struct {
	Operation      bool    `json:"operation"`
	Mode           string  `json:"mode"`
	Temp           float32 `json:"temp"`
	Fan            string  `json:"fan"`
	VerticalVane   string  `json:"vertical_vane"`
	HorizontalVane string  `json:"horizontal_vane"`
}

// Send - Generate & Send IR Signal
func Send(c *Controller) error {
	signal, err := generate(c)
	if err != nil {
		return err
	}

	// TODO: Send
	code := util.SignalToCode(430, signal, 13300)
	logrus.Println("Emitted")
	fmt.Println(code)

	return nil
}

func generate(c *Controller) ([][]int, error) {
	signal := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x24, 0x03, 0x0B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	}

	// Operation
	if c.Operation {
		signal[0][5] = 0x24
	} else {
		signal[0][5] = 0x20
	}

	// Mode
	if c.Mode == "cool" {
		signal[0][6] = 0x03
	} else if c.Mode == "dry" {
		signal[0][6] = 0x02
	} else if c.Mode == "heat" {
		signal[0][6] = 0x01
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	if c.Temp >= 16.0 && c.Temp <= 31.0 {
		signal[0][7] = int(31.0 - c.Temp)
	} else {
		return nil, errors.New("invalid temp provided")
	}

	signal[0][8] = 0x00

	// Fan
	if c.Fan == "auto" {
		signal[0][8] += 0x00
	} else if c.Fan == "low" {
		signal[0][8] += 0x02
	} else if c.Fan == "mid" {
		signal[0][8] += 0x03
	} else if c.Fan == "high" {
		signal[0][8] += 0x05
	} else {
		return nil, errors.New("invalid fan parameters")
	}

	// Vertical Vane
	if c.VerticalVane == "auto" {
		signal[0][8] += 0x00
	} else if c.VerticalVane == "1" {
		signal[0][8] += 0x08
	} else if c.VerticalVane == "2" {
		signal[0][8] += 0x10
	} else if c.VerticalVane == "3" {
		signal[0][8] += 0x18
	} else if c.VerticalVane == "4" {
		signal[0][8] += 0x20
	} else if c.VerticalVane == "5" {
		signal[0][8] += 0x28
	} else {
		return nil, errors.New("invalid vertical_vane parameters")
	}

	// Horizontal Vane
	if c.HorizontalVane == "keep" {
		signal[0][11] = 0x00
	} else if c.HorizontalVane == "swing" {
		signal[0][11] = 0x04
	} else {
		return nil, errors.New("invalid horizontal_vane parameters")
	}
	return signal, nil
}
