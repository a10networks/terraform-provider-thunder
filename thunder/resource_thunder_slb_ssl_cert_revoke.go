package thunder

//Thunder resource SlbSSLCertRevoke

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSSLCertRevoke() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSSLCertRevokeCreate,
		Update: resourceSlbSSLCertRevokeUpdate,
		Read:   resourceSlbSSLCertRevokeRead,
		Delete: resourceSlbSSLCertRevokeDelete,
		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSSLCertRevokeCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSSLCertRevoke (Inside resourceSlbSSLCertRevokeCreate) ")

		data := dataToSlbSSLCertRevoke(d)
		logger.Println("[INFO] received V from method data to SlbSSLCertRevoke --")
		d.SetId("1")
		go_thunder.PostSlbSSLCertRevoke(client.Token, data, client.Host)

		return resourceSlbSSLCertRevokeRead(d, meta)

	}
	return nil
}

func resourceSlbSSLCertRevokeRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbSSLCertRevoke (Inside resourceSlbSSLCertRevokeRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbSSLCertRevoke(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSSLCertRevokeUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSSLCertRevokeRead(d, meta)
}

func resourceSlbSSLCertRevokeDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSSLCertRevokeRead(d, meta)
}

func dataToSlbSSLCertRevoke(d *schema.ResourceData) go_thunder.SSLCertRevoke {
	var vc go_thunder.SSLCertRevoke
	var c go_thunder.SSLCertRevokeInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable37, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable37
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
