package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIcapHttpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_icap_http_oper`: Operational Status for the object icap_http\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbIcapHttpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icap_http_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_200": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_201": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_202": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_203": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_204": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_205": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_206": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_207": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_101": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_102": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_300": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_301": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_302": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_303": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_304": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_305": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_306": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_307": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_400": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_401": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_402": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_403": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_404": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_405": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_406": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_407": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_408": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_409": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_410": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_411": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_412": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_413": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_414": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_415": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_416": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_417": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_418": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_422": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_423": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_424": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_425": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_426": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_449": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_450": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_500": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_501": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_502": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_503": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_504": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_505": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_506": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_507": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_508": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_509": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_510": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_6xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbIcapHttpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIcapHttpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIcapHttpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbIcapHttpOperOper := setObjectSlbIcapHttpOperOper(res)
		d.Set("oper", SlbIcapHttpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbIcapHttpOperOper(ret edpt.DataSlbIcapHttpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"icap_http_cpu_list": setSliceSlbIcapHttpOperOperIcap_httpCpuList(ret.DtSlbIcapHttpOper.Oper.Icap_httpCpuList),
			"cpu_count":          ret.DtSlbIcapHttpOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbIcapHttpOperOperIcap_httpCpuList(d []edpt.SlbIcapHttpOperOperIcap_httpCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["status_2xx"] = item.Status_2xx
		in["status_200"] = item.Status_200
		in["status_201"] = item.Status_201
		in["status_202"] = item.Status_202
		in["status_203"] = item.Status_203
		in["status_204"] = item.Status_204
		in["status_205"] = item.Status_205
		in["status_206"] = item.Status_206
		in["status_207"] = item.Status_207
		in["status_1xx"] = item.Status_1xx
		in["status_100"] = item.Status_100
		in["status_101"] = item.Status_101
		in["status_102"] = item.Status_102
		in["status_3xx"] = item.Status_3xx
		in["status_300"] = item.Status_300
		in["status_301"] = item.Status_301
		in["status_302"] = item.Status_302
		in["status_303"] = item.Status_303
		in["status_304"] = item.Status_304
		in["status_305"] = item.Status_305
		in["status_306"] = item.Status_306
		in["status_307"] = item.Status_307
		in["status_4xx"] = item.Status_4xx
		in["status_400"] = item.Status_400
		in["status_401"] = item.Status_401
		in["status_402"] = item.Status_402
		in["status_403"] = item.Status_403
		in["status_404"] = item.Status_404
		in["status_405"] = item.Status_405
		in["status_406"] = item.Status_406
		in["status_407"] = item.Status_407
		in["status_408"] = item.Status_408
		in["status_409"] = item.Status_409
		in["status_410"] = item.Status_410
		in["status_411"] = item.Status_411
		in["status_412"] = item.Status_412
		in["status_413"] = item.Status_413
		in["status_414"] = item.Status_414
		in["status_415"] = item.Status_415
		in["status_416"] = item.Status_416
		in["status_417"] = item.Status_417
		in["status_418"] = item.Status_418
		in["status_422"] = item.Status_422
		in["status_423"] = item.Status_423
		in["status_424"] = item.Status_424
		in["status_425"] = item.Status_425
		in["status_426"] = item.Status_426
		in["status_449"] = item.Status_449
		in["status_450"] = item.Status_450
		in["status_5xx"] = item.Status_5xx
		in["status_500"] = item.Status_500
		in["status_501"] = item.Status_501
		in["status_502"] = item.Status_502
		in["status_503"] = item.Status_503
		in["status_504"] = item.Status_504
		in["status_505"] = item.Status_505
		in["status_506"] = item.Status_506
		in["status_507"] = item.Status_507
		in["status_508"] = item.Status_508
		in["status_509"] = item.Status_509
		in["status_510"] = item.Status_510
		in["status_6xx"] = item.Status_6xx
		result = append(result, in)
	}
	return result
}

func getObjectSlbIcapHttpOperOper(d []interface{}) edpt.SlbIcapHttpOperOper {

	count1 := len(d)
	var ret edpt.SlbIcapHttpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Icap_httpCpuList = getSliceSlbIcapHttpOperOperIcap_httpCpuList(in["icap_http_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbIcapHttpOperOperIcap_httpCpuList(d []interface{}) []edpt.SlbIcapHttpOperOperIcap_httpCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbIcapHttpOperOperIcap_httpCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbIcapHttpOperOperIcap_httpCpuList
		oi.Status_2xx = in["status_2xx"].(int)
		oi.Status_200 = in["status_200"].(int)
		oi.Status_201 = in["status_201"].(int)
		oi.Status_202 = in["status_202"].(int)
		oi.Status_203 = in["status_203"].(int)
		oi.Status_204 = in["status_204"].(int)
		oi.Status_205 = in["status_205"].(int)
		oi.Status_206 = in["status_206"].(int)
		oi.Status_207 = in["status_207"].(int)
		oi.Status_1xx = in["status_1xx"].(int)
		oi.Status_100 = in["status_100"].(int)
		oi.Status_101 = in["status_101"].(int)
		oi.Status_102 = in["status_102"].(int)
		oi.Status_3xx = in["status_3xx"].(int)
		oi.Status_300 = in["status_300"].(int)
		oi.Status_301 = in["status_301"].(int)
		oi.Status_302 = in["status_302"].(int)
		oi.Status_303 = in["status_303"].(int)
		oi.Status_304 = in["status_304"].(int)
		oi.Status_305 = in["status_305"].(int)
		oi.Status_306 = in["status_306"].(int)
		oi.Status_307 = in["status_307"].(int)
		oi.Status_4xx = in["status_4xx"].(int)
		oi.Status_400 = in["status_400"].(int)
		oi.Status_401 = in["status_401"].(int)
		oi.Status_402 = in["status_402"].(int)
		oi.Status_403 = in["status_403"].(int)
		oi.Status_404 = in["status_404"].(int)
		oi.Status_405 = in["status_405"].(int)
		oi.Status_406 = in["status_406"].(int)
		oi.Status_407 = in["status_407"].(int)
		oi.Status_408 = in["status_408"].(int)
		oi.Status_409 = in["status_409"].(int)
		oi.Status_410 = in["status_410"].(int)
		oi.Status_411 = in["status_411"].(int)
		oi.Status_412 = in["status_412"].(int)
		oi.Status_413 = in["status_413"].(int)
		oi.Status_414 = in["status_414"].(int)
		oi.Status_415 = in["status_415"].(int)
		oi.Status_416 = in["status_416"].(int)
		oi.Status_417 = in["status_417"].(int)
		oi.Status_418 = in["status_418"].(int)
		oi.Status_422 = in["status_422"].(int)
		oi.Status_423 = in["status_423"].(int)
		oi.Status_424 = in["status_424"].(int)
		oi.Status_425 = in["status_425"].(int)
		oi.Status_426 = in["status_426"].(int)
		oi.Status_449 = in["status_449"].(int)
		oi.Status_450 = in["status_450"].(int)
		oi.Status_5xx = in["status_5xx"].(int)
		oi.Status_500 = in["status_500"].(int)
		oi.Status_501 = in["status_501"].(int)
		oi.Status_502 = in["status_502"].(int)
		oi.Status_503 = in["status_503"].(int)
		oi.Status_504 = in["status_504"].(int)
		oi.Status_505 = in["status_505"].(int)
		oi.Status_506 = in["status_506"].(int)
		oi.Status_507 = in["status_507"].(int)
		oi.Status_508 = in["status_508"].(int)
		oi.Status_509 = in["status_509"].(int)
		oi.Status_510 = in["status_510"].(int)
		oi.Status_6xx = in["status_6xx"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbIcapHttpOper(d *schema.ResourceData) edpt.SlbIcapHttpOper {
	var ret edpt.SlbIcapHttpOper

	ret.Oper = getObjectSlbIcapHttpOperOper(d.Get("oper").([]interface{}))
	return ret
}
