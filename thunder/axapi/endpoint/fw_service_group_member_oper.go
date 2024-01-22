

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FwServiceGroupMemberOper struct {
    
    Name string `json:"name"`
    Oper FwServiceGroupMemberOperOper `json:"oper"`
    Port int `json:"port"`

}
type DataFwServiceGroupMemberOper struct {
    DtFwServiceGroupMemberOper FwServiceGroupMemberOper `json:"member"`
}


type FwServiceGroupMemberOperOper struct {
    State string `json:"state"`
    HmKey int `json:"hm-key"`
    HmIndex int `json:"hm-index"`
    DrsList []FwServiceGroupMemberOperOperDrsList `json:"drs-list"`
    AltList []FwServiceGroupMemberOperOperAltList `json:"alt-list"`
}


type FwServiceGroupMemberOperOperDrsList struct {
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


type FwServiceGroupMemberOperOperAltList struct {
    AltName string `json:"alt-name"`
    AltPort int `json:"alt-port"`
    AltState string `json:"alt-state"`
    AltCurrConn int `json:"alt-curr-conn"`
    AltTotalConn int `json:"alt-total-conn"`
    AltRevPkts int `json:"alt-rev-pkts"`
    AltFwdPkts int `json:"alt-fwd-pkts"`
    AltPeakConn int `json:"alt-peak-conn"`
}

func (p *FwServiceGroupMemberOper) GetId() string{
    return "1"
}

func (p *FwServiceGroupMemberOper) getPath() string{
    
    return "fw/service-group/"+p.Name+"/member/"+p.Name+"+" +strconv.Itoa(p.Port)+"/oper"
}

func (p *FwServiceGroupMemberOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwServiceGroupMemberOper,error) {
logger.Println("FwServiceGroupMemberOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwServiceGroupMemberOper
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
