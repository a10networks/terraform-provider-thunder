package thunder

//Thunder resource SystemVeMacScheme

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemVeMacScheme() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemVeMacSchemeCreate,
		UpdateContext: resourceSystemVeMacSchemeUpdate,
		ReadContext:   resourceSystemVeMacSchemeRead,
		DeleteContext: resourceSystemVeMacSchemeDelete,
		Schema: map[string]*schema.Schema{
			"ve_mac_scheme_val": {
				Type:        schema.TypeString,
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

func resourceSystemVeMacSchemeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SystemVeMacScheme (Inside resourceSystemVeMacSchemeCreate) ")

		data := dataToSystemVeMacScheme(d)
		logger.Println("[INFO] received formatted data from method data to SystemVeMacScheme --")
		d.SetId("1")
		err := go_thunder.PostSystemVeMacScheme(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSystemVeMacSchemeRead(ctx, d, meta)

	}
	return diags
}

func resourceSystemVeMacSchemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SystemVeMacScheme (Inside resourceSystemVeMacSchemeRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSystemVeMacScheme(client.Token, client.Host)
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

func resourceSystemVeMacSchemeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSystemVeMacSchemeRead(ctx, d, meta)
}

func resourceSystemVeMacSchemeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSystemVeMacSchemeRead(ctx, d, meta)
}
func dataToSystemVeMacScheme(d *schema.ResourceData) go_thunder.SystemVeMacScheme {
	var vc go_thunder.SystemVeMacScheme
	var c go_thunder.SystemVeMacSchemeInstance
	c.VeMacSchemeVal = d.Get("ve_mac_scheme_val").(string)

	vc.VeMacSchemeVal = c
	return vc
}
