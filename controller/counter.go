package controller

import (
	"JMVC-Middleware/model"
	utilities "JMVC-Middleware/utility"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NeowayLabs/wabbit"
	client "github.com/influxdata/influxdb1-client/v2"
	"gopkg.in/go-playground/validator.v9"
)

type CounterData struct {
	ID          string  `json:"id"`
	Time        string  `json:"time"`
	CarUp       int     `json:"car_up"`
	BusSUp      int     `json:"bus(s)_up`
	BusLUp      int     `json:"bus(l)_up"`
	TruckSUp    int     `json:"truck(s)_up"`
	TruckMUp    int     `json:"truck(m)_up"`
	TruckLUp    int     `json:"truck(l)_up"`
	TruckXLUp   int     `json:"truck(xl)_up"`
	SpeedUp     float64 `json:"speed_up"`
	CarDown     int     `json:"car_down"`
	BusSDown    int     `json:"bus(s)_down"`
	BusLDown    int     `json:"bus(l)_down"`
	TruckSDown  int     `json:"truck(s)_down"`
	TruckMDown  int     `json:"truck(m)_down"`
	TruckLDown  int     `json:"truck(l)_down"`
	TruckXLDown int     `json:"truck(xl)_down"`
	SpeedDown   float64 `json:"speed_down"`
}

//Function that will process sensor data before insert it to InfluxDB
func InsertCounterData(msg wabbit.Delivery, cInflux client.Client) error {
	jsonData := CounterData{}
	//Convert data body to a struct
	err := json.Unmarshal(msg.Body(), &jsonData)
	if err != nil {
		msg.Ack(false)
		utilities.Error.Printf("Error parsing data : %s \n", err)
		return errors.New(fmt.Sprintf("Error parsing data : %s \n", err))
	} else {
		//validate struct data
		validate := validator.New()
		err = validate.Struct(jsonData)

		if err != nil {
			msg.Ack(false)
			utilities.Error.Printf("Error validate data : %s \n", err)
			return errors.New(fmt.Sprintf("Error validate data : %s \n", err))
		} else {
			cm := model.CounterModel{
				CInflux: cInflux,
			}

			err := cm.InsertCounterToInflux(model.CounterData(jsonData))

			if err != nil {
				utilities.Error.Printf("Error inserting data : %s", jsonData.ID, err)
				return err
			}
		}
	}
	return nil
}
