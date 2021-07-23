package thunder

//Thunder resource AccessListStandard

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccessListStandard() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccessListStandardCreate,
		UpdateContext: resourceAccessListStandardUpdate,
		ReadContext:   resourceAccessListStandardRead,
		DeleteContext: resourceAccessListStandardDelete,
		Schema: map[string]*schema.Schema{
			"std": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stdrules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"std_remark": {
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
						"host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"subnet": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"rev_subnet_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"transparent_session_only": {
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

func resourceAccessListStandardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating AccessListStandard (Inside resourceAccessListStandardCreate) ")
		name1 := d.Get("std").(int)
		data := dataToAccessListStandard(d)
		logger.Println("[INFO] received formatted data from method data to AccessListStandard --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostAccessListStandard(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceAccessListStandardRead(ctx, d, meta)

	}
	return diags
}

func resourceAccessListStandardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading AccessListStandard (Inside resourceAccessListStandardRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetAccessListStandard(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceAccessListStandardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying AccessListStandard   (Inside resourceAccessListStandardUpdate) ")
		data := dataToAccessListStandard(d)
		logger.Println("[INFO] received formatted data from method data to AccessListStandard ")
		err := go_thunder.PutAccessListStandard(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceAccessListStandardRead(ctx, d, meta)

	}
	return diags
}

func resourceAccessListStandardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceAccessListStandardDelete) " + name1)
		err := go_thunder.DeleteAccessListStandard(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToAccessListStandard(d *schema.ResourceData) go_thunder.AccessListStandard {
	var vc go_thunder.AccessListStandard
	var c go_thunder.AccessListStandardInstance
	c.Std = d.Get("std").(int)

	StdrulesCount := d.Get("stdrules.#").(int)
	c.SeqNum = make([]go_thunder.AccessListStdrules, 0, StdrulesCount)

	for i := 0; i < StdrulesCount; i++ {
		var obj1 go_thunder.AccessListStdrules
		prefix1 := fmt.Sprintf("stdrules.%d.", i)
		obj1.SeqNum = d.Get(prefix1 + "seq_num").(int)
		obj1.StdRemark = d.Get(prefix1 + "std_remark").(string)
		obj1.Action = d.Get(prefix1 + "action").(string)
		obj1.Any = d.Get(prefix1 + "any").(int)
		obj1.Host = d.Get(prefix1 + "host").(string)
		obj1.Subnet = d.Get(prefix1 + "subnet").(string)
		obj1.RevSubnetMask = d.Get(prefix1 + "rev_subnet_mask").(string)
		obj1.Log = d.Get(prefix1 + "log").(int)
		obj1.TransparentSessionOnly = d.Get(prefix1 + "transparent_session_only").(int)
		c.SeqNum = append(c.SeqNum, obj1)
	}

	vc.Std = c
	return vc
}
