package main

import (
	"io/ioutil"
	"log"

	"github.com/unidoc/unipdf/v3/common/license"
)

func init() {
	content, err := ioutil.ReadFile("./.api-key")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = license.SetMeteredKey(string(content))
	if err != nil {
		log.Fatal("Error: failed to set metered key: ", err.Error(), "\nMake sure you get a valid key from https://cloud.unidoc.io\n")
	}
}