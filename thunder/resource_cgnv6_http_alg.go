package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6HttpAlg() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_http_alg`: HTTP-ALG Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6HttpAlgCreate,
		UpdateContext: resourceCgnv6HttpAlgUpdate,
		ReadContext:   resourceCgnv6HttpAlgRead,
		DeleteContext: resourceCgnv6HttpAlgDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'request-processed': HTTP Request Processed; 'request-insert-msisdn-performed': HTTP MSISDN Insertion Performed; 'request-insert-client-ip-performed': HTTP Client IP Insertion Performed; 'request-insert-msisdn-unavailable': Inserted MSISDN is 0000 (MSISDN Unavailable); 'queued-session-too-many': Queued Session Exceed Drop; 'radius-query-succeed': MSISDN Query Succeed; 'radius-query-failed': MSISDN Query Failed; 'radius-requst-sent': Query Request Sent; 'radius-requst-dropped': Query Request Dropped; 'radius-response-received': Query Response Received; 'radius-response-dropped': Query Response Dropped; 'out-of-memory-dropped': Out-of-Memory Dropped; 'queue-len-exceed-dropped': Queue Length Exceed Dropped; 'out-of-order-dropped': Packet Out-of-Order Dropped; 'buff-resent': Packet Resent from Queue; 'buff-spilt-failed': Buff Split Failed; 'header-insertion-failed': Buff Insertion Failed; 'header-removal-failed': Buff Removal Failed; 'no-queue': No Queue;",
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
func resourceCgnv6HttpAlgCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6HttpAlgCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6HttpAlg(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6HttpAlgRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6HttpAlgUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6HttpAlgUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6HttpAlg(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6HttpAlgRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6HttpAlgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6HttpAlgDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6HttpAlg(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6HttpAlgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6HttpAlgRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6HttpAlg(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6HttpAlgSamplingEnable(d []interface{}) []edpt.Cgnv6HttpAlgSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6HttpAlgSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6HttpAlgSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6HttpAlg(d *schema.ResourceData) edpt.Cgnv6HttpAlg {
	var ret edpt.Cgnv6HttpAlg
	ret.Inst.SamplingEnable = getSliceCgnv6HttpAlgSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
