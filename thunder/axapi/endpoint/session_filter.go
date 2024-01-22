

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SessionFilter struct {
	Inst struct {

    FilterCfg SessionFilterFilterCfg `json:"filter-cfg"`
    Name string `json:"name"`
    Set int `json:"set"`
    Uuid string `json:"uuid"`

	} `json:"session-filter"`
}


type SessionFilterFilterCfg struct {
    SessionType string `json:"session-type"`
    SourceAddr string `json:"source-addr"`
    SourceMask string `json:"source-mask"`
    SourcePort int `json:"source-port"`
    DestAddr string `json:"dest-addr"`
    DestMask string `json:"dest-mask"`
    Dport2 int `json:"dport2"`
    App string `json:"app"`
    AppCategory string `json:"app-category"`
}

func (p *SessionFilter) GetId() string{
    return p.Inst.Name
}

func (p *SessionFilter) getPath() string{
    return "session-filter"
}

func (p *SessionFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilter::Post")
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

func (p *SessionFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilter::Get")
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
func (p *SessionFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilter::Put")
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

func (p *SessionFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
