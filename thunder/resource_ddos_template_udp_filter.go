package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateUdpFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_udp_filter`: UDP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateUdpFilterCreate,
		UpdateContext: resourceDdosTemplateUdpFilterUpdate,
		ReadContext:   resourceDdosTemplateUdpFilterRead,
		DeleteContext: resourceDdosTemplateUdpFilterDelete,

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
func resourceDdosTemplateUdpFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdpFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateUdpFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateUdpFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdpFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateUdpFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateUdpFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdpFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateUdpFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdpFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateUdpFilter(d *schema.ResourceData) edpt.DdosTemplateUdpFilter {
	var ret edpt.DdosTemplateUdpFilter
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
