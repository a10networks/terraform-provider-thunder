package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLoggingConsole() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceLoggingConsoleCreate,
        UpdateContext: resourceLoggingConsoleUpdate,
        ReadContext: resourceLoggingConsoleRead,
        DeleteContext: resourceLoggingConsoleDelete,
        Schema: map[string]*schema.Schema{
            "console_levelname": {
                Type: schema.TypeString, Required: true, Description: "'disable': Do not send log to console; 'emergency': System unusable log messages      (severity=0); 'alert': Action must be taken immediately  (severity=1); 'critical': Critical conditions               (severity=2); 'error': Error conditions                  (severity=3); 'warning': Warning conditions                (severity=4); 'notification': Normal but significant conditions (severity=5); 'information': Informational messages            (severity=6); 'debugging': Debug level messages              (severity=7); ",
                ValidateFunc: validation.StringInSlice([]string{"disable", "emergency", "alert", "critical", "error", "warning", "notification", "information", "debugging"}, false),
            },
            "uuid": {
                Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
            },
        },
    }
}

func resourceLoggingConsoleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingConsoleCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingConsole(d)
        d.SetId(obj.GetId())
        err := PostEndpointLoggingConsole(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingConsoleRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingConsoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingConsoleRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointLoggingConsole(client.Token, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceLoggingConsoleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingConsoleUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingConsole(d)
        err := PutEndpointLoggingConsole(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingConsoleRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingConsoleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingConsoleDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointLoggingConsole(client.Token, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointLoggingConsole(d *schema.ResourceData) EndpointLoggingConsole{
    var ret EndpointLoggingConsole
    ret.Inst.ConsoleLevelname = d.Get("console_levelname").(string)
    //omit uuid
    return ret
}
