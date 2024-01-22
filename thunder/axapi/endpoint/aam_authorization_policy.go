

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthorizationPolicy struct {
	Inst struct {

    AttributeList []AamAuthorizationPolicyAttributeList `json:"attribute-list"`
    AttributeRule string `json:"attribute-rule"`
    ExtendedFilter string `json:"extended-filter"`
    ForwardPolicyAuthorizeOnly int `json:"forward-policy-authorize-only"`
    JwtAuthorization string `json:"jwt-authorization"`
    JwtClaimMapList []AamAuthorizationPolicyJwtClaimMapList `json:"jwt-claim-map-list"`
    Name string `json:"name"`
    Server string `json:"server"`
    ServiceGroup string `json:"service-group"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"policy"`
}


type AamAuthorizationPolicyAttributeList struct {
    AttrNum int `json:"attr-num"`
    AttributeName string `json:"attribute-name"`
    Any int `json:"any"`
    AttrType int `json:"attr-type"`
    StringType int `json:"string-type"`
    IntegerType int `json:"integer-type"`
    IpType int `json:"ip-type"`
    NumberType int `json:"number-type"`
    AttrStr string `json:"attr-str"`
    AttrStrVal string `json:"attr-str-val"`
    AttrInt string `json:"attr-int"`
    AttrIntVal int `json:"attr-int-val"`
    AttrIp string `json:"attr-ip"`
    AttrIpv4 string `json:"attr-ipv4"`
    AttrNumber string `json:"attr-number"`
    AttrNumberVal string `json:"attr-number-val"`
    A10AxAuthUri int `json:"A10-AX-AUTH-URI"`
    CustomAttrType int `json:"custom-attr-type"`
    CustomAttrStr string `json:"custom-attr-str"`
    A10DynamicDefined int `json:"a10-dynamic-defined"`
    Uuid string `json:"uuid"`
}


type AamAuthorizationPolicyJwtClaimMapList struct {
    AttrNum int `json:"attr-num"`
    Claim string `json:"claim"`
    Type int `json:"type"`
    StringType int `json:"string-type"`
    NumberType int `json:"number-type"`
    BooleanType int `json:"boolean-type"`
    StrVal string `json:"str-val"`
    NumVal int `json:"num-val"`
    BoolVal string `json:"bool-val"`
    Uuid string `json:"uuid"`
}

func (p *AamAuthorizationPolicy) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthorizationPolicy) getPath() string{
    return "aam/authorization/policy"
}

func (p *AamAuthorizationPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicy::Post")
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

func (p *AamAuthorizationPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicy::Get")
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
func (p *AamAuthorizationPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicy::Put")
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

func (p *AamAuthorizationPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
