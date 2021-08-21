package powerformula

import (
	"errors"
	"log"
)

// PowerModel data model for Power formula
type PowerModel struct {
	Power float64
	Work  float64
	Time  float64
}

// GetResultingPower calculate a single work and time to get the resulting power
func GetResultingPower(w float64, t float64) (PowerModel, error) {
	newPowerModel := PowerModel{}

	if w <= 0 && t <= 0 {
		return newPowerModel, errors.New("work or time is less than or equal to zero")
	}

	// Power Formula is p=w/t
	p := w / t
	newPowerModel.Work = w
	newPowerModel.Time = t
	newPowerModel.Power = p
	return newPowerModel, nil
}

// GetListOfPower process the list of given work and time. Throws an error if the work and time is less than zero
func GetListOfPower(pm []PowerModel) []PowerModel {
	result := []PowerModel{}

	for _, val := range pm {
		calculateResult, err := GetResultingPower(val.Work, val.Time)
		if err != nil {
			log.Fatalf(err.Error())
		}
		result = append(result, calculateResult)
	}
	return result
}

// GetTransposeTime calculate the time from work and power
func GetTransposeTime(w float64, p float64) (PowerModel, error) {
	t := 0.
	newPowerModel := PowerModel{}

	if w < 0 && p < 0 {
		return newPowerModel, errors.New("given work or power must be greater than zero")
	}
	// time formula is t = w/p
	t = w / p
	newPowerModel.Work = w
	newPowerModel.Power = p
	newPowerModel.Time = t

	return newPowerModel, nil
}

// GetTransposeWork calculates the work from the given time and power
func GetTransposeWork(t float64, p float64) (PowerModel, error) {
	w := 0.
	newPowerModel := PowerModel{}

	if t <= 0 && p <= 0 {
		return newPowerModel, errors.New("given time or power must be greater than zero")
	}

	// work formula is w=p*t
	w = p * t
	newPowerModel.Power = p
	newPowerModel.Work = w
	newPowerModel.Time = t
	return newPowerModel, nil
}
