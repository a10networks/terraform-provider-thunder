

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportToDevice struct {
	Inst struct {

    Device int `json:"device"`
    GlmCert string `json:"glm-cert"`
    GlmLicense string `json:"glm-license"`
    Overwrite int `json:"overwrite"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`
    WebCategoryLicense string `json:"web-category-license"`

	} `json:"to-device"`
}

func (p *ImportToDevice) GetId() string{
    return "1"
}

func (p *ImportToDevice) getPath() string{
    return "import/to-device"
}

func (p *ImportToDevice) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportToDevice::Post")
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

func (p *ImportToDevice) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportToDevice::Get")
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
func (p *ImportToDevice) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportToDevice::Put")
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

func (p *ImportToDevice) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportToDevice::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
