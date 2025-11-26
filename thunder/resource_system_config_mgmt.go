package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemConfigMgmt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_config_mgmt`: Setting of Configuration Management\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemConfigMgmtCreate,
		UpdateContext: resourceSystemConfigMgmtUpdate,
		ReadContext:   resourceSystemConfigMgmtRead,
		DeleteContext: resourceSystemConfigMgmtDelete,

		Schema: map[string]*schema.Schema{
			"delete_referenced_tagged_objects": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Allow deletion of referenced tagged objects. Default option.; 'disable': Block deletion of referenced tagged objects;",
			},
			"mpm": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_workers": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set max workers count. Default is 1",
						},
						"min_idle_workers": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set minimum idle workers count. Default is 1",
						},
						"start_workers": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set starting workers count. Default is 1",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"notification": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"period": {
							Type: schema.TypeInt, Optional: true, Default: 15, Description: "Time interval (seconds) for kafka notification. Default is 15 seconds.",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pu_sync_detection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Time interval (seconds) for detection. Default is 30 seconds.",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable pu-sync-detection feature; 'disable': Disable pu-sync-detection feature;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceSystemConfigMgmtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemConfigMgmtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemConfigMgmtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemConfigMgmtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemConfigMgmtMpm1673(d []interface{}) edpt.SystemConfigMgmtMpm1673 {

	count1 := len(d)
	var ret edpt.SystemConfigMgmtMpm1673
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxWorkers = in["max_workers"].(int)
		ret.MinIdleWorkers = in["min_idle_workers"].(int)
		ret.StartWorkers = in["start_workers"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemConfigMgmtNotification1674(d []interface{}) edpt.SystemConfigMgmtNotification1674 {

	count1 := len(d)
	var ret edpt.SystemConfigMgmtNotification1674
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Period = in["period"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSystemConfigMgmtPuSyncDetection1675(d []interface{}) edpt.SystemConfigMgmtPuSyncDetection1675 {

	count1 := len(d)
	var ret edpt.SystemConfigMgmtPuSyncDetection1675
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointSystemConfigMgmt(d *schema.ResourceData) edpt.SystemConfigMgmt {
	var ret edpt.SystemConfigMgmt
	ret.Inst.DeleteReferencedTaggedObjects = d.Get("delete_referenced_tagged_objects").(string)
	ret.Inst.Mpm = getObjectSystemConfigMgmtMpm1673(d.Get("mpm").([]interface{}))
	ret.Inst.Notification = getObjectSystemConfigMgmtNotification1674(d.Get("notification").([]interface{}))
	ret.Inst.PuSyncDetection = getObjectSystemConfigMgmtPuSyncDetection1675(d.Get("pu_sync_detection").([]interface{}))
	//omit uuid
	return ret
}
