package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateQuicVersionSupportedMalformedCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_quic_version_supported_malformed_check`: Configure and enable malformed packet check\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckCreate,
		UpdateContext: resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckUpdate,
		ReadContext:   resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckRead,
		DeleteContext: resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckDelete,

		Schema: map[string]*schema.Schema{
			"malformed_check_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src;",
			},
			"malformed_check_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take. Overwrites version action",
			},
			"malformed_enable": {
				Type: schema.TypeString, Required: true, Description: "'enable': Enable malformed check;",
			},
			"max_destination_cid_length": {
				Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set the maximum destination CID length",
			},
			"max_source_cid_length": {
				Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set the maximum source CID length",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"version_start": {
				Type: schema.TypeString, Required: true, Description: "VersionStart",
			},
			"version_end": {
				Type: schema.TypeString, Required: true, Description: "VersionEnd",
			},
			"quic_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "QuicTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupportedMalformedCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupportedMalformedCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupportedMalformedCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateQuicVersionSupportedMalformedCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateQuicVersionSupportedMalformedCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateQuicVersionSupportedMalformedCheck(d *schema.ResourceData) edpt.DdosZoneTemplateQuicVersionSupportedMalformedCheck {
	var ret edpt.DdosZoneTemplateQuicVersionSupportedMalformedCheck
	ret.Inst.MalformedCheckAction = d.Get("malformed_check_action").(string)
	ret.Inst.MalformedCheckActionListName = d.Get("malformed_check_action_list_name").(string)
	ret.Inst.MalformedEnable = d.Get("malformed_enable").(string)
	ret.Inst.MaxDestinationCidLength = d.Get("max_destination_cid_length").(int)
	ret.Inst.MaxSourceCidLength = d.Get("max_source_cid_length").(int)
	//omit uuid
	ret.Inst.VersionStart = d.Get("version_start").(string)
	ret.Inst.VersionEnd = d.Get("version_end").(string)
	ret.Inst.QuicTmplName = d.Get("quic_tmpl_name").(string)
	return ret
}
