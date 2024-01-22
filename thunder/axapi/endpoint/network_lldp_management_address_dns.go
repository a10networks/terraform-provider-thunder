

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkLldpManagementAddressDns struct {
	Inst struct {

    Dns string `json:"dns"`
    Interface NetworkLldpManagementAddressDnsInterface `json:"interface"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type NetworkLldpManagementAddressDnsInterface struct {
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Management int `json:"management"`
}

func (p *NetworkLldpManagementAddressDns) GetId() string{
    return p.Inst.Dns
}

func (p *NetworkLldpManagementAddressDns) getPath() string{
    return "network/lldp/management-address/dns"
}

func (p *NetworkLldpManagementAddressDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressDns::Post")
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

func (p *NetworkLldpManagementAddressDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressDns::Get")
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
func (p *NetworkLldpManagementAddressDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressDns::Put")
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

func (p *NetworkLldpManagementAddressDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
