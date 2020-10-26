package thunder

//Thunder resource SlbPerf

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbPerf() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbPerfCreate,
		Update: resourceSlbPerfUpdate,
		Read:   resourceSlbPerfRead,
		Delete: resourceSlbPerfDelete,
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

func resourceSlbPerfCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPerf (Inside resourceSlbPerfCreate) ")

		data := dataToSlbPerf(d)
		logger.Println("[INFO] received formatted data from method data to SlbPerf --")
		d.SetId("1")
		go_thunder.PostSlbPerf(client.Token, data, client.Host)

		return resourceSlbPerfRead(d, meta)

	}
	return nil
}

func resourceSlbPerfRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbPerf (Inside resourceSlbPerfRead)")

	if client.Host != "" {
		name := d.Id()

		data, err := go_thunder.GetSlbPerf(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbPerfUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPerfRead(d, meta)
}

func resourceSlbPerfDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPerfRead(d, meta)
}

func dataToSlbPerf(d *schema.ResourceData) go_thunder.Perf {
	var vc go_thunder.Perf
	var c go_thunder.PerfInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable25, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable25
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
