package concurrency

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/MakMoinee/go-mith/pkg/common"
)

type ConcurrentService struct {
}

type IConcurrent interface {
	ProcessSomething(interface{}) (interface{}, int, error)
	ProcessConcurrentlySliceInt(item []int, occurrences int, myFunc func(i int, end int, wg *sync.WaitGroup, item []int, resultChan chan []int, errChan chan error)) ([]int, []error)
	ProcessConcurrentlySliceStr(item []string, occurrences int, myFunc func(i int, end int, wg *sync.WaitGroup, item []string, resultChan chan []string, errChan chan error)) ([]string, []error)
	GetServiceStruct() ConcurrentService
}

func NewService() IConcurrent {
	svc := ConcurrentService{}

	return &svc
}

// ProcessItem process slice of string
func ProcessItemSliceString(occurrences int, item []string, svc *ConcurrentService) []string {
	if occurrences <= 0 {
		panic("Occurance/s must not be zero")
	} else if len(item) <= 0 {
		panic("Item must not be empty")
	} else if len(item) < occurrences {
		panic("Item must not be less than occurance/s")
	}

	var result []string
	noOfCalls := len(item) / occurrences
	itemLength := len(item)
	errChan := make(chan error, noOfCalls)
	resultChan := make(chan []string, noOfCalls)
	end := itemLength
	var wg sync.WaitGroup

	for i := 0; i < occurrences; i = i + occurrences {
		if itemLength > occurrences {
			end = occurrences
		}
		wg.Add(1)
		go func(i, end int) {
			defer wg.Done()
			list := item[i:end]
			res, flag, err := svc.ProcessSomething(list)

			switch flag {
			case 1:
				resultChan <- res.([]string)
			default:
				resultChan <- []string{}
			}

			errChan <- err
		}(i, end)
		end = end + occurrences
	}
	wg.Wait()
	close(resultChan)
	close(errChan)

	for resultFromChan := range resultChan {
		if len(resultFromChan) != 0 {
			result = resultFromChan
		}
	}
	return result
}

func ProcessItem(occurrences int, item interface{}, svc IConcurrent) (interface{}, error) {
	var result interface{}
	var err []error
	flag := 0
	//initialize maps
	common.InitializeMaps()

	switch t := item.(type) {
	case []int:
		flag = 1
		dataType := reflect.ValueOf(t).Type()
		common.ValueTypes[dataType] = reflect.ValueOf(t)
	case []string:
		flag = 2
		dataType := reflect.ValueOf(t).Type()
		common.ValueTypes[dataType] = reflect.ValueOf(t)
	case int:
		flag = 3
		dataType := reflect.ValueOf(t).Type()
		common.ValueTypes[dataType] = reflect.ValueOf(t)
	case string:
		flag = 4
		dataType := reflect.ValueOf(t).Type()
		common.ValueTypes[dataType] = reflect.ValueOf(t)
	default:
		return nil, errors.New("data type not supported in processSomething()")
	}

	resultSliceInt, resultSliceString, resultInt, resultString := common.GetData(flag)

	if occurrences <= 0 {
		panic("Occurance/s must not be zero")
	} else {
		if len(resultSliceInt) == 0 && len(resultSliceString) == 0 && resultInt == 0 && len(strings.TrimSpace(resultString)) == 0 {
			panic("Item is empty")
		} else if len(resultSliceInt) < occurrences && len(resultSliceString) < occurrences && resultInt < occurrences && len(resultString) < occurrences {
			panic("Item must not be less than occurance/s >>" + fmt.Sprintf("%v %v %v %v", len(resultSliceInt), len(resultSliceString), resultInt, len(resultString)))
		}
	}

	switch flag {
	case 1:
		result, err = svc.ProcessConcurrentlySliceInt(resultSliceInt, occurrences, processItemSliceInt)
		if err != nil {
			return nil, err[0]
		}
	case 2:
		result, err = svc.ProcessConcurrentlySliceStr(resultSliceString, occurrences, processItemSliceString)
		if err != nil {
			return nil, err[0]
		}
	}

	//clear maps
	common.TypeValues = nil
	common.ValueTypes = nil

	return result, nil
}

// ProcessSomething process the data to be used by the concurrency call.
//
// It can be overridden by initializing new ConcurrentService Struct
func (svc *ConcurrentService) ProcessSomething(data interface{}) (interface{}, int, error) {

	switch t := data.(type) {
	case []string:
		return t, 2, nil
	case []int:
		return t, 1, nil
	default:
		return nil, 0, errors.New("data type not supported in processSomething()")
	}

}

func (svc *ConcurrentService) GetServiceStruct() ConcurrentService {
	return *svc
}

// ProcessConcurrentlySliceInt process []int data types concurrently with the predefined function
func (svc *ConcurrentService) ProcessConcurrentlySliceInt(item []int, occurrences int, myFunc func(i int, end int, wg *sync.WaitGroup, item []int, resultChan chan []int, errChan chan error)) ([]int, []error) {
	var result []int
	var err []error
	noOfCalls := len(item) / occurrences
	itemLength := len(item)
	errChan := make(chan error, noOfCalls)
	resultChan := make(chan []int, noOfCalls)
	end := itemLength
	var wg sync.WaitGroup

	for i := 0; i < occurrences; i = i + occurrences {
		if itemLength > occurrences {
			end = occurrences
		}
		wg.Add(1)
		go myFunc(i, end, &wg, item, resultChan, errChan)
		end = end + occurrences
	}
	wg.Wait()
	close(resultChan)
	close(errChan)

	for resultFromChan := range resultChan {
		if len(resultFromChan) != 0 {
			result = resultFromChan
		}
	}

	for errFromChan := range errChan {
		if errFromChan != nil {
			err = append(err, errFromChan)
		}
	}

	return result, err
}

// ProcessConcurrentlySliceStr process []string data types
func (svc *ConcurrentService) ProcessConcurrentlySliceStr(item []string, occurrences int, myFunc func(i int, end int, wg *sync.WaitGroup, item []string, resultChan chan []string, errChan chan error)) ([]string, []error) {
	var result []string
	var err []error
	noOfCalls := len(item) / occurrences
	itemLength := len(item)
	errChan := make(chan error, noOfCalls)
	resultChan := make(chan []string, noOfCalls)
	end := itemLength
	var wg sync.WaitGroup

	for i := 0; i < occurrences; i = i + occurrences {
		if itemLength > occurrences {
			end = occurrences
		}
		wg.Add(1)
		go myFunc(i, end, &wg, item, resultChan, errChan)
		end = end + occurrences
	}
	wg.Wait()
	close(resultChan)
	close(errChan)

	for resultFromChan := range resultChan {
		if len(resultFromChan) != 0 {
			result = resultFromChan
		}
	}

	for errFromChan := range errChan {
		if errFromChan != nil {
			err = append(err, errFromChan)
		}
	}

	return result, err
}

func processItemSliceString(i int, end int, wg *sync.WaitGroup, item []string, resultChan chan []string, errChan chan error) {
	defer wg.Done()
	list := item[i:end]

	svc := NewService()
	res, flag, err := svc.ProcessSomething(list)

	switch flag {
	case 2:
		resultChan <- res.([]string)
	default:
		resultChan <- []string{}
	}

	errChan <- err
}

func processItemSliceInt(i int, end int, wg *sync.WaitGroup, item []int, resultChan chan []int, errChan chan error) {
	defer wg.Done()
	list := item[i:end]

	svc := NewService()
	res, flag, err := svc.ProcessSomething(list)

	switch flag {
	case 1:
		resultChan <- res.([]int)
	default:
		resultChan <- []int{}
	}

	errChan <- err
}
