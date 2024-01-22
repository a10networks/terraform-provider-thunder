

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbCommonCertPinningCandidateListFeedbackOptIn struct {
	Inst struct {

    Daily int `json:"daily"`
    DayTime string `json:"day-time"`
    Enable int `json:"enable"`
    Schedule int `json:"schedule"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`
    WeekDay string `json:"week-day"`
    WeekTime string `json:"week-time"`
    Weekly int `json:"weekly"`

	} `json:"candidate-list-feedback-opt-in"`
}

func (p *SlbCommonCertPinningCandidateListFeedbackOptIn) GetId() string{
    return "1"
}

func (p *SlbCommonCertPinningCandidateListFeedbackOptIn) getPath() string{
    return "slb/common/cert-pinning/candidate-list-feedback-opt-in"
}

func (p *SlbCommonCertPinningCandidateListFeedbackOptIn) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinningCandidateListFeedbackOptIn::Post")
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

func (p *SlbCommonCertPinningCandidateListFeedbackOptIn) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinningCandidateListFeedbackOptIn::Get")
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
func (p *SlbCommonCertPinningCandidateListFeedbackOptIn) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinningCandidateListFeedbackOptIn::Put")
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

func (p *SlbCommonCertPinningCandidateListFeedbackOptIn) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonCertPinningCandidateListFeedbackOptIn::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
