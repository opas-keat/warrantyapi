package common

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func CreateExcel(pathFile string, sheet string, data [][]interface{}) {
	var (
		addr  string
		// sheet = "Sheet1"
		// data = [][]interface{}{
        //     {"ชื่อ-สกุล", "เบอร์โทร", "ทะเบียนรถ", "เลขที่ใบรับประกัน", "ร้านค้าที่ซื้อ", "วันที่ซื้อ", "ประเภทสินค้า", "รายการสินค้า", "จำนวน", "วันที่หมดรับประกันโครงสร้าง", "วันที่หมดรับประกันสี", "วันที่หมดรับประกันคุณภาพ", "หมดรับประกันระยะ", "แคมเปญ", "จำนวน"},
        // }
	)
	// สร้างไฟล์ Excel ใหม่
    f := excelize.NewFile()
	defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()	

    // เพิ่มเวิร์กชีทใหม่
    index, err := f.NewSheet(sheet)
    if err != nil {
        fmt.Println(err)
        return
    }
	// Set active sheet of the workbook.
    f.SetActiveSheet(index)

	// set each cell value
    for r, row := range data {
        if addr, err = excelize.JoinCellName("A", r+1); err != nil {
            fmt.Println(err)
            return
        }
        if err = f.SetSheetRow(sheet, addr, &row); err != nil {
            fmt.Println(err)
            return
        }
    }

	if err := f.SaveAs(pathFile); err != nil {
        fmt.Println(err)
    }

    fmt.Println("สร้างไฟล์ Excel สำเร็จ")
}