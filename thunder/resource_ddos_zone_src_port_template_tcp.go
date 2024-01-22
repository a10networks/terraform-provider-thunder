package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneSrcPortTemplateTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_src_port_template_tcp`: TCP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneSrcPortTemplateTcpCreate,
		UpdateContext: resourceDdosZoneSrcPortTemplateTcpUpdate,
		ReadContext:   resourceDdosZoneSrcPortTemplateTcpRead,
		DeleteContext: resourceDdosZoneSrcPortTemplateTcpDelete,

		Schema: map[string]*schema.Schema{
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"tcp_filter_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"tcp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"tcp_filter_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
						},
						"tcp_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"tcp_filter_action": {
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
				Type: schema.TypeString, Required: true, Description: "",
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
func resourceDdosZoneSrcPortTemplateTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneSrcPortTemplateTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneSrcPortTemplateTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneSrcPortTemplateTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneSrcPortTemplateTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneSrcPortTemplateTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneSrcPortTemplateTcpFilterList(d []interface{}) []edpt.DdosZoneSrcPortTemplateTcpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneSrcPortTemplateTcpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneSrcPortTemplateTcpFilterList
		oi.TcpFilterName = in["tcp_filter_name"].(string)
		oi.TcpFilterSeq = in["tcp_filter_seq"].(int)
		oi.TcpFilterRegex = in["tcp_filter_regex"].(string)
		oi.TcpFilterInverseMatch = in["tcp_filter_inverse_match"].(int)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.TcpFilterActionListName = in["tcp_filter_action_list_name"].(string)
		oi.TcpFilterAction = in["tcp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosZoneSrcPortTemplateTcp(d *schema.ResourceData) edpt.DdosZoneSrcPortTemplateTcp {
	var ret edpt.DdosZoneSrcPortTemplateTcp
	ret.Inst.FilterList = getSliceDdosZoneSrcPortTemplateTcpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.FilterMatchType = d.Get("filter_match_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
