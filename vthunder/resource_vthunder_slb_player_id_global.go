package vthunder

//vThunder resource SlbPlayerIdGlobal

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbPlayerIdGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbPlayerIdGlobalCreate,
		Update: resourceSlbPlayerIdGlobalUpdate,
		Read:   resourceSlbPlayerIdGlobalRead,
		Delete: resourceSlbPlayerIdGlobalDelete,
		Schema: map[string]*schema.Schema{
			"min_expiration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pkt_activity_expiration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_64bit_player_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"abs_max_expiration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
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
			"force_passive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enforcement_timer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbPlayerIdGlobalCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPlayerIdGlobal (Inside resourceSlbPlayerIdGlobalCreate) ")

		data := dataToSlbPlayerIdGlobal(d)
		logger.Println("[INFO] received formatted data from method data to SlbPlayerIdGlobal --")
		d.SetId("1")
		go_vthunder.PostSlbPlayerIdGlobal(client.Token, data, client.Host)

		return resourceSlbPlayerIdGlobalRead(d, meta)

	}
	return nil
}

func resourceSlbPlayerIdGlobalRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbPlayerIdGlobal (Inside resourceSlbPlayerIdGlobalRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbPlayerIdGlobal(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbPlayerIdGlobalUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPlayerIdGlobalRead(d, meta)
}

func resourceSlbPlayerIdGlobalDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPlayerIdGlobalRead(d, meta)
}

func dataToSlbPlayerIdGlobal(d *schema.ResourceData) go_vthunder.PlayerIdGlobal {
	var vc go_vthunder.PlayerIdGlobal
	var c go_vthunder.PlayerIdGlobalInstance
	c.AbsMaxExpiration = d.Get("abs_max_expiration").(int)
	c.PktActivityExpiration = d.Get("pkt_activity_expiration").(int)
	c.Enable64BitPlayerID = d.Get("enable_64bit_player_id").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable27, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable27
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.MinExpiration = d.Get("min_expiration").(int)
	c.ForcePassive = d.Get("force_passive").(int)
	c.EnforcementTimer = d.Get("enforcement_timer").(int)

	vc.Counters1 = c
	return vc
}
