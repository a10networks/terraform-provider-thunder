package thunder

//Thunder resource IpDnsPrimary

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpDnsPrimary() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpDnsPrimaryCreate,
		UpdateContext: resourceIpDnsPrimaryUpdate,
		ReadContext:   resourceIpDnsPrimaryRead,
		DeleteContext: resourceIpDnsPrimaryDelete,
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

func resourceIpDnsPrimaryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpDnsPrimary (Inside resourceIpDnsPrimaryCreate) ")

		data := dataToIpDnsPrimary(d)
		logger.Println("[INFO] received V from method data to IpDnsPrimary --")
		d.SetId("1")
		err := go_thunder.PostIpDnsPrimary(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpDnsPrimaryRead(ctx, d, meta)

	}
	return diags
}

func resourceIpDnsPrimaryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpDnsPrimary (Inside resourceIpDnsPrimaryRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpDnsPrimary(client.Token, client.Host)
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

func resourceIpDnsPrimaryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpDnsPrimaryRead(ctx, d, meta)
}

func resourceIpDnsPrimaryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpDnsPrimaryRead(ctx, d, meta)
}
func dataToIpDnsPrimary(d *schema.ResourceData) go_thunder.DnsPrimary {
	var vc go_thunder.DnsPrimary
	var c go_thunder.DnsPrimaryInstance
	c.IPV4Addr = d.Get("ip_v4_addr").(string)
	c.IPV6Addr = d.Get("ip_v6_addr").(string)

	vc.UUID = c
	return vc
}
