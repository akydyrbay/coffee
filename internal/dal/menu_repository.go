package dal

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"frappuccino/config"
	"frappuccino/models"
)

func ReadMenu() ([]models.MenuItem, error) {
	jsonFile, err := os.Open(config.Path + "menu_items.json")
	items := make([]models.MenuItem, 0)
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

func WriteMenu(items []models.MenuItem) error {
	jsonString, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(config.Path+"menu_items.json", jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func ContentToMenu(content io.ReadCloser) (models.MenuItem, error) {
	byteValue, err := ioutil.ReadAll(content)
	if err != nil {
		return models.MenuItem{}, err
	}

	var item models.MenuItem
	err = json.Unmarshal(byteValue, &item)
	if err != nil {
		return models.MenuItem{}, err
	}
	return item, nil
}
