

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportPeriodicSslCrl struct {
	Inst struct {

    Period int `json:"period"`
    RemoteFile string `json:"remote-file"`
    SslCrl string `json:"ssl-crl"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"ssl-crl"`
}

func (p *ImportPeriodicSslCrl) GetId() string{
    return p.Inst.SslCrl
}

func (p *ImportPeriodicSslCrl) getPath() string{
    return "import-periodic/ssl-crl"
}

func (p *ImportPeriodicSslCrl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicSslCrl::Post")
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

func (p *ImportPeriodicSslCrl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicSslCrl::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *ImportPeriodicSslCrl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicSslCrl::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *ImportPeriodicSslCrl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicSslCrl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
