package thunder

//Thunder resource SlbSvmSourceNat

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSvmSourceNat() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSvmSourceNatCreate,
		UpdateContext: resourceSlbSvmSourceNatUpdate,
		ReadContext:   resourceSlbSvmSourceNatRead,
		DeleteContext: resourceSlbSvmSourceNatDelete,
		Schema: map[string]*schema.Schema{
			"pool": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbSvmSourceNatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSvmSourceNat (Inside resourceSlbSvmSourceNatCreate) ")

		data := dataToSlbSvmSourceNat(d)
		logger.Println("[INFO] received formatted data from method data to SlbSvmSourceNat --")
		d.SetId("1")
		err := go_thunder.PostSlbSvmSourceNat(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSvmSourceNatRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSvmSourceNatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSvmSourceNat (Inside resourceSlbSvmSourceNatRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSvmSourceNat(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbSvmSourceNatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSvmSourceNatRead(ctx, d, meta)
}

func resourceSlbSvmSourceNatDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSvmSourceNatRead(ctx, d, meta)
}

func dataToSlbSvmSourceNat(d *schema.ResourceData) go_thunder.SvmSourceNat {
	var vc go_thunder.SvmSourceNat
	var c go_thunder.SvmSourceNatInstance
	c.Pool = d.Get("pool").(string)

	vc.Pool = c
	return vc
}
