package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIpProtoFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_ip_proto_filter`: Ip-proto Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIpProtoFilterCreate,
		UpdateContext: resourceDdosZoneTemplateIpProtoFilterUpdate,
		ReadContext:   resourceDdosZoneTemplateIpProtoFilterRead,
		DeleteContext: resourceDdosZoneTemplateIpProtoFilterDelete,

		Schema: map[string]*schema.Schema{
			"byte_offset_filter": {
				Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
			},
			"other_filter_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src;",
			},
			"other_filter_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"other_filter_inverse_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
			},
			"other_filter_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"other_filter_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"other_filter_seq": {
				Type: schema.TypeInt, Optional: true, Description: "Sequence number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosZoneTemplateIpProtoFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProtoFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIpProtoFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIpProtoFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProtoFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIpProtoFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIpProtoFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProtoFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIpProtoFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIpProtoFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIpProtoFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateIpProtoFilter(d *schema.ResourceData) edpt.DdosZoneTemplateIpProtoFilter {
	var ret edpt.DdosZoneTemplateIpProtoFilter
	ret.Inst.ByteOffsetFilter = d.Get("byte_offset_filter").(string)
	ret.Inst.OtherFilterAction = d.Get("other_filter_action").(string)
	ret.Inst.OtherFilterActionListName = d.Get("other_filter_action_list_name").(string)
	ret.Inst.OtherFilterInverseMatch = d.Get("other_filter_inverse_match").(int)
	ret.Inst.OtherFilterName = d.Get("other_filter_name").(string)
	ret.Inst.OtherFilterRegex = d.Get("other_filter_regex").(string)
	ret.Inst.OtherFilterSeq = d.Get("other_filter_seq").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
