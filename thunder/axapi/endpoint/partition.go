

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Partition struct {
	Inst struct {

    ApplicationType string `json:"application-type"`
    Id1 int `json:"id"`
    PartitionName string `json:"partition-name"`
    SharedVlan PartitionSharedVlan1089 `json:"shared-vlan"`
    Template PartitionTemplate1092 `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"partition"`
}


type PartitionSharedVlan1089 struct {
    Vlan int `json:"vlan"`
    AllowableIpRange []PartitionSharedVlanAllowableIpRange1090 `json:"allowable-ip-range"`
    AllowableIpv6Range []PartitionSharedVlanAllowableIpv6Range1091 `json:"allowable-ipv6-range"`
    MgmtFloatingIpAddress string `json:"mgmt-floating-ip-address"`
    Vrid int `json:"vrid"`
    Uuid string `json:"uuid"`
}


type PartitionSharedVlanAllowableIpRange1090 struct {
}


type PartitionSharedVlanAllowableIpv6Range1091 struct {
}


type PartitionTemplate1092 struct {
    ResourceAccounting string `json:"resource-accounting"`
    Uuid string `json:"uuid"`
}

func (p *Partition) GetId() string{
    return p.Inst.PartitionName
}

func (p *Partition) getPath() string{
    return "partition"
}

func (p *Partition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Partition::Post")
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

func (p *Partition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Partition::Get")
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
func (p *Partition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Partition::Put")
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

func (p *Partition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Partition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
