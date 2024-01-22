package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgRtsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg_rtsp`: Change Firewall RTSP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgRtspCreate,
		UpdateContext: resourceFwAlgRtspUpdate,
		ReadContext:   resourceFwAlgRtspRead,
		DeleteContext: resourceFwAlgRtspDelete,

		Schema: map[string]*schema.Schema{
			"default_port_disable": {
				Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable RTSP ALG default port 554;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'transport-inserted': Transport Created; 'transport-freed': Transport Freed; 'transport-alloc-failure': Transport Alloc Failure; 'data-session-created': Data Session Created; 'data-session-freed': Data Session Freed; 'ext-creation-failure': Extension Creation Failure; 'transport-add-to-ext': Transport Added to Extension; 'transport-removed-from-ext': Transport Removed from Extension; 'transport-too-many': Too Many Transports for Control Conn; 'transport-already-in-ext': Transport Already in Extension; 'transport-exists': Transport Already Exists; 'transport-link-ext-failure-control': Transport Link to Extension Failure Control; 'transport-link-ext-data': Transport Link to Extension Data; 'transport-link-ext-failure-data': Transport Link to Extension Failure Data; 'transport-inserted-shadow': Transport Inserted Shadow; 'transport-creation-race': Transport Create Race; 'transport-alloc-failure-shadow': Transport Alloc Failure Shadow; 'transport-put-in-del-q': Transport Put in Delete Queue; 'transport-freed-shadow': Transport Freed Shadow; 'transport-acquired-from-control': Transport Acquired Control; 'transport-found-from-prev-control': Transport Found From Prev Control; 'transport-acquire-failure-from-control': Transport Acquire Failure Control; 'transport-released-from-control': Transport Released Control; 'transport-double-release-from-control': Transport Double Release Control; 'transport-acquired-from-data': Transport Acquired Data; 'transport-acquire-failure-from-data': Transport Acquire Failure Data; 'transport-released-from-data': Transport Released Data; 'transport-double-release-from-data': Transport Double Release Data; 'transport-retry-lookup-on-data-free': Transport Retry Lookup Data; 'transport-not-found-on-data-free': Transport Not Found Data; 'data-session-created-shadow': Data Session Created Shadow; 'data-session-freed-shadow': Data Session Freed Shadow; 'ha-control-ext-creation-failure': HA Control Extension Creation Failure; 'ha-control-session-created': HA Control Session Created; 'ha-data-session-created': HA Data Session Created;",
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
func resourceFwAlgRtspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgRtspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgRtsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgRtspRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgRtspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgRtspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgRtsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgRtspRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgRtspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgRtspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgRtsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgRtspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgRtspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgRtsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwAlgRtspSamplingEnable(d []interface{}) []edpt.FwAlgRtspSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwAlgRtspSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgRtspSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwAlgRtsp(d *schema.ResourceData) edpt.FwAlgRtsp {
	var ret edpt.FwAlgRtsp
	ret.Inst.DefaultPortDisable = d.Get("default_port_disable").(string)
	ret.Inst.SamplingEnable = getSliceFwAlgRtspSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
