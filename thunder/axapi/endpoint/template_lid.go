

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type TemplateLid struct {
	Inst struct {

    DdosProtectionFactor int `json:"ddos-protection-factor"`
    DownlinkPps int `json:"downlink-pps"`
    DownlinkThroughput int `json:"downlink-throughput"`
    LidNumber int `json:"lid-number"`
    LimitCps int `json:"limit-cps"`
    LimitRate string `json:"limit-rate"`
    RespondToUserMac int `json:"respond-to-user-mac"`
    SrcIp TemplateLidSrcIp `json:"src-ip"`
    TotalPps int `json:"total-pps"`
    TotalThroughput int `json:"total-throughput"`
    UplinkPps int `json:"uplink-pps"`
    UplinkThroughput int `json:"uplink-throughput"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"lid"`
}


type TemplateLidSrcIp struct {
    ConcurrentSessions int `json:"concurrent-sessions"`
    Log int `json:"log"`
    PrefixLength int `json:"prefix-length"`
    BurstsizeDownlinkPps int `json:"burstsize-downlink-pps"`
    BurstsizeUplinkPps int `json:"burstsize-uplink-pps"`
    BurstsizeTotalPps int `json:"burstsize-total-pps"`
    BurstsizeDownlinkThroughput int `json:"burstsize-downlink-throughput"`
    BurstsizeUplinkThroughput int `json:"burstsize-uplink-throughput"`
    BurstsizeTotalThroughput int `json:"burstsize-total-throughput"`
    BurstsizeCps int `json:"burstsize-cps"`
}

func (p *TemplateLid) GetId() string{
    return strconv.Itoa(p.Inst.LidNumber)
}

func (p *TemplateLid) getPath() string{
    return "template/lid"
}

func (p *TemplateLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLid::Post")
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

func (p *TemplateLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLid::Get")
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
func (p *TemplateLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLid::Put")
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

func (p *TemplateLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
