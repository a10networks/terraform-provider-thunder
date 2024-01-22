

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Bgp struct {
	Inst struct {

    Asdot int `json:"asdot"`
    AsdotPlus int `json:"asdot-plus"`
    DisableAdvertisement BgpDisableAdvertisement `json:"disable-advertisement"`
    ExtendedAsnCap int `json:"extended-asn-cap"`
    NexthopTrigger BgpNexthopTrigger `json:"nexthop-trigger"`
    Uuid string `json:"uuid"`

	} `json:"bgp"`
}


type BgpDisableAdvertisement struct {
    OnBoot int `json:"on-boot"`
}


type BgpNexthopTrigger struct {
    Enable int `json:"enable"`
    Delay int `json:"delay"`
}

func (p *Bgp) GetId() string{
    return "1"
}

func (p *Bgp) getPath() string{
    return "bgp"
}

func (p *Bgp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Bgp::Post")
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

func (p *Bgp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Bgp::Get")
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
func (p *Bgp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Bgp::Put")
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

func (p *Bgp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Bgp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
