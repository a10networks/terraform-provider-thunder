package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerMigrateVipOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_virtual_server_migrate_vip_oper`: Operational Status for the object migrate-vip\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbVirtualServerMigrateVipOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbVirtualServerMigrateVipOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerMigrateVipOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerMigrateVipOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbVirtualServerMigrateVipOperOper := setObjectSlbVirtualServerMigrateVipOperOper(res)
		d.Set("oper", SlbVirtualServerMigrateVipOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbVirtualServerMigrateVipOperOper(ret edpt.DataSlbVirtualServerMigrateVipOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state": ret.DtSlbVirtualServerMigrateVipOper.Oper.State,
		},
	}
}

func getObjectSlbVirtualServerMigrateVipOperOper(d []interface{}) edpt.SlbVirtualServerMigrateVipOperOper {

	count1 := len(d)
	var ret edpt.SlbVirtualServerMigrateVipOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointSlbVirtualServerMigrateVipOper(d *schema.ResourceData) edpt.SlbVirtualServerMigrateVipOper {
	var ret edpt.SlbVirtualServerMigrateVipOper

	ret.Oper = getObjectSlbVirtualServerMigrateVipOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
