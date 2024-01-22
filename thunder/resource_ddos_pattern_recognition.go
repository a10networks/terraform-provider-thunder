package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosPatternRecognition() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_pattern_recognition`: Pattern Recognition\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosPatternRecognitionCreate,
		UpdateContext: resourceDdosPatternRecognitionUpdate,
		ReadContext:   resourceDdosPatternRecognitionRead,
		DeleteContext: resourceDdosPatternRecognitionDelete,

		Schema: map[string]*schema.Schema{
			"capture_backup": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Capture Backup",
			},
			"capturing_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Capturing state timeout in seconds",
			},
			"cpu": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cpu_limit": {
				Type: schema.TypeInt, Optional: true, Description: "CPU Limit",
			},
			"dedicated_cpus": {
				Type: schema.TypeInt, Optional: true, Description: "Configure the number of dedicated cores for Pattern Recognition",
			},
			"disable_app_payload_all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable application payload processing for all ports",
			},
			"error_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Error state timeout in seconds",
			},
			"extracting_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Extracting state timeout in seconds",
			},
			"hardware_filter": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Pattern Recognition hardware filter; 'disable': Disable Pattern Recognition harware filter;",
			},
			"sample_size": {
				Type: schema.TypeInt, Optional: true, Description: "Sample Size",
			},
			"scheduling_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Scheduling state timeout in seconds",
			},
			"sensitivity": {
				Type: schema.TypeString, Optional: true, Description: "'high': High Sensitivity; 'medium': Medium Sensitivity; 'low': Low Sensitivity;",
			},
			"sflow_event_periodic_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Configure the interval in minutes of periodic event (Default: 5 minutes, 0: No periodic updates)",
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Pattern Recognition; 'disable': Disable Pattern Recognition;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosPatternRecognitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosPatternRecognitionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosPatternRecognition(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosPatternRecognitionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosPatternRecognitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosPatternRecognitionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosPatternRecognition(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosPatternRecognitionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosPatternRecognitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosPatternRecognitionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosPatternRecognition(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosPatternRecognitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosPatternRecognitionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosPatternRecognition(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosPatternRecognitionCpu290(d []interface{}) edpt.DdosPatternRecognitionCpu290 {

	var ret edpt.DdosPatternRecognitionCpu290
	return ret
}

func dataToEndpointDdosPatternRecognition(d *schema.ResourceData) edpt.DdosPatternRecognition {
	var ret edpt.DdosPatternRecognition
	ret.Inst.CaptureBackup = d.Get("capture_backup").(int)
	ret.Inst.CapturingTimeout = d.Get("capturing_timeout").(int)
	ret.Inst.Cpu = getObjectDdosPatternRecognitionCpu290(d.Get("cpu").([]interface{}))
	ret.Inst.CpuLimit = d.Get("cpu_limit").(int)
	ret.Inst.DedicatedCpus = d.Get("dedicated_cpus").(int)
	ret.Inst.DisableAppPayloadAll = d.Get("disable_app_payload_all").(int)
	ret.Inst.ErrorTimeout = d.Get("error_timeout").(int)
	ret.Inst.ExtractingTimeout = d.Get("extracting_timeout").(int)
	ret.Inst.HardwareFilter = d.Get("hardware_filter").(string)
	ret.Inst.SampleSize = d.Get("sample_size").(int)
	ret.Inst.SchedulingTimeout = d.Get("scheduling_timeout").(int)
	ret.Inst.Sensitivity = d.Get("sensitivity").(string)
	ret.Inst.SflowEventPeriodicInterval = d.Get("sflow_event_periodic_interval").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
