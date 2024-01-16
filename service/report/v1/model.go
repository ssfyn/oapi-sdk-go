// Package report code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkreport

import (
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

const (
	IncludeDeletedExclude = 0 // 不包括已删除
	IncludeDeletedInclude = 1 // 包括已删除

)

const (
	UserIdTypeUserId  = "user_id"  // 以user_id来识别用户
	UserIdTypeUnionId = "union_id" // 以union_id来识别用户
	UserIdTypeOpenId  = "open_id"  // 以open_id来识别用户
)

const (
	UserIdTypeRemoveRuleViewUserId  = "user_id"  //
	UserIdTypeRemoveRuleViewUnionId = "union_id" //
	UserIdTypeRemoveRuleViewOpenId  = "open_id"  //
)

const (
	UserIdTypeQueryTaskUserId  = "user_id"  // 以user_id来识别用户
	UserIdTypeQueryTaskUnionId = "union_id" // 以union_id来识别用户
	UserIdTypeQueryTaskOpenId  = "open_id"  // 以open_id来识别用户
)

type DepartmentId struct {
	DepartmentId     *string `json:"department_id,omitempty"`      //
	OpenDepartmentId *string `json:"open_department_id,omitempty"` //
}

type DepartmentIdBuilder struct {
	departmentId         string //
	departmentIdFlag     bool
	openDepartmentId     string //
	openDepartmentIdFlag bool
}

func NewDepartmentIdBuilder() *DepartmentIdBuilder {
	builder := &DepartmentIdBuilder{}
	return builder
}

//
//
// 示例值：
func (builder *DepartmentIdBuilder) DepartmentId(departmentId string) *DepartmentIdBuilder {
	builder.departmentId = departmentId
	builder.departmentIdFlag = true
	return builder
}

//
//
// 示例值：
func (builder *DepartmentIdBuilder) OpenDepartmentId(openDepartmentId string) *DepartmentIdBuilder {
	builder.openDepartmentId = openDepartmentId
	builder.openDepartmentIdFlag = true
	return builder
}

func (builder *DepartmentIdBuilder) Build() *DepartmentId {
	req := &DepartmentId{}
	if builder.departmentIdFlag {
		req.DepartmentId = &builder.departmentId

	}
	if builder.openDepartmentIdFlag {
		req.OpenDepartmentId = &builder.openDepartmentId

	}
	return req
}

type FormContent struct {
	FieldId    *string `json:"field_id,omitempty"`    // 表单字段ID
	FieldName  *string `json:"field_name,omitempty"`  // 表单字段名称
	FieldValue *string `json:"field_value,omitempty"` // 表单字段值
}

type FormContentBuilder struct {
	fieldId        string // 表单字段ID
	fieldIdFlag    bool
	fieldName      string // 表单字段名称
	fieldNameFlag  bool
	fieldValue     string // 表单字段值
	fieldValueFlag bool
}

func NewFormContentBuilder() *FormContentBuilder {
	builder := &FormContentBuilder{}
	return builder
}

// 表单字段ID
//
// 示例值：6968626905868156948
func (builder *FormContentBuilder) FieldId(fieldId string) *FormContentBuilder {
	builder.fieldId = fieldId
	builder.fieldIdFlag = true
	return builder
}

// 表单字段名称
//
// 示例值：表单测试
func (builder *FormContentBuilder) FieldName(fieldName string) *FormContentBuilder {
	builder.fieldName = fieldName
	builder.fieldNameFlag = true
	return builder
}

// 表单字段值
//
// 示例值：测试数据
func (builder *FormContentBuilder) FieldValue(fieldValue string) *FormContentBuilder {
	builder.fieldValue = fieldValue
	builder.fieldValueFlag = true
	return builder
}

func (builder *FormContentBuilder) Build() *FormContent {
	req := &FormContent{}
	if builder.fieldIdFlag {
		req.FieldId = &builder.fieldId

	}
	if builder.fieldNameFlag {
		req.FieldName = &builder.fieldName

	}
	if builder.fieldValueFlag {
		req.FieldValue = &builder.fieldValue

	}
	return req
}

type FormField struct {
	Name *string `json:"name,omitempty"` // 字段名称
	Type *string `json:"type,omitempty"` // 字段类型
}

type FormFieldBuilder struct {
	name     string // 字段名称
	nameFlag bool
	type_    string // 字段类型
	typeFlag bool
}

func NewFormFieldBuilder() *FormFieldBuilder {
	builder := &FormFieldBuilder{}
	return builder
}

// 字段名称
//
// 示例值：ou_133f0b6d0f097cf7d7ba00b38fffb112
func (builder *FormFieldBuilder) Name(name string) *FormFieldBuilder {
	builder.name = name
	builder.nameFlag = true
	return builder
}

// 字段类型
//
// 示例值：张三
func (builder *FormFieldBuilder) Type(type_ string) *FormFieldBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

func (builder *FormFieldBuilder) Build() *FormField {
	req := &FormField{}
	if builder.nameFlag {
		req.Name = &builder.name

	}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	return req
}

type Rule struct {
	RuleId                  *string      `json:"rule_id,omitempty"`                    // 规则唯一标识
	Name                    *string      `json:"name,omitempty"`                       // 规则名称
	IconName                *string      `json:"icon_name,omitempty"`                  // 规则图标
	CreatedAt               *int         `json:"created_at,omitempty"`                 // 创建时间
	CreatorUserId           *string      `json:"creator_user_id,omitempty"`            // 创建人ID
	CreatorUserName         *string      `json:"creator_user_name,omitempty"`          // 创建人名称
	OwnerUserId             *string      `json:"owner_user_id,omitempty"`              // 规则所有者ID
	OwnerUserName           *string      `json:"owner_user_name,omitempty"`            // 规则所有者名称
	FormSchema              []*FormField `json:"form_schema,omitempty"`                // 表单定义
	IsDeleted               *int         `json:"is_deleted,omitempty"`                 // 规则是否已删除
	NeedReportUserIds       []string     `json:"need_report_user_ids,omitempty"`       // 需要汇报的用户ID列表
	NeedReportDepartmentIds []string     `json:"need_report_department_ids,omitempty"` // 需要汇报的部门ID列表（如果id为0，表示全员）
	NeedReportChatIds       []string     `json:"need_report_chat_ids,omitempty"`       // 需要汇报的群ID列表
	CcUserIds               []string     `json:"cc_user_ids,omitempty"`                // 抄送用户ID列表
	CcDepartmentIds         []string     `json:"cc_department_ids,omitempty"`          // 抄送部门ID列表
	ToUserIds               []string     `json:"to_user_ids,omitempty"`                // 汇报对象用户ID列表
	ToChatIds               []string     `json:"to_chat_ids,omitempty"`                // 汇报对象群ID列表
	ToLeaders               []int        `json:"to_leaders,omitempty"`                 // 上级汇报对象，0表示第一级，依次类推，最大为5表示第六级
	ToDepartmentOwners      []int        `json:"to_department_owners,omitempty"`       // 部门负责人汇报对象，0表示第一级，依次类推，最大为5表示第六级
	ManagerUserIds          []string     `json:"manager_user_ids,omitempty"`           // 规则管理员用户ID列表
	CcChatIds               []string     `json:"cc_chat_ids,omitempty"`                // 抄送群ID列表
}

type RuleBuilder struct {
	ruleId                      string // 规则唯一标识
	ruleIdFlag                  bool
	name                        string // 规则名称
	nameFlag                    bool
	iconName                    string // 规则图标
	iconNameFlag                bool
	createdAt                   int // 创建时间
	createdAtFlag               bool
	creatorUserId               string // 创建人ID
	creatorUserIdFlag           bool
	creatorUserName             string // 创建人名称
	creatorUserNameFlag         bool
	ownerUserId                 string // 规则所有者ID
	ownerUserIdFlag             bool
	ownerUserName               string // 规则所有者名称
	ownerUserNameFlag           bool
	formSchema                  []*FormField // 表单定义
	formSchemaFlag              bool
	isDeleted                   int // 规则是否已删除
	isDeletedFlag               bool
	needReportUserIds           []string // 需要汇报的用户ID列表
	needReportUserIdsFlag       bool
	needReportDepartmentIds     []string // 需要汇报的部门ID列表（如果id为0，表示全员）
	needReportDepartmentIdsFlag bool
	needReportChatIds           []string // 需要汇报的群ID列表
	needReportChatIdsFlag       bool
	ccUserIds                   []string // 抄送用户ID列表
	ccUserIdsFlag               bool
	ccDepartmentIds             []string // 抄送部门ID列表
	ccDepartmentIdsFlag         bool
	toUserIds                   []string // 汇报对象用户ID列表
	toUserIdsFlag               bool
	toChatIds                   []string // 汇报对象群ID列表
	toChatIdsFlag               bool
	toLeaders                   []int // 上级汇报对象，0表示第一级，依次类推，最大为5表示第六级
	toLeadersFlag               bool
	toDepartmentOwners          []int // 部门负责人汇报对象，0表示第一级，依次类推，最大为5表示第六级
	toDepartmentOwnersFlag      bool
	managerUserIds              []string // 规则管理员用户ID列表
	managerUserIdsFlag          bool
	ccChatIds                   []string // 抄送群ID列表
	ccChatIdsFlag               bool
}

func NewRuleBuilder() *RuleBuilder {
	builder := &RuleBuilder{}
	return builder
}

// 规则唯一标识
//
// 示例值：6894788526240432147
func (builder *RuleBuilder) RuleId(ruleId string) *RuleBuilder {
	builder.ruleId = ruleId
	builder.ruleIdFlag = true
	return builder
}

// 规则名称
//
// 示例值：工作月报
func (builder *RuleBuilder) Name(name string) *RuleBuilder {
	builder.name = name
	builder.nameFlag = true
	return builder
}

// 规则图标
//
// 示例值：日报
func (builder *RuleBuilder) IconName(iconName string) *RuleBuilder {
	builder.iconName = iconName
	builder.iconNameFlag = true
	return builder
}

// 创建时间
//
// 示例值：1622427266
func (builder *RuleBuilder) CreatedAt(createdAt int) *RuleBuilder {
	builder.createdAt = createdAt
	builder.createdAtFlag = true
	return builder
}

// 创建人ID
//
// 示例值：ou_133f0b6d0f097cf7d7ba00b38fffb110
func (builder *RuleBuilder) CreatorUserId(creatorUserId string) *RuleBuilder {
	builder.creatorUserId = creatorUserId
	builder.creatorUserIdFlag = true
	return builder
}

// 创建人名称
//
// 示例值：张三
func (builder *RuleBuilder) CreatorUserName(creatorUserName string) *RuleBuilder {
	builder.creatorUserName = creatorUserName
	builder.creatorUserNameFlag = true
	return builder
}

// 规则所有者ID
//
// 示例值：ou_133f0b6d0f097cf7d7ba00b38fffb111
func (builder *RuleBuilder) OwnerUserId(ownerUserId string) *RuleBuilder {
	builder.ownerUserId = ownerUserId
	builder.ownerUserIdFlag = true
	return builder
}

// 规则所有者名称
//
// 示例值：张三
func (builder *RuleBuilder) OwnerUserName(ownerUserName string) *RuleBuilder {
	builder.ownerUserName = ownerUserName
	builder.ownerUserNameFlag = true
	return builder
}

// 表单定义
//
// 示例值：
func (builder *RuleBuilder) FormSchema(formSchema []*FormField) *RuleBuilder {
	builder.formSchema = formSchema
	builder.formSchemaFlag = true
	return builder
}

// 规则是否已删除
//
// 示例值：0
func (builder *RuleBuilder) IsDeleted(isDeleted int) *RuleBuilder {
	builder.isDeleted = isDeleted
	builder.isDeletedFlag = true
	return builder
}

// 需要汇报的用户ID列表
//
// 示例值：['ou_c04cebc780341ab22bd311ba6902ffsd']
func (builder *RuleBuilder) NeedReportUserIds(needReportUserIds []string) *RuleBuilder {
	builder.needReportUserIds = needReportUserIds
	builder.needReportUserIdsFlag = true
	return builder
}

// 需要汇报的部门ID列表（如果id为0，表示全员）
//
// 示例值：
func (builder *RuleBuilder) NeedReportDepartmentIds(needReportDepartmentIds []string) *RuleBuilder {
	builder.needReportDepartmentIds = needReportDepartmentIds
	builder.needReportDepartmentIdsFlag = true
	return builder
}

// 需要汇报的群ID列表
//
// 示例值：['oc_a7bb9ca5efa68ab8b4fdd2e3b54ffref']
func (builder *RuleBuilder) NeedReportChatIds(needReportChatIds []string) *RuleBuilder {
	builder.needReportChatIds = needReportChatIds
	builder.needReportChatIdsFlag = true
	return builder
}

// 抄送用户ID列表
//
// 示例值：['ou_45454c20ef2c92c173445abf6f4955rf']
func (builder *RuleBuilder) CcUserIds(ccUserIds []string) *RuleBuilder {
	builder.ccUserIds = ccUserIds
	builder.ccUserIdsFlag = true
	return builder
}

// 抄送部门ID列表
//
// 示例值：['od-251480c0bfb8c5c8784ea194ef8b734d']
func (builder *RuleBuilder) CcDepartmentIds(ccDepartmentIds []string) *RuleBuilder {
	builder.ccDepartmentIds = ccDepartmentIds
	builder.ccDepartmentIdsFlag = true
	return builder
}

// 汇报对象用户ID列表
//
// 示例值：['ou_c04cebc780341ab22bd311ba6902fdfe']
func (builder *RuleBuilder) ToUserIds(toUserIds []string) *RuleBuilder {
	builder.toUserIds = toUserIds
	builder.toUserIdsFlag = true
	return builder
}

// 汇报对象群ID列表
//
// 示例值：['oc_a7bb9ca5efa68ab8b4fdd2e3b54fcerf']
func (builder *RuleBuilder) ToChatIds(toChatIds []string) *RuleBuilder {
	builder.toChatIds = toChatIds
	builder.toChatIdsFlag = true
	return builder
}

// 上级汇报对象，0表示第一级，依次类推，最大为5表示第六级
//
// 示例值：[0]
func (builder *RuleBuilder) ToLeaders(toLeaders []int) *RuleBuilder {
	builder.toLeaders = toLeaders
	builder.toLeadersFlag = true
	return builder
}

// 部门负责人汇报对象，0表示第一级，依次类推，最大为5表示第六级
//
// 示例值：[0]
func (builder *RuleBuilder) ToDepartmentOwners(toDepartmentOwners []int) *RuleBuilder {
	builder.toDepartmentOwners = toDepartmentOwners
	builder.toDepartmentOwnersFlag = true
	return builder
}

// 规则管理员用户ID列表
//
// 示例值：['ou_c04cebc780341ab22bd311ba6902sseb']
func (builder *RuleBuilder) ManagerUserIds(managerUserIds []string) *RuleBuilder {
	builder.managerUserIds = managerUserIds
	builder.managerUserIdsFlag = true
	return builder
}

// 抄送群ID列表
//
// 示例值：['oc_a7bb9ca5efa68ab8b4fdd2e3b54fffsf']
func (builder *RuleBuilder) CcChatIds(ccChatIds []string) *RuleBuilder {
	builder.ccChatIds = ccChatIds
	builder.ccChatIdsFlag = true
	return builder
}

func (builder *RuleBuilder) Build() *Rule {
	req := &Rule{}
	if builder.ruleIdFlag {
		req.RuleId = &builder.ruleId

	}
	if builder.nameFlag {
		req.Name = &builder.name

	}
	if builder.iconNameFlag {
		req.IconName = &builder.iconName

	}
	if builder.createdAtFlag {
		req.CreatedAt = &builder.createdAt

	}
	if builder.creatorUserIdFlag {
		req.CreatorUserId = &builder.creatorUserId

	}
	if builder.creatorUserNameFlag {
		req.CreatorUserName = &builder.creatorUserName

	}
	if builder.ownerUserIdFlag {
		req.OwnerUserId = &builder.ownerUserId

	}
	if builder.ownerUserNameFlag {
		req.OwnerUserName = &builder.ownerUserName

	}
	if builder.formSchemaFlag {
		req.FormSchema = builder.formSchema
	}
	if builder.isDeletedFlag {
		req.IsDeleted = &builder.isDeleted

	}
	if builder.needReportUserIdsFlag {
		req.NeedReportUserIds = builder.needReportUserIds
	}
	if builder.needReportDepartmentIdsFlag {
		req.NeedReportDepartmentIds = builder.needReportDepartmentIds
	}
	if builder.needReportChatIdsFlag {
		req.NeedReportChatIds = builder.needReportChatIds
	}
	if builder.ccUserIdsFlag {
		req.CcUserIds = builder.ccUserIds
	}
	if builder.ccDepartmentIdsFlag {
		req.CcDepartmentIds = builder.ccDepartmentIds
	}
	if builder.toUserIdsFlag {
		req.ToUserIds = builder.toUserIds
	}
	if builder.toChatIdsFlag {
		req.ToChatIds = builder.toChatIds
	}
	if builder.toLeadersFlag {
		req.ToLeaders = builder.toLeaders
	}
	if builder.toDepartmentOwnersFlag {
		req.ToDepartmentOwners = builder.toDepartmentOwners
	}
	if builder.managerUserIdsFlag {
		req.ManagerUserIds = builder.managerUserIds
	}
	if builder.ccChatIdsFlag {
		req.CcChatIds = builder.ccChatIds
	}
	return req
}

type Task struct {
	TaskId         *string        `json:"task_id,omitempty"`         // 汇报任务ID
	RuleName       *string        `json:"rule_name,omitempty"`       // 规则名称
	FromUserId     *string        `json:"from_user_id,omitempty"`    // 汇报用户ID
	FromUserName   *string        `json:"from_user_name,omitempty"`  // 汇报用户名称
	DepartmentName *string        `json:"department_name,omitempty"` // 汇报用户部门名称
	CommitTime     *int           `json:"commit_time,omitempty"`     // 提交时间时间戳
	FormContents   []*FormContent `json:"form_contents,omitempty"`   // 汇报表单内容
	RuleId         *string        `json:"rule_id,omitempty"`         // 汇报规则ID
}

type TaskBuilder struct {
	taskId             string // 汇报任务ID
	taskIdFlag         bool
	ruleName           string // 规则名称
	ruleNameFlag       bool
	fromUserId         string // 汇报用户ID
	fromUserIdFlag     bool
	fromUserName       string // 汇报用户名称
	fromUserNameFlag   bool
	departmentName     string // 汇报用户部门名称
	departmentNameFlag bool
	commitTime         int // 提交时间时间戳
	commitTimeFlag     bool
	formContents       []*FormContent // 汇报表单内容
	formContentsFlag   bool
	ruleId             string // 汇报规则ID
	ruleIdFlag         bool
}

func NewTaskBuilder() *TaskBuilder {
	builder := &TaskBuilder{}
	return builder
}

// 汇报任务ID
//
// 示例值：6968793659214921747
func (builder *TaskBuilder) TaskId(taskId string) *TaskBuilder {
	builder.taskId = taskId
	builder.taskIdFlag = true
	return builder
}

// 规则名称
//
// 示例值：工作月报
func (builder *TaskBuilder) RuleName(ruleName string) *TaskBuilder {
	builder.ruleName = ruleName
	builder.ruleNameFlag = true
	return builder
}

// 汇报用户ID
//
// 示例值：ou_c04cebc780341ab22bd311ba6902ffeb
func (builder *TaskBuilder) FromUserId(fromUserId string) *TaskBuilder {
	builder.fromUserId = fromUserId
	builder.fromUserIdFlag = true
	return builder
}

// 汇报用户名称
//
// 示例值：张三
func (builder *TaskBuilder) FromUserName(fromUserName string) *TaskBuilder {
	builder.fromUserName = fromUserName
	builder.fromUserNameFlag = true
	return builder
}

// 汇报用户部门名称
//
// 示例值：部门A
func (builder *TaskBuilder) DepartmentName(departmentName string) *TaskBuilder {
	builder.departmentName = departmentName
	builder.departmentNameFlag = true
	return builder
}

// 提交时间时间戳
//
// 示例值：1622548713
func (builder *TaskBuilder) CommitTime(commitTime int) *TaskBuilder {
	builder.commitTime = commitTime
	builder.commitTimeFlag = true
	return builder
}

// 汇报表单内容
//
// 示例值：
func (builder *TaskBuilder) FormContents(formContents []*FormContent) *TaskBuilder {
	builder.formContents = formContents
	builder.formContentsFlag = true
	return builder
}

// 汇报规则ID
//
// 示例值：6968793659214921747
func (builder *TaskBuilder) RuleId(ruleId string) *TaskBuilder {
	builder.ruleId = ruleId
	builder.ruleIdFlag = true
	return builder
}

func (builder *TaskBuilder) Build() *Task {
	req := &Task{}
	if builder.taskIdFlag {
		req.TaskId = &builder.taskId

	}
	if builder.ruleNameFlag {
		req.RuleName = &builder.ruleName

	}
	if builder.fromUserIdFlag {
		req.FromUserId = &builder.fromUserId

	}
	if builder.fromUserNameFlag {
		req.FromUserName = &builder.fromUserName

	}
	if builder.departmentNameFlag {
		req.DepartmentName = &builder.departmentName

	}
	if builder.commitTimeFlag {
		req.CommitTime = &builder.commitTime

	}
	if builder.formContentsFlag {
		req.FormContents = builder.formContents
	}
	if builder.ruleIdFlag {
		req.RuleId = &builder.ruleId

	}
	return req
}

type View struct {
}

type QueryRuleReqBuilder struct {
	apiReq *larkcore.ApiReq
}

func NewQueryRuleReqBuilder() *QueryRuleReqBuilder {
	builder := &QueryRuleReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 规则名称
//
// 示例值：工作月报
func (builder *QueryRuleReqBuilder) RuleName(ruleName string) *QueryRuleReqBuilder {
	builder.apiReq.QueryParams.Set("rule_name", fmt.Sprint(ruleName))
	return builder
}

// 是否包括已删除，默认未删除
//
// 示例值：0
func (builder *QueryRuleReqBuilder) IncludeDeleted(includeDeleted int) *QueryRuleReqBuilder {
	builder.apiReq.QueryParams.Set("include_deleted", fmt.Sprint(includeDeleted))
	return builder
}

// 此次调用中使用的用户ID的类型
//
// 示例值：
func (builder *QueryRuleReqBuilder) UserIdType(userIdType string) *QueryRuleReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

func (builder *QueryRuleReqBuilder) Build() *QueryRuleReq {
	req := &QueryRuleReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	return req
}

type QueryRuleReq struct {
	apiReq *larkcore.ApiReq
}

type QueryRuleRespData struct {
	Rules []*Rule `json:"rules,omitempty"` // 规则列表
}

type QueryRuleResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *QueryRuleRespData `json:"data"` // 业务数据
}

func (resp *QueryRuleResp) Success() bool {
	return resp.Code == 0
}

type RemoveRuleViewReqBodyBuilder struct {
	userIds     []string // 列表为空删除规则下全用户视图，列表不为空删除指定用户视图，大小限制200。
	userIdsFlag bool
}

func NewRemoveRuleViewReqBodyBuilder() *RemoveRuleViewReqBodyBuilder {
	builder := &RemoveRuleViewReqBodyBuilder{}
	return builder
}

// 列表为空删除规则下全用户视图，列表不为空删除指定用户视图，大小限制200。
//
//示例值：["ou_d6a5b5a55c77ca0b5b6c6ca0dd628c85","ou_d6a5b5a55c77ca0b5b6c6ca0dd628c55"]
func (builder *RemoveRuleViewReqBodyBuilder) UserIds(userIds []string) *RemoveRuleViewReqBodyBuilder {
	builder.userIds = userIds
	builder.userIdsFlag = true
	return builder
}

func (builder *RemoveRuleViewReqBodyBuilder) Build() *RemoveRuleViewReqBody {
	req := &RemoveRuleViewReqBody{}
	if builder.userIdsFlag {
		req.UserIds = builder.userIds
	}
	return req
}

type RemoveRuleViewPathReqBodyBuilder struct {
	userIds     []string
	userIdsFlag bool
}

func NewRemoveRuleViewPathReqBodyBuilder() *RemoveRuleViewPathReqBodyBuilder {
	builder := &RemoveRuleViewPathReqBodyBuilder{}
	return builder
}

// 列表为空删除规则下全用户视图，列表不为空删除指定用户视图，大小限制200。
//
// 示例值：["ou_d6a5b5a55c77ca0b5b6c6ca0dd628c85","ou_d6a5b5a55c77ca0b5b6c6ca0dd628c55"]
func (builder *RemoveRuleViewPathReqBodyBuilder) UserIds(userIds []string) *RemoveRuleViewPathReqBodyBuilder {
	builder.userIds = userIds
	builder.userIdsFlag = true
	return builder
}

func (builder *RemoveRuleViewPathReqBodyBuilder) Build() (*RemoveRuleViewReqBody, error) {
	req := &RemoveRuleViewReqBody{}
	if builder.userIdsFlag {
		req.UserIds = builder.userIds
	}
	return req, nil
}

type RemoveRuleViewReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *RemoveRuleViewReqBody
}

func NewRemoveRuleViewReqBuilder() *RemoveRuleViewReqBuilder {
	builder := &RemoveRuleViewReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 汇报规则ID
//
// 示例值：6894419345318182122
func (builder *RemoveRuleViewReqBuilder) RuleId(ruleId string) *RemoveRuleViewReqBuilder {
	builder.apiReq.PathParams.Set("rule_id", fmt.Sprint(ruleId))
	return builder
}

//
//
// 示例值：
func (builder *RemoveRuleViewReqBuilder) UserIdType(userIdType string) *RemoveRuleViewReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

// 移除规则看板
func (builder *RemoveRuleViewReqBuilder) Body(body *RemoveRuleViewReqBody) *RemoveRuleViewReqBuilder {
	builder.body = body
	return builder
}

func (builder *RemoveRuleViewReqBuilder) Build() *RemoveRuleViewReq {
	req := &RemoveRuleViewReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.body
	return req
}

type RemoveRuleViewReqBody struct {
	UserIds []string `json:"user_ids,omitempty"` // 列表为空删除规则下全用户视图，列表不为空删除指定用户视图，大小限制200。
}

type RemoveRuleViewReq struct {
	apiReq *larkcore.ApiReq
	Body   *RemoveRuleViewReqBody `body:""`
}

type RemoveRuleViewResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
}

func (resp *RemoveRuleViewResp) Success() bool {
	return resp.Code == 0
}

type QueryTaskReqBodyBuilder struct {
	commitStartTime     int // 提交开始时间时间戳
	commitStartTimeFlag bool
	commitEndTime       int // 提交结束时间时间戳
	commitEndTimeFlag   bool
	ruleId              string // 汇报规则ID
	ruleIdFlag          bool
	userId              string // 用户ID
	userIdFlag          bool
	pageToken           string // 分页标识符
	pageTokenFlag       bool
	pageSize            int // 单次分页返回的条数
	pageSizeFlag        bool
}

func NewQueryTaskReqBodyBuilder() *QueryTaskReqBodyBuilder {
	builder := &QueryTaskReqBodyBuilder{}
	return builder
}

// 提交开始时间时间戳
//
//示例值：1622427266
func (builder *QueryTaskReqBodyBuilder) CommitStartTime(commitStartTime int) *QueryTaskReqBodyBuilder {
	builder.commitStartTime = commitStartTime
	builder.commitStartTimeFlag = true
	return builder
}

// 提交结束时间时间戳
//
//示例值：1622427266
func (builder *QueryTaskReqBodyBuilder) CommitEndTime(commitEndTime int) *QueryTaskReqBodyBuilder {
	builder.commitEndTime = commitEndTime
	builder.commitEndTimeFlag = true
	return builder
}

// 汇报规则ID
//
//示例值：6894419345318182932
func (builder *QueryTaskReqBodyBuilder) RuleId(ruleId string) *QueryTaskReqBodyBuilder {
	builder.ruleId = ruleId
	builder.ruleIdFlag = true
	return builder
}

// 用户ID
//
//示例值：ou_133f0b6d0f097cf7d7ba00b38fffb110
func (builder *QueryTaskReqBodyBuilder) UserId(userId string) *QueryTaskReqBodyBuilder {
	builder.userId = userId
	builder.userIdFlag = true
	return builder
}

// 分页标识符
//
//示例值：6895699275733778451
func (builder *QueryTaskReqBodyBuilder) PageToken(pageToken string) *QueryTaskReqBodyBuilder {
	builder.pageToken = pageToken
	builder.pageTokenFlag = true
	return builder
}

// 单次分页返回的条数
//
//示例值：10
func (builder *QueryTaskReqBodyBuilder) PageSize(pageSize int) *QueryTaskReqBodyBuilder {
	builder.pageSize = pageSize
	builder.pageSizeFlag = true
	return builder
}

func (builder *QueryTaskReqBodyBuilder) Build() *QueryTaskReqBody {
	req := &QueryTaskReqBody{}
	if builder.commitStartTimeFlag {
		req.CommitStartTime = &builder.commitStartTime
	}
	if builder.commitEndTimeFlag {
		req.CommitEndTime = &builder.commitEndTime
	}
	if builder.ruleIdFlag {
		req.RuleId = &builder.ruleId
	}
	if builder.userIdFlag {
		req.UserId = &builder.userId
	}
	if builder.pageTokenFlag {
		req.PageToken = &builder.pageToken
	}
	if builder.pageSizeFlag {
		req.PageSize = &builder.pageSize
	}
	return req
}

type QueryTaskPathReqBodyBuilder struct {
	commitStartTime     int
	commitStartTimeFlag bool
	commitEndTime       int
	commitEndTimeFlag   bool
	ruleId              string
	ruleIdFlag          bool
	userId              string
	userIdFlag          bool
	pageToken           string
	pageTokenFlag       bool
	pageSize            int
	pageSizeFlag        bool
}

func NewQueryTaskPathReqBodyBuilder() *QueryTaskPathReqBodyBuilder {
	builder := &QueryTaskPathReqBodyBuilder{}
	return builder
}

// 提交开始时间时间戳
//
// 示例值：1622427266
func (builder *QueryTaskPathReqBodyBuilder) CommitStartTime(commitStartTime int) *QueryTaskPathReqBodyBuilder {
	builder.commitStartTime = commitStartTime
	builder.commitStartTimeFlag = true
	return builder
}

// 提交结束时间时间戳
//
// 示例值：1622427266
func (builder *QueryTaskPathReqBodyBuilder) CommitEndTime(commitEndTime int) *QueryTaskPathReqBodyBuilder {
	builder.commitEndTime = commitEndTime
	builder.commitEndTimeFlag = true
	return builder
}

// 汇报规则ID
//
// 示例值：6894419345318182932
func (builder *QueryTaskPathReqBodyBuilder) RuleId(ruleId string) *QueryTaskPathReqBodyBuilder {
	builder.ruleId = ruleId
	builder.ruleIdFlag = true
	return builder
}

// 用户ID
//
// 示例值：ou_133f0b6d0f097cf7d7ba00b38fffb110
func (builder *QueryTaskPathReqBodyBuilder) UserId(userId string) *QueryTaskPathReqBodyBuilder {
	builder.userId = userId
	builder.userIdFlag = true
	return builder
}

// 分页标识符
//
// 示例值：6895699275733778451
func (builder *QueryTaskPathReqBodyBuilder) PageToken(pageToken string) *QueryTaskPathReqBodyBuilder {
	builder.pageToken = pageToken
	builder.pageTokenFlag = true
	return builder
}

// 单次分页返回的条数
//
// 示例值：10
func (builder *QueryTaskPathReqBodyBuilder) PageSize(pageSize int) *QueryTaskPathReqBodyBuilder {
	builder.pageSize = pageSize
	builder.pageSizeFlag = true
	return builder
}

func (builder *QueryTaskPathReqBodyBuilder) Build() (*QueryTaskReqBody, error) {
	req := &QueryTaskReqBody{}
	if builder.commitStartTimeFlag {
		req.CommitStartTime = &builder.commitStartTime
	}
	if builder.commitEndTimeFlag {
		req.CommitEndTime = &builder.commitEndTime
	}
	if builder.ruleIdFlag {
		req.RuleId = &builder.ruleId
	}
	if builder.userIdFlag {
		req.UserId = &builder.userId
	}
	if builder.pageTokenFlag {
		req.PageToken = &builder.pageToken
	}
	if builder.pageSizeFlag {
		req.PageSize = &builder.pageSize
	}
	return req, nil
}

type QueryTaskReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *QueryTaskReqBody
}

func NewQueryTaskReqBuilder() *QueryTaskReqBuilder {
	builder := &QueryTaskReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 此次调用中使用的用户ID的类型
//
// 示例值：
func (builder *QueryTaskReqBuilder) UserIdType(userIdType string) *QueryTaskReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

// 任务查询
func (builder *QueryTaskReqBuilder) Body(body *QueryTaskReqBody) *QueryTaskReqBuilder {
	builder.body = body
	return builder
}

func (builder *QueryTaskReqBuilder) Build() *QueryTaskReq {
	req := &QueryTaskReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.body
	return req
}

type QueryTaskReqBody struct {
	CommitStartTime *int    `json:"commit_start_time,omitempty"` // 提交开始时间时间戳
	CommitEndTime   *int    `json:"commit_end_time,omitempty"`   // 提交结束时间时间戳
	RuleId          *string `json:"rule_id,omitempty"`           // 汇报规则ID
	UserId          *string `json:"user_id,omitempty"`           // 用户ID
	PageToken       *string `json:"page_token,omitempty"`        // 分页标识符
	PageSize        *int    `json:"page_size,omitempty"`         // 单次分页返回的条数
}

type QueryTaskReq struct {
	apiReq *larkcore.ApiReq
	Body   *QueryTaskReqBody `body:""`
}

type QueryTaskRespData struct {
	Items     []*Task `json:"items,omitempty"`      // 任务列表
	HasMore   *bool   `json:"has_more,omitempty"`   // 是否有下一页数据
	PageToken *string `json:"page_token,omitempty"` // 下一页分页的token
}

type QueryTaskResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *QueryTaskRespData `json:"data"` // 业务数据
}

func (resp *QueryTaskResp) Success() bool {
	return resp.Code == 0
}
