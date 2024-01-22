

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsRoutingBgpAx struct {
	Inst struct {

    Bgpbackwardtransnotification int `json:"bgpBackwardTransNotification"`
    Bgpestablishednotification int `json:"bgpEstablishedNotification"`
    Bgpprefixthresholdclearnotification int `json:"bgpPrefixThresholdClearNotification"`
    Bgpprefixthresholdexceedednotification int `json:"bgpPrefixThresholdExceededNotification"`
    Uuid string `json:"uuid"`

	} `json:"ax"`
}

func (p *SnmpServerEnableTrapsRoutingBgpAx) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsRoutingBgpAx) getPath() string{
    return "snmp-server/enable/traps/routing/bgp/ax"
}

func (p *SnmpServerEnableTrapsRoutingBgpAx) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgpAx::Post")
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

func (p *SnmpServerEnableTrapsRoutingBgpAx) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgpAx::Get")
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
func (p *SnmpServerEnableTrapsRoutingBgpAx) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgpAx::Put")
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

func (p *SnmpServerEnableTrapsRoutingBgpAx) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgpAx::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
