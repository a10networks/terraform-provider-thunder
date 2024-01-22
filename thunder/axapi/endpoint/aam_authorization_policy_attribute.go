

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AamAuthorizationPolicyAttribute struct {
	Inst struct {

    A10AxAuthUri int `json:"A10-AX-AUTH-URI"`
    A10DynamicDefined int `json:"a10-dynamic-defined"`
    Any int `json:"any"`
    AttrInt string `json:"attr-int"`
    AttrIntVal int `json:"attr-int-val"`
    AttrIp string `json:"attr-ip"`
    AttrIpv4 string `json:"attr-ipv4"`
    AttrNum int `json:"attr-num"`
    AttrNumber string `json:"attr-number"`
    AttrNumberVal string `json:"attr-number-val"`
    AttrStr string `json:"attr-str"`
    AttrStrVal string `json:"attr-str-val"`
    AttrType int `json:"attr-type"`
    AttributeName string `json:"attribute-name"`
    CustomAttrStr string `json:"custom-attr-str"`
    CustomAttrType int `json:"custom-attr-type"`
    IntegerType int `json:"integer-type"`
    IpType int `json:"ip-type"`
    NumberType int `json:"number-type"`
    StringType int `json:"string-type"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"attribute"`
}

func (p *AamAuthorizationPolicyAttribute) GetId() string{
    return strconv.Itoa(p.Inst.AttrNum)
}

func (p *AamAuthorizationPolicyAttribute) getPath() string{
    return "aam/authorization/policy/" +p.Inst.Name + "/attribute"
}

func (p *AamAuthorizationPolicyAttribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyAttribute::Post")
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

func (p *AamAuthorizationPolicyAttribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyAttribute::Get")
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
func (p *AamAuthorizationPolicyAttribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyAttribute::Put")
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

func (p *AamAuthorizationPolicyAttribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyAttribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
