package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProbeInfo() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_probe_info`: Harmony controller probe statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerProbeInfoCreate,
		UpdateContext: resourceControllerProbeInfoUpdate,
		ReadContext:   resourceControllerProbeInfoRead,
		DeleteContext: resourceControllerProbeInfoDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'data-showtech-sent': DATA_SHOWTECH samples sent successfully to probe; 'data-showtech-failed': DATA_SHOWTECH samples failed to send to probe; 'data-varlog-sent': DATA_VARLOG samples sent successfully to probe; 'data-varlog-failed': DATA_VARLOG samples failed to send to probe; 'ssh-connection-failed': SSH_CONNECTION failures (ACOS-A10C);",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceControllerProbeInfoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProbeInfoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProbeInfo(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProbeInfoRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerProbeInfoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProbeInfoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProbeInfo(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProbeInfoRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerProbeInfoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProbeInfoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProbeInfo(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerProbeInfoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProbeInfoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProbeInfo(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceControllerProbeInfoSamplingEnable(d []interface{}) []edpt.ControllerProbeInfoSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.ControllerProbeInfoSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ControllerProbeInfoSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointControllerProbeInfo(d *schema.ResourceData) edpt.ControllerProbeInfo {
	var ret edpt.ControllerProbeInfo
	ret.Inst.SamplingEnable = getSliceControllerProbeInfoSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
