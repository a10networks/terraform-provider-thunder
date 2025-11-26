package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_netflow_monitor_oper`: Operational Status for the object monitor\n\n__PLACEHOLDER__",
		ReadContext: resourceNetflowMonitorOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of netflow monitor",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"destination": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_ip_use_mgmt": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"source_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"flow_timeout": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"counter_polling_interval": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"resend_template_per_records": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"resend_template_timeout": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceNetflowMonitorOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetflowMonitorOperOper := setObjectNetflowMonitorOperOper(res)
		d.Set("oper", NetflowMonitorOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetflowMonitorOperOper(ret edpt.DataNetflowMonitorOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"protocol":                    ret.DtNetflowMonitorOper.Oper.Protocol,
			"status":                      ret.DtNetflowMonitorOper.Oper.Status,
			"filter":                      ret.DtNetflowMonitorOper.Oper.Filter,
			"destination":                 ret.DtNetflowMonitorOper.Oper.Destination,
			"source_ip_use_mgmt":          ret.DtNetflowMonitorOper.Oper.SourceIpUseMgmt,
			"source_ip_addr":              ret.DtNetflowMonitorOper.Oper.SourceIpAddr,
			"flow_timeout":                ret.DtNetflowMonitorOper.Oper.FlowTimeout,
			"counter_polling_interval":    ret.DtNetflowMonitorOper.Oper.CounterPollingInterval,
			"resend_template_per_records": ret.DtNetflowMonitorOper.Oper.ResendTemplatePerRecords,
			"resend_template_timeout":     ret.DtNetflowMonitorOper.Oper.ResendTemplateTimeout,
		},
	}
}

func getObjectNetflowMonitorOperOper(d []interface{}) edpt.NetflowMonitorOperOper {

	count1 := len(d)
	var ret edpt.NetflowMonitorOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol = in["protocol"].(string)
		ret.Status = in["status"].(string)
		ret.Filter = in["filter"].(string)
		ret.Destination = in["destination"].(string)
		ret.SourceIpUseMgmt = in["source_ip_use_mgmt"].(string)
		ret.SourceIpAddr = in["source_ip_addr"].(string)
		ret.FlowTimeout = in["flow_timeout"].(string)
		ret.CounterPollingInterval = in["counter_polling_interval"].(string)
		ret.ResendTemplatePerRecords = in["resend_template_per_records"].(string)
		ret.ResendTemplateTimeout = in["resend_template_timeout"].(string)
	}
	return ret
}

func dataToEndpointNetflowMonitorOper(d *schema.ResourceData) edpt.NetflowMonitorOper {
	var ret edpt.NetflowMonitorOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectNetflowMonitorOperOper(d.Get("oper").([]interface{}))
	return ret
}
