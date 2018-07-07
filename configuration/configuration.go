package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

var (
	config configLogsNom
	once   sync.Once
	err    error
	b      []byte
)

type configLogs struct {
	LogFolder string `json:"log_folder"`
}

type configLogsNom struct {
	nombre    string
	configLog configLogs
}

func GetConfig() configLogsNom {
	config.nombre = ""
	once.Do(loadConfig)
	return config
}

func GetConfigNomb(nombre string) configLogsNom {
	config.nombre = nombre
	once.Do(loadConfig)
	return config
}
func loadConfig() {

	if config.nombre == "" {
		b, err = ioutil.ReadFile("./configLogs.json")
	} else {
		nomSplit := strings.Split(config.nombre, ".")
		println("ojo:" + nomSplit[len(nomSplit)-1])
		if nomSplit[len(nomSplit)-1] == "json" {
			b, err = ioutil.ReadFile(config.nombre)
		} else {
			config.nombre += ".json"
			b, err = ioutil.ReadFile(config.nombre)
		}
	}
	if err != nil {
		fmt.Printf("Error al leer archivo %s: %s", config.nombre, err.Error())
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		fmt.Printf("Error al parsear el archivo %s: %s", config.nombre, err.Error())
	}
}
