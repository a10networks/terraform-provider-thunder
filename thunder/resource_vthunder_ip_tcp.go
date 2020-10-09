package thunder

//Thunder resource IpTcp

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpTcp() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpTcpCreate,
		Update: resourceIpTcpUpdate,
		Read:   resourceIpTcpRead,
		Delete: resourceIpTcpDelete,
		Schema: map[string]*schema.Schema{
			"syn_cookie": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpTcpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println(d)
	if client.Host != "" {
		logger.Println("[INFO] Creating IpTcp (Inside resourceIpTcpCreate) ")

		data := dataToIpTcp(d)
		logger.Println("[INFO] received V from method data to IpTcp --")
		d.SetId("1")
		go_thunder.PostIpTcp(client.Token, data, client.Host)

		return resourceIpTcpRead(d, meta)

	}
	return nil
}

func resourceIpTcpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpTcp (Inside resourceIpTcpRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpTcp(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}
func resourceIpTcpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpTcpRead(d, meta)
}

func resourceIpTcpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpTcpRead(d, meta)
}

func dataToIpTcp(d *schema.ResourceData) go_thunder.IpTcp {
	var vc go_thunder.IpTcp
	var c go_thunder.IpTcpInstance
	logger := util.GetLoggerInstance()
	logger.Println(d.Get("syn_cookie"))
	var obj1 go_thunder.SynCookie
	obj1.Threshold = d.Get("syn_cookie.0.threshold").(int)
	//obj1.Threshold = 4
	c.Threshold = obj1

	vc.UUID = c
	return vc
}
