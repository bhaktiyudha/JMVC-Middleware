package config

import (
	utilities "JMVC-Middleware/utility"
	"os"
)

var (
	IS_PRODUCTION bool = false
)

//Change default config data if it's environment variable exist and not empty
func init() {
	if os.Getenv("InfluxDB_Name") != "" {
		InfluxDatabaseName = os.Getenv("InfluxDB_Name")
		utilities.Info.Printf("Influx Database Name : %s\n", InfluxDatabaseName)
	}

	if os.Getenv("InfluxDB_Username") != "" {
		InfluxUsername = os.Getenv("InfluxDB_Username")
		utilities.Info.Printf("Influx Username : %s\n", InfluxUsername)
	}

	if os.Getenv("InfluxDB_Password") != "" {
		InfluxPassword = os.Getenv("InfluxDB_Password")
		utilities.Info.Printf("Influx Password : %s\n", InfluxPassword)
	}

	if os.Getenv("InfluxDB_Address") != "" {
		InfluxAddress = os.Getenv("InfluxDB_Address")
		utilities.Info.Printf("Influx Address : %s\n", InfluxAddress)
	}

	if os.Getenv("InfluxDB_Port") != "" {
		InfluxPort = os.Getenv("InfluxDB_Port")
		utilities.Info.Printf("Influx Port : %s\n", InfluxPort)
	}

	if os.Getenv("Rabbit_Address") != "" {
		RABBIT_ADDRESS = os.Getenv("Rabbit_Address")
	}

	if os.Getenv("Rabbit_Port") != "" {
		RABBIT_PORT = os.Getenv("Rabbit_Port")
	}

	if os.Getenv("Rabbit_Username") != "" {
		RABBIT_USERNAME = os.Getenv("Rabbit_Username")
	}

	if os.Getenv("Rabbit_Password") != "" {
		RABBIT_PASSWORD = os.Getenv("Rabbit_Password")
	}

	if os.Getenv("IS_PRODUCTION") == "true" {
		IS_PRODUCTION = true
	}
}
