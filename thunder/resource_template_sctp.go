package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateSctp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_sctp`: Define a SCTP template\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateSctpCreate,
		UpdateContext: resourceTemplateSctpUpdate,
		ReadContext:   resourceTemplateSctpRead,
		DeleteContext: resourceTemplateSctpDelete,

		Schema: map[string]*schema.Schema{
			"checksum_check": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable Checksum check;",
			},
			"log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload_proto_filtering": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log Payload Protocol IDs Filtered",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SCTP Template Name",
			},
			"permit_payload_protocol": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permit_config_id": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol_id": {
										Type: schema.TypeInt, Optional: true, Description: "Specify SCTP permitted payload protocol IDs",
									},
								},
							},
						},
						"permit_config_name": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol_name": {
										Type: schema.TypeString, Optional: true, Description: "'iua': IUA; 'm2ua': M2UA; 'm3ua': M3UA; 'sua': SUA; 'm2pa': M2PA; 'h.323': H.323;",
									},
								},
							},
						},
					},
				},
			},
			"sctp_half_open_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "Set SCTP half-open timeout (SCTP half-open timeout in seconds (default 4))",
			},
			"sctp_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "SCTP idle timeout in minutes (default 5)",
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
func resourceTemplateSctpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateSctpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateSctp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateSctpRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateSctpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateSctpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateSctp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateSctpRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateSctpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateSctpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateSctp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateSctpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateSctpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateSctp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTemplateSctpLog(d []interface{}) edpt.TemplateSctpLog {

	count1 := len(d)
	var ret edpt.TemplateSctpLog
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PayloadProtoFiltering = in["payload_proto_filtering"].(int)
	}
	return ret
}

func getObjectTemplateSctpPermitPayloadProtocol(d []interface{}) edpt.TemplateSctpPermitPayloadProtocol {

	count1 := len(d)
	var ret edpt.TemplateSctpPermitPayloadProtocol
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PermitConfigId = getSliceTemplateSctpPermitPayloadProtocolPermitConfigId(in["permit_config_id"].([]interface{}))
		ret.PermitConfigName = getSliceTemplateSctpPermitPayloadProtocolPermitConfigName(in["permit_config_name"].([]interface{}))
	}
	return ret
}

func getSliceTemplateSctpPermitPayloadProtocolPermitConfigId(d []interface{}) []edpt.TemplateSctpPermitPayloadProtocolPermitConfigId {

	count1 := len(d)
	ret := make([]edpt.TemplateSctpPermitPayloadProtocolPermitConfigId, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateSctpPermitPayloadProtocolPermitConfigId
		oi.ProtocolId = in["protocol_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTemplateSctpPermitPayloadProtocolPermitConfigName(d []interface{}) []edpt.TemplateSctpPermitPayloadProtocolPermitConfigName {

	count1 := len(d)
	ret := make([]edpt.TemplateSctpPermitPayloadProtocolPermitConfigName, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateSctpPermitPayloadProtocolPermitConfigName
		oi.ProtocolName = in["protocol_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTemplateSctp(d *schema.ResourceData) edpt.TemplateSctp {
	var ret edpt.TemplateSctp
	ret.Inst.ChecksumCheck = d.Get("checksum_check").(string)
	ret.Inst.Log = getObjectTemplateSctpLog(d.Get("log").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PermitPayloadProtocol = getObjectTemplateSctpPermitPayloadProtocol(d.Get("permit_payload_protocol").([]interface{}))
	ret.Inst.SctpHalfOpenIdleTimeout = d.Get("sctp_half_open_idle_timeout").(int)
	ret.Inst.SctpIdleTimeout = d.Get("sctp_idle_timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
