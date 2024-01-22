

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocationFile struct {
	Inst struct {

    ErrorInfo SystemGeolocationFileErrorInfo1573 `json:"error-info"`
    Uuid string `json:"uuid"`

	} `json:"geolocation-file"`
}


type SystemGeolocationFileErrorInfo1573 struct {
    Uuid string `json:"uuid"`
}

func (p *SystemGeolocationFile) GetId() string{
    return "1"
}

func (p *SystemGeolocationFile) getPath() string{
    return "system/geolocation-file"
}

func (p *SystemGeolocationFile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocationFile::Post")
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

func (p *SystemGeolocationFile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocationFile::Get")
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
func (p *SystemGeolocationFile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocationFile::Put")
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

func (p *SystemGeolocationFile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocationFile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
