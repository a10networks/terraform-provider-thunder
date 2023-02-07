package thunder

//Thunder resource IpDnsSuffix

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpDnsSuffix() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpDnsSuffixCreate,
		UpdateContext: resourceIpDnsSuffixUpdate,
		ReadContext:   resourceIpDnsSuffixRead,
		DeleteContext: resourceIpDnsSuffixDelete,
		Schema: map[string]*schema.Schema{
			"domain_name": {
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

func resourceIpDnsSuffixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpDnsSuffix (Inside resourceIpDnsSuffixCreate) ")

		data := dataToIpDnsSuffix(d)
		logger.Println("[INFO] received V from method data to IpDnsSuffix --")
		d.SetId("1")
		err := go_thunder.PostIpDnsSuffix(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpDnsSuffixRead(ctx, d, meta)

	}
	return diags
}

func resourceIpDnsSuffixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpDnsSuffix (Inside resourceIpDnsSuffixRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpDnsSuffix(client.Token, client.Host)
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

func resourceIpDnsSuffixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpDnsSuffixRead(ctx, d, meta)
}

func resourceIpDnsSuffixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpDnsSuffixRead(ctx, d, meta)
}
func dataToIpDnsSuffix(d *schema.ResourceData) go_thunder.DnsSuffix {
	var vc go_thunder.DnsSuffix
	var c go_thunder.DnsSuffixInstance
	c.DomainName = d.Get("domain_name").(string)

	vc.UUID = c
	return vc
}
