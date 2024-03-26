package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/stroomnetwork/keygen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "keygen"
	app.Usage = "Command line tool to securely generate random keys"
	app.DefaultCommand = "generate"
	app.Commands = []*cli.Command{
		{
			Name: "generate",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:      "output",
					Aliases:   []string{"o"},
					Required:  false,
					TakesFile: true,
					Usage:     "File name to write keys output, for example keys.json",
				},
				&cli.BoolFlag{
					Name:     "force",
					Usage:    "Overwrite output file if it exists",
					Required: false,
					Aliases:  []string{"f"},
				},
			},
			Action: generate,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Error: ", err)
	}
}

var ErrFileAlreadyExists = errors.New("file already exists")

func generate(cCtx *cli.Context) error {
	keys, err := keygen.GenerateRandomKeys()
	if err != nil {
		return fmt.Errorf("can't generate keys: %w", err)
	}
	keysJson, err := json.MarshalIndent(keys, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal JSON: %w", err)
	}

	outputFile := cCtx.String("output")
	if outputFile == "" {
		return printToConsole(keysJson)
	}
	// print to file
	if fileExists(outputFile) && !cCtx.Bool("force") {
		return fmt.Errorf("%w: %s", ErrFileAlreadyExists, outputFile)
	}
	err = os.WriteFile(outputFile, keysJson, 0600)
	if err != nil {
		return fmt.Errorf("can't write to file '%s': %w", outputFile, err)
	}
	return nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func printToConsole(keysJson []byte) error {
	fmt.Println("")
	fmt.Println("!!!DO NOT FORGET TO BACKUP THIS KEY!!!")
	fmt.Println("")
	fmt.Println(string(keysJson))
	fmt.Println("")
	fmt.Println("!!!DO NOT FORGET TO BACKUP THIS KEY!!!")
	fmt.Println("")
	return nil
}
