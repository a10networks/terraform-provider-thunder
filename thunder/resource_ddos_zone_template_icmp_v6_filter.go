package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIcmpV6Filter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_icmp_v6_filter`: ICMP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIcmpV6FilterCreate,
		UpdateContext: resourceDdosZoneTemplateIcmpV6FilterUpdate,
		ReadContext:   resourceDdosZoneTemplateIcmpV6FilterRead,
		DeleteContext: resourceDdosZoneTemplateIcmpV6FilterDelete,

		Schema: map[string]*schema.Schema{
			"byte_offset_filter": {
				Type: schema.TypeString, Optional: true, Description: "filter using Berkeley packet filter syntax",
			},
			"icmp_filter_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src;",
			},
			"icmp_filter_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "list to take",
			},
			"icmp_filter_inverse_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of matching",
			},
			"icmp_filter_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"icmp_filter_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"icmp_filter_seq": {
				Type: schema.TypeInt, Optional: true, Description: "sequence number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"icmp_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "IcmpTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateIcmpV6FilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6FilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6Filter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV6FilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV6FilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6FilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6Filter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV6FilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIcmpV6FilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6FilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6Filter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV6FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6FilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6Filter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateIcmpV6Filter(d *schema.ResourceData) edpt.DdosZoneTemplateIcmpV6Filter {
	var ret edpt.DdosZoneTemplateIcmpV6Filter
	ret.Inst.ByteOffsetFilter = d.Get("byte_offset_filter").(string)
	ret.Inst.IcmpFilterAction = d.Get("icmp_filter_action").(string)
	ret.Inst.IcmpFilterActionListName = d.Get("icmp_filter_action_list_name").(string)
	ret.Inst.IcmpFilterInverseMatch = d.Get("icmp_filter_inverse_match").(int)
	ret.Inst.IcmpFilterName = d.Get("icmp_filter_name").(string)
	ret.Inst.IcmpFilterRegex = d.Get("icmp_filter_regex").(string)
	ret.Inst.IcmpFilterSeq = d.Get("icmp_filter_seq").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	return ret
}
