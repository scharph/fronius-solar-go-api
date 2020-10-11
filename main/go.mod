module main

replace github.com/scharph/fronius-solar-go-api/froniussolar => ../froniussolar

go 1.15

require (
	github.com/scharph/fronius-solar-go-api/froniussolar v0.0.0
	golang.org/x/tools v0.0.0-20201011145850-ed2f50202694 // indirect
)
