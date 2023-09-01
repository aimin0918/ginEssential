package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	defaultSheetName = "Sheet1" //默认Sheet名称
	defaultHeight    = 25.0     //默认行高度
)

type lzExcelExport struct {
	file      *excelize.File
	sheetName string //可定义默认sheet名称
}

func NewMyExcel() *lzExcelExport {
	return &lzExcelExport{file: createFile(), sheetName: defaultSheetName}
}

// ExportToPath  导出基本的表格
func (l *lzExcelExport) ExportToPath(params []map[string]string, data []map[string]interface{}, path string) (string, error) {
	l.export(params, data)
	name := createFileName()
	filePath := path + "/" + name
	err := l.file.SaveAs(filePath)
	return filePath, err
}

// ExportToWeb 导出到浏览器。此处使用的gin框架 其他框架可自行修改ctx
func (l *lzExcelExport) ExportToWeb(params []map[string]string, data []map[string]interface{}, c *gin.Context) {
	l.export(params, data)
	buffer, _ := l.file.WriteToBuffer()
	//设置文件类型
	c.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	//设置文件名称
	c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(createFileName()))
	_, _ = c.Writer.Write(buffer.Bytes())
}

// 设置首行
func (l *lzExcelExport) writeTop(params []map[string]string) {
	style := &excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
	topStyle, _ := l.file.NewStyle(style)
	var word = 'A'
	//首行写入
	for _, conf := range params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		line := fmt.Sprintf("%c1", word)
		//设置标题
		_ = l.file.SetCellValue(l.sheetName, line, title)
		//列宽
		_ = l.file.SetColWidth(l.sheetName, fmt.Sprintf("%c", word), fmt.Sprintf("%c", word), width)
		//设置样式
		_ = l.file.SetCellStyle(l.sheetName, line, line, topStyle)
		word++
	}
}

// 写入数据
func (l *lzExcelExport) writeData(params []map[string]string, data []map[string]interface{}) {
	style := &excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
	lineStyle, _ := l.file.NewStyle(style)
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range data {
		//设置行高
		_ = l.file.SetRowHeight(l.sheetName, i+1, defaultHeight)
		//逐列写入
		var word = 'A'
		for _, conf := range params {
			valKey := conf["key"]
			line := fmt.Sprintf("%c%v", word, j)
			isNum := conf["is_num"]

			//设置值
			if isNum != "0" {
				valNum := fmt.Sprintf("'%v", val[valKey])
				_ = l.file.SetCellValue(l.sheetName, line, valNum)
			} else {
				_ = l.file.SetCellValue(l.sheetName, line, val[valKey])
			}

			//设置样式
			_ = l.file.SetCellStyle(l.sheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	//设置行高 尾行
	_ = l.file.SetRowHeight(l.sheetName, len(data)+1, defaultHeight)
}

func (l *lzExcelExport) export(params []map[string]string, data []map[string]interface{}) {
	l.writeTop(params)
	l.writeData(params, data)
}

func createFile() *excelize.File {
	f := excelize.NewFile()
	// 创建一个默认工作表
	sheetName := defaultSheetName
	index, _ := f.NewSheet(sheetName)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	return f
}

func createFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("excle-%v-%v.xlsx", name, rand.Int63n(time.Now().Unix()))
}

// ExportExcelByStruct excel导出(数据源为Struct) []interface{}
func (l *lzExcelExport) ExportExcelByStruct(c *gin.Context, header []interface{}, data [][]interface{}, fileName string) error {
	style := &excelize.Style{
		Font: &excelize.Font{
			Color:  "#666666",
			Size:   13,
			Family: "arial",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}

	//设置表名
	l.file.NewSheet("Sheet1")
	writer, err := l.file.NewStreamWriter("Sheet1")
	if err != nil {
		return err
	}
	rowStyleID, _ := l.file.NewStyle(style)

	//修改列宽
	writer.SetColWidth(1, len(header), 30)
	//设置表头
	writer.SetRow("A1", header, excelize.RowOpts{
		Height: 30,
	})
	for i, v := range data {
		//索引转单元格坐标
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		//添加的数据
		writer.SetRow(cell, v, excelize.RowOpts{
			StyleID: rowStyleID,
		})
	}

	//结束流式写入
	writer.Flush()

	localXlsx := GetCurrentAbPathByCaller() + fileName + ".xlsx"
	l.file.SaveAs(localXlsx)
	xlsx, err := excelize.OpenFile(localXlsx)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(localXlsx)
	}()

	disposition := fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", disposition)
	c.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	return xlsx.Write(c.Writer)
}

// Letter 遍历a-z
func Letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, string(rune('A'+i)))
	}
	return str
}
