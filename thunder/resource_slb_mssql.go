package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMssql() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_mssql`: Configure mssql\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbMssqlCreate,
		UpdateContext: resourceSlbMssqlUpdate,
		ReadContext:   resourceSlbMssqlRead,
		DeleteContext: resourceSlbMssqlDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_proxy': Curr Proxy Conns; 'total_proxy': Total Proxy Conns; 'curr_be_enc': Curr BE Encryption Conns; 'total_be_enc': Total BE Encryption Conns; 'curr_fe_enc': Curr FE Encryption Conns; 'total_fe_enc': Total FE Encryption Conns; 'client_fin': Client FIN; 'server_fin': Server FIN; 'session_err': Session err; 'queries': DB Queries; 'commands': DB commands reply; 'auth_success': Authentication Success; 'auth_failure': Authentication Failure;",
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
func resourceSlbMssqlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMssqlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMssql(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMssqlRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbMssqlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMssqlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMssql(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMssqlRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbMssqlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMssqlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMssql(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbMssqlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMssqlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMssql(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbMssqlSamplingEnable(d []interface{}) []edpt.SlbMssqlSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbMssqlSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbMssqlSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbMssql(d *schema.ResourceData) edpt.SlbMssql {
	var ret edpt.SlbMssql
	ret.Inst.SamplingEnable = getSliceSlbMssqlSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
