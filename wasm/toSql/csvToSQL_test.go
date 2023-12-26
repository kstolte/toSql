package toSql

import (
	"log"
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_TabSeparated_HasHeader(t *testing.T) {
	in := `first_name	last_name	username	test
"Rob"	"Pike"	rob	test
# lines beginning with a # character are ignored
Ken	Thompson	ken	test
"Robert"	"Griesemer"	"gri"	test
`
	got, err := CsvToSql(in, ParseConfig{Delimiter: '\t', FirstLineIsHeader: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)

}

func Test_CommaSeparated_HasHeader(t *testing.T) {
	in := `first_name,last_name,username,test
"Rob","Pike",rob,test
# lines beginning with a # character are ignored
Ken,Thompson,ken,test
"Robert","Griesemer","gri",test
`
	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)
}

func Test_TabSeparated_NoHeaderRow(t *testing.T) {
	in := `first_name	last_name	username	test
"Rob"	"Pike"	rob	test
# lines beginning with a # character are ignored
Ken	Thompson	ken	test
"Robert"	"Griesemer"	"gri"	test
`
	got, err := CsvToSql(in, ParseConfig{Delimiter: '\t', FirstLineIsHeader: false})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)

}

func Test_CommaSeparated_NoHeaderRow(t *testing.T) {
	in := `first_name,last_name,username,test
"Rob","Pike",rob,test
# lines beginning with a # character are ignored
Ken,Thompson,ken,test
"Robert","Griesemer","gri",test
`
	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: false})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)
}

func Test_CommaSeparated_AirQuality(t *testing.T) {
	data, testFileReadError := os.ReadFile("testData/Air_Quality.csv")
	if testFileReadError != nil {
		log.Fatal(testFileReadError)
	}
	in := string(data[:])

	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)
}

func Test_CommaSeparated_BankList(t *testing.T) {
	data, testFileReadError := os.ReadFile("testData/banklist.csv")
	if testFileReadError != nil {
		log.Fatal(testFileReadError)
	}
	in := string(data[:])

	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)
}

func Test_CommaSeparated_bobross(t *testing.T) {
	data, testFileReadError := os.ReadFile("testData/bobross-elements-by-episode.csv")
	if testFileReadError != nil {
		log.Fatal(testFileReadError)
	}
	in := string(data[:])

	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)
}

func Test_CommaSeparated_pulitzer(t *testing.T) {
	data, testFileReadError := os.ReadFile("testData/pulitzer-circulation-data.csv")
	if testFileReadError != nil {
		log.Fatal(testFileReadError)
	}
	in := string(data[:])

	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)
}

func Test_CommaSeparated_NationalObesityByState(t *testing.T) {
	data, testFileReadError := os.ReadFile("testData/National_Obesity_By_State.csv")
	if testFileReadError != nil {
		log.Fatal("testing Read File", testFileReadError)
	}
	in := string(data[:])

	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	snaps.MatchSnapshot(t, got)
}

func Test_TabSeparated_HasDoubleQuotesUnescaped(t *testing.T) {
	in := `first_name	last_name	username	test
Rob	Pike	3" pipe	test
`
	got, err := CsvToSql(in, ParseConfig{Delimiter: '\t', FirstLineIsHeader: true, StrictQuotes: true})
	if err != nil {
		t.Errorf("received error %d", err)
	}
	if got == "" {
		t.Error("encounterd blank string return value")
	}
	// snaps.MatchSnapshot(t, got)

}

func Benchmark_TabSeparated_HasHeader(b *testing.B) {
	in := `first_name	last_name	username	test
"Rob"	"Pike"	rob	test
# lines beginning with a # character are ignored
Ken	Thompson	ken	test
"Robert"	"Griesemer"	"gri"	test
`
	got, err := CsvToSql(in, ParseConfig{Delimiter: '\t', FirstLineIsHeader: true})
	if err != nil {
		b.Errorf("received error %d", err)
	}
	if got == "" {
		b.Error("encounterd blank string return value")
	}
	// snaps.MatchSnapshot(t, got)

}

func Benchmark_CommaSeparated_AirQuality(b *testing.B) {
	data, testFileReadError := os.ReadFile("testData/Air_Quality.csv")
	if testFileReadError != nil {
		log.Fatal(testFileReadError)
	}
	in := string(data[:])

	got, err := CsvToSql(in, ParseConfig{Delimiter: ',', FirstLineIsHeader: true})
	if err != nil {
		b.Errorf("received error %d", err)
	}
	if got == "" {
		b.Error("encounterd blank string return value")
	}
}
