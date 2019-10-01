package database

import (
	"os"

	"github.com/sirupsen/logrus"
)

const FILE_NAME string = "remonpi.json"

func NewDatabase(path string) {

	// Check is exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logrus.WithError(err).Fatal()
		return
	}

	// Check file is exists
	if _, err := os.Stat(path + "/" + FILE_NAME); os.IsNotExist(err) {

	}

	// Create files
}

func (j *JSONDB) createRemote() {

}
