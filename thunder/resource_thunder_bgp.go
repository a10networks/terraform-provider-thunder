package thunder

//Thunder resource Bgp

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceBgp() *schema.Resource {
	return &schema.Resource{
		Create: resourceBgpCreate,
		Update: resourceBgpUpdate,
		Read:   resourceBgpRead,
		Delete: resourceBgpDelete,
		Schema: map[string]*schema.Schema{
			"extended_asn_cap": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"nexthop_trigger": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"delay": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceBgpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating Bgp (Inside resourceBgpCreate) ")

		data := dataToBgp(d)
		logger.Println("[INFO] received formatted data from method data to Bgp --")
		d.SetId("1")
		go_thunder.PostBgp(client.Token, data, client.Host)

		return resourceBgpRead(d, meta)

	}
	return nil
}

func resourceBgpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading Bgp (Inside resourceBgpRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetBgp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceBgpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceBgpRead(d, meta)
}

func resourceBgpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceBgpRead(d, meta)
}
func dataToBgp(d *schema.ResourceData) go_thunder.Bgp1 {
	var vc go_thunder.Bgp1
	var c go_thunder.BgpInstance
	c.ExtendedAsnCap = d.Get("extended_asn_cap").(int)

	var obj1 go_thunder.NexthopTrigger
	prefix := "nexthop_trigger.0."
	obj1.Enable = d.Get(prefix + "enable").(int)
	obj1.Delay = d.Get(prefix + "delay").(int)

	c.Enable = obj1

	vc.ExtendedAsnCap = c
	return vc
}
