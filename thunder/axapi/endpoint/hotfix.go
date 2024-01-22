

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Hotfix struct {
	Inst struct {

    Apply HotfixApply438 `json:"apply"`
    Revoke int `json:"revoke"`

	} `json:"hotfix"`
}


type HotfixApply438 struct {
    UseMgmtPort int `json:"use-mgmt-port"`
    SourceIpAddress string `json:"source-ip-address"`
    FileUrl string `json:"file-url"`
    ImageFile string `json:"image-file"`
}

func (p *Hotfix) GetId() string{
    return "1"
}

func (p *Hotfix) getPath() string{
    return "hotfix"
}

func (p *Hotfix) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Hotfix::Post")
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

func (p *Hotfix) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Hotfix::Get")
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
func (p *Hotfix) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Hotfix::Put")
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

func (p *Hotfix) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Hotfix::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
