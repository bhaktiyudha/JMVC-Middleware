package config

import (
	utilities "JMVC-Middleware/utility"
	"os"
)

var (
	IS_PRODUCTION         bool    = false
	JUMLAH_LAJUR          float64 = 3
	KAPASITAS_DASAR       float64 = 2300
	KAPASITAS_DASAR_MENIT float64 = 2300 / 60
	LEBAR_LAJUR           float64 = 3600
	FCW                   float64 = 1.012
	FCSP                  float64 = 1.000
	KAPASITAS_JALAN       float64 = JUMLAH_LAJUR * KAPASITAS_DASAR * FCW * FCSP
	KAPASITAS_JALAN_MENIT float64 = JUMLAH_LAJUR * KAPASITAS_DASAR_MENIT * FCW * FCSP
	EMV_GOL_1             float64 = 1
	EMV_GOL_SMALL_BUS             = 1.3
	EMV_GOL_BIG_BUS               = 1.5
	EMV_GOL_2                     = 1.3
	EMV_GOL_345           float64 = 2
	BEBAN_RUAS_HARI       float64 = 51439
	PHF                   float64 = 0.06
	BEBAN_RUAS_JAM                = BEBAN_RUAS_HARI * PHF
	BEBAN_RUAS_MENIT              = BEBAN_RUAS_JAM / 60
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
