

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpMapList struct {
	Inst struct {

    File int `json:"file"`
    MappingList []IpMapListMappingList `json:"mapping-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`
    FileHandle string `json:"file-handle"`

	} `json:"map-list"`
}


type IpMapListMappingList struct {
    LocalStartIp string `json:"local-start-ip"`
    GlobalStartIp string `json:"global-start-ip"`
    Count1 int `json:"count1"`
}

func (p *IpMapList) GetId() string{
    return p.Inst.Name
}

func (p *IpMapList) getPath() string{
    return "ip/map-list"
}

func (p *IpMapList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpMapList::Post")
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

func (p *IpMapList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpMapList::Get")
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
func (p *IpMapList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpMapList::Put")
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

func (p *IpMapList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpMapList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
