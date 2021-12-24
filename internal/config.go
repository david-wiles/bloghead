package internal

import (
	"encoding/json"
	"os"
)

// Blog config corresponds to the configuration file used to
// initialize the bh instance. This file also contains saved
// state for the site, such as blueprint file paths sss
type BlogConfig struct {
	// filesystem data
	Root   string `json:"root"`
	Output string `json:"output"`

	// author and site data
	Author   string `json:"Author"`
	Email    string `json:"Email"`
	Domain   string `json:"Domain"`
	Title    string `json:"Title"`
	SubTitle string `json:"SubTitle"`

	// meta data used within bloghead
	Blueprints map[string]string `json:"blueprints"`
	Articles   []string          `json:"articles"`
}

func ReadConfig(filename string) (*BlogConfig, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bc := &BlogConfig{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(bc); err != nil {
		return nil, err
	}

	return bc, nil
}

func SaveConfig(config *BlogConfig, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
