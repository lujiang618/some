package crontab

import (
	"encoding/json"
	"os"
	"some/api/application/objects/request"
	"some/api/pkg/utils"
)

func readData(path string) *request.CostParamCreate {
	if !utils.PathExists(path) {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	data := &request.CostParamCreate{}
	decoder.Decode(data)

	return data
}
