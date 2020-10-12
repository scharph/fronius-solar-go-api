package main

import (
	"context"
	"fmt"

	"github.com/scharph/fronius-solar-go-api/froniussolar"
)

func main() {
	client := froniussolar.NewClient()
	ctx := context.Background()

	info, err := client.GetAPIVersionInfo(ctx)
	if err != nil {
		fmt.Println(err)
	}

	info.Print()

	commonData, _ := client.GetCommonInverterData(ctx, 1)

	fmt.Println("### GetCommonInverterData")
	fmt.Println("Energy Day:", commonData.EnergyDay.ToString())
	fmt.Println("Year Day:", commonData.EnergyYear.ToString())
	fmt.Println("Total Day:", commonData.EnergyTotal.ToString())
	fmt.Println("AC Power:", commonData.PowerAC.ToString())
	fmt.Println("Current DC:", commonData.CurrentDC.ToString())

	cumulationData, _ := client.GetCumulationInverterData(ctx, 1)
	fmt.Println("\n### GetCommonInverterData")

	fmt.Println("Day:", cumulationData.EnergyDay.ToString())
	fmt.Println("Year:", cumulationData.EnergyYear.ToString())
	fmt.Println("Total:", cumulationData.EnergyTotal.ToString())

	minMaxData, _ := client.GetMinMaxInverterData(ctx, 1)

	fmt.Println("\n### GetCommonInverterData")
	fmt.Println("MaxACPowerDay:", minMaxData.MaxACPowerDay.ToString())
	fmt.Println("MaxACPowerYear:", minMaxData.MaxACPowerYear.ToString())
	fmt.Println("MaxACPowerTotal:", minMaxData.MaxACPowerTotal.ToString())

}
