package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureAutomatedCapturesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_packet_capture_automated_captures_oper`: Operational Status for the object automated-captures\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPacketCaptureAutomatedCapturesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capture_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"automated_capture_config": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"automated_capture_config_line": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityPacketCaptureAutomatedCapturesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureAutomatedCapturesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureAutomatedCapturesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPacketCaptureAutomatedCapturesOperOper := setObjectVisibilityPacketCaptureAutomatedCapturesOperOper(res)
		d.Set("oper", VisibilityPacketCaptureAutomatedCapturesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPacketCaptureAutomatedCapturesOperOper(ret edpt.DataVisibilityPacketCaptureAutomatedCapturesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"capture_name":             ret.DtVisibilityPacketCaptureAutomatedCapturesOper.Oper.CaptureName,
			"automated_capture_config": setSliceVisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig(ret.DtVisibilityPacketCaptureAutomatedCapturesOper.Oper.AutomatedCaptureConfig),
		},
	}
}

func setSliceVisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig(d []edpt.VisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["automated_capture_config_line"] = item.AutomatedCaptureConfigLine
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityPacketCaptureAutomatedCapturesOperOper(d []interface{}) edpt.VisibilityPacketCaptureAutomatedCapturesOperOper {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureAutomatedCapturesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureName = in["capture_name"].(string)
		ret.AutomatedCaptureConfig = getSliceVisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig(in["automated_capture_config"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig(d []interface{}) []edpt.VisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureAutomatedCapturesOperOperAutomatedCaptureConfig
		oi.AutomatedCaptureConfigLine = in["automated_capture_config_line"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureAutomatedCapturesOper(d *schema.ResourceData) edpt.VisibilityPacketCaptureAutomatedCapturesOper {
	var ret edpt.VisibilityPacketCaptureAutomatedCapturesOper

	ret.Oper = getObjectVisibilityPacketCaptureAutomatedCapturesOperOper(d.Get("oper").([]interface{}))
	return ret
}
