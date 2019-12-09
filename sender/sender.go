package sender

import (
	"github.com/synchthia/remonpi/sender/hexpi"
	"github.com/synchthia/remonpi/util"
)

// Send - Generate & Send IR Signal
func Send(signal [][]int) error {
	code := util.SignalToCode(430, signal, 13300)
	if err := hexpi.SendIR(code); err != nil {
		return err
	}

	return nil
}
