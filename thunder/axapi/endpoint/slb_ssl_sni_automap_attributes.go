

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslSniAutomapAttributes struct {
	Inst struct {

    SniDeleteFactor int `json:"sni-delete-factor" dval:"50"`
    SniLowerLimit int `json:"sni-lower-limit" dval:"500"`
    SniUpperLimit int `json:"sni-upper-limit" dval:"2000"`
    Uuid string `json:"uuid"`

	} `json:"sni-automap-attributes"`
}

func (p *SlbSslSniAutomapAttributes) GetId() string{
    return "1"
}

func (p *SlbSslSniAutomapAttributes) getPath() string{
    return "slb/ssl/sni-automap-attributes"
}

func (p *SlbSslSniAutomapAttributes) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslSniAutomapAttributes::Post")
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

func (p *SlbSslSniAutomapAttributes) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslSniAutomapAttributes::Get")
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
func (p *SlbSslSniAutomapAttributes) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslSniAutomapAttributes::Put")
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

func (p *SlbSslSniAutomapAttributes) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslSniAutomapAttributes::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
