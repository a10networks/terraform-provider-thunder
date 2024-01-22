package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatRangeList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_range_list`: IP Source NAT Static range list\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatRangeListCreate,
		UpdateContext: resourceCgnv6NatRangeListUpdate,
		ReadContext:   resourceCgnv6NatRangeListRead,
		DeleteContext: resourceCgnv6NatRangeListDelete,

		Schema: map[string]*schema.Schema{
			"global_netmaskv4": {
				Type: schema.TypeString, Optional: true, Description: "Mask for this Address range",
			},
			"global_start_ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Global Start IPv4 Address of this list",
			},
			"local_netmaskv4": {
				Type: schema.TypeString, Optional: true, Description: "Mask for this Address range",
			},
			"local_start_ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Local Start IPv4 Address of this list",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name for this Static List",
			},
			"partition": {
				Type: schema.TypeString, Required: true, Description: "Inside User Partition (Partition Name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_count": {
				Type: schema.TypeInt, Optional: true, Description: "Number of addresses to be translated in this range",
			},
			"v4_vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceCgnv6NatRangeListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatRangeListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatRangeList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatRangeListRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatRangeListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatRangeListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatRangeList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatRangeListRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatRangeListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatRangeListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatRangeList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatRangeListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatRangeListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatRangeList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6NatRangeList(d *schema.ResourceData) edpt.Cgnv6NatRangeList {
	var ret edpt.Cgnv6NatRangeList
	ret.Inst.GlobalNetmaskv4 = d.Get("global_netmaskv4").(string)
	ret.Inst.GlobalStartIpv4Addr = d.Get("global_start_ipv4_addr").(string)
	ret.Inst.LocalNetmaskv4 = d.Get("local_netmaskv4").(string)
	ret.Inst.LocalStartIpv4Addr = d.Get("local_start_ipv4_addr").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	//omit uuid
	ret.Inst.V4Count = d.Get("v4_count").(int)
	ret.Inst.V4Vrid = d.Get("v4_vrid").(int)
	return ret
}
