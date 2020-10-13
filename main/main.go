package main

import (
	"context"
	"fmt"
	"log"

	"github.com/scharph/fronius-solar-go-api/froniussolar"
)

const (
	url = "http://10.0.0.17/solar_api/v1"
)

func main() {
	client := froniussolar.NewClient(url)
	ctx := context.Background()

	// info, err := client.GetAPIVersionInfo(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// info.PrintApiInfo()

	fmt.Println("\n### GetCommonInverterData")
	commonData, err := client.GetCommonInverterData(ctx, 1)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Energy Day:", commonData.EnergyDay.Pretty())
	fmt.Println("Year Day:", commonData.EnergyYear.Pretty())
	fmt.Println("Total Day:", commonData.EnergyTotal.Pretty())
	fmt.Println("AC Power:", commonData.PowerAC.Pretty())
	fmt.Println("Current DC:", commonData.CurrentDC.Pretty())

	fmt.Println("\n### GetCommonInverterData")
	cumulationData, _ := client.GetCumulationInverterData(ctx, 1)

	fmt.Println("Day:", cumulationData.EnergyDay.Pretty())
	fmt.Println("Year:", cumulationData.EnergyYear.Pretty())
	fmt.Println("Total:", cumulationData.EnergyTotal.Pretty())

	fmt.Println("\n### GetCommonInverterData")
	minMaxData, _ := client.GetMinMaxInverterData(ctx, 1)

	fmt.Println("MaxACPowerDay:", minMaxData.MaxACPowerDay.Pretty())
	fmt.Println("MaxACPowerYear:", minMaxData.MaxACPowerYear.Pretty())
	fmt.Println("MaxACPowerTotal:", minMaxData.MaxACPowerTotal.Pretty())

	fmt.Println("\n### GetPowerFlowRealtimeData")
	flowRealtime, _ := client.GetPowerFlowRealtimeData(ctx)

	fmt.Println("MeterLocation:", *flowRealtime.MeterLocation)
	fmt.Println("Mode:", *flowRealtime.Mode)
	fmt.Println("PGrid:", *flowRealtime.PGrid)
	fmt.Println("PLoad:", *flowRealtime.PLoad)

	fmt.Println("\n### GetLoggerInfo")
	loggerInfo, _ := client.GetLoggerInfo(ctx)

	fmt.Println("SWVersion:", *loggerInfo.SWVersion)
	fmt.Println("HWVersion:", *loggerInfo.HWVersion)

}
