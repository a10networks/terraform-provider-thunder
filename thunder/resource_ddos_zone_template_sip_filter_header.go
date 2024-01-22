package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateSipFilterHeader() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_sip_filter_header`: SIP Header Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateSipFilterHeaderCreate,
		UpdateContext: resourceDdosZoneTemplateSipFilterHeaderUpdate,
		ReadContext:   resourceDdosZoneTemplateSipFilterHeaderRead,
		DeleteContext: resourceDdosZoneTemplateSipFilterHeaderDelete,

		Schema: map[string]*schema.Schema{
			"sip_filter_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src; 'reset': Reset client connection(for sip-tcp);",
			},
			"sip_filter_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"sip_filter_header_seq": {
				Type: schema.TypeInt, Optional: true, Description: "Sequence number",
			},
			"sip_filter_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"sip_header_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sip_filter_header_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"sip_filter_header_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
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
			"sip_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "SipTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateSipFilterHeaderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipFilterHeaderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSipFilterHeader(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSipFilterHeaderRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateSipFilterHeaderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipFilterHeaderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSipFilterHeader(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSipFilterHeaderRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateSipFilterHeaderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipFilterHeaderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSipFilterHeader(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateSipFilterHeaderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipFilterHeaderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSipFilterHeader(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateSipFilterHeaderSipHeaderCfg(d []interface{}) edpt.DdosZoneTemplateSipFilterHeaderSipHeaderCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipFilterHeaderSipHeaderCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SipFilterHeaderRegex = in["sip_filter_header_regex"].(string)
		ret.SipFilterHeaderInverseMatch = in["sip_filter_header_inverse_match"].(int)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateSipFilterHeader(d *schema.ResourceData) edpt.DdosZoneTemplateSipFilterHeader {
	var ret edpt.DdosZoneTemplateSipFilterHeader
	ret.Inst.SipFilterAction = d.Get("sip_filter_action").(string)
	ret.Inst.SipFilterActionListName = d.Get("sip_filter_action_list_name").(string)
	ret.Inst.SipFilterHeaderSeq = d.Get("sip_filter_header_seq").(int)
	ret.Inst.SipFilterName = d.Get("sip_filter_name").(string)
	ret.Inst.SipHeaderCfg = getObjectDdosZoneTemplateSipFilterHeaderSipHeaderCfg(d.Get("sip_header_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SipTmplName = d.Get("sip_tmpl_name").(string)
	return ret
}
