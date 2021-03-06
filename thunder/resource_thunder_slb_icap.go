package thunder

//Thunder resource SlbIcapHTTP

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbIcap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbIcapCreate,
		Update: resourceSlbIcapUpdate,
		Read:   resourceSlbIcapRead,
		Delete: resourceSlbIcapDelete,
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

func resourceSlbIcapCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbIcap (Inside resourceSlbIcapCreate) ")

		data := dataToSlbIcap(d)
		logger.Println("[INFO] received V from method data to SlbIcap --")
		d.SetId("1")
		go_thunder.PostSlbIcap(client.Token, data, client.Host)

		return resourceSlbIcapRead(d, meta)

	}
	return nil
}

func resourceSlbIcapRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbIcap (Inside resourceSlbIcapRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbIcap(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}
func resourceSlbIcapUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbIcapRead(d, meta)
}

func resourceSlbIcapDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbIcapRead(d, meta)
}
func dataToSlbIcap(d *schema.ResourceData) go_thunder.Icap {
	var vc go_thunder.Icap
	var c go_thunder.IcapInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable39, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable39
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
