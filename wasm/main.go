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
		if len(args) != 4 {
			return "Invalid number of arguments passed"
		}
		inputJSON := args[0].String()
		if len(inputJSON) < 1 {
			return "empty input field"
		}

		delimiter := []rune(args[1].String())[0]
		hasHeaderLine := args[2].Truthy()
		useStrictQuotes := args[3].Truthy()

		sqlifiedOutPut, err := toSql.CsvToSql(inputJSON, toSql.ParseConfig{Delimiter: delimiter, FirstLineIsHeader: hasHeaderLine, StrictQuotes: useStrictQuotes})
		if err != nil {
			fmt.Printf("unable to parse error thrown %s\n", err)
			return err.Error()
		}

		return sqlifiedOutPut
	})
	return jsonFunc
}

func toInListWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		inputText := args[0].String()
		if len(inputText) < 1 {
			return "empty input field"
		}

		listifiedOutput, err := toSql.ToInList(inputText)
		if err != nil {
			fmt.Printf("unable to parse error thrown %s\n", err)
			return err.Error()
		}

		return listifiedOutput
	})
	return jsonFunc
}

func handleFileInputParsing() {
	document := js.Global().Get("document")

	fileInput := document.Call("getElementById", "fileInput")
	fileOutput := document.Call("getElementById", "sqlOutput")

	fileInput.Set("oninput", js.FuncOf(func(v js.Value, x []js.Value) any {
		var extension string
		extension = fileInput.Get("files").Call("item", 0).Get("type").String()

		fileInput.Get("files").Call("item", 0).Call("arrayBuffer").Call("then", js.FuncOf(func(v js.Value, x []js.Value) any {
			data := js.Global().Get("Uint8Array").New(x[0])
			dst := make([]byte, data.Get("length").Int())
			js.CopyBytesToGo(dst, data)
			hasHeader := document.Call("getElementById", "chkHeaders").Get("checked").Bool()

			var delimeter rune
			if extension == "text/csv" {
				delimeter = ','
			}
			if extension == "tab-separated-values" {
				delimeter = '\t'
			}

			sqlifiedOutPut, err := toSql.CsvToSql(string(dst), toSql.ParseConfig{Delimiter: delimeter, FirstLineIsHeader: hasHeader})
			if err != nil {
				fileOutput.Set("value", fmt.Sprintf("unable to parse error thrown %s", err))
				fmt.Printf("unable to parse error thrown %s\n", err)
				return err.Error()
			}

			fileOutput.Set("value", sqlifiedOutPut)
			return nil
		}))

		return nil
	}))
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	js.Global().Set("tsvToSQLImport", csvToSqlWrapper())
	js.Global().Set("toInList", toInListWrapper())

	handleFileInputParsing()

	<-make(chan bool)
}
