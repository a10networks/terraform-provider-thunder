package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type VeIP struct {
	UUID VeIPInstance `json:"ip,omitempty"`
}

type VeIPAddressList struct {
	Ipv4Address string `json:"ipv4-address,omitempty"`
	Ipv4Netmask string `json:"ipv4-netmask,omitempty"`
}
type VeIPHelperAddressList struct {
	HelperAddress string `json:"helper-address,omitempty"`
}
type VeIPStatefulFirewall struct {
	UUID       string `json:"uuid,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	ACLID      int    `json:"acl-id,omitempty"`
	AccessList int    `json:"access-list,omitempty"`
}
type VeIPReceiveCfg struct {
	Receive int    `json:"receive,omitempty"`
	Version string `json:"version,omitempty"`
}
type VeIPSplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}
type VeIPKeyChain struct {
	KeyChain string `json:"key-chain,omitempty"`
}
type VeIPMode struct {
	Mode string `json:"mode,omitempty"`
}
type VeIPStr struct {
	String string `json:"string,omitempty"`
}
type VeIPAuthentication struct {
	KeyChain VeIPKeyChain `json:"key-chain,omitempty"`
	Mode     VeIPMode     `json:"mode,omitempty"`
	String   VeIPStr      `json:"str,omitempty"`
}
type VeIPSendCfg struct {
	Version string `json:"version,omitempty"`
	Send    int    `json:"send,omitempty"`
}
type VeIPRip struct {
	Receive       VeIPReceiveCfg      `json:"receive-cfg,omitempty"`
	UUID          string              `json:"uuid,omitempty"`
	ReceivePacket int                 `json:"receive-packet,omitempty"`
	State         VeIPSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	KeyChain      VeIPAuthentication  `json:"authentication,omitempty"`
	Version       VeIPSendCfg         `json:"send-cfg,omitempty"`
	SendPacket    int                 `json:"send-packet,omitempty"`
}
type VeIPIsis struct {
	Tag  string `json:"tag,omitempty"`
	UUID string `json:"uuid,omitempty"`
}
type VeIPRouter struct {
	Tag VeIPIsis `json:"isis,omitempty"`
}
type VeIPMd5 struct {
	Md5Value  string `json:"md5-value,omitempty"`
	Encrypted string `json:"encrypted,omitempty"`
}
type VeIPMessageDigestCfg struct {
	MessageDigestKey int     `json:"message-digest-key,omitempty"`
	Md5Value         VeIPMd5 `json:"md5,omitempty"`
}
type VeIPOspfIPList struct {
	DeadInterval       int                    `json:"dead-interval,omitempty"`
	AuthenticationKey  string                 `json:"authentication-key,omitempty"`
	UUID               string                 `json:"uuid,omitempty"`
	MtuIgnore          int                    `json:"mtu-ignore,omitempty"`
	TransmitDelay      int                    `json:"transmit-delay,omitempty"`
	Value              string                 `json:"value,omitempty"`
	Priority           int                    `json:"priority,omitempty"`
	Authentication     int                    `json:"authentication,omitempty"`
	Cost               int                    `json:"cost,omitempty"`
	DatabaseFilter     string                 `json:"database-filter,omitempty"`
	HelloInterval      int                    `json:"hello-interval,omitempty"`
	IPAddr             string                 `json:"ip-addr,omitempty"`
	RetransmitInterval int                    `json:"retransmit-interval,omitempty"`
	MessageDigestKey   []VeIPMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	Out                int                    `json:"out,omitempty"`
}
type VeIPNetwork struct {
	Broadcast         int `json:"broadcast,omitempty"`
	PointToMultipoint int `json:"point-to-multipoint,omitempty"`
	NonBroadcast      int `json:"non-broadcast,omitempty"`
	PointToPoint      int `json:"point-to-point,omitempty"`
	P2MpNbma          int `json:"p2mp-nbma,omitempty"`
}
type VeIPAuthenticationCfg struct {
	Authentication int    `json:"authentication,omitempty"`
	Value          string `json:"value,omitempty"`
}
type VeIPBfdCfg struct {
	Disable int `json:"disable,omitempty"`
	Bfd     int `json:"bfd,omitempty"`
}
type VeIPDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter,omitempty"`
	Out            int    `json:"out,omitempty"`
}
type VeIPOspfGlobal struct {
	Cost               int                    `json:"cost,omitempty"`
	DeadInterval       int                    `json:"dead-interval,omitempty"`
	AuthenticationKey  string                 `json:"authentication-key,omitempty"`
	Broadcast          VeIPNetwork            `json:"network,omitempty"`
	MtuIgnore          int                    `json:"mtu-ignore,omitempty"`
	TransmitDelay      int                    `json:"transmit-delay,omitempty"`
	Authentication     VeIPAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	RetransmitInterval int                    `json:"retransmit-interval,omitempty"`
	Bfd                VeIPBfdCfg             `json:"bfd-cfg,omitempty"`
	Disable            string                 `json:"disable,omitempty"`
	HelloInterval      int                    `json:"hello-interval,omitempty"`
	DatabaseFilter     VeIPDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	Priority           int                    `json:"priority,omitempty"`
	Mtu                int                    `json:"mtu,omitempty"`
	MessageDigestKey   []VeIPMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	UUID               string                 `json:"uuid,omitempty"`
}
type VeIPOspf struct {
	DeadInterval []VeIPOspfIPList `json:"ospf-ip-list,omitempty"`
	Cost         VeIPOspfGlobal   `json:"ospf-global,omitempty"`
}
type VeIPInstance struct {
	UUID                    string                  `json:"uuid,omitempty"`
	GenerateMembershipQuery int                     `json:"generate-membership-query,omitempty"`
	Ipv4Address             []VeIPAddressList       `json:"address-list,omitempty"`
	Inside                  int                     `json:"inside,omitempty"`
	AllowPromiscuousVip     int                     `json:"allow-promiscuous-vip,omitempty"`
	HelperAddress           []VeIPHelperAddressList `json:"helper-address-list,omitempty"`
	MaxRespTime             int                     `json:"max-resp-time,omitempty"`
	QueryInterval           int                     `json:"query-interval,omitempty"`
	Outside                 int                     `json:"outside,omitempty"`
	Client                  int                     `json:"client,omitempty"`
	ClassList               VeIPStatefulFirewall    `json:"stateful-firewall,omitempty"`
	Receive                 VeIPRip                 `json:"rip,omitempty"`
	TTLIgnore               int                     `json:"ttl-ignore,omitempty"`
	Tag                     VeIPRouter              `json:"router,omitempty"`
	Dhcp                    int                     `json:"dhcp,omitempty"`
	Server                  int                     `json:"server,omitempty"`
	DeadInterval            VeIPOspf                `json:"ospf,omitempty"`
	SlbPartitionRedirect    int                     `json:"slb-partition-redirect,omitempty"`
}

func PostInterfaceVeIP(id string, name int, inst VeIP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceVeIP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ve/"+strconv.Itoa(name)+"/ip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VeIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceVeIP REQ RES..........................", m)
			err := check_api_status("PostInterfaceVeIP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceVeIP(id string, name string, host string) (*VeIP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceVeIP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ve/"+name+"/ip", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VeIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceVeIP REQ RES..........................", m)
			err := check_api_status("GetInterfaceVeIP", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
