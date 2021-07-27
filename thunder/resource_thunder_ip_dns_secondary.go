package thunder

//Thunder resource IpDnsSecondary

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpDnsSecondary() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpDnsSecondaryCreate,
		UpdateContext: resourceIpDnsSecondaryUpdate,
		ReadContext:   resourceIpDnsSecondaryRead,
		DeleteContext: resourceIpDnsSecondaryDelete,
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

func resourceIpDnsSecondaryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpDnsSecondary (Inside resourceIpDnsSecondaryCreate) ")

		data := dataToIpDnsSecondary(d)
		logger.Println("[INFO] received V from method data to IpDnsSecondary --")
		d.SetId("1")
		err := go_thunder.PostIpDnsSecondary(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpDnsSecondaryRead(ctx, d, meta)

	}
	return diags
}

func resourceIpDnsSecondaryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpDnsSecondary (Inside resourceIpDnsSecondaryRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpDnsSecondary(client.Token, client.Host)
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

func resourceIpDnsSecondaryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpDnsSecondaryRead(ctx, d, meta)
}

func resourceIpDnsSecondaryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpDnsSecondaryRead(ctx, d, meta)
}
func dataToIpDnsSecondary(d *schema.ResourceData) go_thunder.DnsSecondary {
	var vc go_thunder.DnsSecondary
	var c go_thunder.DnsSecondaryInstance
	c.IPV4Addr = d.Get("ip_v4_addr").(string)
	c.IPV6Addr = d.Get("ip_v6_addr").(string)

	vc.UUID = c
	return vc
}
