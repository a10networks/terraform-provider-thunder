package thunder

//Thunder resource Vrrp common

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnsPrimary() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDnsPrimaryCreate,
		UpdateContext: resourceDnsPrimaryUpdate,
		ReadContext:   resourceDnsPrimaryRead,
		DeleteContext: resourceDnsPrimaryDelete,

		Schema: map[string]*schema.Schema{
			"ip_v6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_v4_addr": {
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

func resourceDnsPrimaryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating DnsPrimary (Inside resourceDnsPrimaryCreate)")

	if client.Host != "" {
		vc := dataToDnsPrimary(d)
		err := go_thunder.PostDnsPrimary(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("1")

		return resourceDnsPrimaryRead(ctx, d, meta)
	}
	return diags
}

func resourceDnsPrimaryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading DnsPrimary (Inside resourceDnsPrimaryRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		vc, err := go_thunder.GetDnsPrimary(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No DnsPrimary found")
			//d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceDnsPrimaryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceDnsPrimaryRead(ctx, d, meta)
}

func resourceDnsPrimaryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceDnsPrimaryRead(ctx, d, meta)
}

//Utility method to instantiate DnsPrimary Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToDnsPrimary(d *schema.ResourceData) go_thunder.Primary {
	var vc go_thunder.Primary

	var c go_thunder.PrimaryInstance

	c.IPV4Addr = d.Get("ip_v4_addr").(string)

	vc.IPV4Addr = c

	return vc
}
