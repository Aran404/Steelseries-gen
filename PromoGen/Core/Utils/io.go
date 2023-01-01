package utils

import (
	tof "PromoGen/Core/SteelSeries"
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

var mutex sync.Mutex

func AppendLine(filepath string, s string) {
	mutex.Lock()
	defer mutex.Unlock()
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		LogPanic("%v", err)
	}

	defer file.Close()

	if _, err = file.WriteString(s + "\n"); err != nil {
		LogPanic("%v", err)
	}
}

func Readlines(path string) []string {
	var Lines []string

	readFile, err := os.Open(path)

	if err != nil {
		LogPanic("%v", err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		Lines = append(Lines, fileScanner.Text())
	}

	return Lines
}

func LoadConfig(file string) tof.Config {
	var config tof.Config

	configFile, err := os.Open(file)

	if err != nil {
		LogPanic("%v", err)
	}

	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&config)

	if err != nil {
		LogPanic("%v", err)
	}
	return config
}
