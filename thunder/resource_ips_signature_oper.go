package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpsSignatureOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ips_signature_oper`: Operational Status for the object signature\n\n__PLACEHOLDER__",
		ReadContext: resourceIpsSignatureOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"default_action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"category": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"revision": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_enabled": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_custom": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"direction": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"attack_target": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"severity": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cve_number": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"message": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pattern_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"is_invert": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ignore_case": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"has_space_bar": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"is_multi_line": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"pattern_type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"pattern_string": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"parameter_list": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fast_pattern_index": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"max_payload_len": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"min_payload_len": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"max_uri_len": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"min_uri_len": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"http_status_code": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"http_method": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dns_query": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"threshold": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"count1": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"interval": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"track_direction": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"detection_filter": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
									"hash": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"sid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"category": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceIpsSignatureOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsSignatureOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsSignatureOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpsSignatureOperOper := setObjectIpsSignatureOperOper(res)
		d.Set("oper", IpsSignatureOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpsSignatureOperOper(ret edpt.DataIpsSignatureOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"signature_list": setSliceIpsSignatureOperOperSignatureList(ret.DtIpsSignatureOper.Oper.SignatureList),
			"sid":            ret.DtIpsSignatureOper.Oper.Sid,
			"category":       ret.DtIpsSignatureOper.Oper.Category,
			"type":           ret.DtIpsSignatureOper.Oper.Type,
		},
	}
}

func setSliceIpsSignatureOperOperSignatureList(d []edpt.IpsSignatureOperOperSignatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["default_action"] = item.DefaultAction
		in["category"] = item.Category
		in["revision"] = item.Revision
		in["is_enabled"] = item.IsEnabled
		in["is_custom"] = item.IsCustom
		in["direction"] = item.Direction
		in["attack_target"] = item.Attack_target
		in["severity"] = item.Severity
		in["type"] = item.Type
		in["cve_number"] = item.CveNumber
		in["message"] = item.Message
		in["pattern_list"] = setSliceIpsSignatureOperOperSignatureListPatternList(item.PatternList)
		in["parameter_list"] = setObjectIpsSignatureOperOperSignatureListParameterList(item.ParameterList)
		in["hash"] = item.Hash
		result = append(result, in)
	}
	return result
}

func setSliceIpsSignatureOperOperSignatureListPatternList(d []edpt.IpsSignatureOperOperSignatureListPatternList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["is_invert"] = item.IsInvert
		in["ignore_case"] = item.IgnoreCase
		in["has_space_bar"] = item.HasSpaceBar
		in["is_multi_line"] = item.IsMultiLine
		in["pattern_type"] = item.PatternType
		in["pattern_string"] = item.PatternString
		result = append(result, in)
	}
	return result
}

func setObjectIpsSignatureOperOperSignatureListParameterList(d edpt.IpsSignatureOperOperSignatureListParameterList) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["fast_pattern_index"] = d.FastPatternIndex

	in["max_payload_len"] = d.MaxPayloadLen

	in["min_payload_len"] = d.MinPayloadLen

	in["max_uri_len"] = d.MaxUriLen

	in["min_uri_len"] = d.MinUriLen

	in["http_status_code"] = d.HttpStatusCode

	in["http_method"] = d.HttpMethod

	in["dns_query"] = d.DnsQuery
	in["threshold"] = setObjectIpsSignatureOperOperSignatureListParameterListThreshold(d.Threshold)
	result = append(result, in)
	return result
}

func setObjectIpsSignatureOperOperSignatureListParameterListThreshold(d edpt.IpsSignatureOperOperSignatureListParameterListThreshold) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["count1"] = d.Count1

	in["interval"] = d.Interval

	in["track_direction"] = d.TrackDirection

	in["detection_filter"] = d.DetectionFilter
	result = append(result, in)
	return result
}

func getObjectIpsSignatureOperOper(d []interface{}) edpt.IpsSignatureOperOper {

	count1 := len(d)
	var ret edpt.IpsSignatureOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignatureList = getSliceIpsSignatureOperOperSignatureList(in["signature_list"].([]interface{}))
		ret.Sid = in["sid"].(int)
		ret.Category = in["category"].(string)
		ret.Type = in["type"].(string)
	}
	return ret
}

func getSliceIpsSignatureOperOperSignatureList(d []interface{}) []edpt.IpsSignatureOperOperSignatureList {

	count1 := len(d)
	ret := make([]edpt.IpsSignatureOperOperSignatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpsSignatureOperOperSignatureList
		oi.Sid = in["sid"].(int)
		oi.DefaultAction = in["default_action"].(string)
		oi.Category = in["category"].(string)
		oi.Revision = in["revision"].(int)
		oi.IsEnabled = in["is_enabled"].(int)
		oi.IsCustom = in["is_custom"].(int)
		oi.Direction = in["direction"].(string)
		oi.Attack_target = in["attack_target"].(string)
		oi.Severity = in["severity"].(int)
		oi.Type = in["type"].(string)
		oi.CveNumber = in["cve_number"].(string)
		oi.Message = in["message"].(string)
		oi.PatternList = getSliceIpsSignatureOperOperSignatureListPatternList(in["pattern_list"].([]interface{}))
		oi.ParameterList = getObjectIpsSignatureOperOperSignatureListParameterList(in["parameter_list"].([]interface{}))
		oi.Hash = in["hash"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpsSignatureOperOperSignatureListPatternList(d []interface{}) []edpt.IpsSignatureOperOperSignatureListPatternList {

	count1 := len(d)
	ret := make([]edpt.IpsSignatureOperOperSignatureListPatternList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpsSignatureOperOperSignatureListPatternList
		oi.IsInvert = in["is_invert"].(int)
		oi.IgnoreCase = in["ignore_case"].(int)
		oi.HasSpaceBar = in["has_space_bar"].(int)
		oi.IsMultiLine = in["is_multi_line"].(int)
		oi.PatternType = in["pattern_type"].(string)
		oi.PatternString = in["pattern_string"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectIpsSignatureOperOperSignatureListParameterList(d []interface{}) edpt.IpsSignatureOperOperSignatureListParameterList {

	count1 := len(d)
	var ret edpt.IpsSignatureOperOperSignatureListParameterList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastPatternIndex = in["fast_pattern_index"].(int)
		ret.MaxPayloadLen = in["max_payload_len"].(int)
		ret.MinPayloadLen = in["min_payload_len"].(int)
		ret.MaxUriLen = in["max_uri_len"].(int)
		ret.MinUriLen = in["min_uri_len"].(int)
		ret.HttpStatusCode = in["http_status_code"].(int)
		ret.HttpMethod = in["http_method"].(string)
		ret.DnsQuery = in["dns_query"].(int)
		ret.Threshold = getObjectIpsSignatureOperOperSignatureListParameterListThreshold(in["threshold"].([]interface{}))
	}
	return ret
}

func getObjectIpsSignatureOperOperSignatureListParameterListThreshold(d []interface{}) edpt.IpsSignatureOperOperSignatureListParameterListThreshold {

	count1 := len(d)
	var ret edpt.IpsSignatureOperOperSignatureListParameterListThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Count1 = in["count1"].(int)
		ret.Interval = in["interval"].(int)
		ret.TrackDirection = in["track_direction"].(string)
		ret.DetectionFilter = in["detection_filter"].(int)
	}
	return ret
}

func dataToEndpointIpsSignatureOper(d *schema.ResourceData) edpt.IpsSignatureOper {
	var ret edpt.IpsSignatureOper

	ret.Oper = getObjectIpsSignatureOperOper(d.Get("oper").([]interface{}))
	return ret
}
