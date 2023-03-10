package downloadDemo

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func PostReportPlatformDownload() {
	columnDefineMap = make(map[string]*ColumnDefine)
	for _, v := range columnDefineList {
		columnDefineMap[v.value] = v
	}

	var items []string
	for _, v := range diemColumnList {
		items = append(items, v.String())
		if d, ok := columnDefineMap[v.as]; ok {
			excelColumnList = append(excelColumnList, d)
		}
	}

	t := time.Now()
	fileName := fmt.Sprintf("platform_report-%s", t.Format("2006-01-02"))
	fileName += ".csv"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("文件创建失败")
	}

	if f != nil {
		defer f.Close()
	}
	csvWriter := csv.NewWriter(f)
	var columnHead []string
	for _, v := range excelColumnList {
		columnHead = append(columnHead, v.label)
		//fmt.Println(v.value)

	}
	csvWriter.Write(columnHead)
	data1 := &ReportPlatformItem{}
	data1.Date = "2023"
	data1.ProjectName = "捕鱼达人"
	data1.StudioName = "灵眸工作室"
	data1.OptimizerName = "王靖勃"
	data1.PlatformId = "广点通"
	data1.AdvertiserId = "12345"
	data1.AdvertiserNameDisplay = "捕鱼广告"
	data1.AdgroupId = "6789"
	data1.AdgroupNameDisplay = "安卓-全版位-男性付费-ROI视频-0118-4"
	data1.AdId = "777"
	data1.AdNameDisplay = "广告名称"
	data1.AdgroupNameDisplay = "捕鱼达人广告"
	data1.Month = "month"
	data1.ProjectId = "id"
	data1.Show = 1
	data1.Click = 2
	data1.Convert = 3
	data1.Active = 4
	result.DataList = append(result.DataList, data1)
	//for _, v := range columnHead {
	//	fmt.Printf(v + " ")
	//}
	//fmt.Println()
	for _, v := range result.DataList {
		csvWriter.Write(getLine(excelColumnList, v))
		//for _, i := range getLine(excelColumnList, v) {
		//	fmt.Print(i + " ")
		//}
	}

	csvWriter.Flush()

	return
}
