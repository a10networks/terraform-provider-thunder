module github.com/a10networks/terraform-provider-thunder

go 1.16

replace github.com/go_thunder/thunder => ./go_thunder/thunder

replace util => ./util/UtilLog

require (
	github.com/clarketm/json v1.15.7
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	util v0.0.0-00010101000000-000000000000

)
