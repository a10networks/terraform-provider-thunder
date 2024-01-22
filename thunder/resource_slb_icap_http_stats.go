package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIcapHttpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_icap_http_stats`: Statistics for the object icap_http\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbIcapHttpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_200": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 200",
						},
						"status_201": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 201",
						},
						"status_202": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 202",
						},
						"status_203": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 203",
						},
						"status_204": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 204",
						},
						"status_205": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 205",
						},
						"status_206": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 206",
						},
						"status_207": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 207",
						},
						"status_100": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 100",
						},
						"status_101": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 101",
						},
						"status_102": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 102",
						},
						"status_300": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 300",
						},
						"status_301": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 301",
						},
						"status_302": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 302",
						},
						"status_303": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 303",
						},
						"status_304": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 304",
						},
						"status_305": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 305",
						},
						"status_306": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 306",
						},
						"status_307": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 307",
						},
						"status_400": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 400",
						},
						"status_401": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 401",
						},
						"status_402": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 402",
						},
						"status_403": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 403",
						},
						"status_404": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 404",
						},
						"status_405": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 405",
						},
						"status_406": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 406",
						},
						"status_407": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 407",
						},
						"status_408": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 408",
						},
						"status_409": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 409",
						},
						"status_410": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 410",
						},
						"status_411": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 411",
						},
						"status_412": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 412",
						},
						"status_413": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 413",
						},
						"status_414": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 414",
						},
						"status_415": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 415",
						},
						"status_416": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 416",
						},
						"status_417": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 417",
						},
						"status_418": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 418",
						},
						"status_422": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 422",
						},
						"status_423": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 423",
						},
						"status_424": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 424",
						},
						"status_425": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 425",
						},
						"status_426": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 426",
						},
						"status_449": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 449",
						},
						"status_450": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 450",
						},
						"status_500": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 500",
						},
						"status_501": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 501",
						},
						"status_502": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 502",
						},
						"status_503": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 503",
						},
						"status_504": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 504",
						},
						"status_505": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 505",
						},
						"status_506": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 506",
						},
						"status_507": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 507",
						},
						"status_508": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 508",
						},
						"status_509": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 509",
						},
						"status_510": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 510",
						},
						"status_1xx": {
							Type: schema.TypeInt, Optional: true, Description: "status code 1XX",
						},
						"status_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "status code 2XX",
						},
						"status_3xx": {
							Type: schema.TypeInt, Optional: true, Description: "status code 3XX",
						},
						"status_4xx": {
							Type: schema.TypeInt, Optional: true, Description: "status code 4XX",
						},
						"status_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "status code 5XX",
						},
						"status_6xx": {
							Type: schema.TypeInt, Optional: true, Description: "status code 6XX",
						},
						"status_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Status code unknown",
						},
					},
				},
			},
		},
	}
}

func resourceSlbIcapHttpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIcapHttpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIcapHttpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbIcapHttpStatsStats := setObjectSlbIcapHttpStatsStats(res)
		d.Set("stats", SlbIcapHttpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbIcapHttpStatsStats(ret edpt.DataSlbIcapHttpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status_200":     ret.DtSlbIcapHttpStats.Stats.Status_200,
			"status_201":     ret.DtSlbIcapHttpStats.Stats.Status_201,
			"status_202":     ret.DtSlbIcapHttpStats.Stats.Status_202,
			"status_203":     ret.DtSlbIcapHttpStats.Stats.Status_203,
			"status_204":     ret.DtSlbIcapHttpStats.Stats.Status_204,
			"status_205":     ret.DtSlbIcapHttpStats.Stats.Status_205,
			"status_206":     ret.DtSlbIcapHttpStats.Stats.Status_206,
			"status_207":     ret.DtSlbIcapHttpStats.Stats.Status_207,
			"status_100":     ret.DtSlbIcapHttpStats.Stats.Status_100,
			"status_101":     ret.DtSlbIcapHttpStats.Stats.Status_101,
			"status_102":     ret.DtSlbIcapHttpStats.Stats.Status_102,
			"status_300":     ret.DtSlbIcapHttpStats.Stats.Status_300,
			"status_301":     ret.DtSlbIcapHttpStats.Stats.Status_301,
			"status_302":     ret.DtSlbIcapHttpStats.Stats.Status_302,
			"status_303":     ret.DtSlbIcapHttpStats.Stats.Status_303,
			"status_304":     ret.DtSlbIcapHttpStats.Stats.Status_304,
			"status_305":     ret.DtSlbIcapHttpStats.Stats.Status_305,
			"status_306":     ret.DtSlbIcapHttpStats.Stats.Status_306,
			"status_307":     ret.DtSlbIcapHttpStats.Stats.Status_307,
			"status_400":     ret.DtSlbIcapHttpStats.Stats.Status_400,
			"status_401":     ret.DtSlbIcapHttpStats.Stats.Status_401,
			"status_402":     ret.DtSlbIcapHttpStats.Stats.Status_402,
			"status_403":     ret.DtSlbIcapHttpStats.Stats.Status_403,
			"status_404":     ret.DtSlbIcapHttpStats.Stats.Status_404,
			"status_405":     ret.DtSlbIcapHttpStats.Stats.Status_405,
			"status_406":     ret.DtSlbIcapHttpStats.Stats.Status_406,
			"status_407":     ret.DtSlbIcapHttpStats.Stats.Status_407,
			"status_408":     ret.DtSlbIcapHttpStats.Stats.Status_408,
			"status_409":     ret.DtSlbIcapHttpStats.Stats.Status_409,
			"status_410":     ret.DtSlbIcapHttpStats.Stats.Status_410,
			"status_411":     ret.DtSlbIcapHttpStats.Stats.Status_411,
			"status_412":     ret.DtSlbIcapHttpStats.Stats.Status_412,
			"status_413":     ret.DtSlbIcapHttpStats.Stats.Status_413,
			"status_414":     ret.DtSlbIcapHttpStats.Stats.Status_414,
			"status_415":     ret.DtSlbIcapHttpStats.Stats.Status_415,
			"status_416":     ret.DtSlbIcapHttpStats.Stats.Status_416,
			"status_417":     ret.DtSlbIcapHttpStats.Stats.Status_417,
			"status_418":     ret.DtSlbIcapHttpStats.Stats.Status_418,
			"status_422":     ret.DtSlbIcapHttpStats.Stats.Status_422,
			"status_423":     ret.DtSlbIcapHttpStats.Stats.Status_423,
			"status_424":     ret.DtSlbIcapHttpStats.Stats.Status_424,
			"status_425":     ret.DtSlbIcapHttpStats.Stats.Status_425,
			"status_426":     ret.DtSlbIcapHttpStats.Stats.Status_426,
			"status_449":     ret.DtSlbIcapHttpStats.Stats.Status_449,
			"status_450":     ret.DtSlbIcapHttpStats.Stats.Status_450,
			"status_500":     ret.DtSlbIcapHttpStats.Stats.Status_500,
			"status_501":     ret.DtSlbIcapHttpStats.Stats.Status_501,
			"status_502":     ret.DtSlbIcapHttpStats.Stats.Status_502,
			"status_503":     ret.DtSlbIcapHttpStats.Stats.Status_503,
			"status_504":     ret.DtSlbIcapHttpStats.Stats.Status_504,
			"status_505":     ret.DtSlbIcapHttpStats.Stats.Status_505,
			"status_506":     ret.DtSlbIcapHttpStats.Stats.Status_506,
			"status_507":     ret.DtSlbIcapHttpStats.Stats.Status_507,
			"status_508":     ret.DtSlbIcapHttpStats.Stats.Status_508,
			"status_509":     ret.DtSlbIcapHttpStats.Stats.Status_509,
			"status_510":     ret.DtSlbIcapHttpStats.Stats.Status_510,
			"status_1xx":     ret.DtSlbIcapHttpStats.Stats.Status_1xx,
			"status_2xx":     ret.DtSlbIcapHttpStats.Stats.Status_2xx,
			"status_3xx":     ret.DtSlbIcapHttpStats.Stats.Status_3xx,
			"status_4xx":     ret.DtSlbIcapHttpStats.Stats.Status_4xx,
			"status_5xx":     ret.DtSlbIcapHttpStats.Stats.Status_5xx,
			"status_6xx":     ret.DtSlbIcapHttpStats.Stats.Status_6xx,
			"status_unknown": ret.DtSlbIcapHttpStats.Stats.Status_unknown,
		},
	}
}

func getObjectSlbIcapHttpStatsStats(d []interface{}) edpt.SlbIcapHttpStatsStats {

	count1 := len(d)
	var ret edpt.SlbIcapHttpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status_200 = in["status_200"].(int)
		ret.Status_201 = in["status_201"].(int)
		ret.Status_202 = in["status_202"].(int)
		ret.Status_203 = in["status_203"].(int)
		ret.Status_204 = in["status_204"].(int)
		ret.Status_205 = in["status_205"].(int)
		ret.Status_206 = in["status_206"].(int)
		ret.Status_207 = in["status_207"].(int)
		ret.Status_100 = in["status_100"].(int)
		ret.Status_101 = in["status_101"].(int)
		ret.Status_102 = in["status_102"].(int)
		ret.Status_300 = in["status_300"].(int)
		ret.Status_301 = in["status_301"].(int)
		ret.Status_302 = in["status_302"].(int)
		ret.Status_303 = in["status_303"].(int)
		ret.Status_304 = in["status_304"].(int)
		ret.Status_305 = in["status_305"].(int)
		ret.Status_306 = in["status_306"].(int)
		ret.Status_307 = in["status_307"].(int)
		ret.Status_400 = in["status_400"].(int)
		ret.Status_401 = in["status_401"].(int)
		ret.Status_402 = in["status_402"].(int)
		ret.Status_403 = in["status_403"].(int)
		ret.Status_404 = in["status_404"].(int)
		ret.Status_405 = in["status_405"].(int)
		ret.Status_406 = in["status_406"].(int)
		ret.Status_407 = in["status_407"].(int)
		ret.Status_408 = in["status_408"].(int)
		ret.Status_409 = in["status_409"].(int)
		ret.Status_410 = in["status_410"].(int)
		ret.Status_411 = in["status_411"].(int)
		ret.Status_412 = in["status_412"].(int)
		ret.Status_413 = in["status_413"].(int)
		ret.Status_414 = in["status_414"].(int)
		ret.Status_415 = in["status_415"].(int)
		ret.Status_416 = in["status_416"].(int)
		ret.Status_417 = in["status_417"].(int)
		ret.Status_418 = in["status_418"].(int)
		ret.Status_422 = in["status_422"].(int)
		ret.Status_423 = in["status_423"].(int)
		ret.Status_424 = in["status_424"].(int)
		ret.Status_425 = in["status_425"].(int)
		ret.Status_426 = in["status_426"].(int)
		ret.Status_449 = in["status_449"].(int)
		ret.Status_450 = in["status_450"].(int)
		ret.Status_500 = in["status_500"].(int)
		ret.Status_501 = in["status_501"].(int)
		ret.Status_502 = in["status_502"].(int)
		ret.Status_503 = in["status_503"].(int)
		ret.Status_504 = in["status_504"].(int)
		ret.Status_505 = in["status_505"].(int)
		ret.Status_506 = in["status_506"].(int)
		ret.Status_507 = in["status_507"].(int)
		ret.Status_508 = in["status_508"].(int)
		ret.Status_509 = in["status_509"].(int)
		ret.Status_510 = in["status_510"].(int)
		ret.Status_1xx = in["status_1xx"].(int)
		ret.Status_2xx = in["status_2xx"].(int)
		ret.Status_3xx = in["status_3xx"].(int)
		ret.Status_4xx = in["status_4xx"].(int)
		ret.Status_5xx = in["status_5xx"].(int)
		ret.Status_6xx = in["status_6xx"].(int)
		ret.Status_unknown = in["status_unknown"].(int)
	}
	return ret
}

func dataToEndpointSlbIcapHttpStats(d *schema.ResourceData) edpt.SlbIcapHttpStats {
	var ret edpt.SlbIcapHttpStats

	ret.Stats = getObjectSlbIcapHttpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
