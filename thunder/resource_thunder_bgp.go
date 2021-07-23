package thunder

//Thunder resource Bgp

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceBgp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBgpCreate,
		UpdateContext: resourceBgpUpdate,
		ReadContext:   resourceBgpRead,
		DeleteContext: resourceBgpDelete,
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

func resourceBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating Bgp (Inside resourceBgpCreate) ")

		data := dataToBgp(d)
		logger.Println("[INFO] received formatted data from method data to Bgp --")
		d.SetId("1")
		err := go_thunder.PostBgp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceBgpRead(ctx, d, meta)

	}
	return diags
}

func resourceBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Bgp (Inside resourceBgpRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetBgp(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceBgpRead(ctx, d, meta)
}

func resourceBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceBgpRead(ctx, d, meta)
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
