

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwServerOper struct {
    
    Name string `json:"name"`
    Oper FwServerOperOper `json:"oper"`
    PortList []FwServerOperPortList `json:"port-list"`

}
type DataFwServerOper struct {
    DtFwServerOper FwServerOper `json:"server"`
}


type FwServerOperOper struct {
    State string `json:"state"`
}


type FwServerOperPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Oper FwServerOperPortListOper `json:"oper"`
}


type FwServerOperPortListOper struct {
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

func (p *FwServerOper) GetId() string{
    return "1"
}

func (p *FwServerOper) getPath() string{
    
    return "fw/server/"+p.Name+"/oper"
}

func (p *FwServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwServerOper,error) {
logger.Println("FwServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwServerOper
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
