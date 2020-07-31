package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/kaduartur/go-cli-spinner/pkg/spinner"
	"github.com/kaduartur/go-cli-spinner/pkg/template"
)

func main() {
	// Run with default configs
	s := spinner.New("Loading...")
	s.Start()
	time.Sleep(time.Second * 2)
	s.Stop()

	fmt.Println()
	// Run with another template and success message
	sp := spinner.New("Loading...")
	sp.SetTemplate(template.Arrow)
	sp.Start()
	time.Sleep(time.Second * 2)
	sp.Success("Success")

	fmt.Println()
	// Run with error message
	spi := spinner.New("Loading...")
	spi.Start()
	time.Sleep(time.Second * 2)
	spi.Error(errors.New("error to load"))
}
