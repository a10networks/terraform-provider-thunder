

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceIpOper struct {
    
    NodeName string `json:"node-name"`
    Oper GslbServiceIpOperOper `json:"oper"`
    PortList []GslbServiceIpOperPortList `json:"port-list"`

}
type DataGslbServiceIpOper struct {
    DtGslbServiceIpOper GslbServiceIpOper `json:"service-ip"`
}


type GslbServiceIpOperOper struct {
    ServiceIp string `json:"service-ip"`
    Ip string `json:"ip"`
    State string `json:"state"`
    PortCount int `json:"port-count"`
    VirtualServer int `json:"virtual-server"`
    Disabled int `json:"disabled"`
    GslbProtocol int `json:"gslb-protocol"`
    LocalProtocol int `json:"local-protocol"`
    ManuallyHealthCheck int `json:"manually-health-check"`
    Use_gslb_state int `json:"use_gslb_state"`
    Dynamic int `json:"dynamic"`
}


type GslbServiceIpOperPortList struct {
    PortNum int `json:"port-num"`
    PortProto string `json:"port-proto"`
    Oper GslbServiceIpOperPortListOper `json:"oper"`
}


type GslbServiceIpOperPortListOper struct {
    ServicePort int `json:"service-port"`
    State string `json:"state"`
    Disabled int `json:"disabled"`
    GslbProtocol int `json:"gslb-protocol"`
    LocalProtocol int `json:"local-protocol"`
    Tcp int `json:"tcp"`
    ManuallyHealthCheck int `json:"manually-health-check"`
    Use_gslb_state int `json:"use_gslb_state"`
    Dynamic int `json:"dynamic"`
}

func (p *GslbServiceIpOper) GetId() string{
    return "1"
}

func (p *GslbServiceIpOper) getPath() string{
    
    return "gslb/service-ip/"+p.NodeName+"/oper"
}

func (p *GslbServiceIpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbServiceIpOper,error) {
logger.Println("GslbServiceIpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbServiceIpOper
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
