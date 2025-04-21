package xlsx_bench

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/anfilat/xlsx-sax"
)

func BenchmarkXlsx(b *testing.B) {
	data, _ := os.ReadFile("testdata/bench.xlsx")
	br := bytes.NewReader(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xlsx, _ := xlsx.New(br, br.Size())
		sheet, _ := xlsx.OpenSheetByOrder(0)

		_ = sheet.SkipRow()

		for sheet.NextRow() {
			item := xlsx1Item{}
			for sheet.NextCell() {
				if sheet.Col == 0 {
					item.Name, _ = sheet.CellValue()
				} else if sheet.Col == 1 {
					item.Offer, _ = sheet.CellFormatValue()
				} else if sheet.Col == 2 {
					item.Name2, _ = sheet.CellValue()
				} else if sheet.Col == 3 {
					item.Count, _ = sheet.CellInt()
				} else if sheet.Col == 4 {
					item.Date, _ = sheet.CellTime()
				}
			}
		}

		_ = sheet.Close()
	}
}

type xlsx1Item struct {
	Name  string
	Offer string
	Name2 string
	Count int
	Date  time.Time
}
