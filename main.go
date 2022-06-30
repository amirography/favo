package main

import (
	"github.com/alexflint/go-arg"
	"github.com/pterm/pterm"
)

var args struct {
	Address string `arg:"positional,required"`
	File    string `arg:"-f,env:FAVO_BOOKMARK" default:"~/bm.md"`
}

func main() {
	arg.MustParse(&args)
	gettingTitle, _ := pterm.DefaultSpinner.Start(pterm.Info.Sprint("Getting the title"))
	a := normalizeAddress(args.Address)
	t := title(a)
	gettingTitle.Success(t)
	err := write(args.File, a, t)
	writingToFile, _ := pterm.DefaultSpinner.Start(pterm.Info.Sprint("Writing to file"))
	if err != nil {
		writingToFile.Fail(err)
		return
	}
	writingToFile.Success("done adding ", t, " to bookmarks")

}
