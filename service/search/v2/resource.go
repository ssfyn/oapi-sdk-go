// Code generated by Lark OpenAPI.

package larksearch

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type V2 struct {
	App            *app            // app
	DataSource     *dataSource     // 数据源
	DataSourceItem *dataSourceItem // 数据项
	Message        *message        // message
	Schema         *schema         // 数据范式
}

func New(config *larkcore.Config) *V2 {
	return &V2{
		App:            &app{config: config},
		DataSource:     &dataSource{config: config},
		DataSourceItem: &dataSourceItem{config: config},
		Message:        &message{config: config},
		Schema:         &schema{config: config},
	}
}

type app struct {
	config *larkcore.Config
}
type dataSource struct {
	config *larkcore.Config
}
type dataSourceItem struct {
	config *larkcore.Config
}
type message struct {
	config *larkcore.Config
}
type schema struct {
	config *larkcore.Config
}

// Create
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=create&project=search&resource=app&version=v2
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/create_app.go
func (a *app) Create(ctx context.Context, req *CreateAppReq, options ...larkcore.RequestOptionFunc) (*CreateAppResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/app"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateAppResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create 创建数据源
//
// - 创建一个数据源
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/create_dataSource.go
func (d *dataSource) Create(ctx context.Context, req *CreateDataSourceReq, options ...larkcore.RequestOptionFunc) (*CreateDataSourceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDataSourceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Delete 删除数据源
//
// - 删除一个已存在的数据源
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/delete_dataSource.go
func (d *dataSource) Delete(ctx context.Context, req *DeleteDataSourceReq, options ...larkcore.RequestOptionFunc) (*DeleteDataSourceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources/:data_source_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDataSourceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Get 获取数据源
//
// - 获取已经创建的数据源
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/get_dataSource.go
func (d *dataSource) Get(ctx context.Context, req *GetDataSourceReq, options ...larkcore.RequestOptionFunc) (*GetDataSourceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources/:data_source_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDataSourceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// List 批量获取数据源
//
// - 批量获取创建的数据源信息
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/list_dataSource.go
func (d *dataSource) List(ctx context.Context, req *ListDataSourceReq, options ...larkcore.RequestOptionFunc) (*ListDataSourceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDataSourceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) ListByIterator(ctx context.Context, req *ListDataSourceReq, options ...larkcore.RequestOptionFunc) (*ListDataSourceIterator, error) {
	return &ListDataSourceIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.List,
		options:  options,
		limit:    req.Limit}, nil
}

// Patch 修改数据源
//
// - 更新一个已经存在的数据源
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/patch
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/patch_dataSource.go
func (d *dataSource) Patch(ctx context.Context, req *PatchDataSourceReq, options ...larkcore.RequestOptionFunc) (*PatchDataSourceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources/:data_source_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchDataSourceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create 索引数据项
//
// - 索引一条数据记录
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/create_dataSourceItem.go
func (d *dataSourceItem) Create(ctx context.Context, req *CreateDataSourceItemReq, options ...larkcore.RequestOptionFunc) (*CreateDataSourceItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources/:data_source_id/items"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDataSourceItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Delete 删除数据项
//
// - 删除数据项
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/delete_dataSourceItem.go
func (d *dataSourceItem) Delete(ctx context.Context, req *DeleteDataSourceItemReq, options ...larkcore.RequestOptionFunc) (*DeleteDataSourceItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources/:data_source_id/items/:item_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDataSourceItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Get 获取数据项
//
// - 获取单个数据记录
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/get_dataSourceItem.go
func (d *dataSourceItem) Get(ctx context.Context, req *GetDataSourceItemReq, options ...larkcore.RequestOptionFunc) (*GetDataSourceItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/data_sources/:data_source_id/items/:item_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDataSourceItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=create&project=search&resource=message&version=v2
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/create_message.go
func (m *message) Create(ctx context.Context, req *CreateMessageReq, options ...larkcore.RequestOptionFunc) (*CreateMessageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/message"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, m.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMessageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, m.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create 创建数据范式
//
// - 创建一个数据范式
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/schema/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/create_schema.go
func (s *schema) Create(ctx context.Context, req *CreateSchemaReq, options ...larkcore.RequestOptionFunc) (*CreateSchemaResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/schemas"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSchemaResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Delete 删除数据范式
//
// - 删除已存在的数据范式
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/schema/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/delete_schema.go
func (s *schema) Delete(ctx context.Context, req *DeleteSchemaReq, options ...larkcore.RequestOptionFunc) (*DeleteSchemaResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/schemas/:schema_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSchemaResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Get 获取数据范式
//
// - 获取单个数据范式
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/schema/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/get_schema.go
func (s *schema) Get(ctx context.Context, req *GetSchemaReq, options ...larkcore.RequestOptionFunc) (*GetSchemaResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/schemas/:schema_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSchemaResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Patch 修改数据范式
//
// - 修改数据范式
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/schema/patch
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/searchv2/patch_schema.go
func (s *schema) Patch(ctx context.Context, req *PatchSchemaReq, options ...larkcore.RequestOptionFunc) (*PatchSchemaResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/search/v2/schemas/:schema_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchSchemaResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
