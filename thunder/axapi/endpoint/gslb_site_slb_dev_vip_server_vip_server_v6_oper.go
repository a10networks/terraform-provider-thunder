

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerV6Oper struct {
    
    Ipv6 string `json:"ipv6"`
    Oper GslbSiteSlbDevVipServerVipServerV6OperOper `json:"oper"`
    SiteName string 
    DeviceName string 

}
type DataGslbSiteSlbDevVipServerVipServerV6Oper struct {
    DtGslbSiteSlbDevVipServerVipServerV6Oper GslbSiteSlbDevVipServerVipServerV6Oper `json:"vip-server-v6"`
}


type GslbSiteSlbDevVipServerVipServerV6OperOper struct {
    Dev_vip_addr string `json:"dev_vip_addr"`
    Dev_vip_state string `json:"dev_vip_state"`
    NodeName string `json:"node-name"`
    ServiceIp string `json:"service-ip"`
    PortCount int `json:"port-count"`
    VirtualServer int `json:"virtual-server"`
    Disabled int `json:"disabled"`
    GslbProtocol int `json:"gslb-protocol"`
    LocalProtocol int `json:"local-protocol"`
    ManuallyHealthCheck int `json:"manually-health-check"`
    Use_gslb_state int `json:"use_gslb_state"`
    Dynamic int `json:"dynamic"`
    Hits int `json:"hits"`
    Recent int `json:"recent"`
    DevVipPortList []GslbSiteSlbDevVipServerVipServerV6OperOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteSlbDevVipServerVipServerV6OperOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}

func (p *GslbSiteSlbDevVipServerVipServerV6Oper) GetId() string{
    return "1"
}

func (p *GslbSiteSlbDevVipServerVipServerV6Oper) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/slb-dev/" +p.DeviceName + "/vip-server/vip-server-v6/"+p.Ipv6+"/oper"
}

func (p *GslbSiteSlbDevVipServerVipServerV6Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteSlbDevVipServerVipServerV6Oper,error) {
logger.Println("GslbSiteSlbDevVipServerVipServerV6Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteSlbDevVipServerVipServerV6Oper
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
