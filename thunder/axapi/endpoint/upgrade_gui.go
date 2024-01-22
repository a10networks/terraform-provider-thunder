

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type UpgradeGui struct {
	Inst struct {

    Delete []UpgradeGuiDelete `json:"delete"`
    FileUrl string `json:"file-url"`
    GuiUpload int `json:"gui-upload"`
    Image string `json:"image"`
    ImageFile string `json:"image-file"`
    Local string `json:"local"`
    RemoteUrl string `json:"remote-url"`
    Rollback string `json:"rollback"`
    SourceIpAddress string `json:"source-ip-address"`
    Upload int `json:"upload"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"gui"`
}


type UpgradeGuiDelete struct {
    File string `json:"file"`
}

func (p *UpgradeGui) GetId() string{
    return "1"
}

func (p *UpgradeGui) getPath() string{
    return "upgrade/gui"
}

func (p *UpgradeGui) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeGui::Post")
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

func (p *UpgradeGui) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeGui::Get")
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
func (p *UpgradeGui) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeGui::Put")
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

func (p *UpgradeGui) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("UpgradeGui::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
