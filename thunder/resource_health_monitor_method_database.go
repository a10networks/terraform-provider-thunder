package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodDatabase() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_database`: DATABASE type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodDatabaseCreate,
		UpdateContext: resourceHealthMonitorMethodDatabaseUpdate,
		ReadContext:   resourceHealthMonitorMethodDatabaseRead,
		DeleteContext: resourceHealthMonitorMethodDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"database": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "DATABASE type",
			},
			"database_name": {
				Type: schema.TypeString, Optional: true, Description: "'mssql': Specify MSSQL database; 'mysql': Specify MySQL database; 'oracle': Specify Oracle database; 'postgresql': Specify PostgreSQL database;",
			},
			"db_column": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the column number for receiving",
			},
			"db_column_integer": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the column number for receiving",
			},
			"db_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the database name",
			},
			"db_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"db_password_str": {
				Type: schema.TypeString, Optional: true, Description: "Configure password",
			},
			"db_receive": {
				Type: schema.TypeString, Optional: true, Description: "Specify the response string",
			},
			"db_receive_integer": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the response integer",
			},
			"db_row": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the row number for receiving",
			},
			"db_row_integer": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the row number for receiving",
			},
			"db_send": {
				Type: schema.TypeString, Optional: true, Description: "Specify the SQL query",
			},
			"db_username": {
				Type: schema.TypeString, Optional: true, Description: "Specify the username",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodDatabaseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDatabaseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDatabase(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodDatabaseRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodDatabaseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDatabaseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDatabase(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodDatabaseRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodDatabaseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDatabaseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDatabase(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodDatabaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodDatabaseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodDatabase(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodDatabase(d *schema.ResourceData) edpt.HealthMonitorMethodDatabase {
	var ret edpt.HealthMonitorMethodDatabase
	ret.Inst.Database = d.Get("database").(int)
	ret.Inst.DatabaseName = d.Get("database_name").(string)
	ret.Inst.DbColumn = d.Get("db_column").(int)
	ret.Inst.DbColumnInteger = d.Get("db_column_integer").(int)
	//omit db_encrypted
	ret.Inst.DbName = d.Get("db_name").(string)
	ret.Inst.DbPassword = d.Get("db_password").(int)
	ret.Inst.DbPasswordStr = d.Get("db_password_str").(string)
	ret.Inst.DbReceive = d.Get("db_receive").(string)
	ret.Inst.DbReceiveInteger = d.Get("db_receive_integer").(int)
	ret.Inst.DbRow = d.Get("db_row").(int)
	ret.Inst.DbRowInteger = d.Get("db_row_integer").(int)
	ret.Inst.DbSend = d.Get("db_send").(string)
	ret.Inst.DbUsername = d.Get("db_username").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
