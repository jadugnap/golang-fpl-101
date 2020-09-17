// Package csv provides ability to store any StructSlice (bool|int|string) into csv
package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

// StructSlice is automatically stored (bool|int|string) into csv (string)
func StructSlice(myStructSlice interface{}, filePrefix string) {
	// recover if myStructSlice is not valid
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover from panic: %+v\n", err)
			fmt.Printf("myStructSlice: %+v\n", myStructSlice)
		}
	}()
	// generate emptyInstance of myStructSlice
	emptyInstance := reflect.Zero(reflect.TypeOf(myStructSlice).Elem())

	// define file and writer
	fileName := fmt.Sprintf("csv_out/%v-%v.csv", filePrefix, time.Now().Format("2006-01-02_15:00:00"))
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("error os.Create():", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write headerNames based on emptyInstance
	headerNames := []string{}
	headerTypes := []reflect.Type{}
	for i := 0; i < emptyInstance.NumField(); i++ {
		headerNames = append(headerNames, emptyInstance.Type().Field(i).Name)
		headerTypes = append(headerTypes, emptyInstance.Type().Field(i).Type)
	}
	err = writer.Write(headerNames)
	if err != nil {
		log.Println("error writer.Write():", err)
		return
	}

	// appendEachRow should append row-by-row based on headerTypes
	row := reflect.ValueOf(myStructSlice)
	for i := 0; i < row.Len(); i++ {
		appendEachRow(writer, headerTypes, row.Index(i).Interface())
	}
}

func appendEachRow(writer *csv.Writer, headerTypes []reflect.Type, row interface{}) {
	boolType := reflect.TypeOf((bool)(false))
	intType := reflect.TypeOf((int)(0))
	stringType := reflect.TypeOf((string)(""))

	// use Field(i) to detect column's type and append as string
	allColumns := []string{}
	v := reflect.ValueOf(row)
	for i := 0; i < v.NumField(); i++ {
		switch headerTypes[i] {
		case boolType:
			str := strconv.FormatBool(v.Field(i).Bool())
			allColumns = append(allColumns, str)
		case intType:
			str := strconv.FormatInt(v.Field(i).Int(), 10)
			allColumns = append(allColumns, str)
		case stringType:
			allColumns = append(allColumns, v.Field(i).String())
		default:
			log.Panic("csv.StructSlice() only accept bool, int, or string. headerType:", headerTypes[i])
		}
	}
	err := writer.Write(allColumns)
	if err != nil {
		log.Println("error writer.Write():", err)
		return
	}
}
