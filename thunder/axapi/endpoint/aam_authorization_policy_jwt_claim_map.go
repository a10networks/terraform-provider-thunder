

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AamAuthorizationPolicyJwtClaimMap struct {
	Inst struct {

    AttrNum int `json:"attr-num"`
    BoolVal string `json:"bool-val"`
    BooleanType int `json:"boolean-type"`
    Claim string `json:"claim"`
    NumVal int `json:"num-val"`
    NumberType int `json:"number-type"`
    StrVal string `json:"str-val"`
    StringType int `json:"string-type"`
    Type int `json:"type"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"jwt-claim-map"`
}

func (p *AamAuthorizationPolicyJwtClaimMap) GetId() string{
    return strconv.Itoa(p.Inst.AttrNum)
}

func (p *AamAuthorizationPolicyJwtClaimMap) getPath() string{
    return "aam/authorization/policy/" +p.Inst.Name + "/jwt-claim-map"
}

func (p *AamAuthorizationPolicyJwtClaimMap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyJwtClaimMap::Post")
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

func (p *AamAuthorizationPolicyJwtClaimMap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyJwtClaimMap::Get")
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
func (p *AamAuthorizationPolicyJwtClaimMap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyJwtClaimMap::Put")
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

func (p *AamAuthorizationPolicyJwtClaimMap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthorizationPolicyJwtClaimMap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
