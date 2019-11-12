package hexpi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type IRCode struct {
	Code []int `json:"code"`
}

func SendIR(code []int) error {
	url := os.Getenv("HEXPI_ADDRESS")
	if len(url) == 0 {
		logrus.Errorf("[HEXPI] Failed get HEXPI_ADDRESS. defined?")
		return errors.New("failed get HEXPI_ADDRESS")
	}

	ir := &IRCode{
		Code: code,
	}

	b, err := json.Marshal(ir)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/v1/ir", url),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
