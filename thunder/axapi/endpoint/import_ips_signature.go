

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type ImportIpsSignature struct {
	Inst struct {

    Password string `json:"password"`
    RemoteFile string `json:"remote-file"`
    SignatureFilename string `json:"signature-filename"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"signature"`
}

func (p *ImportIpsSignature) GetId() string{
    return "1"
}

func (p *ImportIpsSignature) getPath() string{
    return "import/ips/signature"
}

func (p *ImportIpsSignature) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsSignature::Post")
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

func (p *ImportIpsSignature) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsSignature::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) != 0 {
        err = json.Unmarshal(axResp, &p)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
                     }
            }
   }
    return err
}
func (p *ImportIpsSignature) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsSignature::Put")
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

func (p *ImportIpsSignature) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsSignature::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
