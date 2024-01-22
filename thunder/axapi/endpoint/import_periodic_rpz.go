

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type ImportPeriodicRpz struct {
	Inst struct {

    Period int `json:"period"`
    RemoteFile string `json:"remote-file"`
    RemoteFileZoneTransfer string `json:"remote-file-zone-transfer"`
    Rpz string `json:"rpz"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`
    ZoneTransfer string `json:"zone-transfer"`

	} `json:"rpz"`
}

func (p *ImportPeriodicRpz) GetId() string{
    return url.QueryEscape(p.Inst.Rpz)
}

func (p *ImportPeriodicRpz) getPath() string{
    return "import-periodic/rpz"
}

func (p *ImportPeriodicRpz) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicRpz::Post")
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

func (p *ImportPeriodicRpz) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicRpz::Get")
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
func (p *ImportPeriodicRpz) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicRpz::Put")
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

func (p *ImportPeriodicRpz) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicRpz::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
