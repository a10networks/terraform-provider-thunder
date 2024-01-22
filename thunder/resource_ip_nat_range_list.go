package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatRangeList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_range_list`: IP Source NAT Static range list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatRangeListCreate,
		UpdateContext: resourceIpNatRangeListUpdate,
		ReadContext:   resourceIpNatRangeListRead,
		DeleteContext: resourceIpNatRangeListDelete,

		Schema: map[string]*schema.Schema{
			"global_netmaskv4": {
				Type: schema.TypeString, Optional: true, Description: "Mask for this Address range",
			},
			"global_start_ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Global Start IPv4 Address of this list",
			},
			"global_start_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Global Start IPv6 Address of this list",
			},
			"local_netmaskv4": {
				Type: schema.TypeString, Optional: true, Description: "Mask for this Address range",
			},
			"local_start_ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Local Start IPv4 Address of this list",
			},
			"local_start_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Local Start IPv6 Address of this list",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name for this Static List",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_acl_id": {
				Type: schema.TypeInt, Optional: true, Description: "Access list ID",
			},
			"v4_acl_name": {
				Type: schema.TypeString, Optional: true, Description: "Access list name",
			},
			"v4_count": {
				Type: schema.TypeInt, Optional: true, Description: "Number of addresses to be translated in this range",
			},
			"v4_vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
			"v6_acl_name": {
				Type: schema.TypeString, Optional: true, Description: "Access list name",
			},
			"v6_count": {
				Type: schema.TypeInt, Optional: true, Description: "Number of addresses to be translated in this range",
			},
			"v6_vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceIpNatRangeListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatRangeListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatRangeList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatRangeListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatRangeListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatRangeListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatRangeList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatRangeListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatRangeListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatRangeListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatRangeList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatRangeListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatRangeListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatRangeList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatRangeList(d *schema.ResourceData) edpt.IpNatRangeList {
	var ret edpt.IpNatRangeList
	ret.Inst.GlobalNetmaskv4 = d.Get("global_netmaskv4").(string)
	ret.Inst.GlobalStartIpv4Addr = d.Get("global_start_ipv4_addr").(string)
	ret.Inst.GlobalStartIpv6Addr = d.Get("global_start_ipv6_addr").(string)
	ret.Inst.LocalNetmaskv4 = d.Get("local_netmaskv4").(string)
	ret.Inst.LocalStartIpv4Addr = d.Get("local_start_ipv4_addr").(string)
	ret.Inst.LocalStartIpv6Addr = d.Get("local_start_ipv6_addr").(string)
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	ret.Inst.V4AclId = d.Get("v4_acl_id").(int)
	ret.Inst.V4AclName = d.Get("v4_acl_name").(string)
	ret.Inst.V4Count = d.Get("v4_count").(int)
	ret.Inst.V4Vrid = d.Get("v4_vrid").(int)
	ret.Inst.V6AclName = d.Get("v6_acl_name").(string)
	ret.Inst.V6Count = d.Get("v6_count").(int)
	ret.Inst.V6Vrid = d.Get("v6_vrid").(int)
	return ret
}
