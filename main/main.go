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

	info.PrintApiInfo()

	commonData, _ := client.GetCommonInverterData(ctx, 1)

	fmt.Println("### GetCommonInverterData")
	fmt.Println("Energy Day:", commonData.EnergyDay.Pretty())
	fmt.Println("Year Day:", commonData.EnergyYear.Pretty())
	fmt.Println("Total Day:", commonData.EnergyTotal.Pretty())
	fmt.Println("AC Power:", commonData.PowerAC.Pretty())
	fmt.Println("Current DC:", commonData.CurrentDC.Pretty())

	cumulationData, _ := client.GetCumulationInverterData(ctx, 1)
	fmt.Println("\n### GetCommonInverterData")

	fmt.Println("Day:", cumulationData.EnergyDay.Pretty())
	fmt.Println("Year:", cumulationData.EnergyYear.Pretty())
	fmt.Println("Total:", cumulationData.EnergyTotal.Pretty())

	minMaxData, _ := client.GetMinMaxInverterData(ctx, 1)

	fmt.Println("\n### GetCommonInverterData")
	fmt.Println("MaxACPowerDay:", minMaxData.MaxACPowerDay.Pretty())
	fmt.Println("MaxACPowerYear:", minMaxData.MaxACPowerYear.Pretty())
	fmt.Println("MaxACPowerTotal:", minMaxData.MaxACPowerTotal.Pretty())

}
