package dal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"frappuccino/config"
	"frappuccino/models"
)

// Internal server errors
func ReadOrders() ([]models.Order, error) {
	jsonFile, err := os.Open(config.Path + "orders.json")
	items := make([]models.Order, 0)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func WriteOrders(items []models.Order) error {
	jsonString, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(config.Path+"orders.json", jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func ContentToOrder(content io.ReadCloser) (models.Order, error) {
	byteValue, err := ioutil.ReadAll(content)
	if err != nil {
		return models.Order{}, err
	}

	var item models.Order
	err = json.Unmarshal(byteValue, &item)
	if err != nil {
		return models.Order{}, err
	}
	if item.CreatedAt != "" || item.Status != "" {
		return models.Order{}, fmt.Errorf("Status and creation time can not be passed.")
	}
	return item, nil
}
