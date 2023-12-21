package toSql

import (
	"strings"
	"time"
)

// Take a \n separated list and put it into a IN Style Format
func ToInList(input string) (string, error) {
	defer timeTrack(time.Now(), "toInList")
	var ret []string
	arr := strings.Split(input, "\n")
	for _, v := range arr {
		if v == "" {
			continue
		}
		starSan := strings.ReplaceAll(v, "'", "''")
		ret = append(ret, "'"+starSan+"'")
	}
	finalOut := strings.Join(ret, ",\n")

	/// the initRows, ret... seems like a bad idea...
	return "IN (" + finalOut + "\n)\n", nil
}
