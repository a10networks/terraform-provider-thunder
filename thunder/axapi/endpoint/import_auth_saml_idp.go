

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportAuthSamlIdp struct {
	Inst struct {

    Overwrite int `json:"overwrite"`
    Password string `json:"password"`
    RemoteFile string `json:"remote-file"`
    SamlIdpName string `json:"saml-idp-name"`
    UseMgmtPort int `json:"use-mgmt-port"`
    VerifyXmlSignature int `json:"verify-xml-signature"`

	} `json:"auth-saml-idp"`
}

func (p *ImportAuthSamlIdp) GetId() string{
    return "1"
}

func (p *ImportAuthSamlIdp) getPath() string{
    return "import/auth-saml-idp"
}

func (p *ImportAuthSamlIdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportAuthSamlIdp::Post")
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

func (p *ImportAuthSamlIdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportAuthSamlIdp::Get")
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
func (p *ImportAuthSamlIdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportAuthSamlIdp::Put")
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

func (p *ImportAuthSamlIdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportAuthSamlIdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
