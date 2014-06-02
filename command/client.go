package command

import (
	"errors"
	"github.com/pims/spark/spark"
	"io/ioutil"
	"time"
)

const (
	SettingsFileName = ".sparkio"
)

var (
	errNotLoggedIn = errors.New("You should login first.")
	timeout        = time.Duration(10 * time.Second)
)

func AuthenticatedSparkClient(auth bool) (*spark.SparkClient, error) {
	return AuthenticatedSparkClientWithTimeout(auth, timeout)
}

func AuthenticatedSparkClientWithTimeout(auth bool, timeout time.Duration) (*spark.SparkClient, error) {
	c := spark.NewClient(nil, timeout)
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
