

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterLog struct {
	Inst struct {

    File RouterLogFile1296 `json:"file"`
    LogBuffer int `json:"log-buffer" dval:"1"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`
    FileHandle string `json:"file-handle"`

	} `json:"log"`
}


type RouterLogFile1296 struct {
    PerProtocol int `json:"per-protocol"`
    Name string `json:"name"`
    Rotate int `json:"rotate"`
    Size int `json:"size"`
    Uuid string `json:"uuid"`
}

func (p *RouterLog) GetId() string{
    return "1"
}

func (p *RouterLog) getPath() string{
    return "router/log"
}

func (p *RouterLog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLog::Post")
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

func (p *RouterLog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLog::Get")
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
func (p *RouterLog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLog::Put")
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

func (p *RouterLog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterLog::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
