package s3_test

import (
	"testing"

	"github.com/jinzhu/configor"
	"github.com/qor/oss/s3"
	"github.com/qor/oss/tests"
)

type Config struct {
	AccessID  string `env:"QOR_AWS_ACCESS_KEY_ID"`
	AccessKey string `env:"QOR_AWS_SECRET_ACCESS_KEY"`
	Region    string `env:"QOR_AWS_REGION"`
	Bucket    string `env:"QOR_AWS_BUCKET"`
}

var client *s3.Client

func init() {
	config := Config{}
	configor.Load(&config)

	client = s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: config.Bucket})
}

func TestAll(t *testing.T) {
	tests.TestAll(client, t)
}

func TestToRelativePath(t *testing.T) {
	urlMap := map[string]string{
		"https://mybucket.s3.amazonaws.com/myobject.ext": "/myobject.ext",
		"https://qor-example.com/myobject.ext":           "/myobject.ext",
		"//mybucket.s3.amazonaws.com/myobject.ext":       "/myobject.ext",
		"http://mybucket.s3.amazonaws.com/myobject.ext":  "/myobject.ext",
		"myobject.ext":                                   "/myobject.ext",
	}

	for url, path := range urlMap {
		if client.ToRelativePath(url) != path {
			t.Errorf("%v's relative path should be %v, but got %v", url, path, client.ToRelativePath(url))
		}
	}
}
