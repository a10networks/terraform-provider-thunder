

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportHealthExternal struct {
	Inst struct {

    Description string `json:"description"`
    Externalfilename string `json:"externalfilename"`
    Overwrite int `json:"overwrite"`
    Password string `json:"password"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"health-external"`
}

func (p *ImportHealthExternal) GetId() string{
    return "1"
}

func (p *ImportHealthExternal) getPath() string{
    return "import/health-external"
}

func (p *ImportHealthExternal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthExternal::Post")
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

func (p *ImportHealthExternal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthExternal::Get")
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
func (p *ImportHealthExternal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthExternal::Put")
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

func (p *ImportHealthExternal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthExternal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
