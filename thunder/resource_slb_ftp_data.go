package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFtpData() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ftp_data`: Configure FTP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbFtpDataCreate,
		UpdateContext: resourceSlbFtpDataUpdate,
		ReadContext:   resourceSlbFtpDataRead,
		DeleteContext: resourceSlbFtpDataDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sessions_num': Total Data Sessions; 'port_out_of_range': Drop Data Port out of range;",
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
func resourceSlbFtpDataCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpDataCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpData(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFtpDataRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFtpDataUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpDataUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpData(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFtpDataRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbFtpDataDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpDataDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpData(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbFtpDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpDataRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpData(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbFtpDataSamplingEnable(d []interface{}) []edpt.SlbFtpDataSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbFtpDataSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbFtpDataSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbFtpData(d *schema.ResourceData) edpt.SlbFtpData {
	var ret edpt.SlbFtpData
	ret.Inst.SamplingEnable = getSliceSlbFtpDataSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
