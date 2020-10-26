package thunder

//Thunder resource SlbProxy

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbProxyCreate,
		Update: resourceSlbProxyUpdate,
		Read:   resourceSlbProxyRead,
		Delete: resourceSlbProxyDelete,
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

func resourceSlbProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbProxy (Inside resourceSlbProxyCreate) ")

		data := dataToSlbProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbProxy --")
		d.SetId("1")
		go_thunder.PostSlbProxy(client.Token, data, client.Host)

		return resourceSlbProxyRead(d, meta)

	}
	return nil
}

func resourceSlbProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbProxy (Inside resourceSlbProxyRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbProxy(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbProxyRead(d, meta)
}

func resourceSlbProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbProxyRead(d, meta)
}

func dataToSlbProxy(d *schema.ResourceData) go_thunder.Proxy {
	var vc go_thunder.Proxy
	var c go_thunder.ProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable29, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable29
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
