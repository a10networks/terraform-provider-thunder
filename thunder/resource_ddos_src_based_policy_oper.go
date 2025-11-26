package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcBasedPolicyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_src_based_policy_oper`: Operational Status for the object src-based-policy\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSrcBasedPolicyOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_based_policy_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4_total_single_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_total_subnet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_total_single_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_total_subnet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geoloc_ipv4_total_single_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geoloc_ipv4_total_subnet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geoloc_ipv6_total_single_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geoloc_ipv6_total_subnet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geoloc_unexpanded_node": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"class_list_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"geo_location_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"all_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resource_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosSrcBasedPolicyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcBasedPolicyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcBasedPolicyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSrcBasedPolicyOperOper := setObjectDdosSrcBasedPolicyOperOper(res)
		d.Set("oper", DdosSrcBasedPolicyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSrcBasedPolicyOperOper(ret edpt.DataDdosSrcBasedPolicyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"src_based_policy_name":       ret.DtDdosSrcBasedPolicyOper.Oper.SrcBasedPolicyName,
			"ipv4_total_single_ip":        ret.DtDdosSrcBasedPolicyOper.Oper.Ipv4TotalSingleIp,
			"ipv4_total_subnet":           ret.DtDdosSrcBasedPolicyOper.Oper.Ipv4TotalSubnet,
			"ipv6_total_single_ip":        ret.DtDdosSrcBasedPolicyOper.Oper.Ipv6TotalSingleIp,
			"ipv6_total_subnet":           ret.DtDdosSrcBasedPolicyOper.Oper.Ipv6TotalSubnet,
			"geoloc_ipv4_total_single_ip": ret.DtDdosSrcBasedPolicyOper.Oper.GeolocIpv4TotalSingleIp,
			"geoloc_ipv4_total_subnet":    ret.DtDdosSrcBasedPolicyOper.Oper.GeolocIpv4TotalSubnet,
			"geoloc_ipv6_total_single_ip": ret.DtDdosSrcBasedPolicyOper.Oper.GeolocIpv6TotalSingleIp,
			"geoloc_ipv6_total_subnet":    ret.DtDdosSrcBasedPolicyOper.Oper.GeolocIpv6TotalSubnet,
			"geoloc_unexpanded_node":      ret.DtDdosSrcBasedPolicyOper.Oper.GeolocUnexpandedNode,
			"class_list_entries":          setSliceDdosSrcBasedPolicyOperOperClassListEntries(ret.DtDdosSrcBasedPolicyOper.Oper.ClassListEntries),
			"all_entries":                 ret.DtDdosSrcBasedPolicyOper.Oper.AllEntries,
			"resource_usage":              ret.DtDdosSrcBasedPolicyOper.Oper.ResourceUsage,
		},
	}
}

func setSliceDdosSrcBasedPolicyOperOperClassListEntries(d []edpt.DdosSrcBasedPolicyOperOperClassListEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["class_list_name"] = item.ClassListName
		in["geo_location_name"] = item.GeoLocationName
		result = append(result, in)
	}
	return result
}

func getObjectDdosSrcBasedPolicyOperOper(d []interface{}) edpt.DdosSrcBasedPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosSrcBasedPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcBasedPolicyName = in["src_based_policy_name"].(string)
		ret.Ipv4TotalSingleIp = in["ipv4_total_single_ip"].(int)
		ret.Ipv4TotalSubnet = in["ipv4_total_subnet"].(int)
		ret.Ipv6TotalSingleIp = in["ipv6_total_single_ip"].(int)
		ret.Ipv6TotalSubnet = in["ipv6_total_subnet"].(int)
		ret.GeolocIpv4TotalSingleIp = in["geoloc_ipv4_total_single_ip"].(int)
		ret.GeolocIpv4TotalSubnet = in["geoloc_ipv4_total_subnet"].(int)
		ret.GeolocIpv6TotalSingleIp = in["geoloc_ipv6_total_single_ip"].(int)
		ret.GeolocIpv6TotalSubnet = in["geoloc_ipv6_total_subnet"].(int)
		ret.GeolocUnexpandedNode = in["geoloc_unexpanded_node"].(int)
		ret.ClassListEntries = getSliceDdosSrcBasedPolicyOperOperClassListEntries(in["class_list_entries"].([]interface{}))
		ret.AllEntries = in["all_entries"].(int)
		ret.ResourceUsage = in["resource_usage"].(int)
	}
	return ret
}

func getSliceDdosSrcBasedPolicyOperOperClassListEntries(d []interface{}) []edpt.DdosSrcBasedPolicyOperOperClassListEntries {

	count1 := len(d)
	ret := make([]edpt.DdosSrcBasedPolicyOperOperClassListEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcBasedPolicyOperOperClassListEntries
		oi.Address = in["address"].(string)
		oi.ClassListName = in["class_list_name"].(string)
		oi.GeoLocationName = in["geo_location_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSrcBasedPolicyOper(d *schema.ResourceData) edpt.DdosSrcBasedPolicyOper {
	var ret edpt.DdosSrcBasedPolicyOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectDdosSrcBasedPolicyOperOper(d.Get("oper").([]interface{}))
	return ret
}
