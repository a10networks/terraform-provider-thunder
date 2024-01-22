

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryWebReputation struct {
	Inst struct {

    BypassedUrls WebCategoryWebReputationBypassedUrls3658 `json:"bypassed-urls"`
    InterceptedUrls WebCategoryWebReputationInterceptedUrls3659 `json:"intercepted-urls"`
    Url WebCategoryWebReputationUrl3660 `json:"url"`
    Uuid string `json:"uuid"`

	} `json:"web-reputation"`
}


type WebCategoryWebReputationBypassedUrls3658 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryWebReputationInterceptedUrls3659 struct {
    Uuid string `json:"uuid"`
}


type WebCategoryWebReputationUrl3660 struct {
    Uuid string `json:"uuid"`
}

func (p *WebCategoryWebReputation) GetId() string{
    return "1"
}

func (p *WebCategoryWebReputation) getPath() string{
    return "web-category/web-reputation"
}

func (p *WebCategoryWebReputation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryWebReputation::Post")
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

func (p *WebCategoryWebReputation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryWebReputation::Get")
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
func (p *WebCategoryWebReputation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryWebReputation::Put")
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

func (p *WebCategoryWebReputation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryWebReputation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
