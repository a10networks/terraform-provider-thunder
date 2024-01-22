package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbProtocolLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_protocol_limit`: Specify limit for GSLB Message Protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbProtocolLimitCreate,
		UpdateContext: resourceGslbProtocolLimitUpdate,
		ReadContext:   resourceGslbProtocolLimitRead,
		DeleteContext: resourceGslbProtocolLimitDelete,

		Schema: map[string]*schema.Schema{
			"ardt_query": {
				Type: schema.TypeInt, Optional: true, Default: 200, Description: "Query Messages of Active RDT, default is 200 (Number)",
			},
			"ardt_response": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Response Messages of Active RDT, default is 1000 (Number)",
			},
			"ardt_session": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Sessions of Active RDT, default is 32768 (Number)",
			},
			"conn_response": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Response Messages of Connection Load, default is no limit (Number)",
			},
			"message": {
				Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Amount of Messages, default is 10000 (Number)",
			},
			"response": {
				Type: schema.TypeInt, Optional: true, Default: 3600, Description: "Amount of Response Messages, default is 3600 (Number)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbProtocolLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbProtocolLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbProtocolLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbProtocolLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbProtocolLimit(d *schema.ResourceData) edpt.GslbProtocolLimit {
	var ret edpt.GslbProtocolLimit
	ret.Inst.ArdtQuery = d.Get("ardt_query").(int)
	ret.Inst.ArdtResponse = d.Get("ardt_response").(int)
	ret.Inst.ArdtSession = d.Get("ardt_session").(int)
	ret.Inst.ConnResponse = d.Get("conn_response").(int)
	ret.Inst.Message = d.Get("message").(int)
	ret.Inst.Response = d.Get("response").(int)
	//omit uuid
	return ret
}
