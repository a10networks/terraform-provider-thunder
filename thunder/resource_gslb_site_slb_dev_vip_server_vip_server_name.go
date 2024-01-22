package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_site_slb_dev_vip_server_vip_server_name`: Specify a VIP for the SLB device\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSiteSlbDevVipServerVipServerNameCreate,
		UpdateContext: resourceGslbSiteSlbDevVipServerVipServerNameUpdate,
		ReadContext:   resourceGslbSiteSlbDevVipServerVipServerNameRead,
		DeleteContext: resourceGslbSiteSlbDevVipServerVipServerNameDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'dev_vip_hits': Number of times the service-ip was selected; 'dev_vip_recent': Recent hits;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip_name": {
				Type: schema.TypeString, Required: true, Description: "Specify a VIP name for the SLB device",
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
			"device_name": {
				Type: schema.TypeString, Required: true, Description: "DeviceName",
			},
		},
	}
}
func resourceGslbSiteSlbDevVipServerVipServerNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevVipServerVipServerNameRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSiteSlbDevVipServerVipServerNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevVipServerVipServerNameRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSiteSlbDevVipServerVipServerNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSiteSlbDevVipServerVipServerNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbSiteSlbDevVipServerVipServerNameSamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerNameSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerNameSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerNameSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerName(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerName {
	var ret edpt.GslbSiteSlbDevVipServerVipServerName
	ret.Inst.SamplingEnable = getSliceGslbSiteSlbDevVipServerVipServerNameSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.VipName = d.Get("vip_name").(string)
	ret.Inst.SiteName = d.Get("site_name").(string)
	ret.Inst.DeviceName = d.Get("device_name").(string)
	return ret
}
