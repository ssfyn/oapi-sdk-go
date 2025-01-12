// Package verification code generated by oapi sdk gen
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

package larkverification

import (
	"github.com/larksuite/oapi-sdk-go/v3/core"
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

type Verification struct {
	Name            *string `json:"name,omitempty"`             // 企业主体名称
	HasVerification *bool   `json:"has_verification,omitempty"` // 企业是否完成认证； true 表示已经完成认证，false 表示未认证
}

type VerificationBuilder struct {
	name                string // 企业主体名称
	nameFlag            bool
	hasVerification     bool // 企业是否完成认证； true 表示已经完成认证，false 表示未认证
	hasVerificationFlag bool
}

func NewVerificationBuilder() *VerificationBuilder {
	builder := &VerificationBuilder{}
	return builder
}

// 企业主体名称
//
// 示例值：
func (builder *VerificationBuilder) Name(name string) *VerificationBuilder {
	builder.name = name
	builder.nameFlag = true
	return builder
}

// 企业是否完成认证； true 表示已经完成认证，false 表示未认证
//
// 示例值：
func (builder *VerificationBuilder) HasVerification(hasVerification bool) *VerificationBuilder {
	builder.hasVerification = hasVerification
	builder.hasVerificationFlag = true
	return builder
}

func (builder *VerificationBuilder) Build() *Verification {
	req := &Verification{}
	if builder.nameFlag {
		req.Name = &builder.name

	}
	if builder.hasVerificationFlag {
		req.HasVerification = &builder.hasVerification

	}
	return req
}

type GetVerificationRespData struct {
	Verification *Verification `json:"verification,omitempty"` // 认证信息
}

type GetVerificationResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *GetVerificationRespData `json:"data"` // 业务数据
}

func (resp *GetVerificationResp) Success() bool {
	return resp.Code == 0
}
