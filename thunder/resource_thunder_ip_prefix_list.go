package thunder

//Thunder resource IpPrefixList

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceIpPrefixList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpPrefixListCreate,
		UpdateContext: resourceIpPrefixListUpdate,
		ReadContext:   resourceIpPrefixListRead,
		DeleteContext: resourceIpPrefixListDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"any": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipaddr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ge": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"le": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpPrefixListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating IpPrefixList (Inside resourceIpPrefixListCreate) ")
		name1 := d.Get("name").(string)
		data := dataToIpPrefixList(d)
		logger.Println("[INFO] received formatted data from method data to IpPrefixList --")
		d.SetId(name1)
		err := go_thunder.PostIpPrefixList(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpPrefixListRead(ctx, d, meta)

	}
	return diags
}

func resourceIpPrefixListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpPrefixList (Inside resourceIpPrefixListRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetIpPrefixList(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceIpPrefixListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying IpPrefixList   (Inside resourceIpPrefixListUpdate) ")
		data := dataToIpPrefixList(d)
		logger.Println("[INFO] received formatted data from method data to IpPrefixList ")
		err := go_thunder.PutIpPrefixList(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpPrefixListRead(ctx, d, meta)

	}
	return diags
}

func resourceIpPrefixListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceIpPrefixListDelete) " + name1)
		err := go_thunder.DeleteIpPrefixList(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToIpPrefixList(d *schema.ResourceData) go_thunder.IpPrefixList {
	var vc go_thunder.IpPrefixList
	var c go_thunder.IPPrefixListInstance
	c.IPPrefixListInstanceName = d.Get("name").(string)

	IPPrefixListInstanceRulesCount := d.Get("rules.#").(int)
	c.IPPrefixListInstanceRulesSeq = make([]go_thunder.IPPrefixListInstanceRules, 0, IPPrefixListInstanceRulesCount)

	for i := 0; i < IPPrefixListInstanceRulesCount; i++ {
		var obj1 go_thunder.IPPrefixListInstanceRules
		prefix1 := fmt.Sprintf("rules.%d.", i)
		obj1.IPPrefixListInstanceRulesSeq = d.Get(prefix1 + "seq").(int)
		obj1.IPPrefixListInstanceRulesDescription = d.Get(prefix1 + "description").(string)
		obj1.IPPrefixListInstanceRulesAction = d.Get(prefix1 + "action").(string)
		obj1.IPPrefixListInstanceRulesAny = d.Get(prefix1 + "any").(int)
		obj1.IPPrefixListInstanceRulesIpaddr = d.Get(prefix1 + "ipaddr").(string)
		obj1.IPPrefixListInstanceRulesGe = d.Get(prefix1 + "ge").(int)
		obj1.IPPrefixListInstanceRulesLe = d.Get(prefix1 + "le").(int)
		c.IPPrefixListInstanceRulesSeq = append(c.IPPrefixListInstanceRulesSeq, obj1)
	}

	vc.IPPrefixListInstanceName = c
	return vc
}
