package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslOcspOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_ocsp_oper`: Operational Status for the object ssl-ocsp\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslOcspOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entries",
						},
						"cached_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "Cert Name",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "Cert Name",
									},
									"subject": {
										Type: schema.TypeString, Optional: true, Description: "Cert Name",
									},
									"length": {
										Type: schema.TypeInt, Optional: true, Description: "Cert Name",
									},
									"uri": {
										Type: schema.TypeString, Optional: true, Description: "Cert Name",
									},
									"expire": {
										Type: schema.TypeInt, Optional: true, Description: "Cert Name",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Cert Name",
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

func resourceSlbSslOcspOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslOcspOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslOcspOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslOcspOperOper := setObjectSlbSslOcspOperOper(res)
		d.Set("oper", SlbSslOcspOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslOcspOperOper(ret edpt.DataSlbSslOcspOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_entries":  ret.DtSlbSslOcspOper.Oper.TotalEntries,
			"cached_entries": setSliceSlbSslOcspOperOperCachedEntries(ret.DtSlbSslOcspOper.Oper.CachedEntries),
		},
	}
}

func setSliceSlbSslOcspOperOperCachedEntries(d []edpt.SlbSslOcspOperOperCachedEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["status"] = item.Status
		in["subject"] = item.Subject
		in["length"] = item.Length
		in["uri"] = item.Uri
		in["expire"] = item.Expire
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslOcspOperOper(d []interface{}) edpt.SlbSslOcspOperOper {

	count1 := len(d)
	var ret edpt.SlbSslOcspOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalEntries = in["total_entries"].(int)
		ret.CachedEntries = getSliceSlbSslOcspOperOperCachedEntries(in["cached_entries"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslOcspOperOperCachedEntries(d []interface{}) []edpt.SlbSslOcspOperOperCachedEntries {

	count1 := len(d)
	ret := make([]edpt.SlbSslOcspOperOperCachedEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslOcspOperOperCachedEntries
		oi.Name = in["name"].(string)
		oi.Status = in["status"].(string)
		oi.Subject = in["subject"].(string)
		oi.Length = in["length"].(int)
		oi.Uri = in["uri"].(string)
		oi.Expire = in["expire"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslOcspOper(d *schema.ResourceData) edpt.SlbSslOcspOper {
	var ret edpt.SlbSslOcspOper

	ret.Oper = getObjectSlbSslOcspOperOper(d.Get("oper").([]interface{}))
	return ret
}
