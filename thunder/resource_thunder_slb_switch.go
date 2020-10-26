package thunder

//Thunder resource SlbSwitch

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSwitchCreate,
		Update: resourceSlbSwitchUpdate,
		Read:   resourceSlbSwitchRead,
		Delete: resourceSlbSwitchDelete,
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

func resourceSlbSwitchCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSwitch (Inside resourceSlbSwitchCreate) ")
		d.SetId("1")
		data := dataToSlbSwitch(d)
		logger.Println("[INFO] received formatted data from method data to SlbSwitch --")
		go_thunder.PostSlbSwitch(client.Token, data, client.Host)

		return resourceSlbSwitchRead(d, meta)

	}
	return nil
}

func resourceSlbSwitchRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbSwitch (Inside resourceSlbSwitchRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbSwitch(client.Token, client.Host)
		if data == nil {
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSwitchUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSwitchRead(d, meta)
}

func resourceSlbSwitchDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSwitchRead(d, meta)
}

func dataToSlbSwitch(d *schema.ResourceData) go_thunder.Switch {
	var vc go_thunder.Switch
	var c go_thunder.SwitchInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable13, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable13
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
