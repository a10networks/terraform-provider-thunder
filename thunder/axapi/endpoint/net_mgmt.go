

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetMgmt struct {
	Inst struct {

    Snmp NetMgmtSnmp1052 `json:"snmp"`

	} `json:"net-mgmt"`
}


type NetMgmtSnmp1052 struct {
    Engineid NetMgmtSnmpEngineid1053 `json:"engineID"`
    Stats NetMgmtSnmpStats1054 `json:"stats"`
}


type NetMgmtSnmpEngineid1053 struct {
    Uuid string `json:"uuid"`
}


type NetMgmtSnmpStats1054 struct {
    Uuid string `json:"uuid"`
}

func (p *NetMgmt) GetId() string{
    return "1"
}

func (p *NetMgmt) getPath() string{
    return "net-mgmt"
}

func (p *NetMgmt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetMgmt::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *NetMgmt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetMgmt::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *NetMgmt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetMgmt::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *NetMgmt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetMgmt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
