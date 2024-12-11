package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type ErrItem struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Answer string `json:"answer"`
}

var ErrAlreadyExist = errors.New("Error already exits")

func main() {
	errSave := flag.NewFlagSet("save", flag.ExitOnError)
	errName := errSave.String("name", "", "Giving the error a title")
	errValue := errSave.String("value", "", "Add the content of the error")

	if len(os.Args) < 2 {
		fmt.Println("missing subcommands")
		os.Exit(1)

	}

	switch os.Args[1] {
	case "save":
		var errValueInput string
		var errNameInput string
		scanner := bufio.NewScanner(os.Stdin)

		errSave.Parse(os.Args[2:])
		if errSave.NFlag() > 0 {
			if *errName == "" {
				fmt.Println("missing name flag")
				os.Exit(1)

			}
			if *errValue == "" {
				fmt.Println("Need to provide the error text")
				os.Exit(1)

			}
			err := errCheck(*errValue)
			if errors.Is(err, ErrAlreadyExist) {
				fmt.Println(err)
				os.Exit(1)

			}
			err = Save(*errName, *errValue)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			os.Exit(0)
		}
		for scanner.Scan() {
			errValueInput = scanner.Text()
			fmt.Println(errValueInput)
		}
		err := scanner.Err()

		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
			}
			fmt.Println(err, "hello")
			os.Exit(1)

		}
		err = errCheck(errValueInput)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("What is the name of the error")
		_, err = fmt.Scanln(&errNameInput)
		if err != nil {
			fmt.Println(err, "EOF")
		}
		fmt.Println(errNameInput)

	}

}

func errCheck(eValue string) error {
	var errList []ErrItem
	c, err := os.ReadFile("test.json")

	if err != nil {
		fmt.Println("File Error")

	}
	err = json.Unmarshal(c, &errList)
	if err != nil {
		fmt.Println("Error parsing json", err)

	}

	for _, e := range errList {
		if e.Value == eValue {
			fmt.Println("already exists")
			fmt.Printf("%v", e)
			return ErrAlreadyExist

		}
		break

	}
	return nil

}

func Save(eName, eValue string) error {
	var errItem ErrItem
	var errList []ErrItem
	c, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)

	}
	err = json.Unmarshal(c, &errList)

	if err != nil {
		fmt.Println(err)
	}

	errItem.Name = eName
	errItem.Value = eValue
	errItem.Answer = "TODO"

	errList = append(errList, errItem)

	b, err := json.Marshal(errList)
	if err != nil {
		fmt.Println("can't marshall json file")
		return err

	}
	f, err := os.Create("test.json")
	if err != nil {
		fmt.Println("can't open file", err)
		return err
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		fmt.Println("can't write to json file", err)
		return err
	}
	fmt.Println("Succesfully add new error to list")
	return nil

}
