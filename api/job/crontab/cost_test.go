package crontab

import (
	"testing"
)

func TestJson(t *testing.T) {
	t.Log(readData("./load.json"))
	t.Log(readData("./rental.json"))

}
