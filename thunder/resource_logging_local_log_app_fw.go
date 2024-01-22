package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogAppFw() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_local_log_app_fw`: Application firewall\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingLocalLogAppFwCreate,
		UpdateContext: resourceLoggingLocalLogAppFwUpdate,
		ReadContext:   resourceLoggingLocalLogAppFwRead,
		DeleteContext: resourceLoggingLocalLogAppFwDelete,

		Schema: map[string]*schema.Schema{
			"dot_plot": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"top_n": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceLoggingLocalLogAppFwCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAppFwCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAppFw(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLocalLogAppFwRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingLocalLogAppFwUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAppFwUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAppFw(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLocalLogAppFwRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingLocalLogAppFwDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAppFwDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAppFw(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingLocalLogAppFwRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogAppFwRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogAppFw(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectLoggingLocalLogAppFwDotPlot1050(d []interface{}) edpt.LoggingLocalLogAppFwDotPlot1050 {

	var ret edpt.LoggingLocalLogAppFwDotPlot1050
	return ret
}

func getObjectLoggingLocalLogAppFwTopN1051(d []interface{}) edpt.LoggingLocalLogAppFwTopN1051 {

	var ret edpt.LoggingLocalLogAppFwTopN1051
	return ret
}

func dataToEndpointLoggingLocalLogAppFw(d *schema.ResourceData) edpt.LoggingLocalLogAppFw {
	var ret edpt.LoggingLocalLogAppFw
	ret.Inst.DotPlot = getObjectLoggingLocalLogAppFwDotPlot1050(d.Get("dot_plot").([]interface{}))
	ret.Inst.TopN = getObjectLoggingLocalLogAppFwTopN1051(d.Get("top_n").([]interface{}))
	//omit uuid
	return ret
}
