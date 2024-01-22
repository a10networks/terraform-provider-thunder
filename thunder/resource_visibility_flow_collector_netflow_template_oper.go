package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorNetflowTemplateOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_flow_collector_netflow_template_oper`: Operational Status for the object template\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityFlowCollectorNetflowTemplateOperRead,

		Schema: map[string]*schema.Schema{
			"detail": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityFlowCollectorNetflowTemplateOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowTemplateOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowTemplateOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityFlowCollectorNetflowTemplateOperDetail := setObjectVisibilityFlowCollectorNetflowTemplateOperDetail(res)
		d.Set("detail", VisibilityFlowCollectorNetflowTemplateOperDetail)
		VisibilityFlowCollectorNetflowTemplateOperOper := setObjectVisibilityFlowCollectorNetflowTemplateOperOper(res)
		d.Set("oper", VisibilityFlowCollectorNetflowTemplateOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityFlowCollectorNetflowTemplateOperDetail(ret edpt.DataVisibilityFlowCollectorNetflowTemplateOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVisibilityFlowCollectorNetflowTemplateOperDetailOper(ret.DtVisibilityFlowCollectorNetflowTemplateOper.Detail.Oper),
		},
	}
}

func setObjectVisibilityFlowCollectorNetflowTemplateOperDetailOper(d edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["nf_template_list"] = setSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList(d.NfTemplateList)
	result = append(result, in)
	return result
}

func setSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList(d []edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList) []map[string]interface{} {
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
		in["template_field_list"] = setSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList(item.TemplateFieldList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList(d []edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList) []map[string]interface{} {
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

func setObjectVisibilityFlowCollectorNetflowTemplateOperOper(ret edpt.DataVisibilityFlowCollectorNetflowTemplateOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"nf_template_list": setSliceVisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList(ret.DtVisibilityFlowCollectorNetflowTemplateOper.Oper.NfTemplateList),
		},
	}
}

func setSliceVisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList(d []edpt.VisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList) []map[string]interface{} {
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
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityFlowCollectorNetflowTemplateOperDetail(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplateOperDetail {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowTemplateOperDetail
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityFlowCollectorNetflowTemplateOperDetailOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityFlowCollectorNetflowTemplateOperDetailOper(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOper {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NfTemplateList = getSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList(in["nf_template_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList(d []interface{}) []edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateList
		oi.ExporterAddress = in["exporter_address"].(string)
		oi.ObservationDomainId = in["observation_domain_id"].(int)
		oi.TemplateId = in["template_id"].(int)
		oi.NflowVersion = in["nflow_version"].(int)
		oi.FieldCount = in["field_count"].(int)
		oi.SecondsToExpire = in["seconds_to_expire"].(int)
		oi.PartitionId = in["partition_id"].(int)
		oi.TemplateFieldList = getSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList(in["template_field_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList(d []interface{}) []edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowTemplateOperDetailOperNfTemplateListTemplateFieldList
		oi.Id1 = in["id1"].(int)
		oi.Length = in["length"].(int)
		oi.EnterpriseField = in["enterprise_field"].(string)
		oi.VariableLength = in["variable_length"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityFlowCollectorNetflowTemplateOperOper(d []interface{}) edpt.VisibilityFlowCollectorNetflowTemplateOperOper {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowTemplateOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NfTemplateList = getSliceVisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList(in["nf_template_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList(d []interface{}) []edpt.VisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList {

	count1 := len(d)
	ret := make([]edpt.VisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFlowCollectorNetflowTemplateOperOperNfTemplateList
		oi.ExporterAddress = in["exporter_address"].(string)
		oi.ObservationDomainId = in["observation_domain_id"].(int)
		oi.TemplateId = in["template_id"].(int)
		oi.NflowVersion = in["nflow_version"].(int)
		oi.FieldCount = in["field_count"].(int)
		oi.SecondsToExpire = in["seconds_to_expire"].(int)
		oi.PartitionId = in["partition_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityFlowCollectorNetflowTemplateOper(d *schema.ResourceData) edpt.VisibilityFlowCollectorNetflowTemplateOper {
	var ret edpt.VisibilityFlowCollectorNetflowTemplateOper

	ret.Detail = getObjectVisibilityFlowCollectorNetflowTemplateOperDetail(d.Get("detail").([]interface{}))

	ret.Oper = getObjectVisibilityFlowCollectorNetflowTemplateOperOper(d.Get("oper").([]interface{}))
	return ret
}
