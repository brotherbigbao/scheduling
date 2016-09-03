package save

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"strconv"
	"time"
	"os/exec"
	"path/filepath"
	"os"
	"runtime"
)

func SaveExcel(data []interface{}) (string, error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	var fileName string

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for k, v := range data {
		names, boolean := v.([]string)
		if(boolean){
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = strconv.Itoa(k+1)
			for _, v2 := range names {
				cell = row.AddCell()
				cell.Value = v2
			}
		}
	}


	runFile, _ := exec.LookPath(os.Args[0])
	runPath, _ := filepath.Abs(runFile)
	fmt.Println(runPath)
	fmt.Println(runtime.GOOS)

	t := time.Now()
	fileName = t.String() + ".xlsx"

	//if runtime.GOOS == "linux" {
	//	fileName = runPath + "/" + fileName
	//} else {
	//	fileName = runPath + "\\" + fileName
	//}

	err = file.Save(fileName)
	if err != nil {
		fmt.Printf(err.Error())
		return "", err
	}

	return fileName, nil
}