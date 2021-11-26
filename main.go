package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Manufacturer string
	Model        string
	KM           float64
}

func main() {

	//DOSYAYI AÇIKTIK
	csvFile, err := os.Open("/home/mert/Desktop/golangCSVtoJson/carsales.csv")

	// ERRORU BURDA YAKALADIK
	if err != nil {
		fmt.Println(err)
	}
	// TÜM İŞLEMLERDEN SONRA DOSYAYI KAPAT
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	//DOSYAYI OKU
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var emp Employee
	//BOŞ BİR LİSTE OLUŞTUR
	var employees []Employee

	//GELEN CSV DATAYI LOOPA AL
	for _, each := range csvData {
		//DATA'NIN 0.İNDEXE NAME YAZ
		emp.Manufacturer = each[0]
		//DATA'NIN 1.İNDEXE MODELİ YAZ
		emp.Model = each[1]
		//DATA FLOAT OLARAK 2.İNDEXE YAZ
		emp.KM, _ = strconv.ParseFloat(each[2], 64)
		//DATAYI BİR BOŞ BİR LİSTEYE EKLE
		employees = append(employees, emp)
	}

	jsonFile, _ := json.Marshal(employees)

	//ERROR'U YAKALA
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//JSON DOSYAYI CONSOLA YAZ
	fmt.Println(string(jsonFile))
	//JSON DOSYA OLUŞTUR
	jsonData, err := os.Create("/home/mert/Desktop/golangCSVtoJson/JsonData.json")

	//ERROR MESAJI YAKALA
	if err != nil {
		fmt.Println(err)

	}
	//EN SON DOSYAYI KAPAT
	defer jsonData.Close()

	//GELEN DATAYI DOSYAYA YAZ
	jsonData.Write(jsonFile)
	//DOSYAYI KAPAT
	jsonData.Close()

}
