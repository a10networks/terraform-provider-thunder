package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetMgmt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_net_mgmt`: Network Management\n\n__PLACEHOLDER__",
		CreateContext: resourceNetMgmtCreate,
		UpdateContext: resourceNetMgmtUpdate,
		ReadContext:   resourceNetMgmtRead,
		DeleteContext: resourceNetMgmtDelete,

		Schema: map[string]*schema.Schema{
			"snmp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engineid": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceNetMgmtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetMgmtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetMgmt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetMgmtRead(ctx, d, meta)
	}
	return diags
}

func resourceNetMgmtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetMgmtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetMgmt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetMgmtRead(ctx, d, meta)
	}
	return diags
}
func resourceNetMgmtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetMgmtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetMgmt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetMgmtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetMgmtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetMgmt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetMgmtSnmp1052(d []interface{}) edpt.NetMgmtSnmp1052 {

	count1 := len(d)
	var ret edpt.NetMgmtSnmp1052
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Engineid = getObjectNetMgmtSnmpEngineid1053(in["engineid"].([]interface{}))
		ret.Stats = getObjectNetMgmtSnmpStats1054(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectNetMgmtSnmpEngineid1053(d []interface{}) edpt.NetMgmtSnmpEngineid1053 {

	var ret edpt.NetMgmtSnmpEngineid1053
	return ret
}

func getObjectNetMgmtSnmpStats1054(d []interface{}) edpt.NetMgmtSnmpStats1054 {

	var ret edpt.NetMgmtSnmpStats1054
	return ret
}

func dataToEndpointNetMgmt(d *schema.ResourceData) edpt.NetMgmt {
	var ret edpt.NetMgmt
	ret.Inst.Snmp = getObjectNetMgmtSnmp1052(d.Get("snmp").([]interface{}))
	return ret
}
