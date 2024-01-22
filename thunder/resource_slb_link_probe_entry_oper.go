package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbLinkProbeEntryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_link_probe_entry_oper`: Operational Status for the object entry\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbLinkProbeEntryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rserver_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sg_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"probe_template_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"target_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ref_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"curr_probe_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"test_interval": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"probe_interval": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"probes_per_test": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_method": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_status_code": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_avg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples3": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples4": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples5": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples6": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples7": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples8": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples9": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rtt_samples10": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSlbLinkProbeEntryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbLinkProbeEntryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbLinkProbeEntryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbLinkProbeEntryOperOper := setObjectSlbLinkProbeEntryOperOper(res)
		d.Set("oper", SlbLinkProbeEntryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbLinkProbeEntryOperOper(ret edpt.DataSlbLinkProbeEntryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list": setSliceSlbLinkProbeEntryOperOperEntryList(ret.DtSlbLinkProbeEntryOper.Oper.EntryList),
		},
	}
}

func setSliceSlbLinkProbeEntryOperOperEntryList(d []edpt.SlbLinkProbeEntryOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rserver_name"] = item.Rserver_name
		in["sg_name"] = item.Sg_name
		in["probe_template_name"] = item.Probe_template_name
		in["domain_name"] = item.Domain_name
		in["url"] = item.Url
		in["target_type"] = item.Target_type
		in["ip_type"] = item.Ip_type
		in["ipv4_addr"] = item.Ipv4_addr
		in["ipv6_addr"] = item.Ipv6_addr
		in["ref_count"] = item.Ref_count
		in["data_cpu_id"] = item.Data_cpu_id
		in["curr_probe_count"] = item.Curr_probe_count
		in["test_interval"] = item.Test_interval
		in["probe_interval"] = item.Probe_interval
		in["probes_per_test"] = item.Probes_per_test
		in["rtt_method"] = item.Rtt_method
		in["last_status_code"] = item.Last_status_code
		in["rtt_avg"] = item.Rtt_avg
		in["rtt_samples1"] = item.Rtt_samples1
		in["rtt_samples2"] = item.Rtt_samples2
		in["rtt_samples3"] = item.Rtt_samples3
		in["rtt_samples4"] = item.Rtt_samples4
		in["rtt_samples5"] = item.Rtt_samples5
		in["rtt_samples6"] = item.Rtt_samples6
		in["rtt_samples7"] = item.Rtt_samples7
		in["rtt_samples8"] = item.Rtt_samples8
		in["rtt_samples9"] = item.Rtt_samples9
		in["rtt_samples10"] = item.Rtt_samples10
		result = append(result, in)
	}
	return result
}

func getObjectSlbLinkProbeEntryOperOper(d []interface{}) edpt.SlbLinkProbeEntryOperOper {

	count1 := len(d)
	var ret edpt.SlbLinkProbeEntryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceSlbLinkProbeEntryOperOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbLinkProbeEntryOperOperEntryList(d []interface{}) []edpt.SlbLinkProbeEntryOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.SlbLinkProbeEntryOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbLinkProbeEntryOperOperEntryList
		oi.Rserver_name = in["rserver_name"].(string)
		oi.Sg_name = in["sg_name"].(string)
		oi.Probe_template_name = in["probe_template_name"].(string)
		oi.Domain_name = in["domain_name"].(string)
		oi.Url = in["url"].(string)
		oi.Target_type = in["target_type"].(int)
		oi.Ip_type = in["ip_type"].(int)
		oi.Ipv4_addr = in["ipv4_addr"].(string)
		oi.Ipv6_addr = in["ipv6_addr"].(string)
		oi.Ref_count = in["ref_count"].(int)
		oi.Data_cpu_id = in["data_cpu_id"].(int)
		oi.Curr_probe_count = in["curr_probe_count"].(int)
		oi.Test_interval = in["test_interval"].(int)
		oi.Probe_interval = in["probe_interval"].(int)
		oi.Probes_per_test = in["probes_per_test"].(int)
		oi.Rtt_method = in["rtt_method"].(int)
		oi.Last_status_code = in["last_status_code"].(int)
		oi.Rtt_avg = in["rtt_avg"].(int)
		oi.Rtt_samples1 = in["rtt_samples1"].(int)
		oi.Rtt_samples2 = in["rtt_samples2"].(int)
		oi.Rtt_samples3 = in["rtt_samples3"].(int)
		oi.Rtt_samples4 = in["rtt_samples4"].(int)
		oi.Rtt_samples5 = in["rtt_samples5"].(int)
		oi.Rtt_samples6 = in["rtt_samples6"].(int)
		oi.Rtt_samples7 = in["rtt_samples7"].(int)
		oi.Rtt_samples8 = in["rtt_samples8"].(int)
		oi.Rtt_samples9 = in["rtt_samples9"].(int)
		oi.Rtt_samples10 = in["rtt_samples10"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbLinkProbeEntryOper(d *schema.ResourceData) edpt.SlbLinkProbeEntryOper {
	var ret edpt.SlbLinkProbeEntryOper

	ret.Oper = getObjectSlbLinkProbeEntryOperOper(d.Get("oper").([]interface{}))
	return ret
}
