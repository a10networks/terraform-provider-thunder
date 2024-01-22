

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportPeriodicAuthJwks struct {
	Inst struct {

    AuthJwks string `json:"auth-jwks"`
    Period int `json:"period"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"auth-jwks"`
}

func (p *ImportPeriodicAuthJwks) GetId() string{
    return p.Inst.AuthJwks
}

func (p *ImportPeriodicAuthJwks) getPath() string{
    return "import-periodic/auth-jwks"
}

func (p *ImportPeriodicAuthJwks) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicAuthJwks::Post")
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

func (p *ImportPeriodicAuthJwks) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicAuthJwks::Get")
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
func (p *ImportPeriodicAuthJwks) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicAuthJwks::Put")
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

func (p *ImportPeriodicAuthJwks) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicAuthJwks::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
