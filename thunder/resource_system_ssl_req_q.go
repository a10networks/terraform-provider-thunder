package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSslReqQ() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ssl_req_q`: System SSL Request Queue statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSslReqQCreate,
		UpdateContext: resourceSystemSslReqQUpdate,
		ReadContext:   resourceSystemSslReqQRead,
		DeleteContext: resourceSystemSslReqQDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num-ssl-queues': num-ssl-queues; 'ssl-req-q-depth-tot': ssl-req-q-depth-tot; 'ssl-req-q-inuse-tot': ssl-req-q-inuse-tot; 'ssl-hw-q-depth-tot': ssl-hw-q-depth-tot; 'ssl-hw-q-inuse-tot': ssl-hw-q-inuse-tot;",
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
func resourceSystemSslReqQCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslReqQCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslReqQ(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslReqQRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSslReqQUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslReqQUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslReqQ(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslReqQRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSslReqQDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslReqQDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslReqQ(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSslReqQRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslReqQRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslReqQ(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemSslReqQSamplingEnable(d []interface{}) []edpt.SystemSslReqQSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemSslReqQSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemSslReqQSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemSslReqQ(d *schema.ResourceData) edpt.SystemSslReqQ {
	var ret edpt.SystemSslReqQ
	ret.Inst.SamplingEnable = getSliceSystemSslReqQSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
