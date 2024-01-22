

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSyncOper struct {
    
    Oper DdosSyncOperOper `json:"oper"`

}
type DataDdosSyncOper struct {
    DtDdosSyncOper DdosSyncOper `json:"sync"`
}


type DdosSyncOperOper struct {
    IfShowErrorNum int `json:"if-show-error-num"`
    ErrorStr string `json:"error-str"`
    Status string `json:"status"`
    LocalIp string `json:"local-ip"`
    TotalMessageReceived int `json:"total-message-received"`
    TotalMessageSent int `json:"total-message-sent"`
    PeerList []DdosSyncOperOperPeerList `json:"peer-list"`
}


type DdosSyncOperOperPeerList struct {
    PeerIp string `json:"peer-ip"`
    PeerStatus string `json:"peer-status"`
    PeerHeartbeatMissing int `json:"peer-heartbeat-missing"`
    PeerMessageReceived int `json:"peer-message-received"`
    PeerMessageSent int `json:"peer-message-sent"`
}

func (p *DdosSyncOper) GetId() string{
    return "1"
}

func (p *DdosSyncOper) getPath() string{
    return "ddos/sync/oper"
}

func (p *DdosSyncOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosSyncOper,error) {
logger.Println("DdosSyncOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosSyncOper
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
