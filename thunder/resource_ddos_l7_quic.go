package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL7Quic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_l7_quic`: DDOS QUIC Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosL7QuicCreate,
		UpdateContext: resourceDdosL7QuicUpdate,
		ReadContext:   resourceDdosL7QuicRead,
		DeleteContext: resourceDdosL7QuicDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'quic_packet_received': QUIC Packet Received; 'quic_initial_received': QUIC Initial Packet Received; 'quic_version_negotiation_received': QUIC Version Negotiation Received; 'quic_retry_received': QUIC Retry Received; 'quic_0rtt_recevied': QUIC 0RTT Received; 'quic_handshake_received': QUIC Handshake Received; 'quic_version_match_action_taken': QUIC Version Match Action Taken; 'quic_version_match_action_drop': QUIC Version Match Action Drop; 'quic_version_match_action_blacklist': QUIC Version Match Action Blacklist; 'quic_malformed_action_taken': QUIC Malformed Action Taken; 'quic_malformed_action_drop': QUIC Malformed Action Drop; 'quic_malformed_action_blacklist': QUIC Malformed Action Blacklist; 'quic_malformed_dcid_len_max_exceed': QUIC Malformed DCID Len Max Exceed; 'quic_malformed_scid_len_max_exceed': QUIC Malformed SCID Len Max Exceed; 'quic_fixed_bit_not_set': QUIC Fixed Bit Not Set; 'quic_retry_auth_sent': QUIC Retry Auth Sent; 'quic_retry_auth_pass': QUIC Retry Auth Pass; 'quic_retry_auth_fail': QUIC Retry Auth Fail; 'quic_connection_close_sent': QUIC Connection Close Sent; 'quic_invalid_retry_token': QUIC Invalid Retry Token; 'quic_short_header_received': QUIC Short Header Received; 'quic_short_header_action_drop': QUIC Short Header Drop; 'quic_encrypt_fail': QUIC Encrypt Fail; 'quic_decrypt_fail': QUIC Decrypt Fail; 'quic_encrypt_success': QUIC Encrypt Success; 'quic_decrypt_success': QUIC Decrypt Success; 'quic_0rtt_drop': QUIC 0RTT Drop; 'quic_aead_pkt_rate_exceed': QUIC AEAD Packet Rate Exceed; 'quic_dcid_pkt_rate_exceed': QUIC DCID Packet Rate Exceed; 'quic_create_conn_init_only': QUIC Create Connection on Initial Only; 'quic_version_no_match_drop': QUIC Version No Match Drop;",
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
func resourceDdosL7QuicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7QuicCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7Quic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosL7QuicRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosL7QuicUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7QuicUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7Quic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosL7QuicRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosL7QuicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7QuicDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7Quic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosL7QuicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7QuicRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7Quic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosL7QuicSamplingEnable(d []interface{}) []edpt.DdosL7QuicSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosL7QuicSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosL7QuicSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosL7Quic(d *schema.ResourceData) edpt.DdosL7Quic {
	var ret edpt.DdosL7Quic
	ret.Inst.SamplingEnable = getSliceDdosL7QuicSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
