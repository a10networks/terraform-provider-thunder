

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcEntryAppType struct {
	Inst struct {

    Protocol string `json:"protocol"`
    Template DdosSrcEntryAppTypeTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    SrcEntryName string 

	} `json:"app-type"`
}


type DdosSrcEntryAppTypeTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}

func (p *DdosSrcEntryAppType) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosSrcEntryAppType) getPath() string{
    return "ddos/src/entry/" +p.Inst.SrcEntryName + "/app-type"
}

func (p *DdosSrcEntryAppType) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryAppType::Post")
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

func (p *DdosSrcEntryAppType) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryAppType::Get")
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
func (p *DdosSrcEntryAppType) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryAppType::Put")
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

func (p *DdosSrcEntryAppType) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryAppType::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
