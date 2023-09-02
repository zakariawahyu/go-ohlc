package helpers

import (
	"bufio"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
)

func LoadFiles(directory string) ([]NdjsonData, error) {
	result := []NdjsonData{}

	files, err := os.Open(directory)
	if err != nil {
		return nil, errors.Wrap(err, "helpers.LoadFile.ReadDir")
	}
	defer files.Close()

	fileInfos, _ := files.Readdir(-1)

	for _, file := range fileInfos {
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
