package thunder

//Thunder resource SnmpServerEnableTrapsVrrpA

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsVrrpA() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsVrrpACreate,
		UpdateContext: resourceSnmpServerEnableTrapsVrrpAUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsVrrpARead,
		DeleteContext: resourceSnmpServerEnableTrapsVrrpADelete,
		Schema: map[string]*schema.Schema{
			"active": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"standby": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsVrrpACreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsVrrpA (Inside resourceSnmpServerEnableTrapsVrrpACreate) ")

		data := dataToSnmpServerEnableTrapsVrrpA(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsVrrpA --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsVrrpA(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsVrrpARead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsVrrpARead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsVrrpA (Inside resourceSnmpServerEnableTrapsVrrpARead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsVrrpA(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsVrrpAUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsVrrpARead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsVrrpADelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsVrrpARead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsVrrpA(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsVrrpA {
	var vc go_thunder.SnmpServerEnableTrapsVrrpA
	var c go_thunder.SnmpServerEnableTrapsVrrpAInstance
	c.Active = d.Get("active").(int)
	c.Standby = d.Get("standby").(int)
	c.All = d.Get("all").(int)

	vc.Active = c
	return vc
}
