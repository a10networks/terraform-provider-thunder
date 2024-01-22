

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoTcpUdp struct {
	Inst struct {

    Deny int `json:"deny"`
    DropFragPkt int `json:"drop-frag-pkt"`
    GlidCfg DdosDstZoneIpProtoProtoTcpUdpGlidCfg `json:"glid-cfg"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    IpFilteringPolicyOper DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOper211 `json:"ip-filtering-policy-oper"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"proto-tcp-udp"`
}


type DdosDstZoneIpProtoProtoTcpUdpGlidCfg struct {
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    PerAddrGlid string `json:"per-addr-glid"`
    ActionList string `json:"action-list"`
}


type DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOper211 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDstZoneIpProtoProtoTcpUdp) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstZoneIpProtoProtoTcpUdp) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-tcp-udp"
}

func (p *DdosDstZoneIpProtoProtoTcpUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoTcpUdp::Post")
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

func (p *DdosDstZoneIpProtoProtoTcpUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoTcpUdp::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *DdosDstZoneIpProtoProtoTcpUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoTcpUdp::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *DdosDstZoneIpProtoProtoTcpUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoTcpUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
