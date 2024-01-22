

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsUnicastElection struct {
	Inst struct {

    Members VcsUnicastElectionMembers1907 `json:"members"`
    Port int `json:"port" dval:"41217"`
    Uuid string `json:"uuid"`

	} `json:"unicast-election"`
}


type VcsUnicastElectionMembers1907 struct {
    IpAddressCfg []VcsUnicastElectionMembersIpAddressCfg1908 `json:"ip-address-cfg"`
    Ipv6AddressCfg []VcsUnicastElectionMembersIpv6AddressCfg1909 `json:"ipv6-address-cfg"`
    Uuid string `json:"uuid"`
}


type VcsUnicastElectionMembersIpAddressCfg1908 struct {
    IpAddress string `json:"ip-address"`
    UseMgmtPort int `json:"use-mgmt-port"`
}


type VcsUnicastElectionMembersIpv6AddressCfg1909 struct {
    Ipv6Address string `json:"ipv6-address"`
    UseMgmtPort int `json:"use-mgmt-port"`
}

func (p *VcsUnicastElection) GetId() string{
    return "1"
}

func (p *VcsUnicastElection) getPath() string{
    return "vcs/unicast-election"
}

func (p *VcsUnicastElection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElection::Post")
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

func (p *VcsUnicastElection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElection::Get")
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
func (p *VcsUnicastElection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElection::Put")
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

func (p *VcsUnicastElection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsUnicastElection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
