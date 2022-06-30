package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func write(filepath, address, title string) error {
	fp := filepath
	if filepath == "~/bm.md" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		fp = home + "/bm.md"
	}

	_, err := os.Stat(fp)
	if err != nil && os.IsNotExist(err) {
		err := createFile(fp)
		if err != nil {
			return fmt.Errorf("Error in creating file: %v", err)
		}
	}
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("Error in opening file: %v", err)
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "\n\n%s\n%s", title, address)
	if err != nil {
		return fmt.Errorf("Error in writing to file: %v", err)
	}
	return nil
}

func createFile(fp string) error {
	createIt, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show(pterm.FgCyan.Sprint(fp, " does not exist. Should I just create it?"))
	if !createIt {
		return errors.New("Ok, I will crash for now, but if you changed your mind, I'm here all day.")
	}
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}
