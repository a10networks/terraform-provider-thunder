

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Copy struct {
	Inst struct {

    DestProfile string `json:"dest-profile"`
    DestRemoteFile string `json:"dest-remote-file"`
    DestUseMgmtPort int `json:"dest-use-mgmt-port"`
    Profile string `json:"profile"`
    RemoteFile string `json:"remote-file"`
    RunningConfig int `json:"running-config"`
    StartupConfig int `json:"startup-config"`
    ToProfile string `json:"to-profile"`
    ToRunningConfig int `json:"to-running-config"`
    ToStartupConfig int `json:"to-startup-config"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"copy"`
}

func (p *Copy) GetId() string{
    return "1"
}

func (p *Copy) getPath() string{
    return "copy"
}

func (p *Copy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Copy::Post")
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

func (p *Copy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Copy::Get")
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
func (p *Copy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Copy::Put")
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

func (p *Copy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Copy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
