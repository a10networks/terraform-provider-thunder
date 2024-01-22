

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportHealthPostfile struct {
	Inst struct {

    Overwrite int `json:"overwrite"`
    Password string `json:"password"`
    Postfilename string `json:"postfilename"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"health-postfile"`
}

func (p *ImportHealthPostfile) GetId() string{
    return "1"
}

func (p *ImportHealthPostfile) getPath() string{
    return "import/health-postfile"
}

func (p *ImportHealthPostfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthPostfile::Post")
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

func (p *ImportHealthPostfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthPostfile::Get")
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
func (p *ImportHealthPostfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthPostfile::Put")
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

func (p *ImportHealthPostfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportHealthPostfile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
