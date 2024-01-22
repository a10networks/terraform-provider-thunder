

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevOper struct {
    
    DeviceName string `json:"device-name"`
    Oper GslbSiteSlbDevOperOper `json:"oper"`
    VipServer GslbSiteSlbDevOperVipServer `json:"vip-server"`
    SiteName string 

}
type DataGslbSiteSlbDevOper struct {
    DtGslbSiteSlbDevOper GslbSiteSlbDevOper `json:"slb-dev"`
}


type GslbSiteSlbDevOperOper struct {
    Dev_name string `json:"dev_name"`
    Dev_ip string `json:"dev_ip"`
    Dev_attr string `json:"dev_attr"`
    Dev_admin_preference int `json:"dev_admin_preference"`
    Dev_session_num int `json:"dev_session_num"`
    Dev_session_util int `json:"dev_session_util"`
    Dev_gw_state string `json:"dev_gw_state"`
    Dev_ip_cnt int `json:"dev_ip_cnt"`
    DevState string `json:"dev-state"`
    ClientLdnsList []GslbSiteSlbDevOperOperClientLdnsList `json:"client-ldns-list"`
}


type GslbSiteSlbDevOperOperClientLdnsList struct {
    ClientIp string `json:"client-ip"`
    Age int `json:"age"`
    Type string `json:"type"`
    RdtSample1 int `json:"rdt-sample1"`
    RdtSample2 int `json:"rdt-sample2"`
    RdtSample3 int `json:"rdt-sample3"`
    RdtSample4 int `json:"rdt-sample4"`
    RdtSample5 int `json:"rdt-sample5"`
    RdtSample6 int `json:"rdt-sample6"`
    RdtSample7 int `json:"rdt-sample7"`
    RdtSample8 int `json:"rdt-sample8"`
}


type GslbSiteSlbDevOperVipServer struct {
    Oper GslbSiteSlbDevOperVipServerOper `json:"oper"`
    VipServerV4List []GslbSiteSlbDevOperVipServerVipServerV4List `json:"vip-server-v4-list"`
    VipServerV6List []GslbSiteSlbDevOperVipServerVipServerV6List `json:"vip-server-v6-list"`
    VipServerNameList []GslbSiteSlbDevOperVipServerVipServerNameList `json:"vip-server-name-list"`
}


type GslbSiteSlbDevOperVipServerOper struct {
}


type GslbSiteSlbDevOperVipServerVipServerV4List struct {
    Ipv4 string `json:"ipv4"`
    Oper GslbSiteSlbDevOperVipServerVipServerV4ListOper `json:"oper"`
}


type GslbSiteSlbDevOperVipServerVipServerV4ListOper struct {
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
    DevVipPortList []GslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}


type GslbSiteSlbDevOperVipServerVipServerV6List struct {
    Ipv6 string `json:"ipv6"`
    Oper GslbSiteSlbDevOperVipServerVipServerV6ListOper `json:"oper"`
}


type GslbSiteSlbDevOperVipServerVipServerV6ListOper struct {
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
    DevVipPortList []GslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}


type GslbSiteSlbDevOperVipServerVipServerNameList struct {
    VipName string `json:"vip-name"`
    Oper GslbSiteSlbDevOperVipServerVipServerNameListOper `json:"oper"`
}


type GslbSiteSlbDevOperVipServerVipServerNameListOper struct {
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
    DevVipPortList []GslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}

func (p *GslbSiteSlbDevOper) GetId() string{
    return "1"
}

func (p *GslbSiteSlbDevOper) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/slb-dev/"+p.DeviceName+"/oper"
}

func (p *GslbSiteSlbDevOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteSlbDevOper,error) {
logger.Println("GslbSiteSlbDevOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteSlbDevOper
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
