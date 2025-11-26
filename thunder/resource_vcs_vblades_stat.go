package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsVbladesStat() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_vblades_stat`: Show aVCS vBlade box statistics information\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsVbladesStatCreate,
		UpdateContext: resourceVcsVbladesStatUpdate,
		ReadContext:   resourceVcsVbladesStatRead,
		DeleteContext: resourceVcsVbladesStatDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'slave_recv_err': vBlade Receive Errors counter of aVCS election; 'slave_send_err': vBlade Send Errors counter of aVCS election; 'slave_recv_bytes': vBlade Received Bytes counter of aVCS election; 'slave_sent_bytes': vBlade Sent Bytes counter of aVCS election; 'slave_n_recv': vBlade Received Messages counter of aVCS election; 'slave_n_sent': vBlade Sent Messages counter of aVCS election; 'slave_msg_inval': vBlade Invalid Messages counter of aVCS election; 'slave_keepalive': vBlade Received Keepalives counter of aVCS election; 'slave_cfg_upd': vBlade Received Configuration Updates counter of aVCS election; 'slave_cfg_upd_l1_fail': vBlade Local Configuration Update Errors (1) counter of aVCS election; 'slave_cfg_upd_r_fail': vBlade Remote Configuration Update Errors counter of aVCS election; 'slave_cfg_upd_l2_fail': vBlade Local Configuration Update Errors (2) counter of aVCS election; 'slave_cfg_upd_notif_err': vBlade Configuration Update Notif Errors counter of aVCS election; 'slave_cfg_upd_result_err': vBlade Configuration Update Result Errors counter of aVCS election;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vblade_id": {
				Type: schema.TypeInt, Required: true, Description: "vBlade-id",
			},
		},
	}
}
func resourceVcsVbladesStatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVbladesStatCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVbladesStat(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVbladesStatRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsVbladesStatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVbladesStatUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVbladesStat(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVbladesStatRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsVbladesStatDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVbladesStatDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVbladesStat(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsVbladesStatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVbladesStatRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVbladesStat(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVcsVbladesStatSamplingEnable(d []interface{}) []edpt.VcsVbladesStatSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VcsVbladesStatSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVbladesStatSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsVbladesStat(d *schema.ResourceData) edpt.VcsVbladesStat {
	var ret edpt.VcsVbladesStat
	ret.Inst.SamplingEnable = getSliceVcsVbladesStatSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.VbladeId = d.Get("vblade_id").(int)
	return ret
}
