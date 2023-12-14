package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	// "toSql"
	toSql "github.com/kstolte/toSql_core"
)

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}

func csvToSqlWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 3 {
			return "Invalid number of arguments passed"
		}
		inputJSON := args[0].String()
		if len(inputJSON) < 1 {
			return "empty input field"
		}

		delimiter := []rune(args[1].String())[0]
		hasHeaderLine := args[2].Truthy()

		sqlifiedOutPut, err := toSql.CsvToSql(inputJSON, toSql.ParseConfig{Delimiter: delimiter, FirstLineIsHeader: hasHeaderLine})
		if err != nil {
			fmt.Printf("unable to parse error thrown %s\n", err)
			return err.Error()
		}
		return sqlifiedOutPut
	})
	return jsonFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	js.Global().Set("tsvToSQLImport", csvToSqlWrapper())
	<-make(chan bool)
}
