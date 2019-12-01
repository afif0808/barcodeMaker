package services

import (
	"fmt"
	"strconv"
)

func InterfaceArrayToInt64Array(interfaceArray []interface{}) ([]int64, error) {
	int64Array := []int64{}
	for _, v := range interfaceArray {
		int64V, interfaceToInt64Err := strconv.Atoi(fmt.Sprint(v))
		if interfaceToInt64Err != nil {
			return nil, fmt.Errorf("Error : the array contains a non-int data")
		}
		int64Array = append(int64Array, int64(int64V))
	}
	return int64Array, nil
}
