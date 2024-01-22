package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnssecOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_dnssec_oper`: Operational Status for the object dnssec\n\n__PLACEHOLDER__",
		ReadContext: resourceDnssecOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"soa_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"soa_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dnskey_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dnskey_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ds_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ds_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nsec3param_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nsec3param_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nsec_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nsec_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nsec3_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nsec3_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rrsig_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rrsig_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"a_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"a_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aaaa_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aaaa_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ptr_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ptr_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cname_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cname_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ns_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ns_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mx_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mx_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"srv_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"srv_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"txt_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"txt_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"zone_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"zone_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"domain_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"domain_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"table_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"table_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reference_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reference_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"array_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"array_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rrsig2_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rrsig2_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_memory": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_objects": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDnssecOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DnssecOperOper := setObjectDnssecOperOper(res)
		d.Set("oper", DnssecOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDnssecOperOper(ret edpt.DataDnssecOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"soa_memory":         ret.DtDnssecOper.Oper.Soa_memory,
			"soa_objects":        ret.DtDnssecOper.Oper.Soa_objects,
			"dnskey_memory":      ret.DtDnssecOper.Oper.Dnskey_memory,
			"dnskey_objects":     ret.DtDnssecOper.Oper.Dnskey_objects,
			"ds_memory":          ret.DtDnssecOper.Oper.Ds_memory,
			"ds_objects":         ret.DtDnssecOper.Oper.Ds_objects,
			"nsec3param_memory":  ret.DtDnssecOper.Oper.Nsec3param_memory,
			"nsec3param_objects": ret.DtDnssecOper.Oper.Nsec3param_objects,
			"nsec_memory":        ret.DtDnssecOper.Oper.Nsec_memory,
			"nsec_objects":       ret.DtDnssecOper.Oper.Nsec_objects,
			"nsec3_memory":       ret.DtDnssecOper.Oper.Nsec3_memory,
			"nsec3_objects":      ret.DtDnssecOper.Oper.Nsec3_objects,
			"rrsig_memory":       ret.DtDnssecOper.Oper.Rrsig_memory,
			"rrsig_objects":      ret.DtDnssecOper.Oper.Rrsig_objects,
			"a_memory":           ret.DtDnssecOper.Oper.A_memory,
			"a_objects":          ret.DtDnssecOper.Oper.A_objects,
			"aaaa_memory":        ret.DtDnssecOper.Oper.Aaaa_memory,
			"aaaa_objects":       ret.DtDnssecOper.Oper.Aaaa_objects,
			"ptr_memory":         ret.DtDnssecOper.Oper.Ptr_memory,
			"ptr_objects":        ret.DtDnssecOper.Oper.Ptr_objects,
			"cname_memory":       ret.DtDnssecOper.Oper.Cname_memory,
			"cname_objects":      ret.DtDnssecOper.Oper.Cname_objects,
			"ns_memory":          ret.DtDnssecOper.Oper.Ns_memory,
			"ns_objects":         ret.DtDnssecOper.Oper.Ns_objects,
			"mx_memory":          ret.DtDnssecOper.Oper.Mx_memory,
			"mx_objects":         ret.DtDnssecOper.Oper.Mx_objects,
			"srv_memory":         ret.DtDnssecOper.Oper.Srv_memory,
			"srv_objects":        ret.DtDnssecOper.Oper.Srv_objects,
			"txt_memory":         ret.DtDnssecOper.Oper.Txt_memory,
			"txt_objects":        ret.DtDnssecOper.Oper.Txt_objects,
			"zone_memory":        ret.DtDnssecOper.Oper.Zone_memory,
			"zone_objects":       ret.DtDnssecOper.Oper.Zone_objects,
			"domain_memory":      ret.DtDnssecOper.Oper.Domain_memory,
			"domain_objects":     ret.DtDnssecOper.Oper.Domain_objects,
			"table_memory":       ret.DtDnssecOper.Oper.Table_memory,
			"table_objects":      ret.DtDnssecOper.Oper.Table_objects,
			"reference_memory":   ret.DtDnssecOper.Oper.Reference_memory,
			"reference_objects":  ret.DtDnssecOper.Oper.Reference_objects,
			"array_memory":       ret.DtDnssecOper.Oper.Array_memory,
			"array_objects":      ret.DtDnssecOper.Oper.Array_objects,
			"rrsig2_memory":      ret.DtDnssecOper.Oper.Rrsig2_memory,
			"rrsig2_objects":     ret.DtDnssecOper.Oper.Rrsig2_objects,
			"total_memory":       ret.DtDnssecOper.Oper.Total_memory,
			"total_objects":      ret.DtDnssecOper.Oper.Total_objects,
		},
	}
}

func getObjectDnssecOperOper(d []interface{}) edpt.DnssecOperOper {

	count1 := len(d)
	var ret edpt.DnssecOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Soa_memory = in["soa_memory"].(int)
		ret.Soa_objects = in["soa_objects"].(int)
		ret.Dnskey_memory = in["dnskey_memory"].(int)
		ret.Dnskey_objects = in["dnskey_objects"].(int)
		ret.Ds_memory = in["ds_memory"].(int)
		ret.Ds_objects = in["ds_objects"].(int)
		ret.Nsec3param_memory = in["nsec3param_memory"].(int)
		ret.Nsec3param_objects = in["nsec3param_objects"].(int)
		ret.Nsec_memory = in["nsec_memory"].(int)
		ret.Nsec_objects = in["nsec_objects"].(int)
		ret.Nsec3_memory = in["nsec3_memory"].(int)
		ret.Nsec3_objects = in["nsec3_objects"].(int)
		ret.Rrsig_memory = in["rrsig_memory"].(int)
		ret.Rrsig_objects = in["rrsig_objects"].(int)
		ret.A_memory = in["a_memory"].(int)
		ret.A_objects = in["a_objects"].(int)
		ret.Aaaa_memory = in["aaaa_memory"].(int)
		ret.Aaaa_objects = in["aaaa_objects"].(int)
		ret.Ptr_memory = in["ptr_memory"].(int)
		ret.Ptr_objects = in["ptr_objects"].(int)
		ret.Cname_memory = in["cname_memory"].(int)
		ret.Cname_objects = in["cname_objects"].(int)
		ret.Ns_memory = in["ns_memory"].(int)
		ret.Ns_objects = in["ns_objects"].(int)
		ret.Mx_memory = in["mx_memory"].(int)
		ret.Mx_objects = in["mx_objects"].(int)
		ret.Srv_memory = in["srv_memory"].(int)
		ret.Srv_objects = in["srv_objects"].(int)
		ret.Txt_memory = in["txt_memory"].(int)
		ret.Txt_objects = in["txt_objects"].(int)
		ret.Zone_memory = in["zone_memory"].(int)
		ret.Zone_objects = in["zone_objects"].(int)
		ret.Domain_memory = in["domain_memory"].(int)
		ret.Domain_objects = in["domain_objects"].(int)
		ret.Table_memory = in["table_memory"].(int)
		ret.Table_objects = in["table_objects"].(int)
		ret.Reference_memory = in["reference_memory"].(int)
		ret.Reference_objects = in["reference_objects"].(int)
		ret.Array_memory = in["array_memory"].(int)
		ret.Array_objects = in["array_objects"].(int)
		ret.Rrsig2_memory = in["rrsig2_memory"].(int)
		ret.Rrsig2_objects = in["rrsig2_objects"].(int)
		ret.Total_memory = in["total_memory"].(int)
		ret.Total_objects = in["total_objects"].(int)
	}
	return ret
}

func dataToEndpointDnssecOper(d *schema.ResourceData) edpt.DnssecOper {
	var ret edpt.DnssecOper

	ret.Oper = getObjectDnssecOperOper(d.Get("oper").([]interface{}))
	return ret
}
