package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIcmpV4Filter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_icmp_v4_filter`: ICMP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIcmpV4FilterCreate,
		UpdateContext: resourceDdosZoneTemplateIcmpV4FilterUpdate,
		ReadContext:   resourceDdosZoneTemplateIcmpV4FilterRead,
		DeleteContext: resourceDdosZoneTemplateIcmpV4FilterDelete,

		Schema: map[string]*schema.Schema{
			"byte_offset_filter": {
				Type: schema.TypeString, Optional: true, Description: "filter using Berkeley packet filter syntax",
			},
			"icmp_filter_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src;",
			},
			"icmp_filter_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
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
func resourceDdosZoneTemplateIcmpV4FilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4FilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Filter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4FilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4FilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4FilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Filter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4FilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIcmpV4FilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4FilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Filter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4FilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Filter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateIcmpV4Filter(d *schema.ResourceData) edpt.DdosZoneTemplateIcmpV4Filter {
	var ret edpt.DdosZoneTemplateIcmpV4Filter
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
