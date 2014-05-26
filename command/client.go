package command

import (
	"github.com/pims/spark/spark"
	"io/ioutil"
)

const (
	SettingsFileName = ".sparkio"
)

func AuthenticatedSparkClient(auth bool) (*spark.SparkClient, error) {
	c := spark.NewClient(nil)
	if auth {
		bytes, err := ioutil.ReadFile(SettingsFileName)
		if err != nil {
			return nil, err
		}
		c.AuthToken = string(bytes)
	}
	return c, nil
}
