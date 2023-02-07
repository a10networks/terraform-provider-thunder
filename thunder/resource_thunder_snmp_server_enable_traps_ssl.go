package thunder

//Thunder resource SnmpServerEnableTrapsSsl

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSsl() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsSslCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSslUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSslRead,
		DeleteContext: resourceSnmpServerEnableTrapsSslDelete,
		Schema: map[string]*schema.Schema{
			"server_certificate_error": {
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

func resourceSnmpServerEnableTrapsSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSsl (Inside resourceSnmpServerEnableTrapsSslCreate) ")

		data := dataToSnmpServerEnableTrapsSsl(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSsl --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsSsl(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsSslRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSsl (Inside resourceSnmpServerEnableTrapsSslRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSsl(client.Token, client.Host)
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

func resourceSnmpServerEnableTrapsSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSslRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSslRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsSsl(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsSsl {
	var vc go_thunder.SnmpServerEnableTrapsSsl
	var c go_thunder.SnmpServerEnableTrapsSslInstance
	c.ServerCertificateError = d.Get("server_certificate_error").(int)

	vc.ServerCertificateError = c
	return vc
}
