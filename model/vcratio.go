package model

import (
	"JMVC-Middleware/config"
)

//Function that will process sensor data before insert it to InfluxDB
func countVCRatio(counterData CounterData) (float64, float64) {
	kendaraan_up_total := counterData.CarUp + counterData.BusSUp + counterData.BusLUp + counterData.TruckSUp + counterData.TruckMUp + counterData.TruckLUp + counterData.TruckXLUp

	persentase_gol_1_up := float64(counterData.CarUp)/float64(kendaraan_up_total)*100 + float64(counterData.TruckSUp)/float64(kendaraan_up_total)*100
	kapasitas_gol_1_up := config.EMV_GOL_1 * persentase_gol_1_up * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_small_bus_up := float64(counterData.BusSUp) / float64(kendaraan_up_total) * 100
	kapasitas_gol_small_bus_up := config.EMV_GOL_SMALL_BUS * persentase_gol_small_bus_up * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_big_bus_up := float64(counterData.BusLUp) / float64(kendaraan_up_total) * 100
	kapasitas_gol_big_bus_up := config.EMV_GOL_BIG_BUS * persentase_gol_big_bus_up * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_2_up := float64(counterData.TruckMUp) / float64(kendaraan_up_total) * 100
	kapasitas_gol_2_up := config.EMV_GOL_2 * persentase_gol_2_up * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_345_up := float64(counterData.TruckLUp)/float64(kendaraan_up_total)*100 + float64(counterData.TruckXLUp)/float64(kendaraan_up_total)*100
	kapasitas_gol_345_up := config.EMV_GOL_345 * persentase_gol_345_up * config.BEBAN_RUAS_MENIT / 100

	volume_lalulintas_up := kapasitas_gol_small_bus_up + kapasitas_gol_big_bus_up + kapasitas_gol_2_up + kapasitas_gol_345_up + kapasitas_gol_1_up
	vcr_up := volume_lalulintas_up / float64(config.KAPASITAS_JALAN_MENIT)

	kendaraan_Down_total := counterData.CarDown + counterData.BusSDown + counterData.BusLDown + counterData.TruckSDown + counterData.TruckMDown + counterData.TruckLDown + counterData.TruckXLDown

	persentase_gol_1_Down := float64(counterData.CarDown)/float64(kendaraan_Down_total)*100 + float64(counterData.TruckSDown)/float64(kendaraan_Down_total)*100
	kapasitas_gol_1_Down := config.EMV_GOL_1 * persentase_gol_1_Down * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_small_bus_Down := float64(counterData.BusSDown) / float64(kendaraan_Down_total) * 100
	kapasitas_gol_small_bus_Down := config.EMV_GOL_SMALL_BUS * persentase_gol_small_bus_Down * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_big_bus_Down := float64(counterData.BusLDown) / float64(kendaraan_Down_total) * 100
	kapasitas_gol_big_bus_Down := config.EMV_GOL_BIG_BUS * persentase_gol_big_bus_Down * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_2_Down := float64(counterData.TruckMDown) / float64(kendaraan_Down_total) * 100
	kapasitas_gol_2_Down := config.EMV_GOL_2 * persentase_gol_2_Down * config.BEBAN_RUAS_MENIT / 100

	persentase_gol_345_Down := float64(counterData.TruckLDown)/float64(kendaraan_Down_total)*100 + float64(counterData.TruckXLDown)/float64(kendaraan_Down_total)*100
	kapasitas_gol_345_Down := config.EMV_GOL_345 * persentase_gol_345_Down * config.BEBAN_RUAS_MENIT / 100

	volume_lalulintas_Down := kapasitas_gol_small_bus_Down + kapasitas_gol_big_bus_Down + kapasitas_gol_2_Down + kapasitas_gol_345_Down + kapasitas_gol_1_Down
	vcr_Down := volume_lalulintas_Down / float64(config.KAPASITAS_JALAN_MENIT)
	return vcr_up, vcr_Down
}
