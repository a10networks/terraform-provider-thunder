

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetMgmtSnmpStatsOper struct {
    
    Oper NetMgmtSnmpStatsOperOper `json:"oper"`

}
type DataNetMgmtSnmpStatsOper struct {
    DtNetMgmtSnmpStatsOper NetMgmtSnmpStatsOper `json:"stats"`
}


type NetMgmtSnmpStatsOperOper struct {
    BadVersion int `json:"bad-version"`
    UnknownCommunity int `json:"unknown-community"`
    IllegalOperation int `json:"illegal-operation"`
    EncodingError int `json:"encoding-error"`
    UnknownSecurityModels int `json:"unknown-security-models"`
    InvalidId int `json:"invalid-id"`
    InputPackets int `json:"input-packets"`
    NumberOfReqVar int `json:"number-of-req-var"`
    GetReqPdu int `json:"get-req-pdu"`
    GetNextPdu int `json:"get-next-pdu"`
    PacketDrop int `json:"packet-drop"`
    TooBigErrors int `json:"too-big-errors"`
    NoSuchNameErrors int `json:"no-such-name-errors"`
    BadValuesErrors int `json:"bad-values-errors"`
    GeneralErrors int `json:"general-errors"`
    OutputPackets int `json:"output-packets"`
    GetRespPdu int `json:"get-resp-pdu"`
    OutputTraps int `json:"output-traps"`
}

func (p *NetMgmtSnmpStatsOper) GetId() string{
    return "1"
}

func (p *NetMgmtSnmpStatsOper) getPath() string{
    return "net-mgmt/snmp/stats/oper"
}

func (p *NetMgmtSnmpStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetMgmtSnmpStatsOper,error) {
logger.Println("NetMgmtSnmpStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetMgmtSnmpStatsOper
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
