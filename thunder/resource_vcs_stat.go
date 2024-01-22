package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsStat() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_stat`: Show aVCS statistics information\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsStatCreate,
		UpdateContext: resourceVcsStatUpdate,
		ReadContext:   resourceVcsStatRead,
		DeleteContext: resourceVcsStatDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'elect_recv_err': Receive error counter of aVCS election; 'elect_send_err': Send error counter of aVCS election; 'elect_recv_byte': Receive bytes counter of aVCS election; 'elect_send_byte': Send bytes counter of aVCS election; 'elect_pdu_master_recv': Received vMaster-PDU counter of aVCS election; 'elect_pdu_master_cand_recv': Received MC-PDU counter of aVCS election; 'elect_pdu_slave_recv': Received vBlade-PDU counter of aVCS election; 'elect_pdu_master_take_over_recv': Received MTO-PDU counter of aVCS election; 'elect_pdu_unknown_recv': Received Unknown-PDU counter of aVCS election; 'elect_pdu_master_sent': Sent vMaster-PDU counter of aVCS election; 'elect_pdu_master_cand_sent': Sent MC-PDU counter of aVCS election; 'elect_pdu_slave_sent': Sent vBlade-PDU counter of aVCS election; 'elect_pdu_master_take_over_sent': Sent MTO-PDU counter of aVCS election; 'elect_pdu_unknown_sent': Sent Unknown-PDU counter of aVCS election; 'elect_pdu_inval': Invalid PDU counter of aVCS election; 'elect_pdu_hw_mismatch': PDU HW mismatch counter of aVCS election; 'elect_pdu_cluster_mismatch': PDU Chassis-ID mismatch counter of aVCS election; 'elect_pdu_dev_id_collision': PDU Device-ID collision counter of aVCS election; 'elect_mc_discard_master': MC discarded vMaster-PDU counter of aVCS election; 'elect_mc_replace_master': MC replaced vMaster-PDU counter of aVCS election; 'elect_mc_dup_masterr': MC duplicate vMaster-PDU counter of aVCS election; 'elect_mc_reset_timer_by_mc': MC timers reset by MC-PDU counter of aVCS election; 'elect_mc_reset_timer_by_mto': MC timers reset by MTO-PDU counter of aVCS election; 'elect_slave_dup_master': vBlade duplicate vMaster-PDU counter of aVCS election; 'elect_slave_discard_challenger': vBlade discard challenger counter of aVCS election; 'elect_slave_replace_challenger': vBlade replace challenger counter of aVCS election; 'elect_slave_dup_challenger': vBlade duplicate challenger counter of aVCS election; 'elect_slave_discard_neighbour': vBlade discard neighbour counter of aVCS election; 'elect_slave_too_many_neighbour': vBlade too many neighbours counter of aVCS election; 'elect_slave_dup_neighbour': send vBlade duplicate neighbours of aVCS election; 'elect_master_discard_challenger': vMaster discard challenger counter of aVCS election; 'elect_master_new_challenger': vMaster new challenger counter of aVCS election; 'elect_master_replace_challenger': vMaster replace challenger counter of aVCS election; 'elect_master_dup_challenger': vMaster duplicate challenger counter of aVCS election; 'elect_master_discard_neighbour': vMaster discard neighbour counter of aVCS election; 'elect_master_too_many_neighbour': vMaster too many neighbours counter of aVCS election; 'elect_master_dup_neighbour': vMaster duplicate neighbours counter of aVCS election; 'elect_enter_master_cand_stat': Enter MC counter of aVCS election; 'elect_enter_slave': Enter vBlade counter of aVCS election; 'elect_enter_master': Enter vMaster counter of aVCS election; 'elect_enter_master_take_over': Enter MTO counter of aVCS election; 'elect_leave_master_cand': Leave MC counter of aVCS election; 'elect_leave_slave': Leave vBlade counter of aVCS election; 'elect_leave_master': Leave vMaster counter of aVCS election; 'elect_leave_master_take_over': Leave MTO counter of aVCS election; 'master_slave_start_err': vMaster Start vBlade Errors counter of aVCS election; 'master_slave_start': vMaster vBlades Started counter of aVCS election; 'master_slave_stop': vMaster vBlades stopped counter of aVCS election; 'master_cfg_upd': Received vMaster Configuration Updates counter of aVCS election; 'master_cfg_upd_l_fail': vMaster Local Configuration Update Errors counter of aVCS election; 'master_cfg_upd_r_fail': vMaster Remote Configuration Update Errors counter of aVCS election; 'master_cfg_upd_notif_err': vMaster Configuration Update Notif Errors counter of aVCS election; 'master_cfg_upd_result_err': vMaster Configuration Update Result Errors counter of aVCS election; 'slave_recv_err': vBlade Receive Errors counter of aVCS election; 'slave_send_err': vBlade Send Errors counter of aVCS election; 'slave_recv_bytes': vBlade Received Bytes counter of aVCS election; 'slave_sent_bytes': vBlade Sent Bytes counter of aVCS election; 'slave_n_recv': vBlade Received Messages counter of aVCS election; 'slave_n_sent': vBlade Sent Messages counter of aVCS election; 'slave_msg_inval': vBlade Invalid Messages counter of aVCS election; 'slave_keepalive': vBlade Received Keepalives counter of aVCS election; 'slave_cfg_upd': vBlade Received Configuration Updates counter of aVCS election; 'slave_cfg_upd_fail': vBlade Configuration Update Failures counter of aVCS election; 'daemon_n_elec_start': times of aVCS election start; 'daemon_n_elec_stop': times of aVCS election stop; 'daemon_recv_err': counter of aVCS daemon receive error; 'daemon_send_err': counter of aVCS daemon sent error; 'daemon_recv_bytes': bytes of aVCS daemon receive; 'daemon_sent_bytes': bytes of aVCS daemon sent; 'daemon_n_recv': counter of aVCS daemon receive; 'daemon_n_sent': counter of aVCS daemon sent; 'daemon_msg_inval': counter of aVCS daemon invalid message; 'daemon_msg_handle_failure': counter of aVCS daemon message handle failure;",
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
func resourceVcsStatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsStatCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsStat(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsStatRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsStatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsStatUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsStat(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsStatRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsStatDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsStatDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsStat(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsStatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsStatRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsStat(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVcsStatSamplingEnable(d []interface{}) []edpt.VcsStatSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VcsStatSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsStatSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsStat(d *schema.ResourceData) edpt.VcsStat {
	var ret edpt.VcsStat
	ret.Inst.SamplingEnable = getSliceVcsStatSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
