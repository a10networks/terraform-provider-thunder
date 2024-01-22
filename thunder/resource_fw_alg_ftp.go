package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgFtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg_ftp`: Change Firewall FTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgFtpCreate,
		UpdateContext: resourceFwAlgFtpUpdate,
		ReadContext:   resourceFwAlgFtpRead,
		DeleteContext: resourceFwAlgFtpDelete,

		Schema: map[string]*schema.Schema{
			"default_port_disable": {
				Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable FTP ALG default port 21;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'client-port-request': PORT Requests From Client; 'client-eprt-request': EPRT Requests From Client; 'server-pasv-reply': PASV Replies From Server; 'server-epsv-reply': EPSV Replies From Server; 'port-retransmits': PORT Retransmits; 'pasv-retransmits': PASV Retransmits; 'smp-app-type-mismatch': SMP App Type Mismatch; 'retransmit-sanity-check-failure': Retransmit Sanity Check Failure; 'smp-conn-alloc-failure': SMP Helper Conn Alloc Failure; 'port-helper-created': PORT Helper Created; 'pasv-helper-created': PASV Helper Created; 'port-helper-acquire-in-del-q': PORT Helper Acquire In Del Queue; 'port-helper-acquire-already-used': PORT Helper Acquire Already Used; 'pasv-helper-acquire-in-del-q': PASV Helper Acquire In Del Queue; 'pasv-helper-acquire-already-used': PASV Helper Acquire Already Used; 'port-helper-freed-used': PORT Helper Freed Used; 'port-helper-freed-unused': PORT Helper Freed Unused; 'pasv-helper-freed-used': PASV Helper Freed Used; 'pasv-helper-freed-unused': PASV Helper Freed Unused;",
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
func resourceFwAlgFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgFtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgFtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgFtpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgFtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgFtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgFtpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgFtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgFtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgFtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgFtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwAlgFtpSamplingEnable(d []interface{}) []edpt.FwAlgFtpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwAlgFtpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgFtpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwAlgFtp(d *schema.ResourceData) edpt.FwAlgFtp {
	var ret edpt.FwAlgFtp
	ret.Inst.DefaultPortDisable = d.Get("default_port_disable").(string)
	ret.Inst.SamplingEnable = getSliceFwAlgFtpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
