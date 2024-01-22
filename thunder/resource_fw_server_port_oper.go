package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwServerPortOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_server_port_oper`: Operational Status for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceFwServerPortOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vrid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ha_group_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ports_consumed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ports_consumed_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ports_freed_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"alloc_failed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceFwServerPortOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerPortOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServerPortOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwServerPortOperOper := setObjectFwServerPortOperOper(res)
		d.Set("oper", FwServerPortOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwServerPortOperOper(ret edpt.DataFwServerPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                ret.DtFwServerPortOper.Oper.State,
			"ip":                   ret.DtFwServerPortOper.Oper.Ip,
			"ipv6":                 ret.DtFwServerPortOper.Oper.Ipv6,
			"vrid":                 ret.DtFwServerPortOper.Oper.Vrid,
			"ha_group_id":          ret.DtFwServerPortOper.Oper.Ha_group_id,
			"ports_consumed":       ret.DtFwServerPortOper.Oper.Ports_consumed,
			"ports_consumed_total": ret.DtFwServerPortOper.Oper.Ports_consumed_total,
			"ports_freed_total":    ret.DtFwServerPortOper.Oper.Ports_freed_total,
			"alloc_failed":         ret.DtFwServerPortOper.Oper.Alloc_failed,
		},
	}
}

func getObjectFwServerPortOperOper(d []interface{}) edpt.FwServerPortOperOper {

	count1 := len(d)
	var ret edpt.FwServerPortOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Vrid = in["vrid"].(int)
		ret.Ha_group_id = in["ha_group_id"].(int)
		ret.Ports_consumed = in["ports_consumed"].(int)
		ret.Ports_consumed_total = in["ports_consumed_total"].(int)
		ret.Ports_freed_total = in["ports_freed_total"].(int)
		ret.Alloc_failed = in["alloc_failed"].(int)
	}
	return ret
}

func dataToEndpointFwServerPortOper(d *schema.ResourceData) edpt.FwServerPortOper {
	var ret edpt.FwServerPortOper

	ret.Oper = getObjectFwServerPortOperOper(d.Get("oper").([]interface{}))

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
