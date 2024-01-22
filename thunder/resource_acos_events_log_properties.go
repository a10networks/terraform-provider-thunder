package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogProperties() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_log_properties`: Configure Logging related properties\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsLogPropertiesCreate,
		UpdateContext: resourceAcosEventsLogPropertiesUpdate,
		ReadContext:   resourceAcosEventsLogPropertiesRead,
		DeleteContext: resourceAcosEventsLogPropertiesDelete,

		Schema: map[string]*schema.Schema{
			"add_msgid_in_header": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Add Message ID in log messages",
			},
			"enable_8k_tcp_syslog": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 8K size remote TCP syslog",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsLogPropertiesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogPropertiesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogProperties(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogPropertiesRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsLogPropertiesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogPropertiesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogProperties(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogPropertiesRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsLogPropertiesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogPropertiesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogProperties(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsLogPropertiesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogPropertiesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogProperties(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsLogProperties(d *schema.ResourceData) edpt.AcosEventsLogProperties {
	var ret edpt.AcosEventsLogProperties
	ret.Inst.AddMsgidInHeader = d.Get("add_msgid_in_header").(int)
	ret.Inst.Enable8kTcpSyslog = d.Get("enable_8k_tcp_syslog").(int)
	//omit uuid
	return ret
}
