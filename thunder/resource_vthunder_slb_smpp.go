package thunder

//Thunder resource SlbSmpp

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSmpp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSmppCreate,
		Update: resourceSlbSmppUpdate,
		Read:   resourceSlbSmppRead,
		Delete: resourceSlbSmppDelete,
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

func resourceSlbSmppCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSmpp (Inside resourceSlbSmppCreate) ")

		data := dataToSlbSmpp(d)
		logger.Println("[INFO] received formatted data from method data to SlbSmpp --")
		d.SetId("1")
		go_thunder.PostSlbSmpp(client.Token, data, client.Host)

		return resourceSlbSmppRead(d, meta)

	}
	return nil
}

func resourceSlbSmppRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbSmpp (Inside resourceSlbSmppRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSmpp(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSmppUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSmppRead(d, meta)
}

func resourceSlbSmppDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSmppRead(d, meta)
}

func dataToSlbSmpp(d *schema.ResourceData) go_thunder.SlbSmpp {
	var vc go_thunder.SlbSmpp
	var c go_thunder.SlbSmppInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable11, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable11
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
