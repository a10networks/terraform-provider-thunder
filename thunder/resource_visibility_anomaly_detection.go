package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityAnomalyDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_anomaly_detection`: Anomaly detection parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityAnomalyDetectionCreate,
		UpdateContext: resourceVisibilityAnomalyDetectionUpdate,
		ReadContext:   resourceVisibilityAnomalyDetectionRead,
		DeleteContext: resourceVisibilityAnomalyDetectionDelete,

		Schema: map[string]*schema.Schema{
			"feature_status": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable anomaly-detection (Not valid for source-nat-ip and source monitor types); 'disable': Disable anomaly detection (default);",
			},
			"logging": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'per-entity': Enable per entity logging; 'per-metric': Enable per metric logging with threshold details; 'disable': Disable anomaly notifications (Default);",
			},
			"restart_learning_on_anomaly": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relearn anomaly detection parameters after detecting an anomaly",
			},
			"sensitivity": {
				Type: schema.TypeString, Optional: true, Default: "low", Description: "'high': Highly sensitive anomaly detection. Can lead to false positives; 'low': Low sensitivity anomaly detection. Can cause delay in detection and might not detect certain attacks. (default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityAnomalyDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityAnomalyDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityAnomalyDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityAnomalyDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityAnomalyDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityAnomalyDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityAnomalyDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityAnomalyDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityAnomalyDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityAnomalyDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityAnomalyDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityAnomalyDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityAnomalyDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityAnomalyDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityAnomalyDetection(d *schema.ResourceData) edpt.VisibilityAnomalyDetection {
	var ret edpt.VisibilityAnomalyDetection
	ret.Inst.FeatureStatus = d.Get("feature_status").(string)
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.RestartLearningOnAnomaly = d.Get("restart_learning_on_anomaly").(int)
	ret.Inst.Sensitivity = d.Get("sensitivity").(string)
	//omit uuid
	return ret
}
