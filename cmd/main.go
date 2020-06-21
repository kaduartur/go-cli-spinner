package main

import (
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
	// Run with another template
	sp := spinner.New("Loading...")
	sp.SetTemplate(template.Arrow)
	sp.Start()
	time.Sleep(time.Second * 2)
	sp.Stop()
}
