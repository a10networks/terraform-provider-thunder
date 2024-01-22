package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnDataSessionsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_data_sessions_oper`: Operational Status for the object data-sessions\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnDataSessionsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"inside_addr_end": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr_start": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_addr_end": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"inside_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnDataSessionsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnDataSessionsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnDataSessionsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnDataSessionsOperOper := setObjectCgnv6LsnDataSessionsOperOper(res)
		d.Set("oper", Cgnv6LsnDataSessionsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnDataSessionsOperOper(ret edpt.DataCgnv6LsnDataSessionsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status":            ret.DtCgnv6LsnDataSessionsOper.Oper.Status,
			"inside_addr":       ret.DtCgnv6LsnDataSessionsOper.Oper.InsideAddr,
			"inside_addr_start": ret.DtCgnv6LsnDataSessionsOper.Oper.InsideAddrStart,
			"inside_addr_end":   ret.DtCgnv6LsnDataSessionsOper.Oper.InsideAddrEnd,
			"nat_addr":          ret.DtCgnv6LsnDataSessionsOper.Oper.NatAddr,
			"nat_addr_start":    ret.DtCgnv6LsnDataSessionsOper.Oper.NatAddrStart,
			"nat_addr_end":      ret.DtCgnv6LsnDataSessionsOper.Oper.NatAddrEnd,
			"nat_port":          ret.DtCgnv6LsnDataSessionsOper.Oper.NatPort,
			"inside_port":       ret.DtCgnv6LsnDataSessionsOper.Oper.InsidePort,
		},
	}
}

func getObjectCgnv6LsnDataSessionsOperOper(d []interface{}) edpt.Cgnv6LsnDataSessionsOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnDataSessionsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(string)
		ret.InsideAddr = in["inside_addr"].(string)
		ret.InsideAddrStart = in["inside_addr_start"].(string)
		ret.InsideAddrEnd = in["inside_addr_end"].(string)
		ret.NatAddr = in["nat_addr"].(string)
		ret.NatAddrStart = in["nat_addr_start"].(string)
		ret.NatAddrEnd = in["nat_addr_end"].(string)
		ret.NatPort = in["nat_port"].(int)
		ret.InsidePort = in["inside_port"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnDataSessionsOper(d *schema.ResourceData) edpt.Cgnv6LsnDataSessionsOper {
	var ret edpt.Cgnv6LsnDataSessionsOper

	ret.Oper = getObjectCgnv6LsnDataSessionsOperOper(d.Get("oper").([]interface{}))
	return ret
}
