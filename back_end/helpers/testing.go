package helpers

import (
	"fmt"
	"os"
	"reflect"

	"gorm.io/gorm"
)

func CompareReflectedStructFields(firstI, secondI any) error {
	first := reflect.ValueOf(firstI)
	second := reflect.ValueOf(secondI)

	// if first.Type().Name() != second.Type().Name() {
	if first.Kind() != reflect.Struct || first.Type() != second.Type() {
		return fmt.Errorf("struct types did not match got %v -> %v, expected %v -> %v\n", first.Type().Name(), first.Type(), second.Type().Name(), second.Type())
	}

	for i := range first.NumField() {
		firstFieldValue := first.Field(i)
		secondFieldValue := second.Field(i)

		if firstFieldValue.Type().Name() == "Time" {
			continue
		}

		if firstFieldValue.Kind() == reflect.Struct {
			err := CompareReflectedStructFields(firstFieldValue, secondFieldValue)
			if err != nil {
				return fmt.Errorf("inner struct %v, %v", firstFieldValue.Type().Name(), err.Error())
			}
		}

		if !firstFieldValue.Equal(secondFieldValue) {
			return fmt.Errorf("fields did not match: got %v=%[2]v, expected %[1]v=%[3]v", first.Type().Field(i).Name, firstFieldValue, secondFieldValue)
		}
	}
	return nil
}

func LoadFixturesIntoDB(db *gorm.DB, location string, isDir bool) error {
	if isDir {
		fileNames, err := os.ReadDir(location)
		if err != nil {
			return err
		}
		for _, fileName := range fileNames {
			err := LoadFixturesIntoDB(db, location+"/"+fileName.Name(), fileName.IsDir())
			if err != nil {
				return err
			}
		}
	} else {
		file, err := os.ReadFile(location)
		if err != nil {
			return err
		}

		err = db.Exec(string(file)).Error
		if err != nil {
			return err
		}
	}

	return nil
}
