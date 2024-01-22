

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsReload struct {
	Inst struct {

    ClusterDiscovery int `json:"cluster-discovery"`
    Complete int `json:"complete"`
    DbSafe int `json:"db-safe"`
    Device int `json:"device"`
    DisableMerge int `json:"disable-merge"`
    Force int `json:"force"`
    Start int `json:"start"`
    Timeout int `json:"timeout"`

	} `json:"reload"`
}

func (p *VcsReload) GetId() string{
    return "1"
}

func (p *VcsReload) getPath() string{
    return "vcs/reload"
}

func (p *VcsReload) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsReload::Post")
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

func (p *VcsReload) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsReload::Get")
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
func (p *VcsReload) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsReload::Put")
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

func (p *VcsReload) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsReload::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
