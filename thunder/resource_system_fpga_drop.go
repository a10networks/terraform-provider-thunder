package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemFpgaDrop() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_fpga_drop`: \n\n__PLACEHOLDER__",
		CreateContext: resourceSystemFpgaDropCreate,
		UpdateContext: resourceSystemFpgaDropUpdate,
		ReadContext:   resourceSystemFpgaDropRead,
		DeleteContext: resourceSystemFpgaDropDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'mrx-drop': Total MRX Drop; 'hrx-drop': Total HRX Drop; 'siz-drop': Total Size Drop; 'fcs-drop': Total FCS Drop; 'land-drop': Total LAND Attack Drop; 'empty-frag-drop': Total Empty frag Drop; 'mic-frag-drop': Total Micro frag Drop; 'ipv4-opt-drop': Total IPv4 opt Drop; 'ipv4-frag': Total IP frag Drop; 'bad-ip-hdr-len': Total Bad IP hdr len Drop; 'bad-ip-flags-drop': Total Bad IP Flags Drop; 'bad-ip-ttl-drop': Total Bad IP TTL Drop; 'no-ip-payload-drop': Total No IP Payload Drop; 'oversize-ip-payload': Total Oversize IP PL Drop; 'bad-ip-payload-len': Total Bad IP PL len Drop; 'bad-ip-frag-offset': Total Bad IP frag off Drop; 'bad-ip-chksum-drop': Total Bad IP csum Drop; 'icmp-pod-drop': Total ICMP POD Drop; 'tcp-bad-urg-offet': Total TCP bad urg off Drop; 'tcp-short-hdr': Total TCP short hdr Drop; 'tcp-bad-ip-len': Total TCP Bad IP Len Drop; 'tcp-null-flags': Total TCP null flags Drop; 'tcp-null-scan': Total TCP null scan Drop; 'tcp-fin-sin': Total TCP SYN&FIN Drop; 'tcp-xmas-flags': Total TCP XMAS FLAGS Drop; 'tcp-xmas-scan': Total TCP XMAS scan Drop; 'tcp-syn-frag': Total TCP SYN frag Drop; 'tcp-frag-hdr': Total TCP frag header Drop; 'tcp-bad-chksum': Total TCP bad csum Drop; 'udp-short-hdr': Total UDP short hdr Drop; 'udp-bad-ip-len': Total UDP bad leng Drop; 'udp-kb-frags': Total UDP KB frag Drop; 'udp-port-lb': Total UDP port LB Drop; 'udp-bad-chksum': Total UDP bad csum Drop; 'runt-ip-hdr': Total Runt IP hdr Drop; 'runt-tcpudp-hdr': Total Runt TCPUDP hdr Drop; 'tun-mismatch': Total Tun mismatch Drop; 'qdr-drop': Total QDR Drop;",
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
func resourceSystemFpgaDropCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaDropCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaDrop(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemFpgaDropRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemFpgaDropUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaDropUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaDrop(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemFpgaDropRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemFpgaDropDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaDropDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaDrop(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemFpgaDropRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaDropRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaDrop(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemFpgaDropSamplingEnable(d []interface{}) []edpt.SystemFpgaDropSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemFpgaDropSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemFpgaDropSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemFpgaDrop(d *schema.ResourceData) edpt.SystemFpgaDrop {
	var ret edpt.SystemFpgaDrop
	ret.Inst.SamplingEnable = getSliceSystemFpgaDropSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
