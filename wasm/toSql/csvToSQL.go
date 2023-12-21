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

func createInsertLine(recordRow *[]string, maxFieldLen *[]int, tblName *string) string {
	var sanitizedRowValues []string
	for i, v := range *recordRow {
		if (*maxFieldLen)[i] < len(v) {
			(*maxFieldLen)[i] = len(v)
		}
		starSan := strings.ReplaceAll(v, "'", "''")
		sanitizedRowValues = append(sanitizedRowValues, starSan)
	}
	return "INSERT INTO #" + *tblName + " SELECT '" + strings.Join(sanitizedRowValues, "','") + "'"
}

type ParseConfig struct {
	Delimiter         rune `default:"\t"`
	FirstLineIsHeader bool `default:"true"` // when set to true the first line of the read value will be used as the column names, when this is false it will default to columnN notation

}

// Take a input and format it into a CSV document
func CsvToSql(input string, parseConfig ParseConfig) (string, error) {
	defer timeTrack(time.Now(), "CsvToSql")
	tableName := "ttbl"
	var ret []string
	// ret = append(ret, "SET NOCOUNT ON")

	r := csv.NewReader(strings.NewReader(input))
	r.Comma = parseConfig.Delimiter
	r.Comment = '#'

	hasHeaderLine := parseConfig.FirstLineIsHeader
	createStatement := "DROP TABLE IF EXISTS #" + tableName + ";\nGO\nCREATE TABLE #ttbl("
	var columnNames []string

	if hasHeaderLine {
		headerRecords, headerReadErr := r.Read()
		if headerReadErr != nil {
			log.Fatal(headerReadErr)
		}
		for _, v := range headerRecords {
			columnNames = append(columnNames, "["+v+"] ")
		}
	}
	var maxFieldLen []int

	if !hasHeaderLine {
		firstRow, firstRowReadErr := r.Read()
		if firstRowReadErr != nil {
			log.Fatal(firstRowReadErr)
		}
		maxFieldLen = make([]int, r.FieldsPerRecord)
		ret = append(ret, createInsertLine(&firstRow, &maxFieldLen, &tableName))

		for i := 0; i < r.FieldsPerRecord; i++ {
			columnNames = append(columnNames, "[column"+strconv.Itoa(i)+"] ")
		}
	}

	fieldLenInitNeeded := len(maxFieldLen) < 1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if fieldLenInitNeeded {
			maxFieldLen = make([]int, r.FieldsPerRecord)
			fieldLenInitNeeded = false
		}

		rowInsert := createInsertLine(&record, &maxFieldLen, &tableName)
		ret = append(ret, rowInsert)
	}

	for i := range columnNames {
		columnNames[i] = columnNames[i] + "VARCHAR(" + strconv.Itoa(maxFieldLen[i]) + ")"
	}

	ret = append(ret, "SET NOCOUNT OFF;")
	createStatement += strings.Join(columnNames, ",") + ")"

	var initRows = []string{"SET NOCOUNT ON;", createStatement}

	/// the initRows, ret... seems like a bad idea...
	return strings.Join(append(initRows, ret...), "\n"), nil
}
