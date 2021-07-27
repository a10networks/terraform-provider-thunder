module github.com/a10networks/terraform-provider-thunder

go 1.16

replace github.com/go_thunder/thunder => ./go_thunder/thunder

replace util => ./util/UtilLog

require (
	github.com/go_thunder/thunder v0.0.0-00010101000000-000000000000
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	util v0.0.0-00010101000000-000000000000

)
