package thunder

//Thunder resource SlL7session

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlL7session() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbL7sessionCreate,
		Update: resourceSlbL7sessionUpdate,
		Read:   resourceSlbL7sessionRead,
		Delete: resourceSlbL7sessionDelete,
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

func resourceSlbL7sessionCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlL7session (Inside resourceSlL7sessionCreate) ")
		data := dataToSlL7session(d)
		logger.Println("[INFO] received V from method data to SlL7session --")
		d.SetId("1")
		go_thunder.PostSlbL7session(client.Token, data, client.Host)

		return resourceSlbL7sessionRead(d, meta)

	}
	return nil
}

func resourceSlbL7sessionRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlL7session (Inside resourceSlL7sessionRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbL7session(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbL7sessionUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbL7sessionRead(d, meta)
}

func resourceSlbL7sessionDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbL7sessionRead(d, meta)
}

func dataToSlL7session(d *schema.ResourceData) go_thunder.L7session {
	var vc go_thunder.L7session
	var c go_thunder.L7sessionInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable43, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable43
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
