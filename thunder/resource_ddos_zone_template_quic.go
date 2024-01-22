package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateQuic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_quic`: QUIC template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateQuicCreate,
		UpdateContext: resourceDdosZoneTemplateQuicUpdate,
		ReadContext:   resourceDdosZoneTemplateQuicRead,
		DeleteContext: resourceDdosZoneTemplateQuicDelete,

		Schema: map[string]*schema.Schema{
			"fixed_bit_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable fixed-bit malform check",
			},
			"quic_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"version_supported_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version_start": {
							Type: schema.TypeString, Required: true, Description: "Configure versions supported",
						},
						"version_end": {
							Type: schema.TypeString, Required: true, Description: "Version supported range end",
						},
						"version_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"version_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets; 'blacklist-src': Blacklist-src;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"malformed_check": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"malformed_enable": {
										Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable malformed check;",
									},
									"max_source_cid_length": {
										Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set the maximum source CID length",
									},
									"max_destination_cid_length": {
										Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set the maximum destination CID length",
									},
									"malformed_check_action_list_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure action-list to take. Overwrites version action",
									},
									"malformed_check_action": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func resourceDdosZoneTemplateQuicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateQuicUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateQuicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateQuicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneTemplateQuicVersionSupportedList(d []interface{}) []edpt.DdosZoneTemplateQuicVersionSupportedList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateQuicVersionSupportedList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateQuicVersionSupportedList
		oi.VersionStart = in["version_start"].(string)
		oi.VersionEnd = in["version_end"].(string)
		oi.VersionActionListName = in["version_action_list_name"].(string)
		oi.VersionAction = in["version_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.MalformedCheck = getObjectDdosZoneTemplateQuicVersionSupportedListMalformedCheck(in["malformed_check"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateQuicVersionSupportedListMalformedCheck(d []interface{}) edpt.DdosZoneTemplateQuicVersionSupportedListMalformedCheck {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateQuicVersionSupportedListMalformedCheck
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MalformedEnable = in["malformed_enable"].(string)
		ret.MaxSourceCidLength = in["max_source_cid_length"].(int)
		ret.MaxDestinationCidLength = in["max_destination_cid_length"].(int)
		ret.MalformedCheckActionListName = in["malformed_check_action_list_name"].(string)
		ret.MalformedCheckAction = in["malformed_check_action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosZoneTemplateQuic(d *schema.ResourceData) edpt.DdosZoneTemplateQuic {
	var ret edpt.DdosZoneTemplateQuic
	ret.Inst.FixedBitCheckDisable = d.Get("fixed_bit_check_disable").(int)
	ret.Inst.QuicTmplName = d.Get("quic_tmpl_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VersionSupportedList = getSliceDdosZoneTemplateQuicVersionSupportedList(d.Get("version_supported_list").([]interface{}))
	return ret
}
