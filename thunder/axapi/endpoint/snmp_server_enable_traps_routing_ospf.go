

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsRoutingOspf struct {
	Inst struct {

    Ospfifauthfailure int `json:"ospfIfAuthFailure"`
    Ospfifconfigerror int `json:"ospfIfConfigError"`
    Ospfifrxbadpacket int `json:"ospfIfRxBadPacket"`
    Ospfifstatechange int `json:"ospfIfStateChange"`
    Ospflsdbapproachingoverflow int `json:"ospfLsdbApproachingOverflow"`
    Ospflsdboverflow int `json:"ospfLsdbOverflow"`
    Ospfmaxagelsa int `json:"ospfMaxAgeLsa"`
    Ospfnbrstatechange int `json:"ospfNbrStateChange"`
    Ospforiginatelsa int `json:"ospfOriginateLsa"`
    Ospftxretransmit int `json:"ospfTxRetransmit"`
    Ospfvirtifauthfailure int `json:"ospfVirtIfAuthFailure"`
    Ospfvirtifconfigerror int `json:"ospfVirtIfConfigError"`
    Ospfvirtifrxbadpacket int `json:"ospfVirtIfRxBadPacket"`
    Ospfvirtifstatechange int `json:"ospfVirtIfStateChange"`
    Ospfvirtiftxretransmit int `json:"ospfVirtIfTxRetransmit"`
    Ospfvirtnbrstatechange int `json:"ospfVirtNbrStateChange"`
    Uuid string `json:"uuid"`

	} `json:"ospf"`
}

func (p *SnmpServerEnableTrapsRoutingOspf) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsRoutingOspf) getPath() string{
    return "snmp-server/enable/traps/routing/ospf"
}

func (p *SnmpServerEnableTrapsRoutingOspf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingOspf::Post")
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

func (p *SnmpServerEnableTrapsRoutingOspf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingOspf::Get")
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
func (p *SnmpServerEnableTrapsRoutingOspf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingOspf::Put")
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

func (p *SnmpServerEnableTrapsRoutingOspf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingOspf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
