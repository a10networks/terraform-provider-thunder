package thunder

//Thunder resource SlbSportRateLimit

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSportRateLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSportRateLimitCreate,
		Update: resourceSlbSportRateLimitUpdate,
		Read:   resourceSlbSportRateLimitRead,
		Delete: resourceSlbSportRateLimitDelete,
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

func resourceSlbSportRateLimitCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSportRateLimit (Inside resourceSlbSportRateLimitCreate) ")

		data := dataToSlbSportRateLimit(d)
		logger.Println("[INFO] received formatted data from method data to SlbSportRateLimit --")
		d.SetId("1")
		go_thunder.PostSlbSportRateLimit(client.Token, data, client.Host)

		return resourceSlbSportRateLimitRead(d, meta)

	}
	return nil
}

func resourceSlbSportRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbSportRateLimit (Inside resourceSlbSportRateLimitRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSportRateLimit(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSportRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSportRateLimitRead(d, meta)
}

func resourceSlbSportRateLimitDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSportRateLimitRead(d, meta)
}

func dataToSlbSportRateLimit(d *schema.ResourceData) go_thunder.SportRateLimit {
	var vc go_thunder.SportRateLimit
	var c go_thunder.SportRateLimitInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable15, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable15
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
