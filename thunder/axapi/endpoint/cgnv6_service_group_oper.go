

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ServiceGroupOper struct {
    
    MemberList []Cgnv6ServiceGroupOperMemberList `json:"member-list"`
    Name string `json:"name"`
    Oper Cgnv6ServiceGroupOperOper `json:"oper"`

}
type DataCgnv6ServiceGroupOper struct {
    DtCgnv6ServiceGroupOper Cgnv6ServiceGroupOper `json:"service-group"`
}


type Cgnv6ServiceGroupOperMemberList struct {
    Name string `json:"name"`
    Port int `json:"port"`
    Oper Cgnv6ServiceGroupOperMemberListOper `json:"oper"`
}


type Cgnv6ServiceGroupOperMemberListOper struct {
    State string `json:"state"`
    HmKey int `json:"hm-key"`
    HmIndex int `json:"hm-index"`
    DrsList []Cgnv6ServiceGroupOperMemberListOperDrsList `json:"drs-list"`
    AltList []Cgnv6ServiceGroupOperMemberListOperAltList `json:"alt-list"`
}


type Cgnv6ServiceGroupOperMemberListOperDrsList struct {
    DrsName string `json:"drs-name"`
    DrsState string `json:"drs-state"`
    DrsHmKey int `json:"drs-hm-key"`
    DrsHmIndex int `json:"drs-hm-index"`
    DrsPort int `json:"drs-port"`
    DrsPriority int `json:"drs-priority"`
    DrsCurrConn int `json:"drs-curr-conn"`
    DrsPersConn int `json:"drs-pers-conn"`
    DrsTotalConn int `json:"drs-total-conn"`
    DrsCurrReq int `json:"drs-curr-req"`
    DrsTotalReq int `json:"drs-total-req"`
    DrsTotalReqSucc int `json:"drs-total-req-succ"`
    DrsRevPkts int `json:"drs-rev-pkts"`
    DrsFwdPkts int `json:"drs-fwd-pkts"`
    DrsRevBts int `json:"drs-rev-bts"`
    DrsFwdBts int `json:"drs-fwd-bts"`
    DrsPeakConn int `json:"drs-peak-conn"`
    DrsRspTime int `json:"drs-rsp-time"`
    DrsFrspTime int `json:"drs-frsp-time"`
    DrsSrspTime int `json:"drs-srsp-time"`
}


type Cgnv6ServiceGroupOperMemberListOperAltList struct {
    AltName string `json:"alt-name"`
    AltPort int `json:"alt-port"`
    AltState string `json:"alt-state"`
    AltCurrConn int `json:"alt-curr-conn"`
    AltTotalConn int `json:"alt-total-conn"`
    AltRevPkts int `json:"alt-rev-pkts"`
    AltFwdPkts int `json:"alt-fwd-pkts"`
    AltPeakConn int `json:"alt-peak-conn"`
}


type Cgnv6ServiceGroupOperOper struct {
    State string `json:"state"`
    Servers_up int `json:"servers_up"`
    Servers_down int `json:"servers_down"`
    Servers_disable int `json:"servers_disable"`
    Servers_total int `json:"servers_total"`
    Stateless_current_rate int `json:"stateless_current_rate"`
    Stateless_current_usage int `json:"stateless_current_usage"`
    Stateless_state int `json:"stateless_state"`
    Stateless_type int `json:"stateless_type"`
    Hm_dsr_enable_all_vip int `json:"hm_dsr_enable_all_vip"`
    Pri_affinity_priority int `json:"pri_affinity_priority"`
    Filter string `json:"filter"`
    SgmList []Cgnv6ServiceGroupOperOperSgmList `json:"sgm-list"`
}


type Cgnv6ServiceGroupOperOperSgmList struct {
    SgmName string `json:"sgm-name"`
    SgmPort int `json:"sgm-port"`
}

func (p *Cgnv6ServiceGroupOper) GetId() string{
    return "1"
}

func (p *Cgnv6ServiceGroupOper) getPath() string{
    
    return "cgnv6/service-group/"+p.Name+"/oper"
}

func (p *Cgnv6ServiceGroupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ServiceGroupOper,error) {
logger.Println("Cgnv6ServiceGroupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6ServiceGroupOper
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
