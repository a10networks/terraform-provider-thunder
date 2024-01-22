

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PartitionSharedVlan struct {
	Inst struct {

    AllowableIpRange []PartitionSharedVlanAllowableIpRange `json:"allowable-ip-range"`
    AllowableIpv6Range []PartitionSharedVlanAllowableIpv6Range `json:"allowable-ipv6-range"`
    MgmtFloatingIpAddress string `json:"mgmt-floating-ip-address"`
    Uuid string `json:"uuid"`
    Vlan int `json:"vlan"`
    Vrid int `json:"vrid"`
    PartitionName string 

	} `json:"shared-vlan"`
}


type PartitionSharedVlanAllowableIpRange struct {
}


type PartitionSharedVlanAllowableIpv6Range struct {
}

func (p *PartitionSharedVlan) GetId() string{
    return "1"
}

func (p *PartitionSharedVlan) getPath() string{
    return "partition/" +p.Inst.PartitionName + "/shared-vlan"
}

func (p *PartitionSharedVlan) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionSharedVlan::Post")
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

func (p *PartitionSharedVlan) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionSharedVlan::Get")
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
func (p *PartitionSharedVlan) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionSharedVlan::Put")
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

func (p *PartitionSharedVlan) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionSharedVlan::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
