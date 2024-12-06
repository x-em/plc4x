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

// IdentifyReplyCommandLogicalAssignment is the corresponding interface of IdentifyReplyCommandLogicalAssignment
type IdentifyReplyCommandLogicalAssignment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetLogicAssigment returns LogicAssigment (property field)
	GetLogicAssigment() []LogicAssignment
	// IsIdentifyReplyCommandLogicalAssignment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandLogicalAssignment()
	// CreateBuilder creates a IdentifyReplyCommandLogicalAssignmentBuilder
	CreateIdentifyReplyCommandLogicalAssignmentBuilder() IdentifyReplyCommandLogicalAssignmentBuilder
}

// _IdentifyReplyCommandLogicalAssignment is the data-structure of this message
type _IdentifyReplyCommandLogicalAssignment struct {
	IdentifyReplyCommandContract
	LogicAssigment []LogicAssignment
}

var _ IdentifyReplyCommandLogicalAssignment = (*_IdentifyReplyCommandLogicalAssignment)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandLogicalAssignment)(nil)

// NewIdentifyReplyCommandLogicalAssignment factory function for _IdentifyReplyCommandLogicalAssignment
func NewIdentifyReplyCommandLogicalAssignment(logicAssigment []LogicAssignment, numBytes uint8) *_IdentifyReplyCommandLogicalAssignment {
	_result := &_IdentifyReplyCommandLogicalAssignment{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		LogicAssigment:               logicAssigment,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandLogicalAssignmentBuilder is a builder for IdentifyReplyCommandLogicalAssignment
type IdentifyReplyCommandLogicalAssignmentBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(logicAssigment []LogicAssignment) IdentifyReplyCommandLogicalAssignmentBuilder
	// WithLogicAssigment adds LogicAssigment (property field)
	WithLogicAssigment(...LogicAssignment) IdentifyReplyCommandLogicalAssignmentBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() IdentifyReplyCommandBuilder
	// Build builds the IdentifyReplyCommandLogicalAssignment or returns an error if something is wrong
	Build() (IdentifyReplyCommandLogicalAssignment, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandLogicalAssignment
}

// NewIdentifyReplyCommandLogicalAssignmentBuilder() creates a IdentifyReplyCommandLogicalAssignmentBuilder
func NewIdentifyReplyCommandLogicalAssignmentBuilder() IdentifyReplyCommandLogicalAssignmentBuilder {
	return &_IdentifyReplyCommandLogicalAssignmentBuilder{_IdentifyReplyCommandLogicalAssignment: new(_IdentifyReplyCommandLogicalAssignment)}
}

type _IdentifyReplyCommandLogicalAssignmentBuilder struct {
	*_IdentifyReplyCommandLogicalAssignment

	parentBuilder *_IdentifyReplyCommandBuilder

	err *utils.MultiError
}

var _ (IdentifyReplyCommandLogicalAssignmentBuilder) = (*_IdentifyReplyCommandLogicalAssignmentBuilder)(nil)

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) setParent(contract IdentifyReplyCommandContract) {
	b.IdentifyReplyCommandContract = contract
	contract.(*_IdentifyReplyCommand)._SubType = b._IdentifyReplyCommandLogicalAssignment
}

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) WithMandatoryFields(logicAssigment []LogicAssignment) IdentifyReplyCommandLogicalAssignmentBuilder {
	return b.WithLogicAssigment(logicAssigment...)
}

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) WithLogicAssigment(logicAssigment ...LogicAssignment) IdentifyReplyCommandLogicalAssignmentBuilder {
	b.LogicAssigment = logicAssigment
	return b
}

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) Build() (IdentifyReplyCommandLogicalAssignment, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandLogicalAssignment.deepCopy(), nil
}

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) MustBuild() IdentifyReplyCommandLogicalAssignment {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) Done() IdentifyReplyCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewIdentifyReplyCommandBuilder().(*_IdentifyReplyCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) buildForIdentifyReplyCommand() (IdentifyReplyCommand, error) {
	return b.Build()
}

func (b *_IdentifyReplyCommandLogicalAssignmentBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandLogicalAssignmentBuilder().(*_IdentifyReplyCommandLogicalAssignmentBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandLogicalAssignmentBuilder creates a IdentifyReplyCommandLogicalAssignmentBuilder
func (b *_IdentifyReplyCommandLogicalAssignment) CreateIdentifyReplyCommandLogicalAssignmentBuilder() IdentifyReplyCommandLogicalAssignmentBuilder {
	if b == nil {
		return NewIdentifyReplyCommandLogicalAssignmentBuilder()
	}
	return &_IdentifyReplyCommandLogicalAssignmentBuilder{_IdentifyReplyCommandLogicalAssignment: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) GetAttribute() Attribute {
	return Attribute_LogicalAssignment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) GetLogicAssigment() []LogicAssignment {
	return m.LogicAssigment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandLogicalAssignment(structType any) IdentifyReplyCommandLogicalAssignment {
	if casted, ok := structType.(IdentifyReplyCommandLogicalAssignment); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandLogicalAssignment); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetTypeName() string {
	return "IdentifyReplyCommandLogicalAssignment"
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.LogicAssigment) > 0 {
		for _curItem, element := range m.LogicAssigment {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LogicAssigment), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandLogicalAssignment) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandLogicalAssignment IdentifyReplyCommandLogicalAssignment, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandLogicalAssignment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandLogicalAssignment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logicAssigment, err := ReadCountArrayField[LogicAssignment](ctx, "logicAssigment", ReadComplex[LogicAssignment](LogicAssignmentParseWithBuffer, readBuffer), uint64(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicAssigment' field"))
	}
	m.LogicAssigment = logicAssigment

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandLogicalAssignment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandLogicalAssignment")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandLogicalAssignment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandLogicalAssignment")
		}

		if err := WriteComplexTypeArrayField(ctx, "logicAssigment", m.GetLogicAssigment(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'logicAssigment' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandLogicalAssignment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandLogicalAssignment")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandLogicalAssignment) IsIdentifyReplyCommandLogicalAssignment() {}

func (m *_IdentifyReplyCommandLogicalAssignment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandLogicalAssignment) deepCopy() *_IdentifyReplyCommandLogicalAssignment {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandLogicalAssignmentCopy := &_IdentifyReplyCommandLogicalAssignment{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		utils.DeepCopySlice[LogicAssignment, LogicAssignment](m.LogicAssigment),
	}
	_IdentifyReplyCommandLogicalAssignmentCopy.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandLogicalAssignmentCopy
}

func (m *_IdentifyReplyCommandLogicalAssignment) String() string {
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