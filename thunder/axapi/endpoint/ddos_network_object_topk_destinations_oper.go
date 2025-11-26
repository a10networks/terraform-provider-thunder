package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosNetworkObjectTopkDestinationsOper struct {
	Oper DdosNetworkObjectTopkDestinationsOperOper `json:"oper"`

	ObjectName string
}
type DataDdosNetworkObjectTopkDestinationsOper struct {
	DtDdosNetworkObjectTopkDestinationsOper DdosNetworkObjectTopkDestinationsOper `json:"topk-destinations"`
}

type DdosNetworkObjectTopkDestinationsOperOper struct {
	Indicators []DdosNetworkObjectTopkDestinationsOperOperIndicators `json:"indicators"`
	TopkType   int                                                   `json:"topk-type"`
	ResetTime  int                                                   `json:"reset-time"`
}

type DdosNetworkObjectTopkDestinationsOperOperIndicators struct {
	IndicatorName  string                                                            `json:"indicator-name"`
	IndicatorIndex int                                                               `json:"indicator-index"`
	Destinations   []DdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}

type DdosNetworkObjectTopkDestinationsOperOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

func (p *DdosNetworkObjectTopkDestinationsOper) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectTopkDestinationsOper) getPath() string {

	return "ddos/network-object/" + p.ObjectName + "/topk-destinations/oper"
}

func (p *DdosNetworkObjectTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectTopkDestinationsOper, error) {
	logger.Println("DdosNetworkObjectTopkDestinationsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkObjectTopkDestinationsOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
