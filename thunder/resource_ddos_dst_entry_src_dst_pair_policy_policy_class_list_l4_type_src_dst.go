package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_l4_type_src_dst`: DDOS L4 type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstDelete,

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
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "ClassListName",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstTemplate
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

func dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst {
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
