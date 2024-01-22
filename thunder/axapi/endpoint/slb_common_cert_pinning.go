

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbCommonCertPinning struct {
	Inst struct {

    CandidateListFeedbackOptIn SlbCommonCertPinningCandidateListFeedbackOptIn1406 `json:"candidate-list-feedback-opt-in"`
    Ttl int `json:"ttl" dval:"144"`
    Uuid string `json:"uuid"`

	} `json:"cert-pinning"`
}


type SlbCommonCertPinningCandidateListFeedbackOptIn1406 struct {
    Enable int `json:"enable"`
    Schedule int `json:"schedule"`
    Weekly int `json:"weekly"`
    WeekDay string `json:"week-day"`
    WeekTime string `json:"week-time"`
    Daily int `json:"daily"`
    DayTime string `json:"day-time"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`
}

func (p *SlbCommonCertPinning) GetId() string{
    return "1"
}

func (p *SlbCommonCertPinning) getPath() string{
    return "slb/common/cert-pinning"
}

func (p *SlbCommonCertPinning) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinning::Post")
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

func (p *SlbCommonCertPinning) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinning::Get")
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
func (p *SlbCommonCertPinning) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinning::Put")
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

func (p *SlbCommonCertPinning) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinning::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
