package connection

import (
	"JMVC-Middleware/config"
	"errors"
	"fmt"

	"github.com/influxdata/influxdb1-client/models"
	client "github.com/influxdata/influxdb1-client/v2"
)

func QueryDB(cmd string, cInflux client.Client) ([]client.Result, error) {
	q := client.Query{
		Command:  cmd,
		Database: config.InfluxDatabaseName,
	}

	var res []client.Result

	if response, err := cInflux.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

//function to get index of a column from influxdb response
func GetIndexColumn(column string, res models.Row) int {
	for index, v := range res.Columns {
		if v == column {
			return index
		}
	}

	return -1
}

//function to check if all columns in param "columnsName" is exists
func CheckIndexColumn(res models.Row, columnsName ...string) error {
	for _, column := range columnsName {
		indexColumn := GetIndexColumn(column, res)

		if indexColumn < 0 {
			return errors.New(fmt.Sprintf("%s column not found", column))
		}
	}

	return nil
}
