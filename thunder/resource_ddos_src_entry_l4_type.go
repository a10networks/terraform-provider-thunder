package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcEntryL4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_entry_l4_type`: DDOS L4 type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcEntryL4TypeCreate,
		UpdateContext: resourceDdosSrcEntryL4TypeUpdate,
		ReadContext:   resourceDdosSrcEntryL4TypeRead,
		DeleteContext: resourceDdosSrcEntryL4TypeDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'permit': Whitelist incoming packets for protocol; 'deny': Blacklist incoming packets for protocol;",
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
			"src_entry_name": {
				Type: schema.TypeString, Required: true, Description: "SrcEntryName",
			},
		},
	}
}
func resourceDdosSrcEntryL4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryL4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryL4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcEntryL4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcEntryL4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryL4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryL4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcEntryL4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcEntryL4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryL4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryL4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcEntryL4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryL4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryL4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosSrcEntryL4TypeTemplate(d []interface{}) edpt.DdosSrcEntryL4TypeTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcEntryL4TypeTemplate
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

func dataToEndpointDdosSrcEntryL4Type(d *schema.ResourceData) edpt.DdosSrcEntryL4Type {
	var ret edpt.DdosSrcEntryL4Type
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Template = getObjectDdosSrcEntryL4TypeTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SrcEntryName = d.Get("src_entry_name").(string)
	return ret
}
