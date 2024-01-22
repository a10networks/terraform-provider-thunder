package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcDynamicEntryOverflowPolicyL4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_dynamic_entry_overflow_policy_l4_type`: DDOS L4 type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcDynamicEntryOverflowPolicyL4TypeCreate,
		UpdateContext: resourceDdosSrcDynamicEntryOverflowPolicyL4TypeUpdate,
		ReadContext:   resourceDdosSrcDynamicEntryOverflowPolicyL4TypeRead,
		DeleteContext: resourceDdosSrcDynamicEntryOverflowPolicyL4TypeDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
						},
						"udp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
						},
						"other": {
							Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
						},
						"template_icmp_v4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
						},
						"template_icmp_v6": {
							Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
func resourceDdosSrcDynamicEntryOverflowPolicyL4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyL4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicyL4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcDynamicEntryOverflowPolicyL4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcDynamicEntryOverflowPolicyL4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyL4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicyL4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcDynamicEntryOverflowPolicyL4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcDynamicEntryOverflowPolicyL4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyL4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicyL4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcDynamicEntryOverflowPolicyL4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyL4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicyL4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosSrcDynamicEntryOverflowPolicyL4TypeTemplate(d []interface{}) edpt.DdosSrcDynamicEntryOverflowPolicyL4TypeTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcDynamicEntryOverflowPolicyL4TypeTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func dataToEndpointDdosSrcDynamicEntryOverflowPolicyL4Type(d *schema.ResourceData) edpt.DdosSrcDynamicEntryOverflowPolicyL4Type {
	var ret edpt.DdosSrcDynamicEntryOverflowPolicyL4Type
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Template = getObjectDdosSrcDynamicEntryOverflowPolicyL4TypeTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	return ret
}
