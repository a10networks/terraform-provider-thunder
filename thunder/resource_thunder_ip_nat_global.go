package thunder

//Thunder resource IpNatGlobal

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatGlobal() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpNatGlobalCreate,
		UpdateContext: resourceIpNatGlobalUpdate,
		ReadContext:   resourceIpNatGlobalRead,
		DeleteContext: resourceIpNatGlobalDelete,
		Schema: map[string]*schema.Schema{
			"reset_idle_tcp_conn": {
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

func resourceIpNatGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpNatGlobal (Inside resourceIpNatGlobalCreate) ")

		data := dataToIpNatGlobal(d)
		logger.Println("[INFO] received V from method data to IpNatGlobal --")
		d.SetId("1")
		err := go_thunder.PostIpNatGlobal(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpNatGlobalRead(ctx, d, meta)

	}
	return diags
}

func resourceIpNatGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpNatGlobal (Inside resourceIpNatGlobalRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpNatGlobal(client.Token, client.Host)
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

func resourceIpNatGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpNatGlobalRead(ctx, d, meta)
}

func resourceIpNatGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpNatGlobalRead(ctx, d, meta)
}
func dataToIpNatGlobal(d *schema.ResourceData) go_thunder.NatGlobal {
	var vc go_thunder.NatGlobal
	var c go_thunder.NatGlobalInstance

	c.ResetIdleTcpConn = d.Get("reset_idle_tcp_conn").(int)

	vc.UUID = c
	return vc
}
