package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowCollectorHostCustomizedSetting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_collector_host_customized_setting`: Customize export settings for collector\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowCollectorHostCustomizedSettingCreate,
		UpdateContext: resourceSflowCollectorHostCustomizedSettingUpdate,
		ReadContext:   resourceSflowCollectorHostCustomizedSettingRead,
		DeleteContext: resourceSflowCollectorHostCustomizedSettingDelete,

		Schema: map[string]*schema.Schema{
			"a10_proprietary_polling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable counters for ACOS control blocks",
			},
			"counter_polling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable counter polling",
			},
			"event_notification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable event notification",
			},
			"export_enable": {
				Type: schema.TypeString, Required: true, Description: "'export': Customizes export settings for collector;",
			},
			"packet_sampling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet sampling",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"host_name": {
				Type: schema.TypeString, Required: true, Description: "Host_name",
			},
			"port": {
				Type: schema.TypeString, Required: true, Description: "Port",
			},
		},
	}
}
func resourceSflowCollectorHostCustomizedSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostCustomizedSettingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHostCustomizedSetting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorHostCustomizedSettingRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowCollectorHostCustomizedSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostCustomizedSettingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHostCustomizedSetting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorHostCustomizedSettingRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowCollectorHostCustomizedSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostCustomizedSettingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHostCustomizedSetting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowCollectorHostCustomizedSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostCustomizedSettingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHostCustomizedSetting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowCollectorHostCustomizedSetting(d *schema.ResourceData) edpt.SflowCollectorHostCustomizedSetting {
	var ret edpt.SflowCollectorHostCustomizedSetting
	ret.Inst.A10ProprietaryPolling = d.Get("a10_proprietary_polling").(int)
	ret.Inst.CounterPolling = d.Get("counter_polling").(int)
	ret.Inst.EventNotification = d.Get("event_notification").(int)
	ret.Inst.ExportEnable = d.Get("export_enable").(string)
	ret.Inst.PacketSampling = d.Get("packet_sampling").(int)
	//omit uuid
	ret.Inst.Host_name = d.Get("host_name").(string)
	ret.Inst.Port = d.Get("port").(string)
	return ret
}
