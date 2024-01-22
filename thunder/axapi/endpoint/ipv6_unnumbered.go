

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6Unnumbered struct {
	Inst struct {

    UseSourceAcl string `json:"use-source-acl"`
    UseSourceIpv6 Ipv6UnnumberedUseSourceIpv61045 `json:"use-source-ipv6"`
    Uuid string `json:"uuid"`

	} `json:"unnumbered"`
}


type Ipv6UnnumberedUseSourceIpv61045 struct {
    UpdateSourceIpv6 string `json:"update-source-ipv6"`
    Uuid string `json:"uuid"`
}

func (p *Ipv6Unnumbered) GetId() string{
    return "1"
}

func (p *Ipv6Unnumbered) getPath() string{
    return "ipv6/unnumbered"
}

func (p *Ipv6Unnumbered) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Unnumbered::Post")
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

func (p *Ipv6Unnumbered) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Unnumbered::Get")
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
func (p *Ipv6Unnumbered) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Unnumbered::Put")
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

func (p *Ipv6Unnumbered) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Unnumbered::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
