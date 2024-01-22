

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type UpgradeHd struct {
	Inst struct {

    Device int `json:"Device"`
    FileUrl string `json:"file-url"`
    Image string `json:"image"`
    ImageFile string `json:"image-file"`
    Local string `json:"local"`
    RebootAfterUpgrade int `json:"reboot-after-upgrade"`
    Rollback int `json:"rollback"`
    SourceIpAddress string `json:"source-ip-address"`
    StaggeredUpgradeMode int `json:"staggered-upgrade-mode"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"hd"`
}

func (p *UpgradeHd) GetId() string{
    return "1"
}

func (p *UpgradeHd) getPath() string{
    return "upgrade/hd"
}

func (p *UpgradeHd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeHd::Post")
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

func (p *UpgradeHd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeHd::Get")
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
func (p *UpgradeHd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeHd::Put")
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

func (p *UpgradeHd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeHd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
