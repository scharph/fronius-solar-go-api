# Fronius Solar Go Api
> Use the local api from a fronius solar inverter.

## Usage example

```go
package main

import (
	"fmt"
	"github.com/scharph/fronius-solar-go-api/froniussolar"
	"context"
)

func main(){

	client := froniussolar.NewClient("<<INVERTER IP>>")

	ctx := context.Background();

	data, err := client.GetCommonInverterData(ctx, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.EnergyDay.Pretty())

}
```
## Example setup

Describe how to install all development dependencies and how to run an automated test-suite of some kind. Potentially do this for multiple platforms.

```sh
cd main
go run main.go
```

## Release History

* 0.0.1
    * Work in progress

## Meta

Philipp Schartlmüller – [@1scharph](https://twitter.com/1scharph)

[https://github.com/scharph](https://github.com/scharph)

## Contributing

1. Fork it (<https://github.com/scharph/fronius-solar-go-api.git>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
