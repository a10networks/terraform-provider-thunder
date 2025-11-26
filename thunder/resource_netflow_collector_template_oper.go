package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowCollectorTemplateOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_netflow_collector_template_oper`: Operational Status for the object template\n\n__PLACEHOLDER__",
		ReadContext: resourceNetflowCollectorTemplateOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nf_template_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exporter_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"observation_domain_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"template_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"template_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nflow_version": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"field_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"seconds_to_expire": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"partition_id": {
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

func resourceNetflowCollectorTemplateOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCollectorTemplateOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCollectorTemplateOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetflowCollectorTemplateOperOper := setObjectNetflowCollectorTemplateOperOper(res)
		d.Set("oper", NetflowCollectorTemplateOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetflowCollectorTemplateOperOper(ret edpt.DataNetflowCollectorTemplateOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"nf_template_list": setSliceNetflowCollectorTemplateOperOperNfTemplateList(ret.DtNetflowCollectorTemplateOper.Oper.NfTemplateList),
		},
	}
}

func setSliceNetflowCollectorTemplateOperOperNfTemplateList(d []edpt.NetflowCollectorTemplateOperOperNfTemplateList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["exporter_address"] = item.ExporterAddress
		in["observation_domain_id"] = item.ObservationDomainId
		in["template_id"] = item.TemplateId
		in["template_type"] = item.TemplateType
		in["nflow_version"] = item.NflowVersion
		in["field_count"] = item.FieldCount
		in["seconds_to_expire"] = item.SecondsToExpire
		in["partition_id"] = item.PartitionId
		result = append(result, in)
	}
	return result
}

func getObjectNetflowCollectorTemplateOperOper(d []interface{}) edpt.NetflowCollectorTemplateOperOper {

	count1 := len(d)
	var ret edpt.NetflowCollectorTemplateOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NfTemplateList = getSliceNetflowCollectorTemplateOperOperNfTemplateList(in["nf_template_list"].([]interface{}))
	}
	return ret
}

func getSliceNetflowCollectorTemplateOperOperNfTemplateList(d []interface{}) []edpt.NetflowCollectorTemplateOperOperNfTemplateList {

	count1 := len(d)
	ret := make([]edpt.NetflowCollectorTemplateOperOperNfTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowCollectorTemplateOperOperNfTemplateList
		oi.ExporterAddress = in["exporter_address"].(string)
		oi.ObservationDomainId = in["observation_domain_id"].(int)
		oi.TemplateId = in["template_id"].(int)
		oi.TemplateType = in["template_type"].(string)
		oi.NflowVersion = in["nflow_version"].(int)
		oi.FieldCount = in["field_count"].(int)
		oi.SecondsToExpire = in["seconds_to_expire"].(int)
		oi.PartitionId = in["partition_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetflowCollectorTemplateOper(d *schema.ResourceData) edpt.NetflowCollectorTemplateOper {
	var ret edpt.NetflowCollectorTemplateOper

	ret.Oper = getObjectNetflowCollectorTemplateOperOper(d.Get("oper").([]interface{}))
	return ret
}
