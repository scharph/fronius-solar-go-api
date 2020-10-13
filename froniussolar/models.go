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

func (i *Info) PrintApiInfo() {
	if i != nil {
		fmt.Printf("APIVersion: %d, BaseURL: %s, CompatiblityRange: %s\n", *i.APIVersion, *i.BaseURL, *i.CompatiblityRange)
	}
	fmt.Println("error")
}

type DataPoint struct {
	Value *float32 `json:"Value"`
	Unit  *string  `json:"Unit"`
}

func (d *DataPoint) Print() {
	if d != nil {
		fmt.Printf("%.2f %s\n", *d.Value, *d.Unit)
	}
	fmt.Println("nil")

}

// Pretty Return formatted string with Unit
func (d *DataPoint) Pretty() string {
	if d != nil {
		return fmt.Sprintf("%.2f %s", *d.Value, *d.Unit)
	}
	return "nil"

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

type PowerFlowRealtimeData struct {
	MeterLocation      *string  `json:"Meter_Location"`
	Mode               *string  `json:"Mode"`
	PGrid              *float32 `json:"P_Grid"`
	PLoad              *float32 `json:"P_Load"`
	PAkku              *float32 `json:"P_Akku"`
	PPV                *float32 `json:"P_PV"`
	RelSelfConsumption *float32 `json:"rel_SelfConsumption"`
	RelAutonomy        *float32 `json:"rel_Autonomy"`
}

// LoggerInfo This provides information about the logging device which provides this API.
type LoggerInfo struct {
	CO2Factor        *float32 `json:"CO2Factor"`
	CO2Unit          string   `json:"CO2Unit"`
	CashCurrency     string   `json:"CashCurrency"`
	CashFactor       *float32 `json:"CashFactor"`
	DeliveryFactor   *float32 `json:"DeliveryFactor"`
	HWVersion        *string  `json:"HWVersion"`
	PlatformID       *string  `json:"PlatformID"`
	ProductID        *string  `json:"ProductID"`
	SWVersion        *string  `json:"SWVersion"`
	TimezoneLocation *string  `json:"TimezoneLocation"`
	TimezoneName     *string  `json:"TimezoneName"`
	UTCOffset        int      `json:"UTCOffset"`
	UniqueID         *string  `json:"UniqueID"`
}
