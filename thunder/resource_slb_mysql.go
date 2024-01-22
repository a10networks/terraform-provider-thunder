package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMysql() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_mysql`: Configure mysql\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbMysqlCreate,
		UpdateContext: resourceSlbMysqlUpdate,
		ReadContext:   resourceSlbMysqlRead,
		DeleteContext: resourceSlbMysqlDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_proxy': Curr Proxy Conns; 'total_proxy': Total Proxy Conns; 'curr_be_enc': Curr BE Encryption Conns; 'total_be_enc': Total BE Encryption Conns; 'curr_fe_enc': Curr FE Encryption Conns; 'total_fe_enc': Total FE Encryption Conns; 'client_fin': Client FIN; 'server_fin': Server FIN; 'session_err': Session err; 'queries': DB Queries; 'commands': DB commands reply;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbMysqlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMysqlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMysql(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMysqlRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbMysqlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMysqlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMysql(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMysqlRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbMysqlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMysqlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMysql(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbMysqlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMysqlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMysql(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbMysqlSamplingEnable(d []interface{}) []edpt.SlbMysqlSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbMysqlSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbMysqlSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbMysql(d *schema.ResourceData) edpt.SlbMysql {
	var ret edpt.SlbMysql
	ret.Inst.SamplingEnable = getSliceSlbMysqlSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
