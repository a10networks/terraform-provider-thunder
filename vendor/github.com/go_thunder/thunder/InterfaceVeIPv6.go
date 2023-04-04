package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type VeIPv6 struct {
	UUID VeIPv6Instance `json:"ipv6,omitempty"`
}

type VeIPv6AddressList struct {
	AddressType string `json:"address-type,omitempty"`
	Ipv6Addr    string `json:"ipv6-addr,omitempty"`
}
type VeIPv6SplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}
type VeIPv6Rip struct {
	State VeIPv6SplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	UUID  string                `json:"uuid,omitempty"`
}
type VeIPv6StatefulFirewall struct {
	UUID       string `json:"uuid,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	ACLName    string `json:"acl-name,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	AccessList int    `json:"access-list,omitempty"`
}
type VeIPv6Ripng struct {
	UUID string `json:"uuid,omitempty"`
	Rip  int    `json:"rip,omitempty"`
}
type VeIPv6AreaList struct {
	AreaIDAddr string `json:"area-id-addr,omitempty"`
	Tag        string `json:"tag,omitempty"`
	InstanceID int    `json:"instance-id,omitempty"`
	AreaIDNum  int    `json:"area-id-num,omitempty"`
}
type VeIPv6Ospf1 struct {
	AreaIDAddr []VeIPv6AreaList `json:"area-list,omitempty"`
	UUID       string           `json:"uuid,omitempty"`
}
type VeIPv6Isis struct {
	Tag  string `json:"tag,omitempty"`
	UUID string `json:"uuid,omitempty"`
}
type VeIPv6Router struct {
	Rip        VeIPv6Ripng `json:"ripng,omitempty"`
	AreaIDAddr VeIPv6Ospf1 `json:"ospf,omitempty"`
	Tag        VeIPv6Isis  `json:"isis,omitempty"`
}
type VeIPv6CostCfg struct {
	Cost       int `json:"cost,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type VeIPv6PriorityCfg struct {
	Priority   int `json:"priority,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type VeIPv6HelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type VeIPv6MtuIgnoreCfg struct {
	MtuIgnore  int `json:"mtu-ignore,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type VeIPv6RetransmitIntervalCfg struct {
	RetransmitInterval int `json:"retransmit-interval,omitempty"`
	InstanceID         int `json:"instance-id,omitempty"`
}
type VeIPv6TransmitDelayCfg struct {
	TransmitDelay int `json:"transmit-delay,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type VeIPv6NeighborCfg struct {
	NeighborPriority     int    `json:"neighbor-priority,omitempty"`
	NeigInst             int    `json:"neig-inst,omitempty"`
	NeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	NeighborCost         int    `json:"neighbor-cost,omitempty"`
	Neighbor             string `json:"neighbor,omitempty"`
}
type VeIPv6NetworkList struct {
	BroadcastType     string `json:"broadcast-type,omitempty"`
	P2MpNbma          int    `json:"p2mp-nbma,omitempty"`
	NetworkInstanceID int    `json:"network-instance-id,omitempty"`
}
type VeIPv6DeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval,omitempty"`
	InstanceID   int `json:"instance-id,omitempty"`
}
type VeIPv6Ospf struct {
	UUID               string                        `json:"uuid,omitempty"`
	Bfd                int                           `json:"bfd,omitempty"`
	Cost               []VeIPv6CostCfg               `json:"cost-cfg,omitempty"`
	Priority           []VeIPv6PriorityCfg           `json:"priority-cfg,omitempty"`
	HelloInterval      []VeIPv6HelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	MtuIgnore          []VeIPv6MtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	RetransmitInterval []VeIPv6RetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	Disable            int                           `json:"disable,omitempty"`
	TransmitDelay      []VeIPv6TransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	NeighborPriority   []VeIPv6NeighborCfg           `json:"neighbor-cfg,omitempty"`
	BroadcastType      []VeIPv6NetworkList           `json:"network-list,omitempty"`
	DeadInterval       []VeIPv6DeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
}
type VeIPv6PrefixList struct {
	NotAutonomous     int    `json:"not-autonomous,omitempty"`
	ValidLifetime     int    `json:"valid-lifetime,omitempty"`
	NotOnLink         int    `json:"not-on-link,omitempty"`
	Prefix            string `json:"prefix,omitempty"`
	PreferredLifetime int    `json:"preferred-lifetime,omitempty"`
}
type VeIPv6RouterAdver struct {
	MaxInterval              int                `json:"max-interval,omitempty"`
	DefaultLifetime          int                `json:"default-lifetime,omitempty"`
	ReachableTime            int                `json:"reachable-time,omitempty"`
	OtherConfigAction        string             `json:"other-config-action,omitempty"`
	FloatingIPDefaultVrid    string             `json:"floating-ip-default-vrid,omitempty"`
	ManagedConfigAction      string             `json:"managed-config-action,omitempty"`
	MinInterval              int                `json:"min-interval,omitempty"`
	RateLimit                int                `json:"rate-limit,omitempty"`
	AdverMtuDisable          int                `json:"adver-mtu-disable,omitempty"`
	NotAutonomous            []VeIPv6PrefixList `json:"prefix-list,omitempty"`
	FloatingIP               string             `json:"floating-ip,omitempty"`
	AdverVrid                int                `json:"adver-vrid,omitempty"`
	UseFloatingIPDefaultVrid int                `json:"use-floating-ip-default-vrid,omitempty"`
	Action                   string             `json:"action,omitempty"`
	AdverVridDefault         int                `json:"adver-vrid-default,omitempty"`
	AdverMtu                 int                `json:"adver-mtu,omitempty"`
	RetransmitTimer          int                `json:"retransmit-timer,omitempty"`
	HopLimit                 int                `json:"hop-limit,omitempty"`
	UseFloatingIP            int                `json:"use-floating-ip,omitempty"`
}
type VeIPv6Instance struct {
	UUID        string                 `json:"uuid,omitempty"`
	Inbound     int                    `json:"inbound,omitempty"`
	AddressType []VeIPv6AddressList    `json:"address-list,omitempty"`
	Inside      int                    `json:"inside,omitempty"`
	Ipv6Enable  int                    `json:"ipv6-enable,omitempty"`
	State       VeIPv6Rip              `json:"rip,omitempty"`
	Outside     int                    `json:"outside,omitempty"`
	ClassList   VeIPv6StatefulFirewall `json:"stateful-firewall,omitempty"`
	V6ACLName   string                 `json:"v6-acl-name,omitempty"`
	TTLIgnore   int                    `json:"ttl-ignore,omitempty"`
	Router      VeIPv6Router           `json:"router,omitempty"`
	Bfd         VeIPv6Ospf             `json:"ospf,omitempty"`
	MaxInterval VeIPv6RouterAdver      `json:"router-adver,omitempty"`
}

func PostInterfaceVeIPv6(id string, name int, inst VeIPv6, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceVeIPv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ve/"+strconv.Itoa(name)+"/ipv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VeIPv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceVeIPv6 REQ RES..........................", m)
			err := check_api_status("PostInterfaceVeIPv6", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceVeIPv6(id string, name string, host string) (*VeIPv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceVeIPv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ve/"+name+"/ipv6", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VeIPv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceVeIPv6 REQ RES..........................", m)
			err := check_api_status("GetInterfaceVeIPv6", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
