

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Glm struct {
	Inst struct {

    AllocateBandwidth int `json:"allocate-bandwidth"`
    ApplianceName string `json:"appliance-name"`
    Burst int `json:"burst"`
    CheckExpiration int `json:"check-expiration"`
    CreateLicenseRequest GlmCreateLicenseRequest380 `json:"create-license-request"`
    EnableRequests int `json:"enable-requests"`
    Enterprise string `json:"enterprise"`
    EnterpriseHaHostList []GlmEnterpriseHaHostList `json:"enterprise-ha-host-list"`
    EnterpriseRequestType string `json:"enterprise-request-type"`
    Host string `json:"host"`
    Interval int `json:"interval"`
    NewLicense GlmNewLicense381 `json:"new-license"`
    Port int `json:"port"`
    ProxyServer GlmProxyServer382 `json:"proxy-server"`
    Send GlmSend383 `json:"send"`
    ThunderCapacityLicense int `json:"thunder-capacity-license"`
    Token string `json:"token"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"glm"`
}


type GlmCreateLicenseRequest380 struct {
    CreateLicenseRequest int `json:"create-license-request"`
    Uuid string `json:"uuid"`
}


type GlmEnterpriseHaHostList struct {
    HostEntry string `json:"host-entry"`
    Uuid string `json:"uuid"`
}


type GlmNewLicense381 struct {
    ExistingOrg int `json:"existing-org"`
    OrgId int `json:"org-id"`
    ExistingUser int `json:"existing-user"`
    GlmEmail string `json:"glm-email"`
    GlmPassword string `json:"glm-password"`
    NewUser int `json:"new-user"`
    NewEmail string `json:"new-email"`
    NewPassword string `json:"new-password"`
    AccountName string `json:"account-name"`
    FirstName string `json:"first-name"`
    LastName string `json:"last-name"`
    Country string `json:"country"`
    Phone string `json:"phone"`
    Name string `json:"name"`
    Type string `json:"type"`
}


type GlmProxyServer382 struct {
    Host string `json:"host"`
    Port int `json:"port"`
    Username string `json:"username"`
    Password int `json:"password"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Uuid string `json:"uuid"`
}


type GlmSend383 struct {
    LicenseRequest int `json:"license-request"`
    HaStatus int `json:"ha-status"`
}

func (p *Glm) GetId() string{
    return "1"
}

func (p *Glm) getPath() string{
    return "glm"
}

func (p *Glm) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Glm::Post")
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

func (p *Glm) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Glm::Get")
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
func (p *Glm) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Glm::Put")
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

func (p *Glm) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Glm::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
