package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosSrcBasedPolicyOper struct {
	Name string `json:"name"`

	Oper DdosSrcBasedPolicyOperOper `json:"oper"`
}
type DataDdosSrcBasedPolicyOper struct {
	DtDdosSrcBasedPolicyOper DdosSrcBasedPolicyOper `json:"src-based-policy"`
}

type DdosSrcBasedPolicyOperOper struct {
	SrcBasedPolicyName      string                                       `json:"src-based-policy-name"`
	Ipv4TotalSingleIp       int                                          `json:"ipv4-total-single-ip"`
	Ipv4TotalSubnet         int                                          `json:"ipv4-total-subnet"`
	Ipv6TotalSingleIp       int                                          `json:"ipv6-total-single-ip"`
	Ipv6TotalSubnet         int                                          `json:"ipv6-total-subnet"`
	GeolocIpv4TotalSingleIp int                                          `json:"geoloc-ipv4-total-single-ip"`
	GeolocIpv4TotalSubnet   int                                          `json:"geoloc-ipv4-total-subnet"`
	GeolocIpv6TotalSingleIp int                                          `json:"geoloc-ipv6-total-single-ip"`
	GeolocIpv6TotalSubnet   int                                          `json:"geoloc-ipv6-total-subnet"`
	GeolocUnexpandedNode    int                                          `json:"geoloc-unexpanded-node"`
	ClassListEntries        []DdosSrcBasedPolicyOperOperClassListEntries `json:"class-list-entries"`
	AllEntries              int                                          `json:"all-entries"`
	ResourceUsage           int                                          `json:"resource-usage"`
}

type DdosSrcBasedPolicyOperOperClassListEntries struct {
	Address         string `json:"address"`
	ClassListName   string `json:"class-list-name"`
	GeoLocationName string `json:"geo-location-name"`
}

func (p *DdosSrcBasedPolicyOper) GetId() string {
	return "1"
}

func (p *DdosSrcBasedPolicyOper) getPath() string {
	return "ddos/src-based-policy/" + p.Name + "/oper"
}

func (p *DdosSrcBasedPolicyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosSrcBasedPolicyOper, error) {
	logger.Println("DdosSrcBasedPolicyOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosSrcBasedPolicyOper
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
