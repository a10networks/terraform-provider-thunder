package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMlb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_mlb`: Configure mlb\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbMlbCreate,
		UpdateContext: resourceSlbMlbUpdate,
		ReadContext:   resourceSlbMlbRead,
		DeleteContext: resourceSlbMlbDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'client_msg_sent': Client message sent; 'server_msg_received': Server message received; 'server_conn_created': Server connection created; 'server_conn_rst': Server connection reset; 'server_conn_failed': Server connection failed; 'server_conn_closed': Server connection closed; 'client_conn_created': Client connection created; 'client_conn_closed': Client connection closed; 'client_conn_not_found': Client connection not found; 'msg_dropped': Message dropped; 'msg_rerouted': Message rerouted; 'mlb_dcmsg_sent': Dcmsg sent; 'mlb_dcmsg_received': Dcmsg received; 'mlb_dcmsg_error': Dcmsg error; 'mlb_dcmsg_alloc': Dcmsg alloc; 'mlb_dcmsg_free': Dcmsg free; 'mlb_server_probe': Server probe; 'mlb_server_down': Server down;",
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
func resourceSlbMlbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMlbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMlb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMlbRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbMlbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMlbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMlb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbMlbRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbMlbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMlbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMlb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbMlbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMlbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMlb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbMlbSamplingEnable(d []interface{}) []edpt.SlbMlbSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbMlbSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbMlbSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbMlb(d *schema.ResourceData) edpt.SlbMlb {
	var ret edpt.SlbMlb
	ret.Inst.SamplingEnable = getSliceSlbMlbSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
