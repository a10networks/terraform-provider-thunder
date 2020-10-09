package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type EthernetIPv6 struct {
	UUID EthernetIPv6Instance `json:"ipv6,omitempty"`
}

type EthernetIPv6AddressList struct {
	AddressType string `json:"address-type,omitempty"`
	Ipv6Addr    string `json:"ipv6-addr,omitempty"`
}
type EthernetIPv6SplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}
type EthernetIPv6Rip struct {
	State EthernetIPv6SplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	UUID  string                      `json:"uuid,omitempty"`
}
type EthernetIPv6StatefulFirewall struct {
	UUID       string `json:"uuid,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	ACLName    string `json:"acl-name,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	AccessList int    `json:"access-list,omitempty"`
}
type EthernetIPv6Ripng struct {
	UUID string `json:"uuid,omitempty"`
	Rip  int    `json:"rip,omitempty"`
}
type EthernetIPv6AreaList struct {
	AreaIDAddr string `json:"area-id-addr,omitempty"`
	Tag        string `json:"tag,omitempty"`
	InstanceID int    `json:"instance-id,omitempty"`
	AreaIDNum  int    `json:"area-id-num,omitempty"`
}
type EthernetIPv6Ospf1 struct {
	AreaIDAddr []EthernetIPv6AreaList `json:"area-list,omitempty"`
	UUID       string                 `json:"uuid,omitempty"`
}
type EthernetIPv6Isis struct {
	Tag  string `json:"tag,omitempty"`
	UUID string `json:"uuid,omitempty"`
}
type EthernetIPv6Router struct {
	Rip        EthernetIPv6Ripng `json:"ripng,omitempty"`
	AreaIDAddr EthernetIPv6Ospf1 `json:"ospf,omitempty"`
	Tag        EthernetIPv6Isis  `json:"isis,omitempty"`
}
type EthernetIPv6AccessListCfg struct {
	Inbound   int    `json:"inbound,omitempty"`
	V6ACLName string `json:"v6-acl-name,omitempty"`
}
type EthernetIPv6CostCfg struct {
	Cost       int `json:"cost,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type EthernetIPv6PriorityCfg struct {
	Priority   int `json:"priority,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type EthernetIPv6HelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type EthernetIPv6MtuIgnoreCfg struct {
	MtuIgnore  int `json:"mtu-ignore,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type EthernetIPv6RetransmitIntervalCfg struct {
	RetransmitInterval int `json:"retransmit-interval,omitempty"`
	InstanceID         int `json:"instance-id,omitempty"`
}
type EthernetIPv6TransmitDelayCfg struct {
	TransmitDelay int `json:"transmit-delay,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type EthernetIPv6NeighborCfg struct {
	NeighborPriority     int    `json:"neighbor-priority,omitempty"`
	NeigInst             int    `json:"neig-inst,omitempty"`
	NeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	NeighborCost         int    `json:"neighbor-cost,omitempty"`
	Neighbor             string `json:"neighbor,omitempty"`
}
type EthernetIPv6NetworkList struct {
	BroadcastType     string `json:"broadcast-type,omitempty"`
	P2MpNbma          int    `json:"p2mp-nbma,omitempty"`
	NetworkInstanceID int    `json:"network-instance-id,omitempty"`
}
type EthernetIPv6DeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval,omitempty"`
	InstanceID   int `json:"instance-id,omitempty"`
}
type EthernetIPv6Ospf struct {
	UUID               string                              `json:"uuid,omitempty"`
	Bfd                int                                 `json:"bfd,omitempty"`
	Cost               []EthernetIPv6CostCfg               `json:"cost-cfg,omitempty"`
	Priority           []EthernetIPv6PriorityCfg           `json:"priority-cfg,omitempty"`
	HelloInterval      []EthernetIPv6HelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	MtuIgnore          []EthernetIPv6MtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	RetransmitInterval []EthernetIPv6RetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	Disable            int                                 `json:"disable,omitempty"`
	TransmitDelay      []EthernetIPv6TransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	NeighborPriority   []EthernetIPv6NeighborCfg           `json:"neighbor-cfg,omitempty"`
	BroadcastType      []EthernetIPv6NetworkList           `json:"network-list,omitempty"`
	DeadInterval       []EthernetIPv6DeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
}
type EthernetIPv6PrefixList struct {
	NotAutonomous     int    `json:"not-autonomous,omitempty"`
	ValidLifetime     int    `json:"valid-lifetime,omitempty"`
	NotOnLink         int    `json:"not-on-link,omitempty"`
	Prefix            string `json:"prefix,omitempty"`
	PreferredLifetime int    `json:"preferred-lifetime,omitempty"`
}
type EthernetIPv6RouterAdver struct {
	MaxInterval              int                      `json:"max-interval,omitempty"`
	DefaultLifetime          int                      `json:"default-lifetime,omitempty"`
	ReachableTime            int                      `json:"reachable-time,omitempty"`
	OtherConfigAction        string                   `json:"other-config-action,omitempty"`
	FloatingIPDefaultVrid    string                   `json:"floating-ip-default-vrid,omitempty"`
	ManagedConfigAction      string                   `json:"managed-config-action,omitempty"`
	MinInterval              int                      `json:"min-interval,omitempty"`
	RateLimit                int                      `json:"rate-limit,omitempty"`
	AdverMtuDisable          int                      `json:"adver-mtu-disable,omitempty"`
	NotAutonomous            []EthernetIPv6PrefixList `json:"prefix-list,omitempty"`
	FloatingIP               string                   `json:"floating-ip,omitempty"`
	AdverVrid                int                      `json:"adver-vrid,omitempty"`
	UseFloatingIPDefaultVrid int                      `json:"use-floating-ip-default-vrid,omitempty"`
	Action                   string                   `json:"action,omitempty"`
	AdverVridDefault         int                      `json:"adver-vrid-default,omitempty"`
	AdverMtu                 int                      `json:"adver-mtu,omitempty"`
	RetransmitTimer          int                      `json:"retransmit-timer,omitempty"`
	HopLimit                 int                      `json:"hop-limit,omitempty"`
	UseFloatingIP            int                      `json:"use-floating-ip,omitempty"`
}
type EthernetIPv6Instance struct {
	UUID        string                       `json:"uuid,omitempty"`
	AddressType []EthernetIPv6AddressList    `json:"address-list,omitempty"`
	Inside      int                          `json:"inside,omitempty"`
	Ipv6Enable  int                          `json:"ipv6-enable,omitempty"`
	State       EthernetIPv6Rip              `json:"rip,omitempty"`
	Outside     int                          `json:"outside,omitempty"`
	ClassList   EthernetIPv6StatefulFirewall `json:"stateful-firewall,omitempty"`
	TTLIgnore   int                          `json:"ttl-ignore,omitempty"`
	Rip         EthernetIPv6Router           `json:"router,omitempty"`
	Inbound     EthernetIPv6AccessListCfg    `json:"access-list-cfg,omitempty"`
	Bfd         EthernetIPv6Ospf             `json:"ospf,omitempty"`
	MaxInterval EthernetIPv6RouterAdver      `json:"router-adver,omitempty"`
}

func PostInterfaceEthernetIPv6(id string, name int, inst EthernetIPv6, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernetIPv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+strconv.Itoa(name)+"/ipv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetIPv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetInterfaceEthernetIPv6(id string, name string, host string) (*EthernetIPv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernetIPv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name+"/ipv6", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetIPv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}
