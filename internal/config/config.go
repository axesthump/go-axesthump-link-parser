package config

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

type AppData struct {
	Doc *html.Node
}

func NewAppData() (*AppData, error) {
	htmlFileName := flag.String("h", "", "html page file path")

	flag.Parse()

	if len(*htmlFileName) == 0 {
		return nil, fmt.Errorf("need html file")
	}

	file, err := os.Open(*htmlFileName)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	return &AppData{
		Doc: doc,
	}, nil
}
