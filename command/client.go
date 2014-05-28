package command

import (
	"errors"
	"github.com/pims/spark/spark"
	"io/ioutil"
)

const (
	SettingsFileName = ".sparkio"
)

var (
	errNotLoggedIn = errors.New("You should login first.")
)

func AuthenticatedSparkClient(auth bool) (*spark.SparkClient, error) {
	c := spark.NewClient(nil)
	if auth {
		bytes, err := ioutil.ReadFile(SettingsFileName)
		if err != nil {
			return nil, err
		}

		if len(bytes) == 0 {
			return nil, errNotLoggedIn
		}
		c.AuthToken = string(bytes)
	}
	return c, nil
}
