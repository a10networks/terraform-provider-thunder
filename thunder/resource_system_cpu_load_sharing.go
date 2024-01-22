package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuLoadSharing() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_cpu_load_sharing`: Redistribute packets uniformly to all CPUs during overload situations\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemCpuLoadSharingCreate,
		UpdateContext: resourceSystemCpuLoadSharingUpdate,
		ReadContext:   resourceSystemCpuLoadSharingRead,
		DeleteContext: resourceSystemCpuLoadSharingDelete,

		Schema: map[string]*schema.Schema{
			"cpu_usage": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"low": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "CPU usage threshold (percentage) that will restore the normal packet distribution (default: 60)",
						},
						"high": {
							Type: schema.TypeInt, Optional: true, Default: 75, Description: "CPU usage threshold (percentage) that will trigger the redistribution (default: 75)",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU load sharing in overload situations",
			},
			"others": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new non TCP/UDP IP sessions",
			},
			"packets_per_second": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min": {
							Type: schema.TypeInt, Optional: true, Default: 100000, Description: "Minimum packets-per-second threshold (per CPU) before redistribution will take effect (Minimum packets-per-second threshold (per CPU) before redistribution will take effect (default: 100000))",
						},
					},
				},
			},
			"tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new TCP sessions",
			},
			"udp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow redistribution of new UDP sessions",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemCpuLoadSharingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuLoadSharingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuLoadSharing(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuLoadSharingRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemCpuLoadSharingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuLoadSharingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuLoadSharing(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuLoadSharingRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemCpuLoadSharingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuLoadSharingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuLoadSharing(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemCpuLoadSharingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuLoadSharingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuLoadSharing(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemCpuLoadSharingCpuUsage(d []interface{}) edpt.SystemCpuLoadSharingCpuUsage {

	count1 := len(d)
	var ret edpt.SystemCpuLoadSharingCpuUsage
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Low = in["low"].(int)
		ret.High = in["high"].(int)
	}
	return ret
}

func getObjectSystemCpuLoadSharingPacketsPerSecond(d []interface{}) edpt.SystemCpuLoadSharingPacketsPerSecond {

	count1 := len(d)
	var ret edpt.SystemCpuLoadSharingPacketsPerSecond
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Min = in["min"].(int)
	}
	return ret
}

func dataToEndpointSystemCpuLoadSharing(d *schema.ResourceData) edpt.SystemCpuLoadSharing {
	var ret edpt.SystemCpuLoadSharing
	ret.Inst.CpuUsage = getObjectSystemCpuLoadSharingCpuUsage(d.Get("cpu_usage").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.PacketsPerSecond = getObjectSystemCpuLoadSharingPacketsPerSecond(d.Get("packets_per_second").([]interface{}))
	ret.Inst.Tcp = d.Get("tcp").(int)
	ret.Inst.Udp = d.Get("udp").(int)
	//omit uuid
	return ret
}
