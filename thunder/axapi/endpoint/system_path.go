

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SystemPath struct {
	Inst struct {

    IfpairEthEnd int `json:"ifpair-eth-end"`
    IfpairEthStart int `json:"ifpair-eth-start"`
    IfpairTrunkEnd int `json:"ifpair-trunk-end"`
    IfpairTrunkStart int `json:"ifpair-trunk-start"`
    L2hmAttach string `json:"l2hm-attach"`
    L2hmPathName string `json:"l2hm-path-name"`
    L2hmSetupTestApi int `json:"l2hm-setup-test-api"`
    L2hmVlan int `json:"l2hm-vlan"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"path"`
}

func (p *SystemPath) GetId() string{
    return url.QueryEscape(p.Inst.L2hmPathName)
}

func (p *SystemPath) getPath() string{
    return "system/path"
}

func (p *SystemPath) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPath::Post")
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

func (p *SystemPath) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPath::Get")
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
func (p *SystemPath) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPath::Put")
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

func (p *SystemPath) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPath::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
