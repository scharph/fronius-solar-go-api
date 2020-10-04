package froniussolar

import "net/http"

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
	APIVersion        int    `json:"APIVersion"`
	BaseURL           string `json:"BaseURL"`
	CompatiblityRange string `json:"CompatibilityRange"`
}

// InvDataCollection ... Only needed for Scope ”Device”
// Selects the collection of data that should be queried from the device.
type InvDataCollection int

const (
	// CumInvData = CumulationInverterData
	CumInvData InvDataCollection = iota
	// ComInvData = CommonInverterData
	ComInvData
	// Inv3PData = 3PInverterData
	Inv3PData
	// MinMaxInvData = MinMaxInverterData
	MinMaxInvData
)

func (i InvDataCollection) String() string {
	return [...]string{"CumulationInverterData", "CommonInverterData", "3PInverterData", "MinMaxInverterData"}[i]
}

// InverterRealtimeDataArgs .. Parameters
type InverterRealtimeDataArgs struct {
	Scope          *string            `json:"Scope"`
	DeviceID       *int               `json:"DeviceId"`
	DataCollection *InvDataCollection `json:"DataCollection"`
}
