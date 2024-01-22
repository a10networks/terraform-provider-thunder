package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLocalLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_local_logging`: Configure local logging/persistant storage of FW logs\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsLocalLoggingCreate,
		UpdateContext: resourceAcosEventsLocalLoggingUpdate,
		ReadContext:   resourceAcosEventsLocalLoggingRead,
		DeleteContext: resourceAcosEventsLocalLoggingDelete,

		Schema: map[string]*schema.Schema{
			"delete_old_logs_in_disk": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Operational command to delete the old logs stored in disk",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local-logging when FW log servers are down (Default: Not enabled)",
			},
			"max_disk_space": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Configure Max disk space in MB to be used for storing the logs (Default: 100MB)",
			},
			"rate_limit": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Configure number of logs per second to be stored in disk (Default: 1000)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'init-pass': Local logging Init Successful; 'init-fail': Local logging Init Fail; 'freed': Local logging Stopped; 'disk-over-thres': Number of logs Dropped, Disk reached threshold; 'rate-limited': Number of logs Dropped, Rate limited; 'not-inited': Number of logs Dropped, Local logging not inited; 'sent-to-store': Number of logs sent to be stored; 'sent-to-store-fail': Number of Logs sent to be stored Failed; 'store-fail': Number of logs failed to be stored; 'in-logs': Number of logs successfully stored; 'in-bytes': Number of bytes successfully stored; 'in-logs-backlog': Number of backlogs loaded from disk; 'in-bytes-backlog': Number of backlog bytes loaded from disk; 'in-store-fail-no-space': Number of logs Dropped, failed without disk space; 'in-discard-logs': Number of old logs discarded to fit in new logs; 'in-discard-bytes': Number of old bytes discarded to fit in new logs; 'out-logs': Number of logs sent to log servers; 'out-bytes': Number of bytes sent to log-servers; 'out-error': Number of errors during send; 'remaining-logs': Total number of remaining logs yet to be sent; 'remaining-bytes': Total number of remaining bytes yet to be sent; 'moved-to-delq': Local Logging moved to delq to be deleted; 'out-retry': Number of attempted retries to send logs; 'out-retry-fail': Number of retries failed with error; 'curr-total-chunks': Current Number of blocks; 'curr-mem-chunks': Current blocks in memory; 'curr-fs-chunks': Current blocks in file system; 'curr-fs-chunks-up': Current blocks in file system loaded in memory; 'curr-fs-chunks-down': Current blocks in file system not loaded in memory; 'in-logs-agg': Total Aggregate, Number of logs successfully stored; 'in-bytes-agg': Total Aggregate, Number of bytes successfully stored; 'in-logs-backlog-agg': Total Aggregate, Number of backlogs loaded from disk; 'in-bytes-backlog-agg': Total Aggregate, Number of backlog bytes loaded from disk; 'in-store-fail-no-space-agg': Total Aggregate, Number of logs Dropped, failed without disk space; 'in-discard-logs-agg': Total Aggregate, Number of old logs discarded to fit in new logs; 'in-discard-bytes-agg': Total Aggregate, Number of old bytes discarded to fit in new logs; 'out-logs-agg': Total Aggregate, Number of logs sent to log servers; 'out-bytes-agg': Total Aggregate, Number of bytes sent to log-servers; 'out-error-agg': Total Aggregate, Number of errors during send; 'out-retry-agg': Total Aggregate, Number of attempted retries to send logs; 'out-retry-fail-agg': Total Aggregate, Number of retries failed with error; 'in-logs-curr-agg': Current Aggregate, Number of logs successfully stored; 'in-bytes-curr-agg': Current Aggregate, Number of bytes successfully stored; 'in-logs-backlog-curr-agg': Current Aggregate, Number of backlogs loaded from disk; 'in-bytes-backlog-curr-agg': Current Aggregate, Number of backlog bytes loaded from disk; 'in-discard-logs-curr-agg': Current Aggregate, Number of old logs discarded to fit in new logs; 'in-discard-bytes-curr-agg': Current Aggregate, Number of old bytes discarded to fit in new logs; 'out-logs-curr-agg': Current Aggregate, Number of logs sent to log servers; 'out-bytes-curr-agg': Current Aggregate, Number of bytes sent to log-servers;",
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
func resourceAcosEventsLocalLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLocalLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLocalLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLocalLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsLocalLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLocalLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLocalLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLocalLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsLocalLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLocalLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLocalLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsLocalLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLocalLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLocalLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosEventsLocalLoggingSamplingEnable(d []interface{}) []edpt.AcosEventsLocalLoggingSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AcosEventsLocalLoggingSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsLocalLoggingSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsLocalLogging(d *schema.ResourceData) edpt.AcosEventsLocalLogging {
	var ret edpt.AcosEventsLocalLogging
	ret.Inst.DeleteOldLogsInDisk = d.Get("delete_old_logs_in_disk").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.MaxDiskSpace = d.Get("max_disk_space").(int)
	ret.Inst.RateLimit = d.Get("rate_limit").(int)
	ret.Inst.SamplingEnable = getSliceAcosEventsLocalLoggingSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
