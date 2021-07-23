package thunder

//Thunder resource SnmpServerEnableTrapsSnmp

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSnmp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsSnmpCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSnmpUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSnmpRead,
		DeleteContext: resourceSnmpServerEnableTrapsSnmpDelete,
		Schema: map[string]*schema.Schema{
			"linkup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"linkdown": {
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

func resourceSnmpServerEnableTrapsSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSnmp (Inside resourceSnmpServerEnableTrapsSnmpCreate) ")

		data := dataToSnmpServerEnableTrapsSnmp(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSnmp --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsSnmp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsSnmpRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSnmp (Inside resourceSnmpServerEnableTrapsSnmpRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSnmp(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSnmpRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSnmpRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsSnmp(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsSnmp {
	var vc go_thunder.SnmpServerEnableTrapsSnmp
	var c go_thunder.SnmpServerEnableTrapsSnmpInstance
	c.Linkup = d.Get("linkup").(int)
	c.All = d.Get("all").(int)
	c.Linkdown = d.Get("linkdown").(int)

	vc.Linkup = c
	return vc
}
