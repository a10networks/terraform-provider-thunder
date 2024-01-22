

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRadiusServer struct {
	Inst struct {

    AccountingInterimUpdate string `json:"accounting-interim-update" dval:"ignore"`
    AccountingOn string `json:"accounting-on" dval:"ignore"`
    AccountingStart string `json:"accounting-start" dval:"append-entry"`
    AccountingStop string `json:"accounting-stop" dval:"delete-entry"`
    Attribute []Cgnv6LsnRadiusServerAttribute `json:"attribute"`
    AttributeName string `json:"attribute-name"`
    CustomAttributeName string `json:"custom-attribute-name"`
    DisableReply int `json:"disable-reply"`
    Encrypted string `json:"encrypted"`
    ListenPort int `json:"listen-port"`
    Remote Cgnv6LsnRadiusServerRemote `json:"remote"`
    SamplingEnable []Cgnv6LsnRadiusServerSamplingEnable `json:"sampling-enable"`
    Secret int `json:"secret"`
    SecretString string `json:"secret-string"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"server"`
}


type Cgnv6LsnRadiusServerAttribute struct {
    AttributeValue string `json:"attribute-value"`
    PrefixLength string `json:"prefix-length"`
    PrefixVendor int `json:"prefix-vendor"`
    PrefixNumber int `json:"prefix-number"`
    Name string `json:"name"`
    Value string `json:"value"`
    CustomVendor int `json:"custom-vendor"`
    CustomNumber int `json:"custom-number"`
    Vendor int `json:"vendor"`
    Number int `json:"number"`
}


type Cgnv6LsnRadiusServerRemote struct {
    IpList []Cgnv6LsnRadiusServerRemoteIpList `json:"ip-list"`
}


type Cgnv6LsnRadiusServerRemoteIpList struct {
    IpListName string `json:"ip-list-name"`
    IpListSecret int `json:"ip-list-secret"`
    IpListSecretString string `json:"ip-list-secret-string"`
    IpListEncrypted string `json:"ip-list-encrypted"`
}


type Cgnv6LsnRadiusServerSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6LsnRadiusServer) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRadiusServer) getPath() string{
    return "cgnv6/lsn/radius/server"
}

func (p *Cgnv6LsnRadiusServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusServer::Post")
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

func (p *Cgnv6LsnRadiusServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusServer::Get")
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
func (p *Cgnv6LsnRadiusServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusServer::Put")
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

func (p *Cgnv6LsnRadiusServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
