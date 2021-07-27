package thunder

//Thunder resource SnmpServerEnableTrapsVcs

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsVcs() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsVcsCreate,
		UpdateContext: resourceSnmpServerEnableTrapsVcsUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsVcsRead,
		DeleteContext: resourceSnmpServerEnableTrapsVcsDelete,
		Schema: map[string]*schema.Schema{
			"state_change": {
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

func resourceSnmpServerEnableTrapsVcsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsVcs (Inside resourceSnmpServerEnableTrapsVcsCreate) ")

		data := dataToSnmpServerEnableTrapsVcs(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsVcs --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsVcs(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsVcsRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsVcsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsVcs (Inside resourceSnmpServerEnableTrapsVcsRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsVcs(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsVcsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsVcsRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsVcsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsVcsRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsVcs(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsVcs {
	var vc go_thunder.SnmpServerEnableTrapsVcs
	var c go_thunder.SnmpServerEnableTrapsVcsInstance
	c.StateChange = d.Get("state_change").(int)

	vc.StateChange = c
	return vc
}
