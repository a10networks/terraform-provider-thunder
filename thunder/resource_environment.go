package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_environment`: environment status colletion parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceEnvironmentCreate,
		UpdateContext: resourceEnvironmentUpdate,
		ReadContext:   resourceEnvironmentRead,
		DeleteContext: resourceEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"threshold_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"low": {
							Type: schema.TypeInt, Optional: true, Description: "Low threshold value in Celsius - default 1",
						},
						"medium": {
							Type: schema.TypeInt, Optional: true, Description: "Medium threshold value in Celsius - default 1",
						},
						"high": {
							Type: schema.TypeInt, Optional: true, Description: "High threshold value in Celsius - default 1",
						},
					},
				},
			},
			"update_interval": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Interval",
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
func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnvironmentCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnvironment(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnvironmentRead(ctx, d, meta)
	}
	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnvironmentUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnvironment(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnvironmentRead(ctx, d, meta)
	}
	return diags
}
func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnvironmentDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnvironment(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnvironmentRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnvironment(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectEnvironmentThresholdCfg(d []interface{}) edpt.EnvironmentThresholdCfg {

	count1 := len(d)
	var ret edpt.EnvironmentThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Low = in["low"].(int)
		ret.Medium = in["medium"].(int)
		ret.High = in["high"].(int)
	}
	return ret
}

func getObjectEnvironmentUpdateInterval(d []interface{}) edpt.EnvironmentUpdateInterval {

	count1 := len(d)
	var ret edpt.EnvironmentUpdateInterval
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
	}
	return ret
}

func dataToEndpointEnvironment(d *schema.ResourceData) edpt.Environment {
	var ret edpt.Environment
	ret.Inst.ThresholdCfg = getObjectEnvironmentThresholdCfg(d.Get("threshold_cfg").([]interface{}))
	ret.Inst.UpdateInterval = getObjectEnvironmentUpdateInterval(d.Get("update_interval").([]interface{}))
	//omit uuid
	return ret
}
