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

// BACnetSetpointReference is the corresponding interface of BACnetSetpointReference
type BACnetSetpointReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetSetPointReference returns SetPointReference (property field)
	GetSetPointReference() BACnetObjectPropertyReferenceEnclosed
	// IsBACnetSetpointReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetSetpointReference()
	// CreateBuilder creates a BACnetSetpointReferenceBuilder
	CreateBACnetSetpointReferenceBuilder() BACnetSetpointReferenceBuilder
}

// _BACnetSetpointReference is the data-structure of this message
type _BACnetSetpointReference struct {
	SetPointReference BACnetObjectPropertyReferenceEnclosed
}

var _ BACnetSetpointReference = (*_BACnetSetpointReference)(nil)

// NewBACnetSetpointReference factory function for _BACnetSetpointReference
func NewBACnetSetpointReference(setPointReference BACnetObjectPropertyReferenceEnclosed) *_BACnetSetpointReference {
	return &_BACnetSetpointReference{SetPointReference: setPointReference}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetSetpointReferenceBuilder is a builder for BACnetSetpointReference
type BACnetSetpointReferenceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetSetpointReferenceBuilder
	// WithSetPointReference adds SetPointReference (property field)
	WithOptionalSetPointReference(BACnetObjectPropertyReferenceEnclosed) BACnetSetpointReferenceBuilder
	// WithOptionalSetPointReferenceBuilder adds SetPointReference (property field) which is build by the builder
	WithOptionalSetPointReferenceBuilder(func(BACnetObjectPropertyReferenceEnclosedBuilder) BACnetObjectPropertyReferenceEnclosedBuilder) BACnetSetpointReferenceBuilder
	// Build builds the BACnetSetpointReference or returns an error if something is wrong
	Build() (BACnetSetpointReference, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetSetpointReference
}

// NewBACnetSetpointReferenceBuilder() creates a BACnetSetpointReferenceBuilder
func NewBACnetSetpointReferenceBuilder() BACnetSetpointReferenceBuilder {
	return &_BACnetSetpointReferenceBuilder{_BACnetSetpointReference: new(_BACnetSetpointReference)}
}

type _BACnetSetpointReferenceBuilder struct {
	*_BACnetSetpointReference

	err *utils.MultiError
}

var _ (BACnetSetpointReferenceBuilder) = (*_BACnetSetpointReferenceBuilder)(nil)

func (b *_BACnetSetpointReferenceBuilder) WithMandatoryFields() BACnetSetpointReferenceBuilder {
	return b
}

func (b *_BACnetSetpointReferenceBuilder) WithOptionalSetPointReference(setPointReference BACnetObjectPropertyReferenceEnclosed) BACnetSetpointReferenceBuilder {
	b.SetPointReference = setPointReference
	return b
}

func (b *_BACnetSetpointReferenceBuilder) WithOptionalSetPointReferenceBuilder(builderSupplier func(BACnetObjectPropertyReferenceEnclosedBuilder) BACnetObjectPropertyReferenceEnclosedBuilder) BACnetSetpointReferenceBuilder {
	builder := builderSupplier(b.SetPointReference.CreateBACnetObjectPropertyReferenceEnclosedBuilder())
	var err error
	b.SetPointReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetSetpointReferenceBuilder) Build() (BACnetSetpointReference, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetSetpointReference.deepCopy(), nil
}

func (b *_BACnetSetpointReferenceBuilder) MustBuild() BACnetSetpointReference {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetSetpointReferenceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetSetpointReferenceBuilder().(*_BACnetSetpointReferenceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetSetpointReferenceBuilder creates a BACnetSetpointReferenceBuilder
func (b *_BACnetSetpointReference) CreateBACnetSetpointReferenceBuilder() BACnetSetpointReferenceBuilder {
	if b == nil {
		return NewBACnetSetpointReferenceBuilder()
	}
	return &_BACnetSetpointReferenceBuilder{_BACnetSetpointReference: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSetpointReference) GetSetPointReference() BACnetObjectPropertyReferenceEnclosed {
	return m.SetPointReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetSetpointReference(structType any) BACnetSetpointReference {
	if casted, ok := structType.(BACnetSetpointReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSetpointReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSetpointReference) GetTypeName() string {
	return "BACnetSetpointReference"
}

func (m *_BACnetSetpointReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (setPointReference)
	if m.SetPointReference != nil {
		lengthInBits += m.SetPointReference.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetSetpointReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSetpointReferenceParse(ctx context.Context, theBytes []byte) (BACnetSetpointReference, error) {
	return BACnetSetpointReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetSetpointReferenceParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSetpointReference, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSetpointReference, error) {
		return BACnetSetpointReferenceParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetSetpointReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSetpointReference, error) {
	v, err := (&_BACnetSetpointReference{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetSetpointReference) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetSetpointReference BACnetSetpointReference, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSetpointReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSetpointReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var setPointReference BACnetObjectPropertyReferenceEnclosed
	_setPointReference, err := ReadOptionalField[BACnetObjectPropertyReferenceEnclosed](ctx, "setPointReference", ReadComplex[BACnetObjectPropertyReferenceEnclosed](BACnetObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'setPointReference' field"))
	}
	if _setPointReference != nil {
		setPointReference = *_setPointReference
		m.SetPointReference = setPointReference
	}

	if closeErr := readBuffer.CloseContext("BACnetSetpointReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSetpointReference")
	}

	return m, nil
}

func (m *_BACnetSetpointReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSetpointReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSetpointReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSetpointReference")
	}

	if err := WriteOptionalField[BACnetObjectPropertyReferenceEnclosed](ctx, "setPointReference", GetRef(m.GetSetPointReference()), WriteComplex[BACnetObjectPropertyReferenceEnclosed](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'setPointReference' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSetpointReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSetpointReference")
	}
	return nil
}

func (m *_BACnetSetpointReference) IsBACnetSetpointReference() {}

func (m *_BACnetSetpointReference) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetSetpointReference) deepCopy() *_BACnetSetpointReference {
	if m == nil {
		return nil
	}
	_BACnetSetpointReferenceCopy := &_BACnetSetpointReference{
		utils.DeepCopy[BACnetObjectPropertyReferenceEnclosed](m.SetPointReference),
	}
	return _BACnetSetpointReferenceCopy
}

func (m *_BACnetSetpointReference) String() string {
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