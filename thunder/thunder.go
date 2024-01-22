package thunder

import (
	"fmt"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	endpnt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/clarketm/json"
)

type Thunder struct {
	Host      string
	User      string
	Password  string
	Token     string
	Partition string
	log       *axapi.ThunderLog
}

type Credentials struct {
	Inst struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
}

type AuthResponse struct {
	Inst struct {
		Signature   string `json:"signature"`
		Description string `json:"description"`
	} `json:"authresponse"`
}

func (t *Thunder) getAuthToken(endpoint string) error {
	c := Credentials{}
	c.Inst.Username = t.User
	c.Inst.Password = t.Password
	payloadBytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Failed to create auth payload")
	}

	var headers = axapi.GenRequestHeaderNoToken()
	_, axapiResp, err := axapi.SendPost(t.Host, endpoint, payloadBytes, headers, t.log)
	if err == nil {
		var m AuthResponse
		err2 := json.Unmarshal(axapiResp, &m)
		if err2 != nil {
			t.log.Println("Failed to unmarshal auth response")
			return err2
		}
		t.Token = "A10 " + m.Inst.Signature
	}
	return err
}

func (t *Thunder) Connect() error {
	t.log = axapi.CreateLogger(t.Partition)
	if t.Host == "" || t.User == "" || t.Password == "" {
		t.log.Println("Failed to connect THUNDER device: no required arguments")
		return fmt.Errorf("THUNDER provider requires address, username and password")
	}
	t.log.Println("Connect to THUNDER device: host="+t.Host+", partition="+t.Partition, "user="+t.User)
	err := t.getAuthToken("auth")
	if err != nil {
		t.log.Printf("Failed to creating New Token Session %s ", err)
		return fmt.Errorf("Failed to creating New Token Session %s ", err)
	}
	t.log.Println("Connected")
	return nil
}

func (t *Thunder) callPartition() error {
	if t.Partition != "shared" {
		p := endpnt.ActivePartition{}
		p.Inst.ChangeTo = t.Partition
		err := p.ChangeTo(t.Token, t.Host, t.log)
		if err != nil {
			t.log.Println("Fail to switch patition: " + t.Partition)
			return err
		}
	}
	t.log.Println("current partition: " + t.Partition)
	return nil
}
