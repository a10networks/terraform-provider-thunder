

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCmUpdateFileNameRef struct {
	Inst struct {

    Dest_name string `json:"dest_name"`
    Id1 int `json:"id1"`
    Source_name string `json:"source_name"`

	} `json:"cm-update-file-name-ref"`
}

func (p *SystemCmUpdateFileNameRef) GetId() string{
    return "1"
}

func (p *SystemCmUpdateFileNameRef) getPath() string{
    return "system/cm-update-file-name-ref"
}

func (p *SystemCmUpdateFileNameRef) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCmUpdateFileNameRef::Post")
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

func (p *SystemCmUpdateFileNameRef) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCmUpdateFileNameRef::Get")
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
func (p *SystemCmUpdateFileNameRef) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCmUpdateFileNameRef::Put")
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

func (p *SystemCmUpdateFileNameRef) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCmUpdateFileNameRef::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
