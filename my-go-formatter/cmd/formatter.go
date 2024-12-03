package main

import (
	"fmt"
	"my-go-formatter/internal/formatter"
	"my-go-formatter/internal/utils"
	"os"
)

func main() {
	lang := "en"
	if len(os.Args) > 1 && (os.Args[1] == "ru" || os.Args[1] == "en") {
		lang = os.Args[1]
		os.Args = os.Args[2:]
	}

	if len(os.Args) < 1 {
		printUsage(lang)
		return
	}

	command := os.Args[0]
	configFile := ""
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	config, err := formatter.LoadConfig(configFile)
	if err != nil {
		fmt.Println(utils.GetMessage(lang, "error_opening"), err)
		return
	}

	switch command {
	case "edit":
		formatter.EditConfig(configFile, lang)
	default:
		if err := formatter.Format(command, config, lang); err != nil {
			fmt.Println(utils.GetMessage(lang, "error_formatting"), err)
		}
	}
}

func printUsage(lang string) {
	fmt.Println(utils.GetMessage(lang, "usage"))
	fmt.Println(utils.GetMessage(lang, "commands"))
	fmt.Println(utils.GetMessage(lang, "format_file"))
	fmt.Println(utils.GetMessage(lang, "edit_config"))
}
