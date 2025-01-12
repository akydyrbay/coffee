package dal

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"frappuccino/config"
	"frappuccino/models"
)

// Internal server errors
func ReadInventory() ([]models.InventoryItem, error) {
	jsonFile, err := os.Open(config.Path + "inventory.json")
	items := make([]models.InventoryItem, 0)
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

func WriteInventory(items []models.InventoryItem) error {
	jsonString, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(config.Path+"inventory.json", jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func ContentToInventoryItem(content io.ReadCloser) (models.InventoryItem, error) {
	byteValue, err := ioutil.ReadAll(content)
	if err != nil {
		return models.InventoryItem{}, err
	}

	var item models.InventoryItem
	err = json.Unmarshal(byteValue, &item)
	if err != nil {
		return models.InventoryItem{}, err
	}
	return item, nil
}
