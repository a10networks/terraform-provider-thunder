

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteOper struct {
    
    IpServerList []GslbSiteOperIpServerList `json:"ip-server-list"`
    Oper GslbSiteOperOper `json:"oper"`
    SiteName string `json:"site-name"`
    SlbDevList []GslbSiteOperSlbDevList `json:"slb-dev-list"`

}
type DataGslbSiteOper struct {
    DtGslbSiteOper GslbSiteOper `json:"site"`
}


type GslbSiteOperIpServerList struct {
    IpServerName string `json:"ip-server-name"`
    Oper GslbSiteOperIpServerListOper `json:"oper"`
}


type GslbSiteOperIpServerListOper struct {
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
    IpServerPort []GslbSiteOperIpServerListOperIpServerPort `json:"ip-server-port"`
}


type GslbSiteOperIpServerListOperIpServerPort struct {
    Vport int `json:"vport"`
    VportState string `json:"vport-state"`
}


type GslbSiteOperOper struct {
    GslbSite string `json:"gslb-site"`
    State string `json:"state"`
    TypeLast []GslbSiteOperOperTypeLast `json:"type-last"`
    ClientLdnsList []GslbSiteOperOperClientLdnsList `json:"client-ldns-list"`
}


type GslbSiteOperOperTypeLast struct {
    Type string `json:"type"`
    Last string `json:"last"`
}


type GslbSiteOperOperClientLdnsList struct {
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


type GslbSiteOperSlbDevList struct {
    DeviceName string `json:"device-name"`
    Oper GslbSiteOperSlbDevListOper `json:"oper"`
    VipServer GslbSiteOperSlbDevListVipServer `json:"vip-server"`
}


type GslbSiteOperSlbDevListOper struct {
    Dev_name string `json:"dev_name"`
    Dev_ip string `json:"dev_ip"`
    Dev_attr string `json:"dev_attr"`
    Dev_admin_preference int `json:"dev_admin_preference"`
    Dev_session_num int `json:"dev_session_num"`
    Dev_session_util int `json:"dev_session_util"`
    Dev_gw_state string `json:"dev_gw_state"`
    Dev_ip_cnt int `json:"dev_ip_cnt"`
    DevState string `json:"dev-state"`
    ClientLdnsList []GslbSiteOperSlbDevListOperClientLdnsList `json:"client-ldns-list"`
}


type GslbSiteOperSlbDevListOperClientLdnsList struct {
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


type GslbSiteOperSlbDevListVipServer struct {
    Oper GslbSiteOperSlbDevListVipServerOper `json:"oper"`
    VipServerV4List []GslbSiteOperSlbDevListVipServerVipServerV4List `json:"vip-server-v4-list"`
    VipServerV6List []GslbSiteOperSlbDevListVipServerVipServerV6List `json:"vip-server-v6-list"`
    VipServerNameList []GslbSiteOperSlbDevListVipServerVipServerNameList `json:"vip-server-name-list"`
}


type GslbSiteOperSlbDevListVipServerOper struct {
}


type GslbSiteOperSlbDevListVipServerVipServerV4List struct {
    Ipv4 string `json:"ipv4"`
    Oper GslbSiteOperSlbDevListVipServerVipServerV4ListOper `json:"oper"`
}


type GslbSiteOperSlbDevListVipServerVipServerV4ListOper struct {
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
    DevVipPortList []GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}


type GslbSiteOperSlbDevListVipServerVipServerV6List struct {
    Ipv6 string `json:"ipv6"`
    Oper GslbSiteOperSlbDevListVipServerVipServerV6ListOper `json:"oper"`
}


type GslbSiteOperSlbDevListVipServerVipServerV6ListOper struct {
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
    DevVipPortList []GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}


type GslbSiteOperSlbDevListVipServerVipServerNameList struct {
    VipName string `json:"vip-name"`
    Oper GslbSiteOperSlbDevListVipServerVipServerNameListOper `json:"oper"`
}


type GslbSiteOperSlbDevListVipServerVipServerNameListOper struct {
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
    DevVipPortList []GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}

func (p *GslbSiteOper) GetId() string{
    return "1"
}

func (p *GslbSiteOper) getPath() string{
    
    return "gslb/site/"+p.SiteName+"/oper"
}

func (p *GslbSiteOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteOper,error) {
logger.Println("GslbSiteOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteOper
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
