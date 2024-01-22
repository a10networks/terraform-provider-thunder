

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbCommonAflexTableEntrySync struct {
	Inst struct {

    AflexTableEntrySyncEnable int `json:"aflex-table-entry-sync-enable"`
    AflexTableEntrySyncMaxKeyLen int `json:"aflex-table-entry-sync-max-key-len" dval:"1000"`
    AflexTableEntrySyncMaxValueLen int `json:"aflex-table-entry-sync-max-value-len" dval:"1000"`
    AflexTableEntrySyncMinLifetime int `json:"aflex-table-entry-sync-min-lifetime"`
    Uuid string `json:"uuid"`

	} `json:"aflex-table-entry-sync"`
}

func (p *SlbCommonAflexTableEntrySync) GetId() string{
    return "1"
}

func (p *SlbCommonAflexTableEntrySync) getPath() string{
    return "slb/common/aflex-table-entry-sync"
}

func (p *SlbCommonAflexTableEntrySync) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonAflexTableEntrySync::Post")
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

func (p *SlbCommonAflexTableEntrySync) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonAflexTableEntrySync::Get")
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
func (p *SlbCommonAflexTableEntrySync) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonAflexTableEntrySync::Put")
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

func (p *SlbCommonAflexTableEntrySync) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonAflexTableEntrySync::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
