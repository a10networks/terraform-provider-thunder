package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateQuicVersionSupported() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_quic_version_supported`: Configure version support options\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateQuicVersionSupportedCreate,
		UpdateContext: resourceDdosZoneTemplateQuicVersionSupportedUpdate,
		ReadContext:   resourceDdosZoneTemplateQuicVersionSupportedRead,
		DeleteContext: resourceDdosZoneTemplateQuicVersionSupportedDelete,

		Schema: map[string]*schema.Schema{
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"version_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets; 'blacklist-src': Blacklist-src;",
			},
			"version_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"version_end": {
				Type: schema.TypeString, Required: true, Description: "Version supported range end",
			},
			"version_start": {
				Type: schema.TypeString, Required: true, Description: "Configure versions supported",
			},
			"quic_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "QuicTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateQuicVersionSupportedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupported(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicVersionSupportedRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateQuicVersionSupportedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupported(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicVersionSupportedRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateQuicVersionSupportedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupported(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateQuicVersionSupportedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupported(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateQuicVersionSupportedMalformedCheck315(d []interface{}) edpt.DdosZoneTemplateQuicVersionSupportedMalformedCheck315 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateQuicVersionSupportedMalformedCheck315
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

func dataToEndpointDdosZoneTemplateQuicVersionSupported(d *schema.ResourceData) edpt.DdosZoneTemplateQuicVersionSupported {
	var ret edpt.DdosZoneTemplateQuicVersionSupported
	ret.Inst.MalformedCheck = getObjectDdosZoneTemplateQuicVersionSupportedMalformedCheck315(d.Get("malformed_check").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VersionAction = d.Get("version_action").(string)
	ret.Inst.VersionActionListName = d.Get("version_action_list_name").(string)
	ret.Inst.VersionEnd = d.Get("version_end").(string)
	ret.Inst.VersionStart = d.Get("version_start").(string)
	ret.Inst.QuicTmplName = d.Get("quic_tmpl_name").(string)
	return ret
}
