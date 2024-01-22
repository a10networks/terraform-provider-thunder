

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsUnicastElectionMembers struct {
	Inst struct {

    IpAddressCfg []VcsUnicastElectionMembersIpAddressCfg `json:"ip-address-cfg"`
    Ipv6AddressCfg []VcsUnicastElectionMembersIpv6AddressCfg `json:"ipv6-address-cfg"`
    Uuid string `json:"uuid"`

	} `json:"members"`
}


type VcsUnicastElectionMembersIpAddressCfg struct {
    IpAddress string `json:"ip-address"`
    UseMgmtPort int `json:"use-mgmt-port"`
}


type VcsUnicastElectionMembersIpv6AddressCfg struct {
    Ipv6Address string `json:"ipv6-address"`
    UseMgmtPort int `json:"use-mgmt-port"`
}

func (p *VcsUnicastElectionMembers) GetId() string{
    return "1"
}

func (p *VcsUnicastElectionMembers) getPath() string{
    return "vcs/unicast-election/members"
}

func (p *VcsUnicastElectionMembers) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElectionMembers::Post")
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

func (p *VcsUnicastElectionMembers) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElectionMembers::Get")
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
func (p *VcsUnicastElectionMembers) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElectionMembers::Put")
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

func (p *VcsUnicastElectionMembers) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElectionMembers::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
