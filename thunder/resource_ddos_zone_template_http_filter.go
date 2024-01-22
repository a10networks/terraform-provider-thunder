package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateHttpFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_http_filter`: HTTP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateHttpFilterCreate,
		UpdateContext: resourceDdosZoneTemplateHttpFilterUpdate,
		ReadContext:   resourceDdosZoneTemplateHttpFilterRead,
		DeleteContext: resourceDdosZoneTemplateHttpFilterDelete,

		Schema: map[string]*schema.Schema{
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_filter_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Set rate limit",
						},
					},
				},
			},
			"http_agent_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_equals_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_agent_equals": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"agent_contains_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_agent_contains": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"agent_starts_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_agent_starts_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"agent_ends_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_agent_ends_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"http_filter_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src; 'reset': Reset client connection;",
			},
			"http_filter_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"http_filter_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"http_filter_seq": {
				Type: schema.TypeInt, Optional: true, Description: "Sequence number",
			},
			"http_header_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_filter_header_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"http_filter_header_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
					},
				},
			},
			"http_referer_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"referer_equals_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_referer_equals": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"referer_contains_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_referer_contains": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"referer_starts_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_referer_starts_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"referer_ends_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_referer_ends_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"http_uri_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri_equal_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_uri_equals": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"uri_contains_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_uri_contains": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"uri_starts_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_uri_starts_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"uri_ends_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_filter_uri_ends_with": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
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
			"http_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "HttpTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateHttpFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateHttpFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateHttpFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateHttpFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateHttpFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateHttpFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateHttpFilterDst(d []interface{}) edpt.DdosZoneTemplateHttpFilterDst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpFilterRateLimit = in["http_filter_rate_limit"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterHttpAgentCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterHttpAgentCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterHttpAgentCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AgentEqualsCfg = getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentEqualsCfg(in["agent_equals_cfg"].([]interface{}))
		ret.AgentContainsCfg = getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentContainsCfg(in["agent_contains_cfg"].([]interface{}))
		ret.AgentStartsCfg = getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentStartsCfg(in["agent_starts_cfg"].([]interface{}))
		ret.AgentEndsCfg = getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentEndsCfg(in["agent_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentEqualsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentEqualsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentEqualsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentEqualsCfg
		oi.HttpFilterAgentEquals = in["http_filter_agent_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentContainsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentContainsCfg
		oi.HttpFilterAgentContains = in["http_filter_agent_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentStartsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentStartsCfg
		oi.HttpFilterAgentStartsWith = in["http_filter_agent_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpAgentCfgAgentEndsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpAgentCfgAgentEndsCfg
		oi.HttpFilterAgentEndsWith = in["http_filter_agent_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterHttpHeaderCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterHttpHeaderCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterHttpHeaderCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpFilterHeaderRegex = in["http_filter_header_regex"].(string)
		ret.HttpFilterHeaderInverseMatch = in["http_filter_header_inverse_match"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterHttpRefererCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterHttpRefererCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterHttpRefererCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RefererEqualsCfg = getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererEqualsCfg(in["referer_equals_cfg"].([]interface{}))
		ret.RefererContainsCfg = getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererContainsCfg(in["referer_contains_cfg"].([]interface{}))
		ret.RefererStartsCfg = getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererStartsCfg(in["referer_starts_cfg"].([]interface{}))
		ret.RefererEndsCfg = getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererEndsCfg(in["referer_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererEqualsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererEqualsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererEqualsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererEqualsCfg
		oi.HttpFilterRefererEquals = in["http_filter_referer_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererContainsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererContainsCfg
		oi.HttpFilterRefererContains = in["http_filter_referer_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererStartsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererStartsCfg
		oi.HttpFilterRefererStartsWith = in["http_filter_referer_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpRefererCfgRefererEndsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpRefererCfgRefererEndsCfg
		oi.HttpFilterRefererEndsWith = in["http_filter_referer_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateHttpFilterHttpUriCfg(d []interface{}) edpt.DdosZoneTemplateHttpFilterHttpUriCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateHttpFilterHttpUriCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UriEqualCfg = getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriEqualCfg(in["uri_equal_cfg"].([]interface{}))
		ret.UriContainsCfg = getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriContainsCfg(in["uri_contains_cfg"].([]interface{}))
		ret.UriStartsCfg = getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriStartsCfg(in["uri_starts_cfg"].([]interface{}))
		ret.UriEndsCfg = getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriEndsCfg(in["uri_ends_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriEqualCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriEqualCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriEqualCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriEqualCfg
		oi.HttpFilterUriEquals = in["http_filter_uri_equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriContainsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriContainsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriContainsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriContainsCfg
		oi.HttpFilterUriContains = in["http_filter_uri_contains"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriStartsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriStartsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriStartsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriStartsCfg
		oi.HttpFilterUriStartsWith = in["http_filter_uri_starts_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateHttpFilterHttpUriCfgUriEndsCfg(d []interface{}) []edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriEndsCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriEndsCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateHttpFilterHttpUriCfgUriEndsCfg
		oi.HttpFilterUriEndsWith = in["http_filter_uri_ends_with"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateHttpFilter(d *schema.ResourceData) edpt.DdosZoneTemplateHttpFilter {
	var ret edpt.DdosZoneTemplateHttpFilter
	ret.Inst.Dst = getObjectDdosZoneTemplateHttpFilterDst(d.Get("dst").([]interface{}))
	ret.Inst.HttpAgentCfg = getObjectDdosZoneTemplateHttpFilterHttpAgentCfg(d.Get("http_agent_cfg").([]interface{}))
	ret.Inst.HttpFilterAction = d.Get("http_filter_action").(string)
	ret.Inst.HttpFilterActionListName = d.Get("http_filter_action_list_name").(string)
	ret.Inst.HttpFilterName = d.Get("http_filter_name").(string)
	ret.Inst.HttpFilterSeq = d.Get("http_filter_seq").(int)
	ret.Inst.HttpHeaderCfg = getObjectDdosZoneTemplateHttpFilterHttpHeaderCfg(d.Get("http_header_cfg").([]interface{}))
	ret.Inst.HttpRefererCfg = getObjectDdosZoneTemplateHttpFilterHttpRefererCfg(d.Get("http_referer_cfg").([]interface{}))
	ret.Inst.HttpUriCfg = getObjectDdosZoneTemplateHttpFilterHttpUriCfg(d.Get("http_uri_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.HttpTmplName = d.Get("http_tmpl_name").(string)
	return ret
}
