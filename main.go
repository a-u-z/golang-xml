package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// f := excelize.NewFile()

	// @@@@@ 如需要，請更改此觸 @@@@@
	// // 開啟檔案
	fileName := "market.xml"
	xmlFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// 關閉檔案
	defer xmlFile.Close()

	// 讀檔
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// @@@@@ 如需要，請更改此觸 @@@@@
	// 宣告一個根據 xml 結構的型別變數
	m := MarketDescriptions{}

	err = xml.Unmarshal(byteValue, &m)
	if err != nil {
		log.Println("here is err:", err)
	}

	// @@@@@ 如需要，請更改此觸 @@@@@
	// 業務邏輯處理  // 這邊需要改動
	ss := []string{}
	for _, v := range m.Market {
		if strings.Contains(v.Name, "total") || strings.Contains(v.Name, "Total") {
			ss = append(ss, fmt.Sprintf(" \"%v\":  true,\n", v.ID))
		}
	}

	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}
