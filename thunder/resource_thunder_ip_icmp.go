package thunder

//Thunder resource IpIcmp

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpIcmp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpIcmpCreate,
		UpdateContext: resourceIpIcmpUpdate,
		ReadContext:   resourceIpIcmpRead,
		DeleteContext: resourceIpIcmpDelete,
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

func resourceIpIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpIcmp (Inside resourceIpIcmpCreate) ")

		data := dataToIpIcmp(d)
		logger.Println("[INFO] received V from method data to IpIcmp --")
		d.SetId("1")
		err := go_thunder.PostIpIcmp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpIcmpRead(ctx, d, meta)

	}
	return diags
}

func resourceIpIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpIcmp (Inside resourceIpIcmpRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpIcmp(client.Token, client.Host)
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

func resourceIpIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpIcmpRead(ctx, d, meta)
}

func resourceIpIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpIcmpRead(ctx, d, meta)
}
func dataToIpIcmp(d *schema.ResourceData) go_thunder.Icmp {
	var vc go_thunder.Icmp
	var c go_thunder.IcmpInstance
	c.Redirect = d.Get("redirect").(int)
	c.Unreachable = d.Get("unreachable").(int)

	vc.UUID = c
	return vc
}
