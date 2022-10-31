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

	fileName := "market.xml"

	byteValue, err := readXmlFile(fileName)
	if err != nil {
		log.Println("here is err:", err)
	}

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

func readXmlFile(fileName string) (data []byte, err error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return []byte{}, err
	}
	// 關閉檔案
	defer xmlFile.Close()

	// 讀檔
	return ioutil.ReadAll(xmlFile)
}
