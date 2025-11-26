package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemConfigMgmtPuSyncDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_config_mgmt_pu_sync_detection`: pu sync detection feature actions\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemConfigMgmtPuSyncDetectionCreate,
		UpdateContext: resourceSystemConfigMgmtPuSyncDetectionUpdate,
		ReadContext:   resourceSystemConfigMgmtPuSyncDetectionRead,
		DeleteContext: resourceSystemConfigMgmtPuSyncDetectionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable pu-sync-detection feature; 'disable': Disable pu-sync-detection feature;",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Time interval (seconds) for detection. Default is 30 seconds.",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemConfigMgmtPuSyncDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtPuSyncDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtPuSyncDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtPuSyncDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemConfigMgmtPuSyncDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtPuSyncDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtPuSyncDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtPuSyncDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemConfigMgmtPuSyncDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtPuSyncDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtPuSyncDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemConfigMgmtPuSyncDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtPuSyncDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtPuSyncDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemConfigMgmtPuSyncDetection(d *schema.ResourceData) edpt.SystemConfigMgmtPuSyncDetection {
	var ret edpt.SystemConfigMgmtPuSyncDetection
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Interval = d.Get("interval").(int)
	//omit uuid
	return ret
}
