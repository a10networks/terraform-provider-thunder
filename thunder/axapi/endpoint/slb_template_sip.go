

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateSip struct {
	Inst struct {

    AclId int `json:"acl-id"`
    AclNameValue string `json:"acl-name-value"`
    AlgDestNat int `json:"alg-dest-nat"`
    AlgSourceNat int `json:"alg-source-nat"`
    CallIdPersistDisable int `json:"call-id-persist-disable"`
    ClientKeepAlive int `json:"client-keep-alive"`
    ClientRequestHeader []SlbTemplateSipClientRequestHeader `json:"client-request-header"`
    ClientResponseHeader []SlbTemplateSipClientResponseHeader `json:"client-response-header"`
    DialogAware int `json:"dialog-aware"`
    DropWhenClientFail int `json:"drop-when-client-fail"`
    DropWhenServerFail int `json:"drop-when-server-fail"`
    ExcludeTranslation []SlbTemplateSipExcludeTranslation `json:"exclude-translation"`
    FailedClientSelection int `json:"failed-client-selection"`
    FailedClientSelectionMessage string `json:"failed-client-selection-message"`
    FailedServerSelection int `json:"failed-server-selection"`
    FailedServerSelectionMessage string `json:"failed-server-selection-message"`
    InsertClientIp int `json:"insert-client-ip"`
    Interval int `json:"interval" dval:"30"`
    KeepServerIpIfMatchAcl int `json:"keep-server-ip-if-match-acl"`
    Name string `json:"name"`
    PstnGw string `json:"pstn-gw" dval:"pstn"`
    ServerKeepAlive int `json:"server-keep-alive"`
    ServerRequestHeader []SlbTemplateSipServerRequestHeader `json:"server-request-header"`
    ServerResponseHeader []SlbTemplateSipServerResponseHeader `json:"server-response-header"`
    ServerSelectionPerRequest int `json:"server-selection-per-request"`
    ServiceGroup string `json:"service-group"`
    SmpCallIdRtpSession int `json:"smp-call-id-rtp-session"`
    Timeout int `json:"timeout" dval:"30"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"sip"`
}


type SlbTemplateSipClientRequestHeader struct {
    ClientRequestHeaderErase string `json:"client-request-header-erase"`
    ClientRequestEraseAll int `json:"client-request-erase-all"`
    ClientRequestHeaderInsert string `json:"client-request-header-insert"`
    InsertConditionClientRequest string `json:"insert-condition-client-request"`
}


type SlbTemplateSipClientResponseHeader struct {
    ClientResponseHeaderErase string `json:"client-response-header-erase"`
    ClientResponseEraseAll int `json:"client-response-erase-all"`
    ClientResponseHeaderInsert string `json:"client-response-header-insert"`
    InsertConditionClientResponse string `json:"insert-condition-client-response"`
}


type SlbTemplateSipExcludeTranslation struct {
    TranslationValue string `json:"translation-value"`
    HeaderString string `json:"header-string"`
}


type SlbTemplateSipServerRequestHeader struct {
    ServerRequestHeaderErase string `json:"server-request-header-erase"`
    ServerRequestEraseAll int `json:"server-request-erase-all"`
    ServerRequestHeaderInsert string `json:"server-request-header-insert"`
    InsertConditionServerRequest string `json:"insert-condition-server-request"`
}


type SlbTemplateSipServerResponseHeader struct {
    ServerResponseHeaderErase string `json:"server-response-header-erase"`
    ServerResponseEraseAll int `json:"server-response-erase-all"`
    ServerResponseHeaderInsert string `json:"server-response-header-insert"`
    InsertConditionServerResponse string `json:"insert-condition-server-response"`
}

func (p *SlbTemplateSip) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateSip) getPath() string{
    return "slb/template/sip"
}

func (p *SlbTemplateSip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSip::Post")
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

func (p *SlbTemplateSip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSip::Get")
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
func (p *SlbTemplateSip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSip::Put")
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

func (p *SlbTemplateSip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
