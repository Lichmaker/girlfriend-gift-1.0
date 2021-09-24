package v1

import "lichmaker/girlfriend-gift-1/app/http/controllers/v1/index"

type ApiGroup struct {
	IndexApiGroup index.IndexApiGroup
}

var ApiGroupApp = new(ApiGroup)