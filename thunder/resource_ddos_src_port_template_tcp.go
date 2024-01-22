package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcPortTemplateTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_port_template_tcp`: TCP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcPortTemplateTcpCreate,
		UpdateContext: resourceDdosSrcPortTemplateTcpUpdate,
		ReadContext:   resourceDdosSrcPortTemplateTcpRead,
		DeleteContext: resourceDdosSrcPortTemplateTcpDelete,

		Schema: map[string]*schema.Schema{
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_filter_seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"tcp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter Expression using Berkeley Packet Filter syntax",
						},
						"tcp_filter_unmatched": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
						},
						"tcp_filter_action": {
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
func resourceDdosSrcPortTemplateTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcPortTemplateTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcPortTemplateTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcPortTemplateTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcPortTemplateTcpFilterList(d []interface{}) []edpt.DdosSrcPortTemplateTcpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcPortTemplateTcpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcPortTemplateTcpFilterList
		oi.TcpFilterSeq = in["tcp_filter_seq"].(int)
		oi.TcpFilterRegex = in["tcp_filter_regex"].(string)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.TcpFilterUnmatched = in["tcp_filter_unmatched"].(int)
		oi.TcpFilterAction = in["tcp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSrcPortTemplateTcp(d *schema.ResourceData) edpt.DdosSrcPortTemplateTcp {
	var ret edpt.DdosSrcPortTemplateTcp
	ret.Inst.FilterList = getSliceDdosSrcPortTemplateTcpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
