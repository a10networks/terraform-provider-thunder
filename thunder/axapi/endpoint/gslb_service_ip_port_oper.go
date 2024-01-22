

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceIpPortOper struct {
    
    Oper GslbServiceIpPortOperOper `json:"oper"`
    PortNum int `json:"port-num"`
    PortProto string `json:"port-proto"`
    NodeName string 

}
type DataGslbServiceIpPortOper struct {
    DtGslbServiceIpPortOper GslbServiceIpPortOper `json:"port"`
}


type GslbServiceIpPortOperOper struct {
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

func (p *GslbServiceIpPortOper) GetId() string{
    return "1"
}

func (p *GslbServiceIpPortOper) getPath() string{
    
    return "gslb/service-ip/" +p.NodeName + "/port/" +strconv.Itoa(p.PortNum)+"+"+p.PortProto+"/oper"
}

func (p *GslbServiceIpPortOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbServiceIpPortOper,error) {
logger.Println("GslbServiceIpPortOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbServiceIpPortOper
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
