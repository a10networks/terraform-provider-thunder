package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTrunkLoadBalance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_trunk_load_balance`: Configure Trunk load balancing options\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTrunkLoadBalanceCreate,
		UpdateContext: resourceSystemTrunkLoadBalanceUpdate,
		ReadContext:   resourceSystemTrunkLoadBalanceRead,
		DeleteContext: resourceSystemTrunkLoadBalanceDelete,

		Schema: map[string]*schema.Schema{
			"use_l3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer-3 Header based load balancing",
			},
			"use_l4": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer-3/4 Header based load balancing",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemTrunkLoadBalanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkLoadBalanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkLoadBalance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTrunkLoadBalanceRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTrunkLoadBalanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkLoadBalanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkLoadBalance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTrunkLoadBalanceRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTrunkLoadBalanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkLoadBalanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkLoadBalance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTrunkLoadBalanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkLoadBalanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkLoadBalance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTrunkLoadBalance(d *schema.ResourceData) edpt.SystemTrunkLoadBalance {
	var ret edpt.SystemTrunkLoadBalance
	ret.Inst.UseL3 = d.Get("use_l3").(int)
	ret.Inst.UseL4 = d.Get("use_l4").(int)
	//omit uuid
	return ret
}
