package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcPortTemplateUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_port_template_udp`: UDP template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcPortTemplateUdpCreate,
		UpdateContext: resourceDdosSrcPortTemplateUdpUpdate,
		ReadContext:   resourceDdosSrcPortTemplateUdpRead,
		DeleteContext: resourceDdosSrcPortTemplateUdpDelete,

		Schema: map[string]*schema.Schema{
			"drop_ntp_monlist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop NTP monlist request/response",
			},
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_filter_seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"udp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter Expression using Berkeley Packet Filter syntax",
						},
						"udp_filter_unmatched": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
						},
						"udp_filter_action": {
							Type: schema.TypeString, Optional: true, Description: "'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;",
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
			"max_payload_size": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum UDP payload size for each single packet",
			},
			"min_payload_size": {
				Type: schema.TypeInt, Optional: true, Description: "Minimum UDP payload size for each single packet",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS UDP Template Name",
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
func resourceDdosSrcPortTemplateUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcPortTemplateUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcPortTemplateUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcPortTemplateUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcPortTemplateUdpFilterList(d []interface{}) []edpt.DdosSrcPortTemplateUdpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcPortTemplateUdpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcPortTemplateUdpFilterList
		oi.UdpFilterSeq = in["udp_filter_seq"].(int)
		oi.UdpFilterRegex = in["udp_filter_regex"].(string)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.UdpFilterUnmatched = in["udp_filter_unmatched"].(int)
		oi.UdpFilterAction = in["udp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSrcPortTemplateUdp(d *schema.ResourceData) edpt.DdosSrcPortTemplateUdp {
	var ret edpt.DdosSrcPortTemplateUdp
	ret.Inst.DropNtpMonlist = d.Get("drop_ntp_monlist").(int)
	ret.Inst.FilterList = getSliceDdosSrcPortTemplateUdpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.MaxPayloadSize = d.Get("max_payload_size").(int)
	ret.Inst.MinPayloadSize = d.Get("min_payload_size").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
