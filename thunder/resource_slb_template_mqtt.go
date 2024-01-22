package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateMqtt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_mqtt`: MQTT template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateMqttCreate,
		UpdateContext: resourceSlbTemplateMqttUpdate,
		ReadContext:   resourceSlbTemplateMqttRead,
		DeleteContext: resourceSlbTemplateMqttDelete,

		Schema: map[string]*schema.Schema{
			"clientid_hash_first": {
				Type: schema.TypeInt, Optional: true, Description: "Use the begining part of client ID to calculate hash value (client ID string length to calculate hash value)",
			},
			"clientid_hash_last": {
				Type: schema.TypeInt, Optional: true, Description: "Use the end part of Client ID to calculate hash value (Client ID length to calculate hash value)",
			},
			"clientid_hash_offset": {
				Type: schema.TypeInt, Optional: true, Description: "Skip part of Client ID to calculate hash value (Offset of the Client ID)",
			},
			"clientid_hash_persist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Client ID's hash value to select server",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "MQTT Template Name",
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
func resourceSlbTemplateMqttCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMqttCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMqtt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateMqttRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateMqttUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMqttUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMqtt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateMqttRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateMqttDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMqttDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMqtt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateMqttRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMqttRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMqtt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateMqtt(d *schema.ResourceData) edpt.SlbTemplateMqtt {
	var ret edpt.SlbTemplateMqtt
	ret.Inst.ClientidHashFirst = d.Get("clientid_hash_first").(int)
	ret.Inst.ClientidHashLast = d.Get("clientid_hash_last").(int)
	ret.Inst.ClientidHashOffset = d.Get("clientid_hash_offset").(int)
	ret.Inst.ClientidHashPersist = d.Get("clientid_hash_persist").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
