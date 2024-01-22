package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosOutboundPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_outbound_policy`: Configure outbound policy\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosOutboundPolicyCreate,
		UpdateContext: resourceDdosOutboundPolicyUpdate,
		ReadContext:   resourceDdosOutboundPolicyRead,
		DeleteContext: resourceDdosOutboundPolicyDelete,

		Schema: map[string]*schema.Schema{
			"asn_based_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': Configure asn based tracking;",
						},
						"per_asn_glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"packet_rate_triggered": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Triggered by 1/2 packet rate limitation in per-asn-glid.",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"country_based_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': Configure country based tracking;",
						},
						"per_country_glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"packet_rate_triggered": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Triggered by 1/2 packet rate limitation in per-country-glid.",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
			},
			"policy_class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_name": {
							Type: schema.TypeString, Required: true, Description: "Class-list name",
						},
						"class_list_glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
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
			"policy_default_class_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Default class-list configuration",
						},
						"class_list_glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceDdosOutboundPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosOutboundPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosOutboundPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosOutboundPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosOutboundPolicyAsnBasedTracking287(d []interface{}) edpt.DdosOutboundPolicyAsnBasedTracking287 {

	count1 := len(d)
	var ret edpt.DdosOutboundPolicyAsnBasedTracking287
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.PerAsnGlid = in["per_asn_glid"].(string)
		ret.PacketRateTriggered = in["packet_rate_triggered"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosOutboundPolicyCountryBasedTracking288(d []interface{}) edpt.DdosOutboundPolicyCountryBasedTracking288 {

	count1 := len(d)
	var ret edpt.DdosOutboundPolicyCountryBasedTracking288
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.PerCountryGlid = in["per_country_glid"].(string)
		ret.PacketRateTriggered = in["packet_rate_triggered"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosOutboundPolicyPolicyClassListList(d []interface{}) []edpt.DdosOutboundPolicyPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosOutboundPolicyPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosOutboundPolicyPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.ClassListGlid = in["class_list_glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosOutboundPolicyPolicyDefaultClassList289(d []interface{}) edpt.DdosOutboundPolicyPolicyDefaultClassList289 {

	count1 := len(d)
	var ret edpt.DdosOutboundPolicyPolicyDefaultClassList289
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(int)
		ret.ClassListGlid = in["class_list_glid"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosOutboundPolicy(d *schema.ResourceData) edpt.DdosOutboundPolicy {
	var ret edpt.DdosOutboundPolicy
	ret.Inst.AsnBasedTracking = getObjectDdosOutboundPolicyAsnBasedTracking287(d.Get("asn_based_tracking").([]interface{}))
	ret.Inst.CountryBasedTracking = getObjectDdosOutboundPolicyCountryBasedTracking288(d.Get("country_based_tracking").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PolicyClassListList = getSliceDdosOutboundPolicyPolicyClassListList(d.Get("policy_class_list_list").([]interface{}))
	ret.Inst.PolicyDefaultClassList = getObjectDdosOutboundPolicyPolicyDefaultClassList289(d.Get("policy_default_class_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
