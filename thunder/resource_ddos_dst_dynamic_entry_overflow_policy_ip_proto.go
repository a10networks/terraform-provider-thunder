package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDynamicEntryOverflowPolicyIpProto() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_dynamic_entry_overflow_policy_ip_proto`: DDOS IP protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstDynamicEntryOverflowPolicyIpProtoCreate,
		UpdateContext: resourceDdosDstDynamicEntryOverflowPolicyIpProtoUpdate,
		ReadContext:   resourceDdosDstDynamicEntryOverflowPolicyIpProtoRead,
		DeleteContext: resourceDdosDstDynamicEntryOverflowPolicyIpProtoDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Protocol Number",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"other": {
							Type: schema.TypeString, Optional: true, Description: "DDOS other template",
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
			"default_address_type": {
				Type: schema.TypeString, Required: true, Description: "DefaultAddressType",
			},
		},
	}
}
func resourceDdosDstDynamicEntryOverflowPolicyIpProtoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyIpProtoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicyIpProto(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDynamicEntryOverflowPolicyIpProtoRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstDynamicEntryOverflowPolicyIpProtoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyIpProtoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicyIpProto(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDynamicEntryOverflowPolicyIpProtoRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstDynamicEntryOverflowPolicyIpProtoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyIpProtoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicyIpProto(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstDynamicEntryOverflowPolicyIpProtoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyIpProtoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicyIpProto(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstDynamicEntryOverflowPolicyIpProtoTemplate(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyIpProtoTemplate {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyIpProtoTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Other = in["other"].(string)
	}
	return ret
}

func dataToEndpointDdosDstDynamicEntryOverflowPolicyIpProto(d *schema.ResourceData) edpt.DdosDstDynamicEntryOverflowPolicyIpProto {
	var ret edpt.DdosDstDynamicEntryOverflowPolicyIpProto
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Template = getObjectDdosDstDynamicEntryOverflowPolicyIpProtoTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	return ret
}
