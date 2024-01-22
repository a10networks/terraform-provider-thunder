

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbProtocolOper struct {
    
    Oper GslbProtocolOperOper `json:"oper"`

}
type DataGslbProtocolOper struct {
    DtGslbProtocolOper GslbProtocolOper `json:"protocol"`
}


type GslbProtocolOperOper struct {
    SessionList []GslbProtocolOperOperSessionList `json:"session-list"`
}


type GslbProtocolOperOperSessionList struct {
    ProtocolInfo string `json:"protocol-info"`
    State string `json:"state"`
    SessionId int `json:"session-id"`
    ConnectionSucceeded int `json:"connection-succeeded"`
    ConnectionFailed int `json:"connection-failed"`
    OpenPacketSent int `json:"open-packet-sent"`
    OpenPacketReceived int `json:"open-packet-received"`
    OpenSessionSucceeded int `json:"open-session-succeeded"`
    OpenSessionFailed int `json:"open-session-failed"`
    SessionsDropped int `json:"sessions-dropped"`
    Retry int `json:"retry"`
    UpdatePacketSent int `json:"update-packet-sent"`
    UpdatePacketReceived int `json:"update-packet-received"`
    KeepalivePacketSent int `json:"keepalive-packet-sent"`
    KeepalivePacketReceived int `json:"keepalive-packet-received"`
    NotifyPacketSent int `json:"notify-packet-sent"`
    NotifyPacketReceived int `json:"notify-packet-received"`
    MessageHeaderError int `json:"message-header-error"`
    SecureNegotiationSuccess int `json:"secure-negotiation-success"`
    SecureNegotiationFail int `json:"secure-negotiation-fail"`
    SslHandshakeSuccess int `json:"ssl-handshake-success"`
    SslHandshakeFail int `json:"ssl-handshake-fail"`
    Secure_state string `json:"secure_state"`
    Secure_config string `json:"secure_config"`
}

func (p *GslbProtocolOper) GetId() string{
    return "1"
}

func (p *GslbProtocolOper) getPath() string{
    return "gslb/protocol/oper"
}

func (p *GslbProtocolOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbProtocolOper,error) {
logger.Println("GslbProtocolOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbProtocolOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
