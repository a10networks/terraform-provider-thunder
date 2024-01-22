

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAVridLead struct {
	Inst struct {

    Partition VrrpAVridLeadPartition `json:"partition"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VridLeadStr string `json:"vrid-lead-str"`

	} `json:"vrid-lead"`
}


type VrrpAVridLeadPartition struct {
    Partition int `json:"partition"`
    SharedCfg VrrpAVridLeadPartitionSharedCfg `json:"shared-cfg"`
    NameCfg VrrpAVridLeadPartitionNameCfg `json:"name-cfg"`
}


type VrrpAVridLeadPartitionSharedCfg struct {
    Shared int `json:"shared"`
    Vrid int `json:"vrid"`
    VridValue int `json:"vrid-value"`
}


type VrrpAVridLeadPartitionNameCfg struct {
    Name string `json:"name"`
    Vrid int `json:"vrid"`
    VridValue int `json:"vrid-value"`
}

func (p *VrrpAVridLead) GetId() string{
    return p.Inst.VridLeadStr
}

func (p *VrrpAVridLead) getPath() string{
    return "vrrp-a/vrid-lead"
}

func (p *VrrpAVridLead) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridLead::Post")
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

func (p *VrrpAVridLead) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridLead::Get")
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
func (p *VrrpAVridLead) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridLead::Put")
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

func (p *VrrpAVridLead) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridLead::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
