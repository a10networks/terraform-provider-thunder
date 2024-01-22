

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbQuicOper struct {
    
    Oper SlbQuicOperOper `json:"oper"`

}
type DataSlbQuicOper struct {
    DtSlbQuicOper SlbQuicOper `json:"quic"`
}


type SlbQuicOperOper struct {
    SessionList []SlbQuicOperOperSessionList `json:"session-list"`
    TotalSessions int `json:"total-sessions"`
}


type SlbQuicOperOperSessionList struct {
    FwdSource string `json:"fwd-source"`
    FwdSourceCid string `json:"fwd-source-cid"`
    FwdDest string `json:"fwd-dest"`
    FwdDestCid string `json:"fwd-dest-cid"`
    FwdState string `json:"fwd-state"`
    FwdFlags string `json:"fwd-flags"`
    FwdActiveScids []SlbQuicOperOperSessionListFwdActiveScids `json:"fwd-active-scids"`
    FwdAvailableScids []SlbQuicOperOperSessionListFwdAvailableScids `json:"fwd-available-scids"`
    FwdRetiredScids []SlbQuicOperOperSessionListFwdRetiredScids `json:"fwd-retired-scids"`
    FwdActiveDcids []SlbQuicOperOperSessionListFwdActiveDcids `json:"fwd-active-dcids"`
    FwdAvailableDcids []SlbQuicOperOperSessionListFwdAvailableDcids `json:"fwd-available-dcids"`
    FwdRetiredDcids []SlbQuicOperOperSessionListFwdRetiredDcids `json:"fwd-retired-dcids"`
    ReverseTuples []SlbQuicOperOperSessionListReverseTuples `json:"reverse-tuples"`
}


type SlbQuicOperOperSessionListFwdActiveScids struct {
    FwdActiveScid string `json:"fwd-active-scid"`
}


type SlbQuicOperOperSessionListFwdAvailableScids struct {
    FwdAvailableScid string `json:"fwd-available-scid"`
}


type SlbQuicOperOperSessionListFwdRetiredScids struct {
    FwdRetiredScid string `json:"fwd-retired-scid"`
}


type SlbQuicOperOperSessionListFwdActiveDcids struct {
    FwdActiveDcid string `json:"fwd-active-dcid"`
}


type SlbQuicOperOperSessionListFwdAvailableDcids struct {
    FwdAvailableDcid string `json:"fwd-available-dcid"`
}


type SlbQuicOperOperSessionListFwdRetiredDcids struct {
    FwdRetiredDcid string `json:"fwd-retired-dcid"`
}


type SlbQuicOperOperSessionListReverseTuples struct {
    RevSource string `json:"rev-source"`
    RevSourceCid string `json:"rev-source-cid"`
    RevDest string `json:"rev-dest"`
    RevDestCid string `json:"rev-dest-cid"`
    RevState string `json:"rev-state"`
    RevFlags string `json:"rev-flags"`
    RevActiveScids []SlbQuicOperOperSessionListReverseTuplesRevActiveScids `json:"rev-active-scids"`
    RevAvailableScids []SlbQuicOperOperSessionListReverseTuplesRevAvailableScids `json:"rev-available-scids"`
    RevRetiredScids []SlbQuicOperOperSessionListReverseTuplesRevRetiredScids `json:"rev-retired-scids"`
    RevActiveDcids []SlbQuicOperOperSessionListReverseTuplesRevActiveDcids `json:"rev-active-dcids"`
    RevAvailableDcids []SlbQuicOperOperSessionListReverseTuplesRevAvailableDcids `json:"rev-available-dcids"`
    RevRetiredDcids []SlbQuicOperOperSessionListReverseTuplesRevRetiredDcids `json:"rev-retired-dcids"`
}


type SlbQuicOperOperSessionListReverseTuplesRevActiveScids struct {
    RevActiveScid string `json:"rev-active-scid"`
}


type SlbQuicOperOperSessionListReverseTuplesRevAvailableScids struct {
    RevAvailableScid string `json:"rev-available-scid"`
}


type SlbQuicOperOperSessionListReverseTuplesRevRetiredScids struct {
    RevRetiredScid string `json:"rev-retired-scid"`
}


type SlbQuicOperOperSessionListReverseTuplesRevActiveDcids struct {
    RevActiveDcid string `json:"rev-active-dcid"`
}


type SlbQuicOperOperSessionListReverseTuplesRevAvailableDcids struct {
    RevAvailableDcid string `json:"rev-available-dcid"`
}


type SlbQuicOperOperSessionListReverseTuplesRevRetiredDcids struct {
    RevRetiredDcid string `json:"rev-retired-dcid"`
}

func (p *SlbQuicOper) GetId() string{
    return "1"
}

func (p *SlbQuicOper) getPath() string{
    return "slb/quic/oper"
}

func (p *SlbQuicOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbQuicOper,error) {
logger.Println("SlbQuicOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbQuicOper
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
