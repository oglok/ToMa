package tokenmanager

import (
	"fmt"
	"io/ioutil"
)

const defaultConfig = ".tokenConfig"

// ImportConfigFile is,..
func ImportConfigFile(config string) {

	c, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Println("Loading config threw an error: ", err)
	}
	ioutil.WriteFile(defaultConfig, c, 0666)

}

// ExportConfigFile is,..
func ExportConfigFile(configPath string) {
	dc, err := ioutil.ReadFile(defaultConfig)
	if err != nil {
		fmt.Println("Loading config threw an error: ", err)
	}
	ioutil.WriteFile(configPath, dc, 0666)

}
