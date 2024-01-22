package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemBfd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_bfd`: BFD Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemBfdCreate,
		UpdateContext: resourceSystemBfdUpdate,
		ReadContext:   resourceSystemBfdRead,
		DeleteContext: resourceSystemBfdDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip_checksum_error': IP packet checksum errors; 'udp_checksum_error': UDP packet checksum errors; 'session_not_found': Session not found; 'multihop_mismatch': Multihop session or packet mismatch; 'version_mismatch': BFD version mismatch; 'length_too_small': Packets too small; 'data_is_short': Packet data length too short; 'invalid_detect_mult': Invalid detect multiplier; 'invalid_multipoint': Invalid multipoint setting; 'invalid_my_disc': Invalid my descriptor; 'invalid_ttl': Invalid TTL; 'auth_length_invalid': Invalid authentication length; 'auth_mismatch': Authentication mismatch; 'auth_type_mismatch': Authentication type mismatch; 'auth_key_id_mismatch': Authentication key-id mismatch; 'auth_key_mismatch': Authentication key mismatch; 'auth_seqnum_invalid': Invalid authentication sequence number; 'auth_failed': Authentication failures; 'local_state_admin_down': Local admin down session state; 'dest_unreachable': Destination unreachable; 'no_ipv6_enable': No IPv6 enable; 'other_error': Other errors;",
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
func resourceSystemBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBfdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBfd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemBfdRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBfdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBfd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemBfdRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBfdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBfd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBfdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBfd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemBfdSamplingEnable(d []interface{}) []edpt.SystemBfdSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemBfdSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemBfdSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemBfd(d *schema.ResourceData) edpt.SystemBfd {
	var ret edpt.SystemBfd
	ret.Inst.SamplingEnable = getSliceSystemBfdSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
