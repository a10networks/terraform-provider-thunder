package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbImapProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_imap_proxy`: Configure IMAP Proxy global\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbImapProxyCreate,
		UpdateContext: resourceSlbImapProxyUpdate,
		ReadContext:   resourceSlbImapProxyRead,
		DeleteContext: resourceSlbImapProxyDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num': Num; 'curr': Current proxy conns; 'total': Total proxy conns; 'svrsel_fail': Server selection failure; 'no_route': no route failure; 'snat_fail': source nat failure; 'feat': feat packet; 'cc': clear ctrl port packet; 'data_ssl': data ssl force; 'line_too_long': line too long; 'line_mem_freed': request line freed; 'invalid_start_line': invalid start line; 'auth_tls': auth tls cmd; 'prot': prot cmd; 'pbsz': pbsz cmd; 'pasv': pasv cmd; 'port': port cmd; 'request_dont_care': other cmd; 'client_auth_tls': client auth tls; 'cant_find_pasv': cant find pasv; 'pasv_addr_ne_server': psv addr not equal to svr; 'smp_create_fail': smp create fail; 'data_server_conn_fail': data svr conn fail; 'data_send_fail': data send fail; 'epsv': epsv command; 'cant_find_epsv': cant find epsv; 'data_curr': Current Data Proxy; 'data_total': Total Data Proxy; 'auth_unsupported': Unsupported auth; 'adat': adat cmd; 'unsupported_pbsz_value': Unsupported PBSZ; 'unsupported_prot_value': Unsupported PROT; 'unsupported_command': Unsupported cmd; 'control_to_clear': Control chn clear txt; 'control_to_ssl': Control chn ssl; 'bad_sequence': Bad Sequence; 'rsv_persist_conn_fail': Serv Sel Persist fail; 'smp_v6_fail': Serv Sel SMPv6 fail; 'smp_v4_fail': Serv Sel SMPv4 fail; 'insert_tuple_fail': Serv Sel insert tuple fail; 'cl_est_err': Client EST state erro; 'ser_connecting_err': Serv CTNG state error; 'server_response_err': Serv RESP state error; 'cl_request_err': Client RQ state error; 'data_conn_start_err': Data Start state error; 'data_serv_connecting_err': Data Serv CTNG error; 'data_serv_connected_err': Data Serv CTED error; 'request': Total FTP Request; 'capability': Capability cmd; 'start_tls': Total Start TLS cmd; 'login': Total Login cmd; 'realloc_error': Realloc error; 'alloc_error': Alloc error; 'boundary_error': Boundary error; 'negative_error': Negative error;",
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
func resourceSlbImapProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbImapProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbImapProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbImapProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbImapProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbImapProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbImapProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbImapProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbImapProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbImapProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbImapProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbImapProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbImapProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbImapProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbImapProxySamplingEnable(d []interface{}) []edpt.SlbImapProxySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbImapProxySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbImapProxySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbImapProxy(d *schema.ResourceData) edpt.SlbImapProxy {
	var ret edpt.SlbImapProxy
	ret.Inst.SamplingEnable = getSliceSlbImapProxySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
