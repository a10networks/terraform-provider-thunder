

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsRoutingIsis struct {
	Inst struct {

    Isisadjacencychange int `json:"isisAdjacencyChange"`
    Isisareamismatch int `json:"isisAreaMismatch"`
    Isisattempttoexceedmaxsequence int `json:"isisAttemptToExceedMaxSequence"`
    Isisauthenticationfailure int `json:"isisAuthenticationFailure"`
    Isisauthenticationtypefailure int `json:"isisAuthenticationTypeFailure"`
    Isiscorruptedlspdetected int `json:"isisCorruptedLSPDetected"`
    Isisdatabaseoverload int `json:"isisDatabaseOverload"`
    Isisidlenmismatch int `json:"isisIDLenMismatch"`
    Isislsperrordetected int `json:"isisLSPErrorDetected"`
    Isislsptoolargetopropagate int `json:"isisLSPTooLargeToPropagate"`
    Isismanualaddressdrops int `json:"isisManualAddressDrops"`
    Isismaxareaaddressesmismatch int `json:"isisMaxAreaAddressesMismatch"`
    Isisoriginatinglspbuffersizemismatch int `json:"isisOriginatingLSPBufferSizeMismatch"`
    Isisownlsppurge int `json:"isisOwnLSPPurge"`
    Isisprotocolssupportedmismatch int `json:"isisProtocolsSupportedMismatch"`
    Isisrejectedadjacency int `json:"isisRejectedAdjacency"`
    Isissequencenumberskip int `json:"isisSequenceNumberSkip"`
    Isisversionskew int `json:"isisVersionSkew"`
    Uuid string `json:"uuid"`

	} `json:"isis"`
}

func (p *SnmpServerEnableTrapsRoutingIsis) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsRoutingIsis) getPath() string{
    return "snmp-server/enable/traps/routing/isis"
}

func (p *SnmpServerEnableTrapsRoutingIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingIsis::Post")
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

func (p *SnmpServerEnableTrapsRoutingIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingIsis::Get")
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
func (p *SnmpServerEnableTrapsRoutingIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingIsis::Put")
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

func (p *SnmpServerEnableTrapsRoutingIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsRoutingIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
