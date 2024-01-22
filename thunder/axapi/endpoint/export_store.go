

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ExportStore struct {
	Inst struct {

    Create int `json:"create"`
    Delete int `json:"delete"`
    Name string `json:"name"`
    RemoteFile string `json:"remote-file"`

	} `json:"store"`
}

func (p *ExportStore) GetId() string{
    return "1"
}

func (p *ExportStore) getPath() string{
    return "export/store"
}

func (p *ExportStore) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ExportStore::Post")
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

func (p *ExportStore) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ExportStore::Get")
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
func (p *ExportStore) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ExportStore::Put")
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

func (p *ExportStore) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ExportStore::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
