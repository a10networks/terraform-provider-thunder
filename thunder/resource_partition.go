package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_partition`: Create/unload a Network partition\n\n__PLACEHOLDER__",
		CreateContext: resourcePartitionCreate,
		UpdateContext: resourcePartitionUpdate,
		ReadContext:   resourcePartitionRead,
		DeleteContext: resourcePartitionDelete,

		Schema: map[string]*schema.Schema{
			"application_type": {
				Type: schema.TypeString, Optional: true, Description: "'adc': Application type ADC; 'cgnv6': Application type CGNv6;",
			},
			"id1": {
				Type: schema.TypeInt, Optional: true, Description: "Specify unique Partition id",
			},
			"partition_name": {
				Type: schema.TypeString, Required: true, Description: "Object partition name",
			},
			"shared_vlan": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
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
						"vrid": {
							Type: schema.TypeInt, Optional: true, Description: "Specify VRRP-A vrid",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_accounting": {
							Type: schema.TypeString, Optional: true, Description: "Attach a resource accounting template (Name of the template)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourcePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionRead(ctx, d, meta)
	}
	return diags
}

func resourcePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePartitionRead(ctx, d, meta)
	}
	return diags
}
func resourcePartitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePartitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectPartitionSharedVlan1089(d []interface{}) edpt.PartitionSharedVlan1089 {

	count1 := len(d)
	var ret edpt.PartitionSharedVlan1089
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vlan = in["vlan"].(int)
		ret.AllowableIpRange = getSlicePartitionSharedVlanAllowableIpRange1090(in["allowable_ip_range"].([]interface{}))
		ret.AllowableIpv6Range = getSlicePartitionSharedVlanAllowableIpv6Range1091(in["allowable_ipv6_range"].([]interface{}))
		ret.MgmtFloatingIpAddress = in["mgmt_floating_ip_address"].(string)
		ret.Vrid = in["vrid"].(int)
		//omit uuid
	}
	return ret
}

func getSlicePartitionSharedVlanAllowableIpRange1090(d []interface{}) []edpt.PartitionSharedVlanAllowableIpRange1090 {

	count1 := len(d)
	ret := make([]edpt.PartitionSharedVlanAllowableIpRange1090, 0, count1)
	var oi edpt.PartitionSharedVlanAllowableIpRange1090
	ret = append(ret, oi)
	return ret
}

func getSlicePartitionSharedVlanAllowableIpv6Range1091(d []interface{}) []edpt.PartitionSharedVlanAllowableIpv6Range1091 {

	count1 := len(d)
	ret := make([]edpt.PartitionSharedVlanAllowableIpv6Range1091, 0, count1)
	var oi edpt.PartitionSharedVlanAllowableIpv6Range1091
	ret = append(ret, oi)
	return ret
}

func getObjectPartitionTemplate1092(d []interface{}) edpt.PartitionTemplate1092 {

	count1 := len(d)
	var ret edpt.PartitionTemplate1092
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ResourceAccounting = in["resource_accounting"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointPartition(d *schema.ResourceData) edpt.Partition {
	var ret edpt.Partition
	ret.Inst.ApplicationType = d.Get("application_type").(string)
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.SharedVlan = getObjectPartitionSharedVlan1089(d.Get("shared_vlan").([]interface{}))
	ret.Inst.Template = getObjectPartitionTemplate1092(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
