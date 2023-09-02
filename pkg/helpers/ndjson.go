package helpers

import (
	"bufio"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func LoadFiles() ([]NdjsonData, error) {
	directory := "../../subsetdata"
	result := []NdjsonData{}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, errors.Wrap(err, "helpers.LoadFile.ReadDir")
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".ndjson" {
			filePath := filepath.Join(directory, file.Name())
			result, err = convertFileToStruct(filePath)
			if err != nil {
				return nil, errors.Wrap(err, "helpers.ConvertFileToStruct")
			}
		}
	}

	return result, nil
}

func convertFileToStruct(filePath string) ([]NdjsonData, error) {
	result := []NdjsonData{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "helpers.osOpen")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		data := NdjsonData{}
		if err = json.Unmarshal([]byte(line), &data); err != nil {
			return nil, errors.Wrap(err, "helpers.jsonUnmarshal")
			continue
		}

		eq, _ := strconv.Atoi(data.ExecutedQuantity)
		if eq > 0 {
			data.Quantity = data.ExecutedQuantity
		}

		ep, _ := strconv.Atoi(data.ExecutionPrice)
		if ep > 0 {
			data.Price = data.ExecutionPrice
		}

		result = append(result, NdjsonData{
			StockCode:   data.StockCode,
			OrderNumber: data.OrderNumber,
			Type:        data.Type,
			Quantity:    data.Quantity,
			Price:       data.Price,
		})
	}

	if err = scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "helpers.scannerErr")
	}

	return result, nil
}
