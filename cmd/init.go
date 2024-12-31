package cmd

import (
	"fmt"
)

type Init struct {
	Directory string `help:"Directory for config file"`
}

// TODO: clone all repos in list to config dir

//TODO: Create a sqlite db for this config dir

func (i *Init) Run() error {

	fmt.Println(`Welcome to errCLI a cli to help document the obscure or home grown tools you use during your day to day`)

	if i.Directory == "" {
		fmt.Println("directory flag is empty creating directory in home directory")
	}

	d, err := createConfigDir(i.Directory)
	if err != nil {
		return err
	}
	err = createConfigCode(d)
	if err != nil {
		return err
	}

	return nil
}
