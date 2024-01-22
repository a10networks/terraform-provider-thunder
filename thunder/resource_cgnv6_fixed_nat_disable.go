package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatDisable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_fixed_nat_disable`: Disable fixed-nat configuration (Operation)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6FixedNatDisableCreate,
		UpdateContext: resourceCgnv6FixedNatDisableUpdate,
		ReadContext:   resourceCgnv6FixedNatDisableRead,
		DeleteContext: resourceCgnv6FixedNatDisableDelete,

		Schema: map[string]*schema.Schema{
			"clear_session": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Clear all sessions",
			},
			"inside_end_v4address": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Inside User End Address",
			},
			"inside_end_v6address": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Inside User End Address",
			},
			"inside_start_v4address": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Inside User Start Address",
			},
			"inside_start_v6address": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Inside User Start Address",
			},
			"ip_list": {
				Type: schema.TypeString, Optional: true, Description: "Name of IP List used to specify Inside Users",
			},
			"partition": {
				Type: schema.TypeString, Optional: true, Description: "Inside User Partition (Partition Name)",
			},
			"v4_netmask": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Netmask",
			},
			"v6_netmask": {
				Type: schema.TypeInt, Optional: true, Description: "Inside User IPv6 Netmask",
			},
		},
	}
}
func resourceCgnv6FixedNatDisableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatDisableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatDisable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatDisableRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6FixedNatDisableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatDisableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatDisable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatDisableRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6FixedNatDisableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatDisableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatDisable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6FixedNatDisableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatDisableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatDisable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6FixedNatDisable(d *schema.ResourceData) edpt.Cgnv6FixedNatDisable {
	var ret edpt.Cgnv6FixedNatDisable
	ret.Inst.ClearSession = d.Get("clear_session").(int)
	ret.Inst.InsideEndV4address = d.Get("inside_end_v4address").(string)
	ret.Inst.InsideEndV6address = d.Get("inside_end_v6address").(string)
	ret.Inst.InsideStartV4address = d.Get("inside_start_v4address").(string)
	ret.Inst.InsideStartV6address = d.Get("inside_start_v6address").(string)
	ret.Inst.IpList = d.Get("ip_list").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.V4Netmask = d.Get("v4_netmask").(string)
	ret.Inst.V6Netmask = d.Get("v6_netmask").(int)
	return ret
}
