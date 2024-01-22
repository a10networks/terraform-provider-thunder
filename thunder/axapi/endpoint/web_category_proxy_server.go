

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryProxyServer struct {
	Inst struct {

    AuthType string `json:"auth-type" dval:"ntlm"`
    Domain string `json:"domain"`
    Encrypted string `json:"encrypted"`
    HttpPort int `json:"http-port"`
    HttpsPort int `json:"https-port"`
    Password int `json:"password"`
    ProxyHost string `json:"proxy-host"`
    SecretString string `json:"secret-string"`
    Username string `json:"username"`
    Uuid string `json:"uuid"`

	} `json:"proxy-server"`
}

func (p *WebCategoryProxyServer) GetId() string{
    return "1"
}

func (p *WebCategoryProxyServer) getPath() string{
    return "web-category/proxy-server"
}

func (p *WebCategoryProxyServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryProxyServer::Post")
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

func (p *WebCategoryProxyServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryProxyServer::Get")
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
func (p *WebCategoryProxyServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryProxyServer::Put")
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

func (p *WebCategoryProxyServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryProxyServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
