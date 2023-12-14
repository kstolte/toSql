// test
// test
package toSql

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func createInsertLine(recordRow *[]string) string {
	var sanitizedRowValues []string
	for _, v := range *recordRow {
		sanitizedRowValues = append(sanitizedRowValues, strings.ReplaceAll(v, "'", "''"))
	}
	return "INSERT INTO #ttbl SELECT '" + strings.Join(sanitizedRowValues, "','") + "'"
}

type ParseConfig struct {
	Delimiter         rune `default:"\t"`
	FirstLineIsHeader bool `default:"true"` // when set to true the first line of the read value will be used as the column names, when this is false it will default to columnN notation

}

// Take a input and format it into a CSV document
func CsvToSql(input string, parseConfig ParseConfig) (string, error) {
	defer timeTrack(time.Now(), "CsvToSql")

	var ret []string
	// ret = append(ret, "SET NOCOUNT ON")

	r := csv.NewReader(strings.NewReader(input))
	r.Comma = parseConfig.Delimiter
	r.Comment = '#'

	hasHeaderLine := parseConfig.FirstLineIsHeader
	createStatement := "DROP TABLE IF EXISTS #ttbl; GO; CREATE TABLE #ttbl("
	var columnNames []string

	if hasHeaderLine {
		headerRecords, headerReadErr := r.Read()
		if headerReadErr != nil {
			log.Fatal(headerReadErr)
		}
		for _, v := range headerRecords {
			columnNames = append(columnNames, "["+v+"] VARCHAR(512)")
		}
	}
	if !hasHeaderLine {
		firstRow, firstRowReadErr := r.Read()
		if firstRowReadErr != nil {
			log.Fatal(firstRowReadErr)
		}
		ret = append(ret, createInsertLine(&firstRow))

		for i := 0; i < r.FieldsPerRecord; i++ {
			columnNames = append(columnNames, "[column"+strconv.Itoa(i)+"] VARCHAR(512)")
		}
	}
	createStatement += strings.Join(columnNames, ",") + ")"
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		rowInsert := createInsertLine(&record)
		ret = append(ret, rowInsert)
	}
	ret = append(ret, "SET NOCOUNT OFF;")
	var initRows = []string{"SET NOCOUNT ON;", createStatement}

	/// the initRows, ret... seems like a bad idea...
	return strings.Join(append(initRows, ret...), "\n"), nil
}
