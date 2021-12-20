package model

import (
	"JMVC-Middleware/config"
	utilities "JMVC-Middleware/utility"
	"errors"
	"fmt"
	"strings"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
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

type CounterModel struct {
	CInflux client.Client
}

func (cm CounterModel) InsertCounterToInflux(counterData CounterData) error {
	ID := counterData.ID

	//Create a batch for inserting sensor data
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.InfluxDatabaseName,
		Precision: "ns",
	})

	utilities.Info.Println(counterData)

	//Check if sensorID is empty or not
	if strings.TrimSpace(ID) == "" {
		return errors.New("sensor ID is empty")
	}

	tags := map[string]string{
		"ID": ID,
	}
	layoutFormat := "2006-01-02 15:04:05"
	sendTime, _ := time.Parse(layoutFormat, counterData.Time)

	pt, err := client.NewPoint("counter_data", tags, map[string]interface{}{
		"car_up":         counterData.CarUp,
		"bus(s)_up":      counterData.BusSUp,
		"bus(l)_up":      counterData.BusLUp,
		"truck(s)_up":    counterData.TruckSUp,
		"truck(m)_up":    counterData.TruckMUp,
		"truck(l)_up":    counterData.TruckLUp,
		"truck(xl)_up":   counterData.TruckXLUp,
		"speed_up":       counterData.SpeedUp,
		"car_down":       counterData.CarDown,
		"bus(s)_down":    counterData.BusSDown,
		"bus(l)_down":    counterData.BusLDown,
		"truck(s)_down":  counterData.TruckSDown,
		"truck(m)_down":  counterData.TruckMDown,
		"truck(l)_down":  counterData.TruckLDown,
		"truck(xl)_down": counterData.TruckXLDown,
		"speed_down":     counterData.SpeedDown,
	}, sendTime)

	if err != nil {
		return fmt.Errorf("error create point for inserting type parameter : %s", err)
	}

	bp.AddPoint(pt)

	if err := cm.CInflux.Write(bp); err != nil {
		return fmt.Errorf("error inserting type parameter : %s", err)
	}

	return nil
}
