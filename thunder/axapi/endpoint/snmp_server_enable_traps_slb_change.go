

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsSlbChange struct {
	Inst struct {

    All int `json:"all"`
    ConnectionResourceEvent int `json:"connection-resource-event"`
    ResourceUsageWarning int `json:"resource-usage-warning"`
    Server int `json:"server"`
    ServerPort int `json:"server-port"`
    SslCertChange int `json:"ssl-cert-change"`
    SslCertExpire int `json:"ssl-cert-expire"`
    SystemThreshold int `json:"system-threshold"`
    Uuid string `json:"uuid"`
    Vip int `json:"vip"`
    VipPort int `json:"vip-port"`

	} `json:"slb-change"`
}

func (p *SnmpServerEnableTrapsSlbChange) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsSlbChange) getPath() string{
    return "snmp-server/enable/traps/slb-change"
}

func (p *SnmpServerEnableTrapsSlbChange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlbChange::Post")
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

func (p *SnmpServerEnableTrapsSlbChange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlbChange::Get")
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
func (p *SnmpServerEnableTrapsSlbChange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlbChange::Put")
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

func (p *SnmpServerEnableTrapsSlbChange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSlbChange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
