package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHardwareAccelerateSlb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_hardware_accelerate_slb`: \n\n__PLACEHOLDER__",
		CreateContext: resourceSystemHardwareAccelerateSlbCreate,
		UpdateContext: resourceSystemHardwareAccelerateSlbUpdate,
		ReadContext:   resourceSystemHardwareAccelerateSlbRead,
		DeleteContext: resourceSystemHardwareAccelerateSlbDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'entry-create': Hardware Entries Created; 'entry-create-failure': Hardware Entries Created; 'entry-create-fail-server-down': Hardware Entries Created; 'entry-create-fail-max-entry': Hardware Entries Created; 'entry-free': Hardware Entries Freed; 'entry-free-opp-entry': Hardware Entries Free due to opposite tuple entry ageout event; 'entry-free-no-hw-prog': Hardware Entry Free no hw prog; 'entry-free-no-conn': Hardware Entry Free no matched conn; 'entry-free-no-sw-entry': Hardware Entry Free no software entry; 'entry-counter': Hardware Entry Count; 'entry-age-out': Hardware Entries Aged Out; 'entry-age-out-idle': Hardware Entries Aged Out Idle; 'entry-age-out-tcp-fin': Hardware Entries Aged Out TCP FIN; 'entry-age-out-tcp-rst': Hardware Entries Aged Out TCP RST; 'entry-age-out-invalid-dst': Hardware Entries Aged Out invalid dst; 'entry-force-hw-invalidate': Hardware Entries Force HW Invalidate; 'entry-invalidate-server-down': Hardware Entries Invalidate due to server down; 'tcam-create': TCAM Entries Created; 'tcam-free': TCAM Entries Freed; 'tcam-counter': TCAM Entry Count;",
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
func resourceSystemHardwareAccelerateSlbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateSlbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerateSlb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHardwareAccelerateSlbRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemHardwareAccelerateSlbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateSlbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerateSlb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHardwareAccelerateSlbRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemHardwareAccelerateSlbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateSlbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerateSlb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemHardwareAccelerateSlbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateSlbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerateSlb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemHardwareAccelerateSlbSamplingEnable(d []interface{}) []edpt.SystemHardwareAccelerateSlbSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemHardwareAccelerateSlbSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareAccelerateSlbSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemHardwareAccelerateSlb(d *schema.ResourceData) edpt.SystemHardwareAccelerateSlb {
	var ret edpt.SystemHardwareAccelerateSlb
	ret.Inst.SamplingEnable = getSliceSystemHardwareAccelerateSlbSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
