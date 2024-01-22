

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPingSweepDetection struct {
	Inst struct {

    Events int `json:"events" dval:"10"`
    Interval int `json:"interval" dval:"60"`
    Uuid string `json:"uuid"`
    V4List string `json:"v4-list"`
    V6List string `json:"v6-list"`

	} `json:"ping-sweep-detection"`
}

func (p *VisibilityPingSweepDetection) GetId() string{
    return "1"
}

func (p *VisibilityPingSweepDetection) getPath() string{
    return "visibility/ping-sweep-detection"
}

func (p *VisibilityPingSweepDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPingSweepDetection::Post")
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

func (p *VisibilityPingSweepDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPingSweepDetection::Get")
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
func (p *VisibilityPingSweepDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPingSweepDetection::Put")
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

func (p *VisibilityPingSweepDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPingSweepDetection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
