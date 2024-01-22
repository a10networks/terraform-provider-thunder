

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterLogFile struct {
	Inst struct {

    Name string `json:"name"`
    PerProtocol int `json:"per-protocol"`
    Rotate int `json:"rotate"`
    Size int `json:"size"`
    Uuid string `json:"uuid"`

	} `json:"file"`
}

func (p *RouterLogFile) GetId() string{
    return "1"
}

func (p *RouterLogFile) getPath() string{
    return "router/log/file"
}

func (p *RouterLogFile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLogFile::Post")
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

func (p *RouterLogFile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLogFile::Get")
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
func (p *RouterLogFile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLogFile::Put")
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

func (p *RouterLogFile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLogFile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
