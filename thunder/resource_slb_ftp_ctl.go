package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFtpCtl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ftp_ctl`: Configure FTP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbFtpCtlCreate,
		UpdateContext: resourceSlbFtpCtlUpdate,
		ReadContext:   resourceSlbFtpCtlRead,
		DeleteContext: resourceSlbFtpCtlDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sessions_num': Total Control Sessions; 'alg_pkts_num': Total ALG packets; 'alg_pkts_xmitted_num': ALG packets rexmitted; 'alg_port_helper_created': Total PORT helper sessions; 'alg_pasv_helper_created': Total PASV helper sessions; 'alg_port_helper_freed_unused': PORT helper freed unused; 'alg_pasv_helper_freed_unused': PASV helper freed unused; 'alg_port_helper_nat_free': PORT helper NAT free;",
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
func resourceSlbFtpCtlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpCtlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpCtl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFtpCtlRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFtpCtlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpCtlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpCtl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFtpCtlRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbFtpCtlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpCtlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpCtl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbFtpCtlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpCtlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpCtl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbFtpCtlSamplingEnable(d []interface{}) []edpt.SlbFtpCtlSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbFtpCtlSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbFtpCtlSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbFtpCtl(d *schema.ResourceData) edpt.SlbFtpCtl {
	var ret edpt.SlbFtpCtl
	ret.Inst.SamplingEnable = getSliceSlbFtpCtlSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
