package thunder

//Thunder resource Ipv6Icmpv6

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6Icmpv6() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpv6Icmpv6Create,
		UpdateContext: resourceIpv6Icmpv6Update,
		ReadContext:   resourceIpv6Icmpv6Read,
		DeleteContext: resourceIpv6Icmpv6Delete,
		Schema: map[string]*schema.Schema{
			"redirect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"unreachable": {
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

func resourceIpv6Icmpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating Ipv6Icmpv6 (Inside resourceIpv6Icmpv6Create) ")

		data := dataToIpv6Icmpv6(d)
		logger.Println("[INFO] received V from method data to Ipv6Icmpv6 --")
		d.SetId("1")
		err := go_thunder.PostIpv6Icmpv6(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpv6Icmpv6Read(ctx, d, meta)

	}
	return diags
}

func resourceIpv6Icmpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Ipv6Icmpv6 (Inside resourceIpv6Icmpv6Read)")

	if client.Host != "" {
		data, err := go_thunder.GetIpv6Icmpv6(client.Token, client.Host)
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

func resourceIpv6Icmpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpv6Icmpv6Read(ctx, d, meta)
}

func resourceIpv6Icmpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpv6Icmpv6Read(ctx, d, meta)
}
func dataToIpv6Icmpv6(d *schema.ResourceData) go_thunder.Icmpv6 {
	var vc go_thunder.Icmpv6
	var c go_thunder.Icmpv6Instance
	c.Redirect = d.Get("redirect").(int)
	c.Unreachable = d.Get("unreachable").(int)

	vc.UUID = c
	return vc
}
