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

	fmt.Println("Baseurl:", info.BaseURL, "Apiversion", info.APIVersion)

	args := froniussolar.InverterRealtimeDataArgs{nil, nil, nil}

	client.GetInverterRealtimeData(ctx, args)

}
