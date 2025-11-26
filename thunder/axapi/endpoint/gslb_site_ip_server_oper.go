package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type GslbSiteIpServerOper struct {
	IpServerName string `json:"ip-server-name"`

	Oper GslbSiteIpServerOperOper `json:"oper"`

	SiteName string
}
type DataGslbSiteIpServerOper struct {
	DtGslbSiteIpServerOper GslbSiteIpServerOper `json:"ip-server"`
}

type GslbSiteIpServerOperOper struct {
	IpServer            string                                 `json:"ip-server"`
	IpAddress           string                                 `json:"ip-address"`
	Desc                string                                 `json:"desc"`
	State               string                                 `json:"state"`
	ServiceIp           string                                 `json:"service-ip"`
	PortCount           int                                    `json:"port-count"`
	VirtualServer       int                                    `json:"virtual-server"`
	Disabled            int                                    `json:"disabled"`
	GslbProtocol        int                                    `json:"gslb-protocol"`
	LocalProtocol       int                                    `json:"local-protocol"`
	ManuallyHealthCheck int                                    `json:"manually-health-check"`
	Use_gslb_state      int                                    `json:"use_gslb_state"`
	Dynamic             int                                    `json:"dynamic"`
	Hits                int                                    `json:"hits"`
	Recent              int                                    `json:"recent"`
	DrsList             []GslbSiteIpServerOperOperDrsList      `json:"drs-list"`
	IpServerPort        []GslbSiteIpServerOperOperIpServerPort `json:"ip-server-port"`
}

type GslbSiteIpServerOperOperDrsList struct {
	DrsName                string                                   `json:"drs-name"`
	DrsIpAddress           string                                   `json:"drs-ip-address"`
	DrsFqdnName            string                                   `json:"drs-fqdn-name"`
	DrsState               string                                   `json:"drs-state"`
	DrsServiceIp           string                                   `json:"drs-service-ip"`
	DrsPortCount           int                                      `json:"drs-port-count"`
	DrsVirtualServer       int                                      `json:"drs-virtual-server"`
	DrsDisabled            int                                      `json:"drs-disabled"`
	DrsGslbProtocol        int                                      `json:"drs-gslb-protocol"`
	DrsLocalProtocol       int                                      `json:"drs-local-protocol"`
	DrsManuallyHealthCheck int                                      `json:"drs-manually-health-check"`
	DrsUse_gslb_state      int                                      `json:"drs-use_gslb_state"`
	DrsDynamic             int                                      `json:"drs-dynamic"`
	DrsHits                int                                      `json:"drs-hits"`
	DrsRecent              int                                      `json:"drs-recent"`
	DrsPort                []GslbSiteIpServerOperOperDrsListDrsPort `json:"drs-port"`
}

type GslbSiteIpServerOperOperDrsListDrsPort struct {
	Vport         int    `json:"vport"`
	VportProtocol string `json:"vport-protocol"`
	VportState    string `json:"vport-state"`
	ServiceName   string `json:"service-name"`
}

type GslbSiteIpServerOperOperIpServerPort struct {
	Vport         int    `json:"vport"`
	VportProtocol string `json:"vport-protocol"`
	VportState    string `json:"vport-state"`
	ServiceName   string `json:"service-name"`
}

func (p *GslbSiteIpServerOper) GetId() string {
	return "1"
}

func (p *GslbSiteIpServerOper) getPath() string {

	return "gslb/site/" + p.SiteName + "/ip-server/" + p.IpServerName + "/oper"
}

func (p *GslbSiteIpServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteIpServerOper, error) {
	logger.Println("GslbSiteIpServerOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataGslbSiteIpServerOper
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
