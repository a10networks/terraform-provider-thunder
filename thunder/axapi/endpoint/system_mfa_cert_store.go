

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemMfaCertStore struct {
	Inst struct {

    CertHost string `json:"cert-host"`
    CertStorePath string `json:"cert-store-path"`
    Encrypted string `json:"encrypted"`
    PasswdString string `json:"passwd-string"`
    Protocol string `json:"protocol"`
    Username string `json:"username"`
    Uuid string `json:"uuid"`

	} `json:"mfa-cert-store"`
}

func (p *SystemMfaCertStore) GetId() string{
    return "1"
}

func (p *SystemMfaCertStore) getPath() string{
    return "system/mfa-cert-store"
}

func (p *SystemMfaCertStore) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMfaCertStore::Post")
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

func (p *SystemMfaCertStore) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMfaCertStore::Get")
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
func (p *SystemMfaCertStore) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMfaCertStore::Put")
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

func (p *SystemMfaCertStore) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMfaCertStore::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
