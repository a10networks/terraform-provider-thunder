

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterDbConfig struct {
	Inst struct {

    BrokenDetectTimeout int `json:"broken-detect-timeout" dval:"12000"`
    ClientRecvTimeout int `json:"client-recv-timeout" dval:"13000"`
    Clientport int `json:"clientPort"`
    ElectConnTimeout int `json:"elect-conn-timeout" dval:"1200"`
    Initlimit int `json:"initLimit"`
    LoopbackIntfSupport int `json:"loopback-intf-support" dval:"1"`
    Maxsessiontimeout int `json:"maxSessionTimeout" dval:"30000"`
    Minsessiontimeout int `json:"minSessionTimeout" dval:"100"`
    MoreElectionPacket int `json:"more-election-packet" dval:"1"`
    Synclimit int `json:"syncLimit"`
    Ticktime int `json:"tickTime"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"db-config"`
}

func (p *ScaleoutClusterDbConfig) GetId() string{
    return "1"
}

func (p *ScaleoutClusterDbConfig) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/db-config"
}

func (p *ScaleoutClusterDbConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDbConfig::Post")
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

func (p *ScaleoutClusterDbConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDbConfig::Get")
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
func (p *ScaleoutClusterDbConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDbConfig::Put")
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

func (p *ScaleoutClusterDbConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutClusterDbConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
