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

	data, err := client.GetCumulationInverterData(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Total.Print()
	data.Day.Print()

	data2, err := client.GetMinMaxInverterData(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	data2.MaxACVoltageTotal.Print()
	data2.MaxACPowerYear.Print()

}
