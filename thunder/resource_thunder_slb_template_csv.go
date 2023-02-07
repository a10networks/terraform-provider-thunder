package thunder

//Thunder resource SlbTemplateCSV

import (
	"context"
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateCSV() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateCSVCreate,
		UpdateContext: resourceSlbTemplateCSVUpdate,
		ReadContext:   resourceSlbTemplateCSVRead,
		DeleteContext: resourceSlbTemplateCSVDelete,
		Schema: map[string]*schema.Schema{
			"delim_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"multiple_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"csv_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"csv_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"delim_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateCSVCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateCSV (Inside resourceSlbTemplateCSVCreate) ")
		name := d.Get("csv_name").(string)
		data := dataToSlbTemplateCSV(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateCSV --")
		d.SetId(name)
		err := go_thunder.PostSlbTemplateCSV(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateCSVRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateCSVRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplateCSV (Inside resourceSlbTemplateCSVRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateCSV(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbTemplateCSVUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateCSV   (Inside resourceSlbTemplateCSVUpdate) ")
		name := d.Get("csv_name").(string)
		data := dataToSlbTemplateCSV(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateCSV ")
		d.SetId(name)
		err := go_thunder.PutSlbTemplateCSV(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateCSVRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateCSVDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateCSVDelete) " + name)
		err := go_thunder.DeleteSlbTemplateCSV(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateCSV(d *schema.ResourceData) go_thunder.CSV {
	var vc go_thunder.CSV
	var c go_thunder.CSVInstance

	c.CsvName = d.Get("csv_name").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Ipv6Enable = d.Get("ipv6_enable").(int)
	c.DelimNum = d.Get("delim_num").(int)
	c.DelimChar = d.Get("delim_char").(string)

	MultipleFieldsCount := d.Get("multiple_fields.#").(int)
	c.Field = make([]go_thunder.MultipleFields, 0, MultipleFieldsCount)

	for i := 0; i < MultipleFieldsCount; i++ {
		var obj1 go_thunder.MultipleFields
		prefix := fmt.Sprintf("multiple_fields.%d.", i)
		obj1.Field = d.Get(prefix + "field").(int)
		obj1.CsvType = d.Get(prefix + "csv_type").(string)
		c.Field = append(c.Field, obj1)
	}

	vc.UUID = c
	return vc
}
