package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgEsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg_esp`: Change Firewall ESP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgEspCreate,
		UpdateContext: resourceFwAlgEspUpdate,
		ReadContext:   resourceFwAlgEspRead,
		DeleteContext: resourceFwAlgEspDelete,

		Schema: map[string]*schema.Schema{
			"default_port_disable": {
				Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable ESP ALG default port 500;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'session-created': ESP Sessions Created; 'helper-created': ESP Helper Sessions Created; 'helper-freed': ESP Helper Sessions Freed; 'helper-freed-used': ESP Helper Sessions freed used; 'helper-freed-unused': ESP Helper Sessions freed unused; 'helper-already-used': ESP Helper Session already used; 'helper-in-rml': ESP Helper Session in Remove List;",
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
func resourceFwAlgEspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgEspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgEsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgEspRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgEspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgEspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgEsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgEspRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgEspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgEspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgEsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgEspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgEspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgEsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwAlgEspSamplingEnable(d []interface{}) []edpt.FwAlgEspSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwAlgEspSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgEspSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwAlgEsp(d *schema.ResourceData) edpt.FwAlgEsp {
	var ret edpt.FwAlgEsp
	ret.Inst.DefaultPortDisable = d.Get("default_port_disable").(string)
	ret.Inst.SamplingEnable = getSliceFwAlgEspSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
