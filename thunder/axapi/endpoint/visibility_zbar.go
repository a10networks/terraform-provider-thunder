

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityZbar struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    Dest VisibilityZbarDest3138 `json:"dest"`
    Truples VisibilityZbarTruples3140 `json:"truples"`
    Uuid string `json:"uuid"`

	} `json:"zbar"`
}


type VisibilityZbarDest3138 struct {
    Uuid string `json:"uuid"`
    BadSources VisibilityZbarDestBadSources3139 `json:"bad-sources"`
}


type VisibilityZbarDestBadSources3139 struct {
    Uuid string `json:"uuid"`
}


type VisibilityZbarTruples3140 struct {
    Uuid string `json:"uuid"`
}

func (p *VisibilityZbar) GetId() string{
    return "1"
}

func (p *VisibilityZbar) getPath() string{
    return "visibility/zbar"
}

func (p *VisibilityZbar) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityZbar::Post")
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

func (p *VisibilityZbar) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityZbar::Get")
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
func (p *VisibilityZbar) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityZbar::Put")
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

func (p *VisibilityZbar) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityZbar::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
