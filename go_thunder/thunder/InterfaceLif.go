package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceLif struct {
	Ifname InterfaceLifInstance `json:"lif,omitempty"`
}

type InterfaceLifInstance struct {
	AclID          InterfaceAccessList       `json:"access-list,omitempty"`
	Action         string                    `json:"action,omitempty"`
	Authentication InterfaceBfd              `json:"bfd,omitempty"`
	Dhcp           InterfaceIP               `json:"ip,omitempty"`
	Ifname         string                    `json:"ifname,omitempty"`
	AddressList    InterfaceIpv6             `json:"ipv6,omitempty"`
	SendOnlyList   InterfaceIsis             `json:"isis,omitempty"`
	Mtu            int                       `json:"mtu,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Counters1      []InterfaceSamplingEnable `json:"sampling-enable,omitempty"`
	UUID           string                    `json:"uuid,omitempty"`
	UserTag        string                    `json:"user-tag,omitempty"`
}

type InterfaceAccessList struct {
	AclID   int    `json:"acl-id,omitempty"`
	AclName string `json:"acl-name,omitempty"`
}

type InterfaceBfd struct {
	KeyID    InterfaceAuthentication `json:"authentication,omitempty"`
	Demand   int                     `json:"demand,omitempty"`
	Echo     int                     `json:"echo,omitempty"`
	Interval InterfaceIntervalCfg    `json:"interval-cfg,omitempty"`
	UUID     string                  `json:"uuid,omitempty"`
}

type InterfaceIP struct {
	Ipv6Addr                []InterfaceLifIPAddressList `json:"address-list,omitempty"`
	AllowPromiscuousVip     int                         `json:"allow-promiscuous-vip,omitempty"`
	CacheSpoofingPort       int                         `json:"cache-spoofing-port,omitempty"`
	Dhcp                    int                         `json:"dhcp,omitempty"`
	GenerateMembershipQuery int                         `json:"generate-membership-query,omitempty"`
	Inside                  int                         `json:"inside,omitempty"`
	MaxRespTime             int                         `json:"max-resp-time,omitempty"`
	OspfGlobal              InterfaceLifIpOspf          `json:"ospf,omitempty"`
	Outside                 int                         `json:"outside,omitempty"`
	QueryInterval           int                         `json:"query-interval,omitempty"`
	Authentication          InterfaceRip                `json:"rip,omitempty"`
	Isis                    InterfaceLifIPRouter        `json:"router,omitempty"`
	UUID                    string                      `json:"uuid,omitempty"`
}

type InterfaceIpv6 struct {
	Ipv6Addr    []InterfaceLifIPAddressList `json:"address-list,omitempty"`
	Inside      int                         `json:"inside,omitempty"`
	Ipv6Enable  int                         `json:"ipv6-enable,omitempty"`
	NetworkList InterfaceOspf               `json:"ospf,omitempty"`
	Outside     int                         `json:"outside,omitempty"`
	Ripng       InterfaceRouter             `json:"router,omitempty"`
	UUID        string                      `json:"uuid,omitempty"`
}

type InterfaceIsis struct {
	SendOnlyList         InterfaceLifIsisAuthentication      `json:"authentication,omitempty"`
	Bfd                  InterfaceBfdCfg                     `json:"bfd-cfg,omitempty"`
	CircuitType          string                              `json:"circuit-type,omitempty"`
	CsnpInterval         []InterfaceCsnpIntervalList         `json:"csnp-interval-list,omitempty"`
	HelloInterval        []InterfaceHelloIntervalList        `json:"hello-interval-list,omitempty"`
	HelloIntervalMinimal []InterfaceHelloIntervalMinimalList `json:"hello-interval-minimal-list,omitempty"`
	HelloMultiplier      []InterfaceHelloMultiplierList      `json:"hello-multiplier-list,omitempty"`
	LspInterval          int                                 `json:"lsp-interval,omitempty"`
	Value                InterfaceMeshGroup                  `json:"mesh-group,omitempty"`
	Metric               []InterfaceMetricList               `json:"metric-list,omitempty"`
	Network              string                              `json:"network,omitempty"`
	Padding              int                                 `json:"padding,omitempty"`
	Password             []InterfacePasswordList             `json:"password-list,omitempty"`
	Priority             []InterfacePriorityList             `json:"priority-list,omitempty"`
	RetransmitInterval   int                                 `json:"retransmit-interval,omitempty"`
	UUID                 string                              `json:"uuid,omitempty"`
	WideMetric           []InterfaceWideMetricList           `json:"wide-metric-list,omitempty"`
}

type InterfaceSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

type InterfaceAuthentication struct {
	KeyID    int    `json:"key-id,omitempty"`
	Method   string `json:"method,omitempty"`
	Password string `json:"password,omitempty"`
}

type InterfaceIntervalCfg struct {
	Interval   int `json:"interval,omitempty"`
	MinRx      int `json:"min-rx,omitempty"`
	Multiplier int `json:"multiplier,omitempty"`
}

type InterfaceLifIPAddressList struct {
	Anycast   int    `json:"anycast,omitempty"`
	Ipv6Addr  string `json:"ipv6-addr,omitempty"`
	LinkLocal int    `json:"link-local,omitempty"`
}

type InterfaceLifIpOspf struct {
	AuthenticationCfg InterfaceOspfGlobal   `json:"ospf-global,omitempty"`
	IPAddr            []InterfaceOspfIPList `json:"ospf-ip-list,omitempty"`
}

type InterfaceRip struct {
	Str           InterfaceLifIpripAuthentication `json:"authentication,omitempty"`
	Receive       InterfaceReceiveCfg             `json:"receive-cfg,omitempty"`
	ReceivePacket int                             `json:"receive-packet,omitempty"`
	Send          InterfaceSendCfg                `json:"send-cfg,omitempty"`
	SendPacket    int                             `json:"send-packet,omitempty"`
	State         InterfaceSplitHorizonCfg        `json:"split-horizon-cfg,omitempty"`
	UUID          string                          `json:"uuid,omitempty"`
}

type InterfaceLifIPRouter struct {
	Tag InterfaceLifIPRipIsis `json:"isis,omitempty"`
}

type InterfaceOspf struct {
	Bfd                int                              `json:"bfd,omitempty"`
	Cost               []InterfaceCostCfg               `json:"cost-cfg,omitempty"`
	DeadInterval       []InterfaceDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
	Disable            int                              `json:"disable,omitempty"`
	HelloInterval      []InterfaceHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	MtuIgnore          []InterfaceMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	Neighbor           []InterfaceNeighborCfg           `json:"neighbor-cfg,omitempty"`
	BroadcastType      []InterfaceNetworkList           `json:"network-list,omitempty"`
	Priority           []InterfacePriorityCfg           `json:"priority-cfg,omitempty"`
	RetransmitInterval []InterfaceRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	TransmitDelay      []InterfaceTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	UUID               string                           `json:"uuid,omitempty"`
}

type InterfaceRouter struct {
	Tag      InterfaceLifIPRipIsis      `json:"isis,omitempty"`
	AreaList InterfaceLifIpv6RouterOspf `json:"ospf,omitempty"`
	Rip      InterfaceRipng             `json:"ripng,omitempty"`
}

type InterfaceLifIsisAuthentication struct {
	KeyChain []InterfaceKeyChainList `json:"key-chain-list,omitempty"`
	Mode     []InterfaceModeList     `json:"mode-list,omitempty"`
	SendOnly []InterfaceSendOnlyList `json:"send-only-list,omitempty"`
}

type InterfaceBfdCfg struct {
	Bfd     int `json:"bfd,omitempty"`
	Disable int `json:"disable,omitempty"`
}

type InterfaceCsnpIntervalList struct {
	CsnpInterval int    `json:"csnp-interval,omitempty"`
	Level        string `json:"level,omitempty"`
}

type InterfaceHelloIntervalList struct {
	HelloInterval int    `json:"hello-interval,omitempty"`
	Level         string `json:"level,omitempty"`
}

type InterfaceHelloIntervalMinimalList struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal,omitempty"`
	Level                string `json:"level,omitempty"`
}

type InterfaceHelloMultiplierList struct {
	HelloMultiplier int    `json:"hello-multiplier,omitempty"`
	Level           string `json:"level,omitempty"`
}

type InterfaceMeshGroup struct {
	Blocked int `json:"blocked,omitempty"`
	Value   int `json:"value,omitempty"`
}

type InterfaceMetricList struct {
	Level  string `json:"level,omitempty"`
	Metric int    `json:"metric,omitempty"`
}

type InterfacePasswordList struct {
	Level    string `json:"level,omitempty"`
	Password string `json:"password,omitempty"`
}

type InterfacePriorityList struct {
	Level    string `json:"level,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type InterfaceWideMetricList struct {
	Level      string `json:"level,omitempty"`
	WideMetric int    `json:"wide-metric,omitempty"`
}

type InterfaceOspfGlobal struct {
	Authentication     InterfaceAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	AuthenticationKey  string                      `json:"authentication-key,omitempty"`
	Bfd                InterfaceBfdCfg             `json:"bfd-cfg,omitempty"`
	Cost               int                         `json:"cost,omitempty"`
	DatabaseFilter     InterfaceDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	DeadInterval       int                         `json:"dead-interval,omitempty"`
	Disable            string                      `json:"disable,omitempty"`
	HelloInterval      int                         `json:"hello-interval,omitempty"`
	MessageDigestKey   []InterfaceMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	Mtu                int                         `json:"mtu,omitempty"`
	MtuIgnore          int                         `json:"mtu-ignore,omitempty"`
	Broadcast          InterfaceNetwork            `json:"network,omitempty"`
	Priority           int                         `json:"priority,omitempty"`
	RetransmitInterval int                         `json:"retransmit-interval,omitempty"`
	TransmitDelay      int                         `json:"transmit-delay,omitempty"`
	UUID               string                      `json:"uuid,omitempty"`
}

type InterfaceOspfIPList struct {
	Authentication     int                         `json:"authentication,omitempty"`
	AuthenticationKey  string                      `json:"authentication-key,omitempty"`
	Cost               int                         `json:"cost,omitempty"`
	DatabaseFilter     string                      `json:"database-filter,omitempty"`
	DeadInterval       int                         `json:"dead-interval,omitempty"`
	HelloInterval      int                         `json:"hello-interval,omitempty"`
	IPAddr             string                      `json:"ip-addr,omitempty"`
	MessageDigestKey   []InterfaceMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	MtuIgnore          int                         `json:"mtu-ignore,omitempty"`
	Out                int                         `json:"out,omitempty"`
	Priority           int                         `json:"priority,omitempty"`
	RetransmitInterval int                         `json:"retransmit-interval,omitempty"`
	TransmitDelay      int                         `json:"transmit-delay,omitempty"`
	UUID               string                      `json:"uuid,omitempty"`
	Value              string                      `json:"value,omitempty"`
}

type InterfaceLifIpripAuthentication struct {
	KeyChain InterfaceKeyChain `json:"key-chain,omitempty"`
	Mode     InterfaceMode     `json:"mode,omitempty"`
	String   InterfaceStr      `json:"str,omitempty"`
}

type InterfaceReceiveCfg struct {
	Receive int    `json:"receive,omitempty"`
	Version string `json:"version,omitempty"`
}

type InterfaceSendCfg struct {
	Send    int    `json:"send,omitempty"`
	Version string `json:"version,omitempty"`
}

type InterfaceSplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}

type InterfaceLifIPRipIsis struct {
	Tag  string `json:"tag,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type InterfaceCostCfg struct {
	Cost       int `json:"cost,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}

type InterfaceDeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval,omitempty"`
	InstanceID   int `json:"instance-id,omitempty"`
}

type InterfaceHelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}

type InterfaceMtuIgnoreCfg struct {
	InstanceID int `json:"instance-id,omitempty"`
	MtuIgnore  int `json:"mtu-ignore,omitempty"`
}

type InterfaceNeighborCfg struct {
	NeigInst             int    `json:"neig-inst,omitempty"`
	Neighbor             string `json:"neighbor,omitempty"`
	NeighborCost         int    `json:"neighbor-cost,omitempty"`
	NeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	NeighborPriority     int    `json:"neighbor-priority,omitempty"`
}

type InterfaceNetworkList struct {
	BroadcastType     string `json:"broadcast-type,omitempty"`
	NetworkInstanceID int    `json:"network-instance-id,omitempty"`
	P2MpNbma          int    `json:"p2mp-nbma,omitempty"`
}

type InterfacePriorityCfg struct {
	InstanceID int `json:"instance-id,omitempty"`
	Priority   int `json:"priority,omitempty"`
}

type InterfaceRetransmitIntervalCfg struct {
	InstanceID         int `json:"instance-id,omitempty"`
	RetransmitInterval int `json:"retransmit-interval,omitempty"`
}

type InterfaceTransmitDelayCfg struct {
	InstanceID    int `json:"instance-id,omitempty"`
	TransmitDelay int `json:"transmit-delay,omitempty"`
}

type InterfaceLifIpv6RouterOspf struct {
	AreaIDNum []InterfaceAreaList `json:"area-list,omitempty"`
	UUID      string              `json:"uuid,omitempty"`
}

type InterfaceRipng struct {
	Rip  int    `json:"rip,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type InterfaceKeyChainList struct {
	KeyChain string `json:"key-chain,omitempty"`
	Level    string `json:"level,omitempty"`
}

type InterfaceModeList struct {
	Level string `json:"level,omitempty"`
	Mode  string `json:"mode,omitempty"`
}

type InterfaceSendOnlyList struct {
	Level    string `json:"level,omitempty"`
	SendOnly int    `json:"send-only,omitempty"`
}

type InterfaceAuthenticationCfg struct {
	Authentication int    `json:"authentication,omitempty"`
	Value          string `json:"value,omitempty"`
}

type InterfaceDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter,omitempty"`
	Out            int    `json:"out,omitempty"`
}

type InterfaceMessageDigestCfg struct {
	Md5Value         string `json:"md5-value,omitempty"`
	MessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceNetwork struct {
	Broadcast         int `json:"broadcast,omitempty"`
	NonBroadcast      int `json:"non-broadcast,omitempty"`
	P2MpNbma          int `json:"p2mp-nbma,omitempty"`
	PointToMultipoint int `json:"point-to-multipoint,omitempty"`
	PointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceKeyChain struct {
	KeyChain string `json:"key-chain,omitempty"`
}

type InterfaceMode struct {
	Mode string `json:"mode,omitempty"`
}

type InterfaceStr struct {
	String string `json:"string,omitempty"`
}

type InterfaceAreaList struct {
	AreaIDAddr string `json:"area-id-addr,omitempty"`
	AreaIDNum  int    `json:"area-id-num,omitempty"`
	InstanceID int    `json:"instance-id,omitempty"`
	Tag        string `json:"tag,omitempty"`
}

func PostInterfaceLif(id string, inst InterfaceLif, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceLif")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/lif", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceLif", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceLif(id string, name1 string, host string) (*InterfaceLif, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceLif")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/lif/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceLif", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutInterfaceLif(id string, name1 string, inst InterfaceLif, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutInterfaceLif")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/lif/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutInterfaceLif", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteInterfaceLif(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteInterfaceLif")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/interface/lif/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteInterfaceLif", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
