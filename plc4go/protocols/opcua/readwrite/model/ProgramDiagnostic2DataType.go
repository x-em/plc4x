/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ProgramDiagnostic2DataType is the corresponding interface of ProgramDiagnostic2DataType
type ProgramDiagnostic2DataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetCreateSessionId returns CreateSessionId (property field)
	GetCreateSessionId() NodeId
	// GetCreateClientName returns CreateClientName (property field)
	GetCreateClientName() PascalString
	// GetInvocationCreationTime returns InvocationCreationTime (property field)
	GetInvocationCreationTime() int64
	// GetLastTransitionTime returns LastTransitionTime (property field)
	GetLastTransitionTime() int64
	// GetLastMethodCall returns LastMethodCall (property field)
	GetLastMethodCall() PascalString
	// GetLastMethodSessionId returns LastMethodSessionId (property field)
	GetLastMethodSessionId() NodeId
	// GetLastMethodInputArguments returns LastMethodInputArguments (property field)
	GetLastMethodInputArguments() []Argument
	// GetLastMethodOutputArguments returns LastMethodOutputArguments (property field)
	GetLastMethodOutputArguments() []Argument
	// GetLastMethodInputValues returns LastMethodInputValues (property field)
	GetLastMethodInputValues() []Variant
	// GetLastMethodOutputValues returns LastMethodOutputValues (property field)
	GetLastMethodOutputValues() []Variant
	// GetLastMethodCallTime returns LastMethodCallTime (property field)
	GetLastMethodCallTime() int64
	// GetLastMethodReturnStatus returns LastMethodReturnStatus (property field)
	GetLastMethodReturnStatus() StatusCode
	// IsProgramDiagnostic2DataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsProgramDiagnostic2DataType()
	// CreateBuilder creates a ProgramDiagnostic2DataTypeBuilder
	CreateProgramDiagnostic2DataTypeBuilder() ProgramDiagnostic2DataTypeBuilder
}

// _ProgramDiagnostic2DataType is the data-structure of this message
type _ProgramDiagnostic2DataType struct {
	ExtensionObjectDefinitionContract
	CreateSessionId           NodeId
	CreateClientName          PascalString
	InvocationCreationTime    int64
	LastTransitionTime        int64
	LastMethodCall            PascalString
	LastMethodSessionId       NodeId
	LastMethodInputArguments  []Argument
	LastMethodOutputArguments []Argument
	LastMethodInputValues     []Variant
	LastMethodOutputValues    []Variant
	LastMethodCallTime        int64
	LastMethodReturnStatus    StatusCode
}

var _ ProgramDiagnostic2DataType = (*_ProgramDiagnostic2DataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ProgramDiagnostic2DataType)(nil)

// NewProgramDiagnostic2DataType factory function for _ProgramDiagnostic2DataType
func NewProgramDiagnostic2DataType(createSessionId NodeId, createClientName PascalString, invocationCreationTime int64, lastTransitionTime int64, lastMethodCall PascalString, lastMethodSessionId NodeId, lastMethodInputArguments []Argument, lastMethodOutputArguments []Argument, lastMethodInputValues []Variant, lastMethodOutputValues []Variant, lastMethodCallTime int64, lastMethodReturnStatus StatusCode) *_ProgramDiagnostic2DataType {
	if createSessionId == nil {
		panic("createSessionId of type NodeId for ProgramDiagnostic2DataType must not be nil")
	}
	if createClientName == nil {
		panic("createClientName of type PascalString for ProgramDiagnostic2DataType must not be nil")
	}
	if lastMethodCall == nil {
		panic("lastMethodCall of type PascalString for ProgramDiagnostic2DataType must not be nil")
	}
	if lastMethodSessionId == nil {
		panic("lastMethodSessionId of type NodeId for ProgramDiagnostic2DataType must not be nil")
	}
	if lastMethodReturnStatus == nil {
		panic("lastMethodReturnStatus of type StatusCode for ProgramDiagnostic2DataType must not be nil")
	}
	_result := &_ProgramDiagnostic2DataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		CreateSessionId:                   createSessionId,
		CreateClientName:                  createClientName,
		InvocationCreationTime:            invocationCreationTime,
		LastTransitionTime:                lastTransitionTime,
		LastMethodCall:                    lastMethodCall,
		LastMethodSessionId:               lastMethodSessionId,
		LastMethodInputArguments:          lastMethodInputArguments,
		LastMethodOutputArguments:         lastMethodOutputArguments,
		LastMethodInputValues:             lastMethodInputValues,
		LastMethodOutputValues:            lastMethodOutputValues,
		LastMethodCallTime:                lastMethodCallTime,
		LastMethodReturnStatus:            lastMethodReturnStatus,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ProgramDiagnostic2DataTypeBuilder is a builder for ProgramDiagnostic2DataType
type ProgramDiagnostic2DataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(createSessionId NodeId, createClientName PascalString, invocationCreationTime int64, lastTransitionTime int64, lastMethodCall PascalString, lastMethodSessionId NodeId, lastMethodInputArguments []Argument, lastMethodOutputArguments []Argument, lastMethodInputValues []Variant, lastMethodOutputValues []Variant, lastMethodCallTime int64, lastMethodReturnStatus StatusCode) ProgramDiagnostic2DataTypeBuilder
	// WithCreateSessionId adds CreateSessionId (property field)
	WithCreateSessionId(NodeId) ProgramDiagnostic2DataTypeBuilder
	// WithCreateSessionIdBuilder adds CreateSessionId (property field) which is build by the builder
	WithCreateSessionIdBuilder(func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithCreateClientName adds CreateClientName (property field)
	WithCreateClientName(PascalString) ProgramDiagnostic2DataTypeBuilder
	// WithCreateClientNameBuilder adds CreateClientName (property field) which is build by the builder
	WithCreateClientNameBuilder(func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithInvocationCreationTime adds InvocationCreationTime (property field)
	WithInvocationCreationTime(int64) ProgramDiagnostic2DataTypeBuilder
	// WithLastTransitionTime adds LastTransitionTime (property field)
	WithLastTransitionTime(int64) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodCall adds LastMethodCall (property field)
	WithLastMethodCall(PascalString) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodCallBuilder adds LastMethodCall (property field) which is build by the builder
	WithLastMethodCallBuilder(func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodSessionId adds LastMethodSessionId (property field)
	WithLastMethodSessionId(NodeId) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodSessionIdBuilder adds LastMethodSessionId (property field) which is build by the builder
	WithLastMethodSessionIdBuilder(func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodInputArguments adds LastMethodInputArguments (property field)
	WithLastMethodInputArguments(...Argument) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodOutputArguments adds LastMethodOutputArguments (property field)
	WithLastMethodOutputArguments(...Argument) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodInputValues adds LastMethodInputValues (property field)
	WithLastMethodInputValues(...Variant) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodOutputValues adds LastMethodOutputValues (property field)
	WithLastMethodOutputValues(...Variant) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodCallTime adds LastMethodCallTime (property field)
	WithLastMethodCallTime(int64) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodReturnStatus adds LastMethodReturnStatus (property field)
	WithLastMethodReturnStatus(StatusCode) ProgramDiagnostic2DataTypeBuilder
	// WithLastMethodReturnStatusBuilder adds LastMethodReturnStatus (property field) which is build by the builder
	WithLastMethodReturnStatusBuilder(func(StatusCodeBuilder) StatusCodeBuilder) ProgramDiagnostic2DataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ProgramDiagnostic2DataType or returns an error if something is wrong
	Build() (ProgramDiagnostic2DataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ProgramDiagnostic2DataType
}

// NewProgramDiagnostic2DataTypeBuilder() creates a ProgramDiagnostic2DataTypeBuilder
func NewProgramDiagnostic2DataTypeBuilder() ProgramDiagnostic2DataTypeBuilder {
	return &_ProgramDiagnostic2DataTypeBuilder{_ProgramDiagnostic2DataType: new(_ProgramDiagnostic2DataType)}
}

type _ProgramDiagnostic2DataTypeBuilder struct {
	*_ProgramDiagnostic2DataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ProgramDiagnostic2DataTypeBuilder) = (*_ProgramDiagnostic2DataTypeBuilder)(nil)

func (b *_ProgramDiagnostic2DataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ProgramDiagnostic2DataType
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithMandatoryFields(createSessionId NodeId, createClientName PascalString, invocationCreationTime int64, lastTransitionTime int64, lastMethodCall PascalString, lastMethodSessionId NodeId, lastMethodInputArguments []Argument, lastMethodOutputArguments []Argument, lastMethodInputValues []Variant, lastMethodOutputValues []Variant, lastMethodCallTime int64, lastMethodReturnStatus StatusCode) ProgramDiagnostic2DataTypeBuilder {
	return b.WithCreateSessionId(createSessionId).WithCreateClientName(createClientName).WithInvocationCreationTime(invocationCreationTime).WithLastTransitionTime(lastTransitionTime).WithLastMethodCall(lastMethodCall).WithLastMethodSessionId(lastMethodSessionId).WithLastMethodInputArguments(lastMethodInputArguments...).WithLastMethodOutputArguments(lastMethodOutputArguments...).WithLastMethodInputValues(lastMethodInputValues...).WithLastMethodOutputValues(lastMethodOutputValues...).WithLastMethodCallTime(lastMethodCallTime).WithLastMethodReturnStatus(lastMethodReturnStatus)
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithCreateSessionId(createSessionId NodeId) ProgramDiagnostic2DataTypeBuilder {
	b.CreateSessionId = createSessionId
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithCreateSessionIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(b.CreateSessionId.CreateNodeIdBuilder())
	var err error
	b.CreateSessionId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithCreateClientName(createClientName PascalString) ProgramDiagnostic2DataTypeBuilder {
	b.CreateClientName = createClientName
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithCreateClientNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(b.CreateClientName.CreatePascalStringBuilder())
	var err error
	b.CreateClientName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithInvocationCreationTime(invocationCreationTime int64) ProgramDiagnostic2DataTypeBuilder {
	b.InvocationCreationTime = invocationCreationTime
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastTransitionTime(lastTransitionTime int64) ProgramDiagnostic2DataTypeBuilder {
	b.LastTransitionTime = lastTransitionTime
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodCall(lastMethodCall PascalString) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodCall = lastMethodCall
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodCallBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(b.LastMethodCall.CreatePascalStringBuilder())
	var err error
	b.LastMethodCall, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodSessionId(lastMethodSessionId NodeId) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodSessionId = lastMethodSessionId
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodSessionIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(b.LastMethodSessionId.CreateNodeIdBuilder())
	var err error
	b.LastMethodSessionId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodInputArguments(lastMethodInputArguments ...Argument) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodInputArguments = lastMethodInputArguments
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodOutputArguments(lastMethodOutputArguments ...Argument) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodOutputArguments = lastMethodOutputArguments
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodInputValues(lastMethodInputValues ...Variant) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodInputValues = lastMethodInputValues
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodOutputValues(lastMethodOutputValues ...Variant) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodOutputValues = lastMethodOutputValues
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodCallTime(lastMethodCallTime int64) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodCallTime = lastMethodCallTime
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodReturnStatus(lastMethodReturnStatus StatusCode) ProgramDiagnostic2DataTypeBuilder {
	b.LastMethodReturnStatus = lastMethodReturnStatus
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) WithLastMethodReturnStatusBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) ProgramDiagnostic2DataTypeBuilder {
	builder := builderSupplier(b.LastMethodReturnStatus.CreateStatusCodeBuilder())
	var err error
	b.LastMethodReturnStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_ProgramDiagnostic2DataTypeBuilder) Build() (ProgramDiagnostic2DataType, error) {
	if b.CreateSessionId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'createSessionId' not set"))
	}
	if b.CreateClientName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'createClientName' not set"))
	}
	if b.LastMethodCall == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastMethodCall' not set"))
	}
	if b.LastMethodSessionId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastMethodSessionId' not set"))
	}
	if b.LastMethodReturnStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastMethodReturnStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ProgramDiagnostic2DataType.deepCopy(), nil
}

func (b *_ProgramDiagnostic2DataTypeBuilder) MustBuild() ProgramDiagnostic2DataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ProgramDiagnostic2DataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ProgramDiagnostic2DataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ProgramDiagnostic2DataTypeBuilder) DeepCopy() any {
	_copy := b.CreateProgramDiagnostic2DataTypeBuilder().(*_ProgramDiagnostic2DataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateProgramDiagnostic2DataTypeBuilder creates a ProgramDiagnostic2DataTypeBuilder
func (b *_ProgramDiagnostic2DataType) CreateProgramDiagnostic2DataTypeBuilder() ProgramDiagnostic2DataTypeBuilder {
	if b == nil {
		return NewProgramDiagnostic2DataTypeBuilder()
	}
	return &_ProgramDiagnostic2DataTypeBuilder{_ProgramDiagnostic2DataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ProgramDiagnostic2DataType) GetExtensionId() int32 {
	return int32(24035)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ProgramDiagnostic2DataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ProgramDiagnostic2DataType) GetCreateSessionId() NodeId {
	return m.CreateSessionId
}

func (m *_ProgramDiagnostic2DataType) GetCreateClientName() PascalString {
	return m.CreateClientName
}

func (m *_ProgramDiagnostic2DataType) GetInvocationCreationTime() int64 {
	return m.InvocationCreationTime
}

func (m *_ProgramDiagnostic2DataType) GetLastTransitionTime() int64 {
	return m.LastTransitionTime
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodCall() PascalString {
	return m.LastMethodCall
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodSessionId() NodeId {
	return m.LastMethodSessionId
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodInputArguments() []Argument {
	return m.LastMethodInputArguments
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodOutputArguments() []Argument {
	return m.LastMethodOutputArguments
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodInputValues() []Variant {
	return m.LastMethodInputValues
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodOutputValues() []Variant {
	return m.LastMethodOutputValues
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodCallTime() int64 {
	return m.LastMethodCallTime
}

func (m *_ProgramDiagnostic2DataType) GetLastMethodReturnStatus() StatusCode {
	return m.LastMethodReturnStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastProgramDiagnostic2DataType(structType any) ProgramDiagnostic2DataType {
	if casted, ok := structType.(ProgramDiagnostic2DataType); ok {
		return casted
	}
	if casted, ok := structType.(*ProgramDiagnostic2DataType); ok {
		return *casted
	}
	return nil
}

func (m *_ProgramDiagnostic2DataType) GetTypeName() string {
	return "ProgramDiagnostic2DataType"
}

func (m *_ProgramDiagnostic2DataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (createSessionId)
	lengthInBits += m.CreateSessionId.GetLengthInBits(ctx)

	// Simple field (createClientName)
	lengthInBits += m.CreateClientName.GetLengthInBits(ctx)

	// Simple field (invocationCreationTime)
	lengthInBits += 64

	// Simple field (lastTransitionTime)
	lengthInBits += 64

	// Simple field (lastMethodCall)
	lengthInBits += m.LastMethodCall.GetLengthInBits(ctx)

	// Simple field (lastMethodSessionId)
	lengthInBits += m.LastMethodSessionId.GetLengthInBits(ctx)

	// Implicit Field (noOfLastMethodInputArguments)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodInputArguments) > 0 {
		for _curItem, element := range m.LastMethodInputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodInputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfLastMethodOutputArguments)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodOutputArguments) > 0 {
		for _curItem, element := range m.LastMethodOutputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodOutputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfLastMethodInputValues)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodInputValues) > 0 {
		for _curItem, element := range m.LastMethodInputValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodInputValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfLastMethodOutputValues)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodOutputValues) > 0 {
		for _curItem, element := range m.LastMethodOutputValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodOutputValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (lastMethodCallTime)
	lengthInBits += 64

	// Simple field (lastMethodReturnStatus)
	lengthInBits += m.LastMethodReturnStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ProgramDiagnostic2DataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ProgramDiagnostic2DataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__programDiagnostic2DataType ProgramDiagnostic2DataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ProgramDiagnostic2DataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ProgramDiagnostic2DataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	createSessionId, err := ReadSimpleField[NodeId](ctx, "createSessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'createSessionId' field"))
	}
	m.CreateSessionId = createSessionId

	createClientName, err := ReadSimpleField[PascalString](ctx, "createClientName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'createClientName' field"))
	}
	m.CreateClientName = createClientName

	invocationCreationTime, err := ReadSimpleField(ctx, "invocationCreationTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'invocationCreationTime' field"))
	}
	m.InvocationCreationTime = invocationCreationTime

	lastTransitionTime, err := ReadSimpleField(ctx, "lastTransitionTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastTransitionTime' field"))
	}
	m.LastTransitionTime = lastTransitionTime

	lastMethodCall, err := ReadSimpleField[PascalString](ctx, "lastMethodCall", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodCall' field"))
	}
	m.LastMethodCall = lastMethodCall

	lastMethodSessionId, err := ReadSimpleField[NodeId](ctx, "lastMethodSessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodSessionId' field"))
	}
	m.LastMethodSessionId = lastMethodSessionId

	noOfLastMethodInputArguments, err := ReadImplicitField[int32](ctx, "noOfLastMethodInputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodInputArguments' field"))
	}
	_ = noOfLastMethodInputArguments

	lastMethodInputArguments, err := ReadCountArrayField[Argument](ctx, "lastMethodInputArguments", ReadComplex[Argument](ExtensionObjectDefinitionParseWithBufferProducer[Argument]((int32)(int32(298))), readBuffer), uint64(noOfLastMethodInputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodInputArguments' field"))
	}
	m.LastMethodInputArguments = lastMethodInputArguments

	noOfLastMethodOutputArguments, err := ReadImplicitField[int32](ctx, "noOfLastMethodOutputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodOutputArguments' field"))
	}
	_ = noOfLastMethodOutputArguments

	lastMethodOutputArguments, err := ReadCountArrayField[Argument](ctx, "lastMethodOutputArguments", ReadComplex[Argument](ExtensionObjectDefinitionParseWithBufferProducer[Argument]((int32)(int32(298))), readBuffer), uint64(noOfLastMethodOutputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodOutputArguments' field"))
	}
	m.LastMethodOutputArguments = lastMethodOutputArguments

	noOfLastMethodInputValues, err := ReadImplicitField[int32](ctx, "noOfLastMethodInputValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodInputValues' field"))
	}
	_ = noOfLastMethodInputValues

	lastMethodInputValues, err := ReadCountArrayField[Variant](ctx, "lastMethodInputValues", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfLastMethodInputValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodInputValues' field"))
	}
	m.LastMethodInputValues = lastMethodInputValues

	noOfLastMethodOutputValues, err := ReadImplicitField[int32](ctx, "noOfLastMethodOutputValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodOutputValues' field"))
	}
	_ = noOfLastMethodOutputValues

	lastMethodOutputValues, err := ReadCountArrayField[Variant](ctx, "lastMethodOutputValues", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfLastMethodOutputValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodOutputValues' field"))
	}
	m.LastMethodOutputValues = lastMethodOutputValues

	lastMethodCallTime, err := ReadSimpleField(ctx, "lastMethodCallTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodCallTime' field"))
	}
	m.LastMethodCallTime = lastMethodCallTime

	lastMethodReturnStatus, err := ReadSimpleField[StatusCode](ctx, "lastMethodReturnStatus", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodReturnStatus' field"))
	}
	m.LastMethodReturnStatus = lastMethodReturnStatus

	if closeErr := readBuffer.CloseContext("ProgramDiagnostic2DataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ProgramDiagnostic2DataType")
	}

	return m, nil
}

func (m *_ProgramDiagnostic2DataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ProgramDiagnostic2DataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ProgramDiagnostic2DataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ProgramDiagnostic2DataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "createSessionId", m.GetCreateSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'createSessionId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "createClientName", m.GetCreateClientName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'createClientName' field")
		}

		if err := WriteSimpleField[int64](ctx, "invocationCreationTime", m.GetInvocationCreationTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'invocationCreationTime' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastTransitionTime", m.GetLastTransitionTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastTransitionTime' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "lastMethodCall", m.GetLastMethodCall(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodCall' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "lastMethodSessionId", m.GetLastMethodSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodSessionId' field")
		}
		noOfLastMethodInputArguments := int32(utils.InlineIf(bool((m.GetLastMethodInputArguments()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetLastMethodInputArguments()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfLastMethodInputArguments", noOfLastMethodInputArguments, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodInputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodInputArguments", m.GetLastMethodInputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodInputArguments' field")
		}
		noOfLastMethodOutputArguments := int32(utils.InlineIf(bool((m.GetLastMethodOutputArguments()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetLastMethodOutputArguments()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfLastMethodOutputArguments", noOfLastMethodOutputArguments, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodOutputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodOutputArguments", m.GetLastMethodOutputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodOutputArguments' field")
		}
		noOfLastMethodInputValues := int32(utils.InlineIf(bool((m.GetLastMethodInputValues()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetLastMethodInputValues()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfLastMethodInputValues", noOfLastMethodInputValues, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodInputValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodInputValues", m.GetLastMethodInputValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodInputValues' field")
		}
		noOfLastMethodOutputValues := int32(utils.InlineIf(bool((m.GetLastMethodOutputValues()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetLastMethodOutputValues()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfLastMethodOutputValues", noOfLastMethodOutputValues, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodOutputValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodOutputValues", m.GetLastMethodOutputValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodOutputValues' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastMethodCallTime", m.GetLastMethodCallTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodCallTime' field")
		}

		if err := WriteSimpleField[StatusCode](ctx, "lastMethodReturnStatus", m.GetLastMethodReturnStatus(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodReturnStatus' field")
		}

		if popErr := writeBuffer.PopContext("ProgramDiagnostic2DataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ProgramDiagnostic2DataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ProgramDiagnostic2DataType) IsProgramDiagnostic2DataType() {}

func (m *_ProgramDiagnostic2DataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ProgramDiagnostic2DataType) deepCopy() *_ProgramDiagnostic2DataType {
	if m == nil {
		return nil
	}
	_ProgramDiagnostic2DataTypeCopy := &_ProgramDiagnostic2DataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.CreateSessionId),
		utils.DeepCopy[PascalString](m.CreateClientName),
		m.InvocationCreationTime,
		m.LastTransitionTime,
		utils.DeepCopy[PascalString](m.LastMethodCall),
		utils.DeepCopy[NodeId](m.LastMethodSessionId),
		utils.DeepCopySlice[Argument, Argument](m.LastMethodInputArguments),
		utils.DeepCopySlice[Argument, Argument](m.LastMethodOutputArguments),
		utils.DeepCopySlice[Variant, Variant](m.LastMethodInputValues),
		utils.DeepCopySlice[Variant, Variant](m.LastMethodOutputValues),
		m.LastMethodCallTime,
		utils.DeepCopy[StatusCode](m.LastMethodReturnStatus),
	}
	_ProgramDiagnostic2DataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ProgramDiagnostic2DataTypeCopy
}

func (m *_ProgramDiagnostic2DataType) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}