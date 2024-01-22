

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VcsVbladesStat struct {
	Inst struct {

    SamplingEnable []VcsVbladesStatSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    VbladeId int `json:"vblade-id"`

	} `json:"stat"`
}


type VcsVbladesStatSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *VcsVbladesStat) GetId() string{
    return strconv.Itoa(p.Inst.VbladeId)
}

func (p *VcsVbladesStat) getPath() string{
    return "vcs-vblades/stat"
}

func (p *VcsVbladesStat) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVbladesStat::Post")
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

func (p *VcsVbladesStat) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVbladesStat::Get")
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
func (p *VcsVbladesStat) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVbladesStat::Put")
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

func (p *VcsVbladesStat) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVbladesStat::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
