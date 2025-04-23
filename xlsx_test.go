package xlsx_bench

import (
	"bytes"
	"os"
	"strconv"
	"testing"
	"time"

	sax "github.com/anfilat/xlsx-sax"
	tealeg "github.com/tealeg/xlsx/v3"
	"github.com/xuri/excelize/v2"
)

type xlsxItem struct {
	Name  string
	Offer string
	Name2 string
	Count int
	Date  time.Time
}

func BenchmarkXlsxSAX(b *testing.B) {
	data, _ := os.ReadFile("testdata/bench.xlsx")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		br := bytes.NewReader(data)
		xlsx, _ := sax.New(br, br.Size())
		sheet, _ := xlsx.OpenSheetByOrder(0)

		_ = sheet.SkipRow()

		for sheet.NextRow() {
			item := xlsxItem{}
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

func BenchmarkTealegXlsx(b *testing.B) {
	data, _ := os.ReadFile("testdata/bench.xlsx")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		br := bytes.NewReader(data)
		xlsx, _ := tealeg.OpenReaderAt(br, br.Size())
		sheet := xlsx.Sheets[0]

		_ = sheet.ForEachRow(func(r *tealeg.Row) error {
			if r.GetCoordinate() == 0 {
				return nil
			}

			item := xlsxItem{}
			return r.ForEachCell(func(c *tealeg.Cell) error {
				col, _ := c.GetCoordinates()
				if col == 0 {
					item.Name = c.Value
				} else if col == 1 {
					item.Offer, _ = c.FormattedValue()
				} else if col == 2 {
					item.Name2 = c.Value
				} else if col == 3 {
					item.Count, _ = c.Int()
				} else if col == 4 {
					item.Date, _ = c.GetTime(false)
				}
				return nil
			})
		})
	}
}

func BenchmarkExcelize(b *testing.B) {
	data, _ := os.ReadFile("testdata/bench.xlsx")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		br := bytes.NewReader(data)
		xlsx, _ := excelize.OpenReader(br)
		rows, _ := xlsx.GetRows("b")

		for _, row := range rows[1:] {
			item := xlsxItem{}
			for col, colCell := range row {
				if col == 0 {
					item.Name = colCell
				} else if col == 1 {
					item.Offer = colCell
				} else if col == 2 {
					item.Name2 = colCell
				} else if col == 3 {
					item.Count, _ = strconv.Atoi(colCell)
				} else if col == 4 {
					item.Date, _ = time.Parse(time.DateTime, colCell)
				}
			}
		}

		_ = xlsx.Close()
	}
}
