

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsRoutingBgp struct {
	Inst struct {

    Ax SnmpServerEnableTrapsRoutingBgpAx1478 `json:"ax"`
    Bgpbackwardtransnotification int `json:"bgpBackwardTransNotification"`
    Bgpestablishednotification int `json:"bgpEstablishedNotification"`
    Uuid string `json:"uuid"`

	} `json:"bgp"`
}


type SnmpServerEnableTrapsRoutingBgpAx1478 struct {
    Bgpestablishednotification int `json:"bgpEstablishedNotification"`
    Bgpbackwardtransnotification int `json:"bgpBackwardTransNotification"`
    Bgpprefixthresholdexceedednotification int `json:"bgpPrefixThresholdExceededNotification"`
    Bgpprefixthresholdclearnotification int `json:"bgpPrefixThresholdClearNotification"`
    Uuid string `json:"uuid"`
}

func (p *SnmpServerEnableTrapsRoutingBgp) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsRoutingBgp) getPath() string{
    return "snmp-server/enable/traps/routing/bgp"
}

func (p *SnmpServerEnableTrapsRoutingBgp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgp::Post")
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

func (p *SnmpServerEnableTrapsRoutingBgp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgp::Get")
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
func (p *SnmpServerEnableTrapsRoutingBgp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgp::Put")
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

func (p *SnmpServerEnableTrapsRoutingBgp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingBgp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
