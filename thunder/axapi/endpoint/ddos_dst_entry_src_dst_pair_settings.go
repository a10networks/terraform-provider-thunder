

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairSettings struct {
	Inst struct {

    Age int `json:"age" dval:"5"`
    AllTypes string `json:"all-types"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    EnableClassListOverflow int `json:"enable-class-list-overflow"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairSettingsL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    SrcPrefixLen int `json:"src-prefix-len"`
    UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"src-dst-pair-settings"`
}


type DdosDstEntrySrcDstPairSettingsL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstEntrySrcDstPairSettings) GetId() string{
    return p.Inst.AllTypes
}

func (p *DdosDstEntrySrcDstPairSettings) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-settings"
}

func (p *DdosDstEntrySrcDstPairSettings) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairSettings::Post")
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

func (p *DdosDstEntrySrcDstPairSettings) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairSettings::Get")
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
func (p *DdosDstEntrySrcDstPairSettings) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairSettings::Put")
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

func (p *DdosDstEntrySrcDstPairSettings) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairSettings::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
