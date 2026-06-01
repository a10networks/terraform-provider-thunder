package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosNetworkObjectOper struct {
	ObjectName string `json:"object-name"`

	Oper DdosNetworkObjectOperOper `json:"oper"`

	TopkDestinations DdosNetworkObjectOperTopkDestinations `json:"topk-destinations"`
}
type DataDdosNetworkObjectOper struct {
	DtDdosNetworkObjectOper DdosNetworkObjectOper `json:"network-object"`
}

type DdosNetworkObjectOperOper struct {
	EntryList                 []DdosNetworkObjectOperOperEntryList `json:"entry-list"`
	EntryCount                int                                  `json:"entry-count"`
	Details                   int                                  `json:"details"`
	TrustedDetails            int                                  `json:"trusted-details"`
	TotalDetails              int                                  `json:"total-details"`
	VictimList                int                                  `json:"victim-list"`
	DiscoveredList            int                                  `json:"discovered-list"`
	SrcServiceList            int                                  `json:"src-service-list"`
	SubnetIpAddr              string                               `json:"subnet-ip-addr"`
	SubnetIpv6Addr            string                               `json:"subnet-ipv6-addr"`
	Ipv4                      string                               `json:"ipv4"`
	DiscoveredIpList          int                                  `json:"discovered-ip-list"`
	AnomalyIpList             int                                  `json:"anomaly-ip-list"`
	SrcPort                   int                                  `json:"src-port"`
	PortStart                 int                                  `json:"port-start"`
	PortEnd                   int                                  `json:"port-end"`
	Protocol                  int                                  `json:"protocol"`
	SingleLayerDiscoveredList int                                  `json:"single-layer-discovered-list"`
	AgentGroupDetails         int                                  `json:"agent-group-details"`
	AggregatedDetails         int                                  `json:"aggregated-details"`
}

type DdosNetworkObjectOperOperEntryList struct {
	Address                 string                                         `json:"address"`
	ChildrenCount           int                                            `json:"children-count"`
	PortRangeStart          int                                            `json:"port-range-start"`
	PortRangeEnd            int                                            `json:"port-range-end"`
	Port                    int                                            `json:"port"`
	ServiceProtocol         string                                         `json:"service-protocol"`
	AgentGroupName          string                                         `json:"agent-group-name"`
	Indicators              []DdosNetworkObjectOperOperEntryListIndicators `json:"indicators"`
	IsAnomaly               int                                            `json:"is-anomaly"`
	IsLearningDone          int                                            `json:"is-learning-done"`
	IsHistogramLearningDone int                                            `json:"is-histogram-learning-done"`
	OperationalMode         int                                            `json:"operational-mode"`
	HistogramMode           int                                            `json:"histogram-mode"`
	DisplayFilter           string                                         `json:"display-filter"`
	EsTimestamp             string                                         `json:"es-timestamp"`
	DeEsTimestamp           string                                         `json:"de-es-timestamp"`
	Estimated_bit_rate      string                                         `json:"estimated_bit_rate"`
}

type DdosNetworkObjectOperOperEntryListIndicators struct {
	IndicatorName  string                                              `json:"indicator-name"`
	IndicatorIndex int                                                 `json:"indicator-index"`
	Value          []DdosNetworkObjectOperOperEntryListIndicatorsValue `json:"value"`
	IsAnomaly      int                                                 `json:"is-anomaly"`
}

type DdosNetworkObjectOperOperEntryListIndicatorsValue struct {
	Current   string `json:"current"`
	Threshold string `json:"threshold"`
	Max       string `json:"max"`
}

type DdosNetworkObjectOperTopkDestinations struct {
	Oper DdosNetworkObjectOperTopkDestinationsOper `json:"oper"`
}

type DdosNetworkObjectOperTopkDestinationsOper struct {
	Indicators []DdosNetworkObjectOperTopkDestinationsOperIndicators `json:"indicators"`
	TopkType   int                                                   `json:"topk-type"`
	ResetTime  int                                                   `json:"reset-time"`
}

type DdosNetworkObjectOperTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                            `json:"indicator-name"`
	IndicatorIndex int                                                               `json:"indicator-index"`
	Destinations   []DdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

func (p *DdosNetworkObjectOper) GetId() string {
	return "1"
}

func (p *DdosNetworkObjectOper) getPath() string {
	return "ddos/network-object/" + p.ObjectName + "/oper"
}

func (p *DdosNetworkObjectOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectOper, error) {
	logger.Println("DdosNetworkObjectOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosNetworkObjectOper
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
