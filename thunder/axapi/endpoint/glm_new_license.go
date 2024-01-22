

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GlmNewLicense struct {
	Inst struct {

    AccountName string `json:"account-name"`
    Country string `json:"country"`
    ExistingOrg int `json:"existing-org"`
    ExistingUser int `json:"existing-user"`
    FirstName string `json:"first-name"`
    GlmEmail string `json:"glm-email"`
    GlmPassword string `json:"glm-password"`
    LastName string `json:"last-name"`
    Name string `json:"name"`
    NewEmail string `json:"new-email"`
    NewPassword string `json:"new-password"`
    NewUser int `json:"new-user"`
    OrgId int `json:"org-id"`
    Phone string `json:"phone"`
    Type string `json:"type"`

	} `json:"new-license"`
}

func (p *GlmNewLicense) GetId() string{
    return "1"
}

func (p *GlmNewLicense) getPath() string{
    return "glm/new-license"
}

func (p *GlmNewLicense) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GlmNewLicense::Post")
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

func (p *GlmNewLicense) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GlmNewLicense::Get")
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
func (p *GlmNewLicense) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GlmNewLicense::Put")
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

func (p *GlmNewLicense) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GlmNewLicense::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
