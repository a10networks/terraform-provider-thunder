

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugOspf struct {
	Inst struct {

    All DebugOspfAll334 `json:"all"`
    Bfd DebugOspfBfd335 `json:"bfd"`
    Dumy int `json:"dumy"`
    Events DebugOspfEvents336 `json:"events"`
    Ifsm DebugOspfIfsm337 `json:"ifsm"`
    Lsa DebugOspfLsa338 `json:"lsa"`
    Nfsm DebugOspfNfsm339 `json:"nfsm"`
    Nsm DebugOspfNsm340 `json:"nsm"`
    Packet DebugOspfPacket341 `json:"packet"`
    Route DebugOspfRoute342 `json:"route"`

	} `json:"ospf"`
}


type DebugOspfAll334 struct {
    Dumy int `json:"dumy"`
    Uuid string `json:"uuid"`
}


type DebugOspfBfd335 struct {
    Dumy int `json:"dumy"`
    Uuid string `json:"uuid"`
}


type DebugOspfEvents336 struct {
    Abr int `json:"abr"`
    Asbr int `json:"asbr"`
    Os int `json:"os"`
    Router int `json:"router"`
    Vlink int `json:"vlink"`
    Uuid string `json:"uuid"`
}


type DebugOspfIfsm337 struct {
    Events int `json:"events"`
    Status int `json:"status"`
    Timers int `json:"timers"`
    Uuid string `json:"uuid"`
}


type DebugOspfLsa338 struct {
    Flooding int `json:"flooding"`
    Gererate int `json:"gererate"`
    Install int `json:"install"`
    Maxage int `json:"maxage"`
    Refresh int `json:"refresh"`
    Uuid string `json:"uuid"`
}


type DebugOspfNfsm339 struct {
    Events int `json:"events"`
    Status int `json:"status"`
    Timers int `json:"timers"`
    Uuid string `json:"uuid"`
}


type DebugOspfNsm340 struct {
    Interface int `json:"interface"`
    Redistribute int `json:"redistribute"`
    Uuid string `json:"uuid"`
}


type DebugOspfPacket341 struct {
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


type DebugOspfRoute342 struct {
    Ase int `json:"ase"`
    Ia int `json:"ia"`
    Install int `json:"install"`
    Spf int `json:"spf"`
    Uuid string `json:"uuid"`
}

func (p *DebugOspf) GetId() string{
    return "1"
}

func (p *DebugOspf) getPath() string{
    return "debug/ospf"
}

func (p *DebugOspf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspf::Post")
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

func (p *DebugOspf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspf::Get")
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
func (p *DebugOspf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspf::Put")
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

func (p *DebugOspf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
