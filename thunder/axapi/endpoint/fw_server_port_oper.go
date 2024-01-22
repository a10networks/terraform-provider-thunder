

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FwServerPortOper struct {
    
    Oper FwServerPortOperOper `json:"oper"`
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Name string 

}
type DataFwServerPortOper struct {
    DtFwServerPortOper FwServerPortOper `json:"port"`
}


type FwServerPortOperOper struct {
    State string `json:"state"`
    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
    Vrid int `json:"vrid"`
    Ha_group_id int `json:"ha_group_id"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}

func (p *FwServerPortOper) GetId() string{
    return "1"
}

func (p *FwServerPortOper) getPath() string{
    
    return "fw/server/" +p.Name + "/port/" +strconv.Itoa(p.PortNumber)+"+"+p.Protocol+"/oper"
}

func (p *FwServerPortOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwServerPortOper,error) {
logger.Println("FwServerPortOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwServerPortOper
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
