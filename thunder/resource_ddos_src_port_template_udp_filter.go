package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcPortTemplateUdpFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_port_template_udp_filter`: UDP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcPortTemplateUdpFilterCreate,
		UpdateContext: resourceDdosSrcPortTemplateUdpFilterUpdate,
		ReadContext:   resourceDdosSrcPortTemplateUdpFilterRead,
		DeleteContext: resourceDdosSrcPortTemplateUdpFilterDelete,

		Schema: map[string]*schema.Schema{
			"byte_offset_filter": {
				Type: schema.TypeString, Optional: true, Description: "Filter Expression using Berkeley Packet Filter syntax",
			},
			"udp_filter_action": {
				Type: schema.TypeString, Optional: true, Description: "'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;",
			},
			"udp_filter_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"udp_filter_seq": {
				Type: schema.TypeInt, Required: true, Description: "Sequence number",
			},
			"udp_filter_unmatched": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
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
func resourceDdosSrcPortTemplateUdpFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdpFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateUdpFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcPortTemplateUdpFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdpFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateUdpFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcPortTemplateUdpFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdpFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcPortTemplateUdpFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateUdpFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateUdpFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosSrcPortTemplateUdpFilter(d *schema.ResourceData) edpt.DdosSrcPortTemplateUdpFilter {
	var ret edpt.DdosSrcPortTemplateUdpFilter
	ret.Inst.ByteOffsetFilter = d.Get("byte_offset_filter").(string)
	ret.Inst.UdpFilterAction = d.Get("udp_filter_action").(string)
	ret.Inst.UdpFilterRegex = d.Get("udp_filter_regex").(string)
	ret.Inst.UdpFilterSeq = d.Get("udp_filter_seq").(int)
	ret.Inst.UdpFilterUnmatched = d.Get("udp_filter_unmatched").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
