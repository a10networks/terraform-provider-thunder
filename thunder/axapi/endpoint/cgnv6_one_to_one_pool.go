

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOnePool struct {
	Inst struct {

    EndAddress string `json:"end-address"`
    Group string `json:"group"`
    Netmask string `json:"netmask"`
    Partition string `json:"partition"`
    PoolName string `json:"pool-name"`
    Shared int `json:"shared"`
    StartAddress string `json:"start-address"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"pool"`
}

func (p *Cgnv6OneToOnePool) GetId() string{
    return url.QueryEscape(p.Inst.PoolName)
}

func (p *Cgnv6OneToOnePool) getPath() string{
    return "cgnv6/one-to-one/pool"
}

func (p *Cgnv6OneToOnePool) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePool::Post")
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

func (p *Cgnv6OneToOnePool) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePool::Get")
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
func (p *Cgnv6OneToOnePool) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePool::Put")
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

func (p *Cgnv6OneToOnePool) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePool::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
