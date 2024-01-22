package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneSrcPortTemplateUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_src_port_template_udp`: UDP template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneSrcPortTemplateUdpCreate,
		UpdateContext: resourceDdosZoneSrcPortTemplateUdpUpdate,
		ReadContext:   resourceDdosZoneSrcPortTemplateUdpRead,
		DeleteContext: resourceDdosZoneSrcPortTemplateUdpDelete,

		Schema: map[string]*schema.Schema{
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"udp_filter_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"udp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"udp_filter_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
						},
						"udp_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"udp_filter_action": {
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
			"max_payload_size_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_payload_size": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum UDP payload size for each single packet",
						},
						"max_payload_size_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for max-payload-size exceed",
						},
						"max_payload_size_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for max-payload-size exceed (Default); 'blacklist-src': Blacklist-src for max-payload-size exceed; 'ignore': Do nothing for max-payload-size exceed;",
						},
					},
				},
			},
			"min_payload_size_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_payload_size": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum UDP payload size for each single packet",
						},
						"min_payload_size_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for min-payload-size exceed",
						},
						"min_payload_size_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for min-payload-size (Default); 'blacklist-src': Blacklist-src for min-payload-size; 'ignore': Do nothing for min-payload-size exceed;",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS UDP Template Name",
			},
			"ntp_monlist_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ntp_monlist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take action for ntp monlist request/response",
						},
						"ntp_monlist_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for ntp-monlist",
						},
						"ntp_monlist_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for ntp-monlist (Default); 'blacklist-src': Blacklist-src for ntp-monlist; 'ignore': Ignore ntp-monlist;",
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
		},
	}
}
func resourceDdosZoneSrcPortTemplateUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneSrcPortTemplateUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneSrcPortTemplateUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneSrcPortTemplateUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneSrcPortTemplateUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneSrcPortTemplateUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneSrcPortTemplateUdpFilterList(d []interface{}) []edpt.DdosZoneSrcPortTemplateUdpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneSrcPortTemplateUdpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneSrcPortTemplateUdpFilterList
		oi.UdpFilterName = in["udp_filter_name"].(string)
		oi.UdpFilterSeq = in["udp_filter_seq"].(int)
		oi.UdpFilterRegex = in["udp_filter_regex"].(string)
		oi.UdpFilterInverseMatch = in["udp_filter_inverse_match"].(int)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.UdpFilterActionListName = in["udp_filter_action_list_name"].(string)
		oi.UdpFilterAction = in["udp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneSrcPortTemplateUdpMaxPayloadSizeCfg(d []interface{}) edpt.DdosZoneSrcPortTemplateUdpMaxPayloadSizeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneSrcPortTemplateUdpMaxPayloadSizeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxPayloadSize = in["max_payload_size"].(int)
		ret.MaxPayloadSizeActionListName = in["max_payload_size_action_list_name"].(string)
		ret.MaxPayloadSizeAction = in["max_payload_size_action"].(string)
	}
	return ret
}

func getObjectDdosZoneSrcPortTemplateUdpMinPayloadSizeCfg(d []interface{}) edpt.DdosZoneSrcPortTemplateUdpMinPayloadSizeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneSrcPortTemplateUdpMinPayloadSizeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinPayloadSize = in["min_payload_size"].(int)
		ret.MinPayloadSizeActionListName = in["min_payload_size_action_list_name"].(string)
		ret.MinPayloadSizeAction = in["min_payload_size_action"].(string)
	}
	return ret
}

func getObjectDdosZoneSrcPortTemplateUdpNtpMonlistCfg(d []interface{}) edpt.DdosZoneSrcPortTemplateUdpNtpMonlistCfg {

	count1 := len(d)
	var ret edpt.DdosZoneSrcPortTemplateUdpNtpMonlistCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NtpMonlist = in["ntp_monlist"].(int)
		ret.NtpMonlistActionListName = in["ntp_monlist_action_list_name"].(string)
		ret.NtpMonlistAction = in["ntp_monlist_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneSrcPortTemplateUdp(d *schema.ResourceData) edpt.DdosZoneSrcPortTemplateUdp {
	var ret edpt.DdosZoneSrcPortTemplateUdp
	ret.Inst.FilterList = getSliceDdosZoneSrcPortTemplateUdpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.FilterMatchType = d.Get("filter_match_type").(string)
	ret.Inst.MaxPayloadSizeCfg = getObjectDdosZoneSrcPortTemplateUdpMaxPayloadSizeCfg(d.Get("max_payload_size_cfg").([]interface{}))
	ret.Inst.MinPayloadSizeCfg = getObjectDdosZoneSrcPortTemplateUdpMinPayloadSizeCfg(d.Get("min_payload_size_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NtpMonlistCfg = getObjectDdosZoneSrcPortTemplateUdpNtpMonlistCfg(d.Get("ntp_monlist_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
