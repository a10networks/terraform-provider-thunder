package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateDnsDns64() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_dns_dns64`: Enable DNS64\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateDnsDns64Create,
		UpdateContext: resourceCgnv6TemplateDnsDns64Update,
		ReadContext:   resourceCgnv6TemplateDnsDns64Read,
		DeleteContext: resourceCgnv6TemplateDnsDns64Delete,

		Schema: map[string]*schema.Schema{
			"answer_only_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Only translate the Answer Section",
			},
			"auth_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set AA flag in DNS Response",
			},
			"cache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a cached A-query response to provide AAAA query responses for the same hostname",
			},
			"change_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always change incoming AAAA DNS Query to A",
			},
			"compress_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Always try DNS Compression",
			},
			"deep_check_qr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check DNS Question Record",
			},
			"deep_check_rr_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Check DNS Response Records",
			},
			"drop_cname_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Drop DNS CNAME Response",
			},
			"edns_append": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append EDNS Record when send A Query to server",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS64 (Need to config this option before config any other dns64 options)",
			},
			"fast_append": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Append translated Records when original Response only has Answer Section",
			},
			"ignore_rcode3_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Ignore DNS error Response with rcode 3",
			},
			"max_qr_length": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Max Question Record Length, default is 128",
			},
			"parallel_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward AAAA Query & generate A Query in parallel",
			},
			"passive_query_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Generate A query upon empty or error Response",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Retry count, default is 3 (Retry Number)",
			},
			"single_response_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Single Response which is used to avoid ambiguity",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Timeout to send additional Queries, unit: second, default is 1",
			},
			"trans_ptr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Translate DNS PTR Records",
			},
			"trans_ptr_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Translate DNS PTR Query",
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Max TTL in DNS Response, unit: second",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6TemplateDnsDns64Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsDns64Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDnsDns64(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateDnsDns64Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateDnsDns64Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsDns64Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDnsDns64(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateDnsDns64Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateDnsDns64Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsDns64Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDnsDns64(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateDnsDns64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateDnsDns64Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateDnsDns64(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6TemplateDnsDns64(d *schema.ResourceData) edpt.Cgnv6TemplateDnsDns64 {
	var ret edpt.Cgnv6TemplateDnsDns64
	ret.Inst.AnswerOnlyDisable = d.Get("answer_only_disable").(int)
	ret.Inst.AuthData = d.Get("auth_data").(int)
	ret.Inst.Cache = d.Get("cache").(int)
	ret.Inst.ChangeQuery = d.Get("change_query").(int)
	ret.Inst.CompressDisable = d.Get("compress_disable").(int)
	ret.Inst.DeepCheckQr = d.Get("deep_check_qr").(int)
	ret.Inst.DeepCheckRrDisable = d.Get("deep_check_rr_disable").(int)
	ret.Inst.DropCnameDisable = d.Get("drop_cname_disable").(int)
	ret.Inst.EdnsAppend = d.Get("edns_append").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.FastAppend = d.Get("fast_append").(int)
	ret.Inst.IgnoreRcode3Disable = d.Get("ignore_rcode3_disable").(int)
	ret.Inst.MaxQrLength = d.Get("max_qr_length").(int)
	ret.Inst.ParallelQuery = d.Get("parallel_query").(int)
	ret.Inst.PassiveQueryDisable = d.Get("passive_query_disable").(int)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.SingleResponseDisable = d.Get("single_response_disable").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.TransPtr = d.Get("trans_ptr").(int)
	ret.Inst.TransPtrQuery = d.Get("trans_ptr_query").(int)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
