

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerMigrateVip struct {
	Inst struct {

    CancelMigration int `json:"cancel-migration"`
    FinishMigration int `json:"finish-migration"`
    TargetDataCpu int `json:"target-data-cpu"`
    TargetFloatingIpv4 string `json:"target-floating-ipv4"`
    TargetFloatingIpv6 string `json:"target-floating-ipv6"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"migrate-vip"`
}

func (p *SlbVirtualServerMigrateVip) GetId() string{
    return "1"
}

func (p *SlbVirtualServerMigrateVip) getPath() string{
    return "slb/virtual-server/" +p.Inst.Name + "/migrate-vip"
}

func (p *SlbVirtualServerMigrateVip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerMigrateVip::Post")
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

func (p *SlbVirtualServerMigrateVip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerMigrateVip::Get")
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
func (p *SlbVirtualServerMigrateVip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerMigrateVip::Put")
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

func (p *SlbVirtualServerMigrateVip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerMigrateVip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
