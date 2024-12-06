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

// ClassSegment is the corresponding interface of ClassSegment
type ClassSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetPathSegmentType returns PathSegmentType (property field)
	GetPathSegmentType() uint8
	// GetLogicalSegmentType returns LogicalSegmentType (property field)
	GetLogicalSegmentType() uint8
	// GetLogicalSegmentFormat returns LogicalSegmentFormat (property field)
	GetLogicalSegmentFormat() uint8
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() uint8
	// IsClassSegment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsClassSegment()
	// CreateBuilder creates a ClassSegmentBuilder
	CreateClassSegmentBuilder() ClassSegmentBuilder
}

// _ClassSegment is the data-structure of this message
type _ClassSegment struct {
	PathSegmentType      uint8
	LogicalSegmentType   uint8
	LogicalSegmentFormat uint8
	ClassSegment         uint8
}

var _ ClassSegment = (*_ClassSegment)(nil)

// NewClassSegment factory function for _ClassSegment
func NewClassSegment(pathSegmentType uint8, logicalSegmentType uint8, logicalSegmentFormat uint8, classSegment uint8) *_ClassSegment {
	return &_ClassSegment{PathSegmentType: pathSegmentType, LogicalSegmentType: logicalSegmentType, LogicalSegmentFormat: logicalSegmentFormat, ClassSegment: classSegment}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ClassSegmentBuilder is a builder for ClassSegment
type ClassSegmentBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pathSegmentType uint8, logicalSegmentType uint8, logicalSegmentFormat uint8, classSegment uint8) ClassSegmentBuilder
	// WithPathSegmentType adds PathSegmentType (property field)
	WithPathSegmentType(uint8) ClassSegmentBuilder
	// WithLogicalSegmentType adds LogicalSegmentType (property field)
	WithLogicalSegmentType(uint8) ClassSegmentBuilder
	// WithLogicalSegmentFormat adds LogicalSegmentFormat (property field)
	WithLogicalSegmentFormat(uint8) ClassSegmentBuilder
	// WithClassSegment adds ClassSegment (property field)
	WithClassSegment(uint8) ClassSegmentBuilder
	// Build builds the ClassSegment or returns an error if something is wrong
	Build() (ClassSegment, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ClassSegment
}

// NewClassSegmentBuilder() creates a ClassSegmentBuilder
func NewClassSegmentBuilder() ClassSegmentBuilder {
	return &_ClassSegmentBuilder{_ClassSegment: new(_ClassSegment)}
}

type _ClassSegmentBuilder struct {
	*_ClassSegment

	err *utils.MultiError
}

var _ (ClassSegmentBuilder) = (*_ClassSegmentBuilder)(nil)

func (b *_ClassSegmentBuilder) WithMandatoryFields(pathSegmentType uint8, logicalSegmentType uint8, logicalSegmentFormat uint8, classSegment uint8) ClassSegmentBuilder {
	return b.WithPathSegmentType(pathSegmentType).WithLogicalSegmentType(logicalSegmentType).WithLogicalSegmentFormat(logicalSegmentFormat).WithClassSegment(classSegment)
}

func (b *_ClassSegmentBuilder) WithPathSegmentType(pathSegmentType uint8) ClassSegmentBuilder {
	b.PathSegmentType = pathSegmentType
	return b
}

func (b *_ClassSegmentBuilder) WithLogicalSegmentType(logicalSegmentType uint8) ClassSegmentBuilder {
	b.LogicalSegmentType = logicalSegmentType
	return b
}

func (b *_ClassSegmentBuilder) WithLogicalSegmentFormat(logicalSegmentFormat uint8) ClassSegmentBuilder {
	b.LogicalSegmentFormat = logicalSegmentFormat
	return b
}

func (b *_ClassSegmentBuilder) WithClassSegment(classSegment uint8) ClassSegmentBuilder {
	b.ClassSegment = classSegment
	return b
}

func (b *_ClassSegmentBuilder) Build() (ClassSegment, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ClassSegment.deepCopy(), nil
}

func (b *_ClassSegmentBuilder) MustBuild() ClassSegment {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ClassSegmentBuilder) DeepCopy() any {
	_copy := b.CreateClassSegmentBuilder().(*_ClassSegmentBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateClassSegmentBuilder creates a ClassSegmentBuilder
func (b *_ClassSegment) CreateClassSegmentBuilder() ClassSegmentBuilder {
	if b == nil {
		return NewClassSegmentBuilder()
	}
	return &_ClassSegmentBuilder{_ClassSegment: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClassSegment) GetPathSegmentType() uint8 {
	return m.PathSegmentType
}

func (m *_ClassSegment) GetLogicalSegmentType() uint8 {
	return m.LogicalSegmentType
}

func (m *_ClassSegment) GetLogicalSegmentFormat() uint8 {
	return m.LogicalSegmentFormat
}

func (m *_ClassSegment) GetClassSegment() uint8 {
	return m.ClassSegment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastClassSegment(structType any) ClassSegment {
	if casted, ok := structType.(ClassSegment); ok {
		return casted
	}
	if casted, ok := structType.(*ClassSegment); ok {
		return *casted
	}
	return nil
}

func (m *_ClassSegment) GetTypeName() string {
	return "ClassSegment"
}

func (m *_ClassSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (pathSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentFormat)
	lengthInBits += 2

	// Simple field (classSegment)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ClassSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ClassSegmentParse(ctx context.Context, theBytes []byte) (ClassSegment, error) {
	return ClassSegmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ClassSegmentParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ClassSegment, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ClassSegment, error) {
		return ClassSegmentParseWithBuffer(ctx, readBuffer)
	}
}

func ClassSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ClassSegment, error) {
	v, err := (&_ClassSegment{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ClassSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__classSegment ClassSegment, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClassSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClassSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pathSegmentType, err := ReadSimpleField(ctx, "pathSegmentType", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pathSegmentType' field"))
	}
	m.PathSegmentType = pathSegmentType

	logicalSegmentType, err := ReadSimpleField(ctx, "logicalSegmentType", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicalSegmentType' field"))
	}
	m.LogicalSegmentType = logicalSegmentType

	logicalSegmentFormat, err := ReadSimpleField(ctx, "logicalSegmentFormat", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicalSegmentFormat' field"))
	}
	m.LogicalSegmentFormat = logicalSegmentFormat

	classSegment, err := ReadSimpleField(ctx, "classSegment", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classSegment' field"))
	}
	m.ClassSegment = classSegment

	if closeErr := readBuffer.CloseContext("ClassSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClassSegment")
	}

	return m, nil
}

func (m *_ClassSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClassSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ClassSegment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ClassSegment")
	}

	if err := WriteSimpleField[uint8](ctx, "pathSegmentType", m.GetPathSegmentType(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'pathSegmentType' field")
	}

	if err := WriteSimpleField[uint8](ctx, "logicalSegmentType", m.GetLogicalSegmentType(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'logicalSegmentType' field")
	}

	if err := WriteSimpleField[uint8](ctx, "logicalSegmentFormat", m.GetLogicalSegmentFormat(), WriteUnsignedByte(writeBuffer, 2)); err != nil {
		return errors.Wrap(err, "Error serializing 'logicalSegmentFormat' field")
	}

	if err := WriteSimpleField[uint8](ctx, "classSegment", m.GetClassSegment(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'classSegment' field")
	}

	if popErr := writeBuffer.PopContext("ClassSegment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ClassSegment")
	}
	return nil
}

func (m *_ClassSegment) IsClassSegment() {}

func (m *_ClassSegment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ClassSegment) deepCopy() *_ClassSegment {
	if m == nil {
		return nil
	}
	_ClassSegmentCopy := &_ClassSegment{
		m.PathSegmentType,
		m.LogicalSegmentType,
		m.LogicalSegmentFormat,
		m.ClassSegment,
	}
	return _ClassSegmentCopy
}

func (m *_ClassSegment) String() string {
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
