package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpv6RouteStaticBfdTrunk() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceIpv6RouteStaticBfdTrunkCreate,
        UpdateContext: resourceIpv6RouteStaticBfdTrunkUpdate,
        ReadContext: resourceIpv6RouteStaticBfdTrunkRead,
        DeleteContext: resourceIpv6RouteStaticBfdTrunkDelete,
        Schema: map[string]*schema.Schema{
            "trunk_num": {
                Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Trunk interface",
                ValidateFunc: validation.IntBetween(1, 4096),
            },
            "nexthop_ipv6_ll": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Nexthop IPv6 address (Link_local)",
                ValidateFunc: validation.IsIPv6Address,
            },
            "template": {
                Type: schema.TypeString, Optional: true, Description: "Configure tracking template (bind tracking template name)",
                ValidateFunc: validation.StringLenBetween(1, 63),
            },
            "threshold": {
                Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
                ValidateFunc: validation.IntBetween(1, 255),
            },
            "action": {
                Type: schema.TypeString, Optional: true, Description: "'down': BFD down;  (BFD state)",
                ValidateFunc: validation.StringInSlice([]string{"down"}, false),
            },
            "uuid": {
                Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
            },
        },
    }
}

func resourceIpv6RouteStaticBfdTrunkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdTrunkCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpv6RouteStaticBfdTrunk(d)
        d.SetId(obj.GetId())
        err := PostEndpointIpv6RouteStaticBfdTrunk(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpv6RouteStaticBfdTrunkRead(ctx, d, meta)
    }
    return diags
}

func resourceIpv6RouteStaticBfdTrunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdTrunkRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointIpv6RouteStaticBfdTrunk(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceIpv6RouteStaticBfdTrunkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdTrunkUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpv6RouteStaticBfdTrunk(d)
        err := PutEndpointIpv6RouteStaticBfdTrunk(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpv6RouteStaticBfdTrunkRead(ctx, d, meta)
    }
    return diags
}

func resourceIpv6RouteStaticBfdTrunkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdTrunkDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointIpv6RouteStaticBfdTrunk(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointIpv6RouteStaticBfdTrunk(d *schema.ResourceData) EndpointIpv6RouteStaticBfdTrunk{
    var ret EndpointIpv6RouteStaticBfdTrunk
    ret.Inst.TrunkNum = d.Get("trunk_num").(int)
    ret.Inst.NexthopIpv6Ll = d.Get("nexthop_ipv6_ll").(string)
    ret.Inst.Template = d.Get("template").(string)
    ret.Inst.Threshold = d.Get("threshold").(int)
    ret.Inst.Action = d.Get("action").(string)
    //omit uuid
    return ret
}
