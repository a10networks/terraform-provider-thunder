package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorNetflowTemplateDetailOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_flow_collector_netflow_template_detail_oper`: Operational Status for the object detail\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityFlowCollectorNetflowTemplateDetailOperRead,

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
									"template_field_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id1": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"length": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"enterprise_field": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"variable_length": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
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

func resourceVisibilityFlowCollectorNetflowTemplateDetailOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowTemplateDetailOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowTemplateDetailOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityFlowCollectorNetflowTemplateDetailOperOper := setObjectVisibilityFlowCollectorNetflowTemplateDetailOperOper(res)
		d.Set("oper", VisibilityFlowCollectorNetflowTemplateDetailOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityFlowCollectorNetflowTemplateDetailOperOper(ret edpt.DataVisibilityFlowCollectorNetflowTemplateDetailOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"nf_template_list": setSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList(ret.DtVisibilityFlowCollectorNetflowTemplateDetailOper.Oper.NfTemplateList),
		},
	}
}

func setSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList(d []edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["exporter_address"] = item.ExporterAddress
		in["observation_domain_id"] = item.ObservationDomainId
		in["template_id"] = item.TemplateId
		in["nflow_version"] = item.NflowVersion
		in["field_count"] = item.FieldCount
		in["seconds_to_expire"] = item.SecondsToExpire
		in["partition_id"] = item.PartitionId
		in["template_field_list"] = setSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList(item.TemplateFieldList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList(d []edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["id1"] = item.Id1
		in["length"] = item.Length
		in["enterprise_field"] = item.EnterpriseField
		in["variable_length"] = item.VariableLength
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityFlowCollectorNetflowTemplateDetailOperOper(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOper {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NfTemplateList = getSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList(in["nf_template_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList(d []interface{}) []edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateList
		oi.ExporterAddress = in["exporter_address"].(string)
		oi.ObservationDomainId = in["observation_domain_id"].(int)
		oi.TemplateId = in["template_id"].(int)
		oi.NflowVersion = in["nflow_version"].(int)
		oi.FieldCount = in["field_count"].(int)
		oi.SecondsToExpire = in["seconds_to_expire"].(int)
		oi.PartitionId = in["partition_id"].(int)
		oi.TemplateFieldList = getSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList(in["template_field_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList(d []interface{}) []edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowTemplateDetailOperOperNfTemplateListTemplateFieldList
		oi.Id1 = in["id1"].(int)
		oi.Length = in["length"].(int)
		oi.EnterpriseField = in["enterprise_field"].(string)
		oi.VariableLength = in["variable_length"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityFlowCollectorNetflowTemplateDetailOper(d *schema.ResourceData) edpt.VisibilityFlowCollectorNetflowTemplateDetailOper {
	var ret edpt.VisibilityFlowCollectorNetflowTemplateDetailOper

	ret.Oper = getObjectVisibilityFlowCollectorNetflowTemplateDetailOperOper(d.Get("oper").([]interface{}))
	return ret
}
