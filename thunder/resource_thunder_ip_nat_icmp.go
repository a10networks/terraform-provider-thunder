package thunder

//Thunder resource IpNatIcmp

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatIcmp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpNatIcmpCreate,
		UpdateContext: resourceIpNatIcmpUpdate,
		ReadContext:   resourceIpNatIcmpRead,
		DeleteContext: resourceIpNatIcmpDelete,
		Schema: map[string]*schema.Schema{
			"respond_to_ping": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"always_source_nat_errors": {
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

func resourceIpNatIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpNatIcmp (Inside resourceIpNatIcmpCreate) ")

		data := dataToIpNatIcmp(d)
		logger.Println("[INFO] received V from method data to IpNatIcmp --")
		d.SetId("1")
		err := go_thunder.PostIpNatIcmp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpNatIcmpRead(ctx, d, meta)

	}
	return diags
}

func resourceIpNatIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpNatIcmp (Inside resourceIpNatIcmpRead)")

	if client.Host != "" {

		data, err := go_thunder.GetIpNatIcmp(client.Token, client.Host)
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

func resourceIpNatIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpNatIcmpRead(ctx, d, meta)
}

func resourceIpNatIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpNatIcmpRead(ctx, d, meta)
}
func dataToIpNatIcmp(d *schema.ResourceData) go_thunder.NatIcmp {
	var vc go_thunder.NatIcmp
	var c go_thunder.NatIcmpInstance
	c.RespondToPing = d.Get("respond_to_ping").(int)
	c.AlwaysSourceNatErrors = d.Get("always_source_nat_errors").(int)

	vc.UUID = c
	return vc
}
