package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_monitor`: System monitor configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceMonitorCreate,
		UpdateContext: resourceMonitorUpdate,
		ReadContext:   resourceMonitorRead,
		DeleteContext: resourceMonitorDelete,

		Schema: map[string]*schema.Schema{
			"buffer_drop": {
				Type: schema.TypeInt, Optional: true, Default: 4000, Description: "Monitor buffer drop threshold (Threshold value)",
			},
			"buffer_usage": {
				Type: schema.TypeInt, Optional: true, Description: "Monitor IO buffer usage threshold (Threshold value)",
			},
			"conn_type0": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "Conn resource type 0 (Threshold value, default 32767)",
			},
			"conn_type1": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "Conn resource type 1 (Threshold value, default 32767)",
			},
			"conn_type2": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "Conn resource type 2 (Threshold value, default 32767)",
			},
			"conn_type3": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "Conn resource type 3 (Threshold value, default 32767)",
			},
			"conn_type4": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "Conn resource type 4 (Threshold value, default 32767)",
			},
			"ctrl_cpu": {
				Type: schema.TypeInt, Optional: true, Default: 90, Description: "Monitor control CPU threshold (Threshold value in percentage, default 90)",
			},
			"data_cpu": {
				Type: schema.TypeInt, Optional: true, Default: 90, Description: "Monitor data CPU threshold (Threshold value in percentage, default 90)",
			},
			"disk": {
				Type: schema.TypeInt, Optional: true, Default: 85, Description: "Monitor hard disk usage threshold (Threshold value in percentage, default 85)",
			},
			"memory": {
				Type: schema.TypeInt, Optional: true, Default: 95, Description: "Monitor memory usage threshold (Threshold value in percentage, default 95)",
			},
			"smp_type0": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "SMP resource type 0 (Threshold value, default 32767)",
			},
			"smp_type1": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "SMP resource type 1 (Threshold value, default 32767)",
			},
			"smp_type2": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "SMP resource type 2 (Threshold value, default 32767)",
			},
			"smp_type3": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "SMP resource type 3 (Threshold value, default 32767)",
			},
			"smp_type4": {
				Type: schema.TypeInt, Optional: true, Default: 32767, Description: "SMP resource type 4 (Threshold value, default 32767)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"warn_temp": {
				Type: schema.TypeInt, Optional: true, Description: "Monitor warning system temperature threshold (Threshold value in Celsius, default 1)",
			},
		},
	}
}
func resourceMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointMonitor(d *schema.ResourceData) edpt.Monitor {
	var ret edpt.Monitor
	ret.Inst.BufferDrop = d.Get("buffer_drop").(int)
	ret.Inst.BufferUsage = d.Get("buffer_usage").(int)
	ret.Inst.ConnType0 = d.Get("conn_type0").(int)
	ret.Inst.ConnType1 = d.Get("conn_type1").(int)
	ret.Inst.ConnType2 = d.Get("conn_type2").(int)
	ret.Inst.ConnType3 = d.Get("conn_type3").(int)
	ret.Inst.ConnType4 = d.Get("conn_type4").(int)
	ret.Inst.CtrlCpu = d.Get("ctrl_cpu").(int)
	ret.Inst.DataCpu = d.Get("data_cpu").(int)
	ret.Inst.Disk = d.Get("disk").(int)
	ret.Inst.Memory = d.Get("memory").(int)
	ret.Inst.SmpType0 = d.Get("smp_type0").(int)
	ret.Inst.SmpType1 = d.Get("smp_type1").(int)
	ret.Inst.SmpType2 = d.Get("smp_type2").(int)
	ret.Inst.SmpType3 = d.Get("smp_type3").(int)
	ret.Inst.SmpType4 = d.Get("smp_type4").(int)
	//omit uuid
	ret.Inst.WarnTemp = d.Get("warn_temp").(int)
	return ret
}
