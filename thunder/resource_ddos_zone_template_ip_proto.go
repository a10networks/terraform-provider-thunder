package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIpProto() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_ip_proto`: Ip-proto template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIpProtoCreate,
		UpdateContext: resourceDdosZoneTemplateIpProtoUpdate,
		ReadContext:   resourceDdosZoneTemplateIpProtoRead,
		DeleteContext: resourceDdosZoneTemplateIpProtoDelete,

		Schema: map[string]*schema.Schema{
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"other_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"other_filter_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"other_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"other_filter_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
						},
						"other_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"other_filter_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"filter_match_type": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Stop matching on drop/blacklist action; 'stop-on-first-match': Stop matching on first match;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS Ip-proto Template Name",
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
func resourceDdosZoneTemplateIpProtoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProto(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIpProtoRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIpProtoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProto(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIpProtoRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIpProtoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProto(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIpProtoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProto(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneTemplateIpProtoFilterList(d []interface{}) []edpt.DdosZoneTemplateIpProtoFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIpProtoFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIpProtoFilterList
		oi.OtherFilterName = in["other_filter_name"].(string)
		oi.OtherFilterSeq = in["other_filter_seq"].(int)
		oi.OtherFilterRegex = in["other_filter_regex"].(string)
		oi.OtherFilterInverseMatch = in["other_filter_inverse_match"].(int)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.OtherFilterActionListName = in["other_filter_action_list_name"].(string)
		oi.OtherFilterAction = in["other_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateIpProto(d *schema.ResourceData) edpt.DdosZoneTemplateIpProto {
	var ret edpt.DdosZoneTemplateIpProto
	ret.Inst.FilterList = getSliceDdosZoneTemplateIpProtoFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.FilterMatchType = d.Get("filter_match_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
