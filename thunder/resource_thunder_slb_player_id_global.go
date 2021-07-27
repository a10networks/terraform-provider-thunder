package thunder

//Thunder resource SlbPlayerIdGlobal

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPlayerIdGlobal() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbPlayerIdGlobalCreate,
		UpdateContext: resourceSlbPlayerIdGlobalUpdate,
		ReadContext:   resourceSlbPlayerIdGlobalRead,
		DeleteContext: resourceSlbPlayerIdGlobalDelete,
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

func resourceSlbPlayerIdGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPlayerIdGlobal (Inside resourceSlbPlayerIdGlobalCreate) ")

		data := dataToSlbPlayerIdGlobal(d)
		logger.Println("[INFO] received formatted data from method data to SlbPlayerIdGlobal --")
		d.SetId("1")
		err := go_thunder.PostSlbPlayerIdGlobal(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbPlayerIdGlobalRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbPlayerIdGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbPlayerIdGlobal (Inside resourceSlbPlayerIdGlobalRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbPlayerIdGlobal(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbPlayerIdGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPlayerIdGlobalRead(ctx, d, meta)
}

func resourceSlbPlayerIdGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPlayerIdGlobalRead(ctx, d, meta)
}

func dataToSlbPlayerIdGlobal(d *schema.ResourceData) go_thunder.PlayerIdGlobal {
	var vc go_thunder.PlayerIdGlobal
	var c go_thunder.PlayerIdGlobalInstance
	c.AbsMaxExpiration = d.Get("abs_max_expiration").(int)
	c.PktActivityExpiration = d.Get("pkt_activity_expiration").(int)
	c.Enable64BitPlayerID = d.Get("enable_64bit_player_id").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable27, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable27
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
