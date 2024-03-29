

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetMgmtSnmpEngineidOper struct {
    
    Oper NetMgmtSnmpEngineidOperOper `json:"oper"`

}
type DataNetMgmtSnmpEngineidOper struct {
    DtNetMgmtSnmpEngineidOper NetMgmtSnmpEngineidOper `json:"engineID"`
}


type NetMgmtSnmpEngineidOperOper struct {
    Engineid string `json:"engineid"`
}

func (p *NetMgmtSnmpEngineidOper) GetId() string{
    return "1"
}

func (p *NetMgmtSnmpEngineidOper) getPath() string{
    return "net-mgmt/snmp/engineID/oper"
}

func (p *NetMgmtSnmpEngineidOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetMgmtSnmpEngineidOper,error) {
logger.Println("NetMgmtSnmpEngineidOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetMgmtSnmpEngineidOper
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
