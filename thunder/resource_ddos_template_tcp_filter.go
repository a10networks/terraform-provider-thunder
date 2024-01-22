package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcpFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp_filter`: TCP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpFilterCreate,
		UpdateContext: resourceDdosTemplateTcpFilterUpdate,
		ReadContext:   resourceDdosTemplateTcpFilterRead,
		DeleteContext: resourceDdosTemplateTcpFilterDelete,

		Schema: map[string]*schema.Schema{
			"byte_offset_filter": {
				Type: schema.TypeString, Optional: true, Description: "Filter Expression using Berkeley Packet Filter syntax",
			},
			"tcp_filter_action": {
				Type: schema.TypeString, Optional: true, Description: "'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;",
			},
			"tcp_filter_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"tcp_filter_seq": {
				Type: schema.TypeInt, Required: true, Description: "Sequence number",
			},
			"tcp_filter_unmatched": {
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
func resourceDdosTemplateTcpFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateTcpFilter(d *schema.ResourceData) edpt.DdosTemplateTcpFilter {
	var ret edpt.DdosTemplateTcpFilter
	ret.Inst.ByteOffsetFilter = d.Get("byte_offset_filter").(string)
	ret.Inst.TcpFilterAction = d.Get("tcp_filter_action").(string)
	ret.Inst.TcpFilterRegex = d.Get("tcp_filter_regex").(string)
	ret.Inst.TcpFilterSeq = d.Get("tcp_filter_seq").(int)
	ret.Inst.TcpFilterUnmatched = d.Get("tcp_filter_unmatched").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
