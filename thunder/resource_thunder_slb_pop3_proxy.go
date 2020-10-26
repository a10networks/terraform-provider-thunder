package thunder

//Thunder resource SlbPop3Proxy

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbPop3Proxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbPop3ProxyCreate,
		Update: resourceSlbPop3ProxyUpdate,
		Read:   resourceSlbPop3ProxyRead,
		Delete: resourceSlbPop3ProxyDelete,
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

func resourceSlbPop3ProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPop3Proxy (Inside resourceSlbPop3ProxyCreate) ")

		data := dataToSlbPop3Proxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbPop3Proxy --")
		d.SetId("1")
		go_thunder.PostSlbPop3Proxy(client.Token, data, client.Host)

		return resourceSlbPop3ProxyRead(d, meta)

	}
	return nil
}

func resourceSlbPop3ProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbPop3Proxy (Inside resourceSlbPop3ProxyRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbPop3Proxy(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbPop3ProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPop3ProxyRead(d, meta)
}

func resourceSlbPop3ProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPop3ProxyRead(d, meta)
}

func dataToSlbPop3Proxy(d *schema.ResourceData) go_thunder.Pop3Proxy {
	var vc go_thunder.Pop3Proxy
	var c go_thunder.Pop3ProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable28, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable28
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
