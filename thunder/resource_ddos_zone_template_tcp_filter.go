package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateTcpFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_tcp_filter`: TCP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateTcpFilterCreate,
		UpdateContext: resourceDdosZoneTemplateTcpFilterUpdate,
		ReadContext:   resourceDdosZoneTemplateTcpFilterRead,
		DeleteContext: resourceDdosZoneTemplateTcpFilterDelete,

		Schema: map[string]*schema.Schema{
			"byte_offset_filter": {
				Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
			},
			"tcp_filter_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src;",
			},
			"tcp_filter_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"tcp_filter_inverse_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
			},
			"tcp_filter_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"tcp_filter_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"tcp_filter_seq": {
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
func resourceDdosZoneTemplateTcpFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateTcpFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateTcpFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateTcpFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateTcpFilter(d *schema.ResourceData) edpt.DdosZoneTemplateTcpFilter {
	var ret edpt.DdosZoneTemplateTcpFilter
	ret.Inst.ByteOffsetFilter = d.Get("byte_offset_filter").(string)
	ret.Inst.TcpFilterAction = d.Get("tcp_filter_action").(string)
	ret.Inst.TcpFilterActionListName = d.Get("tcp_filter_action_list_name").(string)
	ret.Inst.TcpFilterInverseMatch = d.Get("tcp_filter_inverse_match").(int)
	ret.Inst.TcpFilterName = d.Get("tcp_filter_name").(string)
	ret.Inst.TcpFilterRegex = d.Get("tcp_filter_regex").(string)
	ret.Inst.TcpFilterSeq = d.Get("tcp_filter_seq").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
