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

// BACnetFaultParameterFaultExtendedParametersEntryReference is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryReference
type BACnetFaultParameterFaultExtendedParametersEntryReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetReference returns Reference (property field)
	GetReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// IsBACnetFaultParameterFaultExtendedParametersEntryReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryReference()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder() BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryReference is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryReference struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	Reference BACnetDeviceObjectPropertyReferenceEnclosed
}

var _ BACnetFaultParameterFaultExtendedParametersEntryReference = (*_BACnetFaultParameterFaultExtendedParametersEntryReference)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryReference)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryReference factory function for _BACnetFaultParameterFaultExtendedParametersEntryReference
func NewBACnetFaultParameterFaultExtendedParametersEntryReference(peekedTagHeader BACnetTagHeader, reference BACnetDeviceObjectPropertyReferenceEnclosed) *_BACnetFaultParameterFaultExtendedParametersEntryReference {
	if reference == nil {
		panic("reference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetFaultParameterFaultExtendedParametersEntryReference must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryReference{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		Reference: reference,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryReference
type BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder
	// WithReference adds Reference (property field)
	WithReference(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder
	// WithReferenceBuilder adds Reference (property field) which is build by the builder
	WithReferenceBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryReference or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryReference, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryReference
}

// NewBACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder() BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder{_BACnetFaultParameterFaultExtendedParametersEntryReference: new(_BACnetFaultParameterFaultExtendedParametersEntryReference)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryReference

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryReference
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) WithMandatoryFields(reference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder {
	return b.WithReference(reference)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) WithReference(reference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder {
	b.Reference = reference
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) WithReferenceBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder {
	builder := builderSupplier(b.Reference.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	b.Reference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryReference, error) {
	if b.Reference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'reference' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryReference.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryReference {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryReference) CreateBACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder() BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryReferenceBuilder{_BACnetFaultParameterFaultExtendedParametersEntryReference: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.Reference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryReference(structType any) BACnetFaultParameterFaultExtendedParametersEntryReference {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryReference"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (reference)
	lengthInBits += m.Reference.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryReference BACnetFaultParameterFaultExtendedParametersEntryReference, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "reference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reference' field"))
	}
	m.Reference = reference

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryReference")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryReference")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "reference", m.GetReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reference' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryReference")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) IsBACnetFaultParameterFaultExtendedParametersEntryReference() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryReference {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryReferenceCopy := &_BACnetFaultParameterFaultExtendedParametersEntryReference{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectPropertyReferenceEnclosed](m.Reference),
	}
	_BACnetFaultParameterFaultExtendedParametersEntryReferenceCopy.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryReferenceCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) String() string {
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