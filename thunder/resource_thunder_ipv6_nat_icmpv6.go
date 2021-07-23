package thunder

//Thunder resource Ipv6NatIcmpv6

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatIcmpv6() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpv6NatIcmpv6Create,
		UpdateContext: resourceIpv6NatIcmpv6Update,
		ReadContext:   resourceIpv6NatIcmpv6Read,
		DeleteContext: resourceIpv6NatIcmpv6Delete,
		Schema: map[string]*schema.Schema{
			"respond_to_ping": {
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

func resourceIpv6NatIcmpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating Ipv6NatIcmpv6 (Inside resourceIpv6NatIcmpv6Create) ")

		data := dataToIpv6NatIcmpv6(d)
		logger.Println("[INFO] received V from method data to Ipv6NatIcmpv6 --")
		d.SetId("1")
		err := go_thunder.PostIpv6NatIcmpv6(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpv6NatIcmpv6Read(ctx, d, meta)

	}
	return diags
}

func resourceIpv6NatIcmpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Ipv6NatIcmpv6 (Inside resourceIpv6NatIcmpv6Read)")

	if client.Host != "" {
		data, err := go_thunder.GetIpv6NatIcmpv6(client.Token, client.Host)
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

func resourceIpv6NatIcmpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpv6NatIcmpv6Read(ctx, d, meta)
}

func resourceIpv6NatIcmpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpv6NatIcmpv6Read(ctx, d, meta)
}
func dataToIpv6NatIcmpv6(d *schema.ResourceData) go_thunder.NatIcmpv6 {
	var vc go_thunder.NatIcmpv6
	var c go_thunder.NatIcmpv6Instance
	c.RespondToPing = d.Get("respond_to_ping").(int)

	vc.UUID = c
	return vc
}
