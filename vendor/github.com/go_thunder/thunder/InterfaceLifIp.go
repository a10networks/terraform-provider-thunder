package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceLifIp struct {
	Dhcp InterfaceLifIPInstance `json:"ip,omitempty"`
}

type InterfaceLifIPInstance struct {
	Ipv4Address             []InterfaceLifAddressList `json:"address-list,omitempty"`
	AllowPromiscuousVip     int                       `json:"allow-promiscuous-vip,omitempty"`
	CacheSpoofingPort       int                       `json:"cache-spoofing-port,omitempty"`
	Dhcp                    int                       `json:"dhcp,omitempty"`
	GenerateMembershipQuery int                       `json:"generate-membership-query,omitempty"`
	Inside                  int                       `json:"inside,omitempty"`
	MaxRespTime             int                       `json:"max-resp-time,omitempty"`
	OspfGlobal              InterfaceLifOspf          `json:"ospf,omitempty"`
	Outside                 int                       `json:"outside,omitempty"`
	QueryInterval           int                       `json:"query-interval,omitempty"`
	Authentication          InterfaceLifRip           `json:"rip,omitempty"`
	Isis                    InterfaceLifRouter        `json:"router,omitempty"`
	UUID                    string                    `json:"uuid,omitempty"`
}

type InterfaceLifAddressList struct {
	Ipv4Address string `json:"ipv4-address,omitempty"`
	Ipv4Netmask string `json:"ipv4-netmask,omitempty"`
}

type InterfaceLifOspf struct {
	AuthenticationCfg InterfaceLifOspfGlobal   `json:"ospf-global,omitempty"`
	IPAddr            []InterfaceLifOspfIPList `json:"ospf-ip-list,omitempty"`
}

type InterfaceLifRip struct {
	Str           InterfaceLifAuthentication  `json:"authentication,omitempty"`
	Receive       InterfaceLifReceiveCfg      `json:"receive-cfg,omitempty"`
	ReceivePacket int                         `json:"receive-packet,omitempty"`
	Send          InterfaceLifSendCfg         `json:"send-cfg,omitempty"`
	SendPacket    int                         `json:"send-packet,omitempty"`
	State         InterfaceLifSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	UUID          string                      `json:"uuid,omitempty"`
}

type InterfaceLifRouter struct {
	Tag InterfaceLifIsis `json:"isis,omitempty"`
}

type InterfaceLifOspfGlobal struct {
	Authentication     InterfaceLifAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	AuthenticationKey  string                         `json:"authentication-key,omitempty"`
	Bfd                InterfaceLifBfdCfg             `json:"bfd-cfg,omitempty"`
	Cost               int                            `json:"cost,omitempty"`
	DatabaseFilter     InterfaceLifDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	DeadInterval       int                            `json:"dead-interval,omitempty"`
	Disable            string                         `json:"disable,omitempty"`
	HelloInterval      int                            `json:"hello-interval,omitempty"`
	MessageDigestKey   []InterfaceLifMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	Mtu                int                            `json:"mtu,omitempty"`
	MtuIgnore          int                            `json:"mtu-ignore,omitempty"`
	Broadcast          InterfaceLifNetwork            `json:"network,omitempty"`
	Priority           int                            `json:"priority,omitempty"`
	RetransmitInterval int                            `json:"retransmit-interval,omitempty"`
	TransmitDelay      int                            `json:"transmit-delay,omitempty"`
	UUID               string                         `json:"uuid,omitempty"`
}

type InterfaceLifOspfIPList struct {
	Authentication     int                            `json:"authentication,omitempty"`
	AuthenticationKey  string                         `json:"authentication-key,omitempty"`
	Cost               int                            `json:"cost,omitempty"`
	DatabaseFilter     string                         `json:"database-filter,omitempty"`
	DeadInterval       int                            `json:"dead-interval,omitempty"`
	HelloInterval      int                            `json:"hello-interval,omitempty"`
	IPAddr             string                         `json:"ip-addr,omitempty"`
	MessageDigestKey   []InterfaceLifMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	MtuIgnore          int                            `json:"mtu-ignore,omitempty"`
	Out                int                            `json:"out,omitempty"`
	Priority           int                            `json:"priority,omitempty"`
	RetransmitInterval int                            `json:"retransmit-interval,omitempty"`
	TransmitDelay      int                            `json:"transmit-delay,omitempty"`
	UUID               string                         `json:"uuid,omitempty"`
	Value              string                         `json:"value,omitempty"`
}

type InterfaceLifAuthentication struct {
	KeyChain InterfaceLifKeyChain `json:"key-chain,omitempty"`
	Mode     InterfaceLifMode     `json:"mode,omitempty"`
	String   InterfaceLifStr      `json:"str,omitempty"`
}

type InterfaceLifReceiveCfg struct {
	Receive int    `json:"receive,omitempty"`
	Version string `json:"version,omitempty"`
}

type InterfaceLifSendCfg struct {
	Send    int    `json:"send,omitempty"`
	Version string `json:"version,omitempty"`
}

type InterfaceLifSplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}

type InterfaceLifIsis struct {
	Tag  string `json:"tag,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type InterfaceLifAuthenticationCfg struct {
	Authentication int    `json:"authentication,omitempty"`
	Value          string `json:"value,omitempty"`
}

type InterfaceLifBfdCfg struct {
	Bfd     int `json:"bfd,omitempty"`
	Disable int `json:"disable,omitempty"`
}

type InterfaceLifDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter,omitempty"`
	Out            int    `json:"out,omitempty"`
}

type InterfaceLifMessageDigestCfg struct {
	Md5Value         string `json:"md5-value,omitempty"`
	MessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceLifNetwork struct {
	Broadcast         int `json:"broadcast,omitempty"`
	NonBroadcast      int `json:"non-broadcast,omitempty"`
	P2MpNbma          int `json:"p2mp-nbma,omitempty"`
	PointToMultipoint int `json:"point-to-multipoint,omitempty"`
	PointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceLifKeyChain struct {
	KeyChain string `json:"key-chain,omitempty"`
}

type InterfaceLifMode struct {
	Mode string `json:"mode,omitempty"`
}

type InterfaceLifStr struct {
	String string `json:"string,omitempty"`
}

func PostInterfaceLifIp(id string, name1 string, inst InterfaceLifIp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceLifIp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/lif/"+name1+"/ip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLifIp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceLifIp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceLifIp(id string, name1 string, host string) (*InterfaceLifIp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceLifIp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/lif/"+name1+"/ip", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLifIp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceLifIp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
