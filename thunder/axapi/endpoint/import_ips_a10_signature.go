

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type ImportIpsA10Signature struct {
	Inst struct {

    A10SignatureFilename string `json:"a10-signature-filename"`
    Password string `json:"password"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"a10-signature"`
}

func (p *ImportIpsA10Signature) GetId() string{
    return "1"
}

func (p *ImportIpsA10Signature) getPath() string{
    return "import/ips/a10-signature"
}

func (p *ImportIpsA10Signature) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsA10Signature::Post")
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

func (p *ImportIpsA10Signature) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsA10Signature::Get")
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
func (p *ImportIpsA10Signature) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsA10Signature::Put")
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

func (p *ImportIpsA10Signature) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportIpsA10Signature::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
