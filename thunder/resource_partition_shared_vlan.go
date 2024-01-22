package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartitionSharedVlan() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_partition_shared_vlan`: Enable shared-vlan feature in Network partition\n\n__PLACEHOLDER__",
		CreateContext: resourcePartitionSharedVlanCreate,
		UpdateContext: resourcePartitionSharedVlanUpdate,
		ReadContext:   resourcePartitionSharedVlanRead,
		DeleteContext: resourcePartitionSharedVlanDelete,

		Schema: map[string]*schema.Schema{
			"allowable_ip_range": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"allowable_ipv6_range": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"mgmt_floating_ip_address": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Address for Shared VLAN Mgmt IP Address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Specify VRRP-A vrid",
			},
			"partition_name": {
				Type: schema.TypeString, Required: true, Description: "PartitionName",
			},
		},
	}
}
func resourcePartitionSharedVlanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionSharedVlanCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionSharedVlan(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionSharedVlanRead(ctx, d, meta)
	}
	return diags
}

func resourcePartitionSharedVlanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionSharedVlanUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionSharedVlan(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionSharedVlanRead(ctx, d, meta)
	}
	return diags
}
func resourcePartitionSharedVlanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionSharedVlanDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionSharedVlan(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePartitionSharedVlanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionSharedVlanRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionSharedVlan(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSlicePartitionSharedVlanAllowableIpRange(d []interface{}) []edpt.PartitionSharedVlanAllowableIpRange {

	count1 := len(d)
	ret := make([]edpt.PartitionSharedVlanAllowableIpRange, 0, count1)
	var oi edpt.PartitionSharedVlanAllowableIpRange
	ret = append(ret, oi)
	return ret
}

func getSlicePartitionSharedVlanAllowableIpv6Range(d []interface{}) []edpt.PartitionSharedVlanAllowableIpv6Range {

	count1 := len(d)
	ret := make([]edpt.PartitionSharedVlanAllowableIpv6Range, 0, count1)
	var oi edpt.PartitionSharedVlanAllowableIpv6Range
	ret = append(ret, oi)
	return ret
}

func dataToEndpointPartitionSharedVlan(d *schema.ResourceData) edpt.PartitionSharedVlan {
	var ret edpt.PartitionSharedVlan
	ret.Inst.AllowableIpRange = getSlicePartitionSharedVlanAllowableIpRange(d.Get("allowable_ip_range").([]interface{}))
	ret.Inst.AllowableIpv6Range = getSlicePartitionSharedVlanAllowableIpv6Range(d.Get("allowable_ipv6_range").([]interface{}))
	ret.Inst.MgmtFloatingIpAddress = d.Get("mgmt_floating_ip_address").(string)
	//omit uuid
	ret.Inst.Vlan = d.Get("vlan").(int)
	ret.Inst.Vrid = d.Get("vrid").(int)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	return ret
}
