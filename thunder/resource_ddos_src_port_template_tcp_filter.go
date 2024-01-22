package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcPortTemplateTcpFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_port_template_tcp_filter`: TCP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcPortTemplateTcpFilterCreate,
		UpdateContext: resourceDdosSrcPortTemplateTcpFilterUpdate,
		ReadContext:   resourceDdosSrcPortTemplateTcpFilterRead,
		DeleteContext: resourceDdosSrcPortTemplateTcpFilterDelete,

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
func resourceDdosSrcPortTemplateTcpFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcpFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateTcpFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcPortTemplateTcpFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcpFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateTcpFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcPortTemplateTcpFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcpFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcPortTemplateTcpFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcpFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosSrcPortTemplateTcpFilter(d *schema.ResourceData) edpt.DdosSrcPortTemplateTcpFilter {
	var ret edpt.DdosSrcPortTemplateTcpFilter
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
