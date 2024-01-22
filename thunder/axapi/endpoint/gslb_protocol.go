

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbProtocol struct {
	Inst struct {

    AutoDetect int `json:"auto-detect"`
    DisableNewGslbSync int `json:"disable-new-gslb-sync"`
    EnableList []GslbProtocolEnableList `json:"enable-list"`
    Limit GslbProtocolLimit394 `json:"limit"`
    MsgFormatAcos2x int `json:"msg-format-acos-2x"`
    PingSite string `json:"ping-site"`
    Secure GslbProtocolSecure395 `json:"secure"`
    StatusInterval int `json:"status-interval"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UseMgmtPortForAllPartitions int `json:"use-mgmt-port-for-all-partitions"`
    Uuid string `json:"uuid"`

	} `json:"protocol"`
}


type GslbProtocolEnableList struct {
    Type string `json:"type"`
    Uuid string `json:"uuid"`
}


type GslbProtocolLimit394 struct {
    ArdtQuery int `json:"ardt-query" dval:"200"`
    ArdtResponse int `json:"ardt-response" dval:"1000"`
    ArdtSession int `json:"ardt-session" dval:"32768"`
    ConnResponse int `json:"conn-response"`
    Response int `json:"response" dval:"3600"`
    Message int `json:"message" dval:"10000"`
    Uuid string `json:"uuid"`
}


type GslbProtocolSecure395 struct {
    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`
}

func (p *GslbProtocol) GetId() string{
    return "1"
}

func (p *GslbProtocol) getPath() string{
    return "gslb/protocol"
}

func (p *GslbProtocol) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocol::Post")
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

func (p *GslbProtocol) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocol::Get")
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
func (p *GslbProtocol) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocol::Put")
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

func (p *GslbProtocol) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocol::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
