package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTerminal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_terminal`: Set Terminal Startup Parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceTerminalCreate,
		UpdateContext: resourceTerminalUpdate,
		ReadContext:   resourceTerminalRead,
		DeleteContext: resourceTerminalDelete,

		Schema: map[string]*schema.Schema{
			"auto_size": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable terminal length and width automatically (not work if width or length set to 0)",
			},
			"editing": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable command line editing",
			},
			"gslb_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_prompt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "The gslb status prompt function set",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Group status show disable",
						},
						"group_role": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show GSLB group role on CLI prompt",
						},
						"symbol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show \"gslb\" symbol on CLI prompt",
						},
					},
				},
			},
			"history_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable terminal history",
						},
						"size": {
							Type: schema.TypeInt, Optional: true, Default: 256, Description: "Set history buffer size (Size of history buffer, default is 256)",
						},
					},
				},
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "Set interval for closing connection when there is no input detected (Timeout in minutes, 0 means never timeout, default is 15)",
			},
			"length": {
				Type: schema.TypeInt, Optional: true, Default: 24, Description: "Set number of lines on a screen(0 for no pausing) (Number of lines on screen, 0 for no pausing, default is 24)",
			},
			"prompt_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prompt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure the normal prompt format",
						},
						"ha_status": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Display HA status in prompt, eg. Active, Standby, ForcedStandby",
						},
						"hostname": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Display hostname in prompt",
						},
						"vcs_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vcs_status": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Display VCS status in prompt, eg. vMaster, vBlade",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"width": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "Set width of the display terminal (Number of characters on a screen line, 0 means infinite, default is 80)",
			},
		},
	}
}
func resourceTerminalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTerminalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTerminal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTerminalRead(ctx, d, meta)
	}
	return diags
}

func resourceTerminalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTerminalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTerminal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTerminalRead(ctx, d, meta)
	}
	return diags
}
func resourceTerminalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTerminalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTerminal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTerminalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTerminalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTerminal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTerminalGslbCfg(d []interface{}) edpt.TerminalGslbCfg {

	count1 := len(d)
	var ret edpt.TerminalGslbCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbPrompt = in["gslb_prompt"].(int)
		ret.Disable = in["disable"].(int)
		ret.GroupRole = in["group_role"].(int)
		ret.Symbol = in["symbol"].(int)
	}
	return ret
}

func getObjectTerminalHistoryCfg(d []interface{}) edpt.TerminalHistoryCfg {

	count1 := len(d)
	var ret edpt.TerminalHistoryCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Size = in["size"].(int)
	}
	return ret
}

func getObjectTerminalPromptCfg(d []interface{}) edpt.TerminalPromptCfg {

	count1 := len(d)
	var ret edpt.TerminalPromptCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Prompt = in["prompt"].(int)
		ret.HaStatus = in["ha_status"].(int)
		ret.Hostname = in["hostname"].(int)
		ret.VcsCfg = getObjectTerminalPromptCfgVcsCfg(in["vcs_cfg"].([]interface{}))
	}
	return ret
}

func getObjectTerminalPromptCfgVcsCfg(d []interface{}) edpt.TerminalPromptCfgVcsCfg {

	count1 := len(d)
	var ret edpt.TerminalPromptCfgVcsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VcsStatus = in["vcs_status"].(int)
	}
	return ret
}

func dataToEndpointTerminal(d *schema.ResourceData) edpt.Terminal {
	var ret edpt.Terminal
	ret.Inst.AutoSize = d.Get("auto_size").(int)
	ret.Inst.Editing = d.Get("editing").(int)
	ret.Inst.GslbCfg = getObjectTerminalGslbCfg(d.Get("gslb_cfg").([]interface{}))
	ret.Inst.HistoryCfg = getObjectTerminalHistoryCfg(d.Get("history_cfg").([]interface{}))
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.PromptCfg = getObjectTerminalPromptCfg(d.Get("prompt_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Width = d.Get("width").(int)
	return ret
}
