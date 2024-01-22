

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopbackIpOspfOspfIp struct {
	Inst struct {

    Authentication int `json:"authentication"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    IpAddr string `json:"ip-addr"`
    MessageDigestCfg []InterfaceLoopbackIpOspfOspfIpMessageDigestCfg `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Out int `json:"out"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
    Value string `json:"value"`
    Ifnum string 

	} `json:"ospf-ip"`
}


type InterfaceLoopbackIpOspfOspfIpMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}

func (p *InterfaceLoopbackIpOspfOspfIp) GetId() string{
    return p.Inst.IpAddr
}

func (p *InterfaceLoopbackIpOspfOspfIp) getPath() string{
    return "interface/loopback/" +p.Inst.Ifnum + "/ip/ospf/ospf-ip"
}

func (p *InterfaceLoopbackIpOspfOspfIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpOspfOspfIp::Post")
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

func (p *InterfaceLoopbackIpOspfOspfIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpOspfOspfIp::Get")
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
func (p *InterfaceLoopbackIpOspfOspfIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpOspfOspfIp::Put")
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

func (p *InterfaceLoopbackIpOspfOspfIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpOspfOspfIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
