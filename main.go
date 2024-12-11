package main

import (
	"os"

	"github.com/storskegg/caesarcrack/application/caesarcrack"
)

func main() {
	if err := caesarcrack.Execute(); err != nil {
		os.Exit(1)
	}
}
