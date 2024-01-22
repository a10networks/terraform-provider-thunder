package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCrlSrcipOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_crl_srcip_oper`: Operational Status for the object crl-srcip\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbCrlSrcipOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crl_srcip_lockedout_ips": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"start": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"end": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"lockedout_ips_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbCrlSrcipOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCrlSrcipOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCrlSrcipOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbCrlSrcipOperOper := setObjectSlbCrlSrcipOperOper(res)
		d.Set("oper", SlbCrlSrcipOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbCrlSrcipOperOper(ret edpt.DataSlbCrlSrcipOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"crl_srcip_lockedout_ips": setSliceSlbCrlSrcipOperOperCrl_srcip_lockedout_ips(ret.DtSlbCrlSrcipOper.Oper.Crl_srcip_lockedout_ips),
			"lockedout_ips_count":     ret.DtSlbCrlSrcipOper.Oper.Lockedout_ips_count,
		},
	}
}

func setSliceSlbCrlSrcipOperOperCrl_srcip_lockedout_ips(d []edpt.SlbCrlSrcipOperOperCrl_srcip_lockedout_ips) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["client_ip"] = item.Client_ip
		in["start"] = item.Start
		in["end"] = item.End
		in["drops"] = item.Drops
		in["active"] = item.Active
		result = append(result, in)
	}
	return result
}

func getObjectSlbCrlSrcipOperOper(d []interface{}) edpt.SlbCrlSrcipOperOper {

	count1 := len(d)
	var ret edpt.SlbCrlSrcipOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Crl_srcip_lockedout_ips = getSliceSlbCrlSrcipOperOperCrl_srcip_lockedout_ips(in["crl_srcip_lockedout_ips"].([]interface{}))
		ret.Lockedout_ips_count = in["lockedout_ips_count"].(int)
	}
	return ret
}

func getSliceSlbCrlSrcipOperOperCrl_srcip_lockedout_ips(d []interface{}) []edpt.SlbCrlSrcipOperOperCrl_srcip_lockedout_ips {

	count1 := len(d)
	ret := make([]edpt.SlbCrlSrcipOperOperCrl_srcip_lockedout_ips, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbCrlSrcipOperOperCrl_srcip_lockedout_ips
		oi.Client_ip = in["client_ip"].(string)
		oi.Start = in["start"].(string)
		oi.End = in["end"].(string)
		oi.Drops = in["drops"].(int)
		oi.Active = in["active"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbCrlSrcipOper(d *schema.ResourceData) edpt.SlbCrlSrcipOper {
	var ret edpt.SlbCrlSrcipOper

	ret.Oper = getObjectSlbCrlSrcipOperOper(d.Get("oper").([]interface{}))
	return ret
}
