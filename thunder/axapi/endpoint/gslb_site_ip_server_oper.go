

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteIpServerOper struct {
    
    IpServerName string `json:"ip-server-name"`
    Oper GslbSiteIpServerOperOper `json:"oper"`
    SiteName string 

}
type DataGslbSiteIpServerOper struct {
    DtGslbSiteIpServerOper GslbSiteIpServerOper `json:"ip-server"`
}


type GslbSiteIpServerOperOper struct {
    IpServer string `json:"ip-server"`
    IpAddress string `json:"ip-address"`
    State string `json:"state"`
    ServiceIp string `json:"service-ip"`
    PortCount int `json:"port-count"`
    VirtualServer int `json:"virtual-server"`
    Disabled int `json:"disabled"`
    GslbProtocol int `json:"gslb-protocol"`
    LocalProtocol int `json:"local-protocol"`
    ManuallyHealthCheck int `json:"manually-health-check"`
    Use_gslb_state int `json:"use_gslb_state"`
    Dynamic int `json:"dynamic"`
    IpServerPort []GslbSiteIpServerOperOperIpServerPort `json:"ip-server-port"`
}


type GslbSiteIpServerOperOperIpServerPort struct {
    Vport int `json:"vport"`
    VportState string `json:"vport-state"`
}

func (p *GslbSiteIpServerOper) GetId() string{
    return "1"
}

func (p *GslbSiteIpServerOper) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/ip-server/"+p.IpServerName+"/oper"
}

func (p *GslbSiteIpServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteIpServerOper,error) {
logger.Println("GslbSiteIpServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteIpServerOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
