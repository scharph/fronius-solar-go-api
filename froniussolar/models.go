package froniussolar

import (
	"fmt"
	"net/http"
)

//Client ...
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Info ...
type Info struct {
	APIVersion        *int    `json:"APIVersion"`
	BaseURL           *string `json:"BaseURL"`
	CompatiblityRange *string `json:"CompatibilityRange"`
}

func (i *Info) Print() {
	fmt.Printf("APIVersion: %d, BaseURL: %s, CompatiblityRange: %s\n", *i.APIVersion, *i.BaseURL, *i.CompatiblityRange)
}

type DataPoint struct {
	Value *float32 `json:"Value"`
	Unit  *string  `json:"Unit"`
}

func (d *DataPoint) Print() {
	fmt.Printf("%.2f %s\n", *d.Value, *d.Unit)
}

func (d *DataPoint) ToString() string {
	return fmt.Sprintf("%.2f %s", *d.Value, *d.Unit)
}

type CumulationInverterData struct {
	EnergyDay   *DataPoint `json:"DAY_ENERGY"`
	EnergyTotal *DataPoint `json:"TOTAL_ENERGY"`
	EnergyYear  *DataPoint `json:"YEAR_ENERGY"`
	PowerAC     *DataPoint `json:"PAC"`
}

type CommonInverterData struct {
	VoltageDC   *DataPoint `json:"UDC"`
	VoltageAC   *DataPoint `json:"UAC"`
	CurrentDC   *DataPoint `json:"IDC"`
	CurrentAC   *DataPoint `json:"IAC"`
	EnergyDay   *DataPoint `json:"DAY_ENERGY"`
	EnergyTotal *DataPoint `json:"TOTAL_ENERGY"`
	EnergyYear  *DataPoint `json:"YEAR_ENERGY"`
	PowerAC     *DataPoint `json:"PAC"`
	FrequencyAC *DataPoint `json:"FAC"`
}

type MinMaxInverterData struct {
	MaxACPowerDay   *DataPoint `json:"DAY_PMAX"`
	MaxACVoltageDay *DataPoint `json:"DAY_UACMAX"`
	MinACVoltageDay *DataPoint `json:"DAY_UACMIN"`
	MaxDCVoltageDay *DataPoint `json:"DAY_UDCMAX"`

	MaxACPowerYear   *DataPoint `json:"YEAR_PMAX"`
	MaxACVoltageYear *DataPoint `json:"YEAR_UACMAX"`
	MinACVoltageYear *DataPoint `json:"YEAR_UACMIN"`
	MaxDCVoltageYear *DataPoint `json:"YEAR_UDCMAX"`

	MaxACPowerTotal   *DataPoint `json:"TOTAL_PMAX"`
	MaxACVoltageTotal *DataPoint `json:"TOTAL_UACMAX"`
	MinACVoltageTotal *DataPoint `json:"TOTAL_UACMIN"`
	MaxDCVoltageTotal *DataPoint `json:"TOTAL_UDCMAX"`
}
