package thunder

//Thunder resource SnmpServerDisableTraps

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerDisableTraps() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerDisableTrapsCreate,
		UpdateContext: resourceSnmpServerDisableTrapsUpdate,
		ReadContext:   resourceSnmpServerDisableTrapsRead,
		DeleteContext: resourceSnmpServerDisableTrapsDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"vrrp_a": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snmp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gslb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerDisableTrapsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerDisableTraps (Inside resourceSnmpServerDisableTrapsCreate) ")

		data := dataToSnmpServerDisableTraps(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerDisableTraps --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerDisableTraps(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerDisableTrapsRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerDisableTrapsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerDisableTraps (Inside resourceSnmpServerDisableTrapsRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerDisableTraps(client.Token, client.Host)
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

func resourceSnmpServerDisableTrapsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerDisableTrapsRead(ctx, d, meta)
}

func resourceSnmpServerDisableTrapsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerDisableTrapsRead(ctx, d, meta)
}
func dataToSnmpServerDisableTraps(d *schema.ResourceData) go_thunder.SnmpServerDisableTraps {
	var vc go_thunder.SnmpServerDisableTraps
	var c go_thunder.SnmpServerDisableTrapsInstance
	c.All = d.Get("all").(int)
	c.SlbChange = d.Get("slb_change").(int)
	c.VrrpA = d.Get("vrrp_a").(int)
	c.Snmp = d.Get("snmp").(int)
	c.Gslb = d.Get("gslb").(int)
	c.Slb = d.Get("slb").(int)

	vc.All = c
	return vc
}
