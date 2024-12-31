package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/unidoc/unipdf/v3/model"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SourceFilepath string `yaml:"source-filepath"`
	OutputFilename string `yaml:"output-filename"`
	Dist string `yaml:"dist"`
	Splitter []int `yaml:"splitter"`
}
func ReadConfig() *Config {

	config := new(Config)
	yamlFile, err := ioutil.ReadFile("./.config")
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return config
}

func main() {
ReadConfig()

	config := ReadConfig()

	reader, _, err := model.NewPdfReaderFromFile(config.SourceFilepath, nil)

	if err != nil {
		panic("Error: fail to read input file: " + err.Error())
	}


	numPages, err := reader.GetNumPages()
	if err != nil { 
		panic("Error: failed to get number of pages: " + err.Error())
	}

	if len(config.Splitter) == 0 {
		panic("Error: The splitter array is empty. You have no splitting task. Please ensure your config is right.")
	}

	// clean previous dist
	_, err = os.Stat(config.Dist)

	if err == nil {
		os.RemoveAll(config.Dist)
	}

	err = os.MkdirAll(config.Dist, 0777)
	if err != nil {
		log.Fatal("Error in creating directory: ", err.Error())
	}


	seperator := [][]int{}

	trailer := 0
	for _, s := range(append(config.Splitter, numPages)) {
		serial := []int{trailer + 1, s}
		seperator = append(seperator, serial)
		trailer = s
	}


	for i, r := range(seperator) {
		writer := model.NewPdfWriter()
		for j := r[0]; j <= r[1]; j++ {
			fmt.Println(i, j)
			page, err := reader.GetPage(j)
			if err != nil { 
				panic("Error: failed to GetPage: " + err.Error())
			}


			writer.AddPage(page)
		}
		writer.WriteToFile(fmt.Sprintf("%s/%s-%d.pdf", config.Dist, config.OutputFilename, i+1))
	}

}