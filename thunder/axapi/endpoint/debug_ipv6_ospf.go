

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugIpv6Ospf struct {
	Inst struct {

    All DebugIpv6OspfAll325 `json:"all"`
    Bfd DebugIpv6OspfBfd326 `json:"bfd"`
    Dumy int `json:"dumy"`
    Events DebugIpv6OspfEvents327 `json:"events"`
    Ifsm DebugIpv6OspfIfsm328 `json:"ifsm"`
    Lsa DebugIpv6OspfLsa329 `json:"lsa"`
    Nfsm DebugIpv6OspfNfsm330 `json:"nfsm"`
    Nsm DebugIpv6OspfNsm331 `json:"nsm"`
    Packet DebugIpv6OspfPacket332 `json:"packet"`
    Route DebugIpv6OspfRoute333 `json:"route"`

	} `json:"ospf"`
}


type DebugIpv6OspfAll325 struct {
    Dumy int `json:"dumy"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfBfd326 struct {
    Dumy int `json:"dumy"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfEvents327 struct {
    Abr int `json:"abr"`
    Asbr int `json:"asbr"`
    Os int `json:"os"`
    Router int `json:"router"`
    Vlink int `json:"vlink"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfIfsm328 struct {
    Events int `json:"events"`
    Status int `json:"status"`
    Timers int `json:"timers"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfLsa329 struct {
    Flooding int `json:"flooding"`
    Gererate int `json:"gererate"`
    Install int `json:"install"`
    Maxage int `json:"maxage"`
    Refresh int `json:"refresh"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfNfsm330 struct {
    Events int `json:"events"`
    Status int `json:"status"`
    Timers int `json:"timers"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfNsm331 struct {
    Interface int `json:"interface"`
    Redistribute int `json:"redistribute"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfPacket332 struct {
    Dd int `json:"dd"`
    Detail int `json:"detail"`
    Hello int `json:"hello"`
    LsAck int `json:"ls-ack"`
    LsRequest int `json:"ls-request"`
    LsUpdate int `json:"ls-update"`
    Recv int `json:"recv"`
    Send int `json:"send"`
    Uuid string `json:"uuid"`
}


type DebugIpv6OspfRoute333 struct {
    Ase int `json:"ase"`
    Ia int `json:"ia"`
    Install int `json:"install"`
    Spf int `json:"spf"`
    Uuid string `json:"uuid"`
}

func (p *DebugIpv6Ospf) GetId() string{
    return "1"
}

func (p *DebugIpv6Ospf) getPath() string{
    return "debug/ipv6/ospf"
}

func (p *DebugIpv6Ospf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6Ospf::Post")
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

func (p *DebugIpv6Ospf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6Ospf::Get")
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
func (p *DebugIpv6Ospf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6Ospf::Put")
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

func (p *DebugIpv6Ospf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6Ospf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
