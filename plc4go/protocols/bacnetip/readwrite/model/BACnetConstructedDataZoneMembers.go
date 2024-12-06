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

// BACnetConstructedDataZoneMembers is the corresponding interface of BACnetConstructedDataZoneMembers
type BACnetConstructedDataZoneMembers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMembers returns Members (property field)
	GetMembers() []BACnetDeviceObjectReference
	// IsBACnetConstructedDataZoneMembers is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataZoneMembers()
	// CreateBuilder creates a BACnetConstructedDataZoneMembersBuilder
	CreateBACnetConstructedDataZoneMembersBuilder() BACnetConstructedDataZoneMembersBuilder
}

// _BACnetConstructedDataZoneMembers is the data-structure of this message
type _BACnetConstructedDataZoneMembers struct {
	BACnetConstructedDataContract
	Members []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataZoneMembers = (*_BACnetConstructedDataZoneMembers)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataZoneMembers)(nil)

// NewBACnetConstructedDataZoneMembers factory function for _BACnetConstructedDataZoneMembers
func NewBACnetConstructedDataZoneMembers(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, members []BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataZoneMembers {
	_result := &_BACnetConstructedDataZoneMembers{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Members:                       members,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataZoneMembersBuilder is a builder for BACnetConstructedDataZoneMembers
type BACnetConstructedDataZoneMembersBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(members []BACnetDeviceObjectReference) BACnetConstructedDataZoneMembersBuilder
	// WithMembers adds Members (property field)
	WithMembers(...BACnetDeviceObjectReference) BACnetConstructedDataZoneMembersBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataZoneMembers or returns an error if something is wrong
	Build() (BACnetConstructedDataZoneMembers, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataZoneMembers
}

// NewBACnetConstructedDataZoneMembersBuilder() creates a BACnetConstructedDataZoneMembersBuilder
func NewBACnetConstructedDataZoneMembersBuilder() BACnetConstructedDataZoneMembersBuilder {
	return &_BACnetConstructedDataZoneMembersBuilder{_BACnetConstructedDataZoneMembers: new(_BACnetConstructedDataZoneMembers)}
}

type _BACnetConstructedDataZoneMembersBuilder struct {
	*_BACnetConstructedDataZoneMembers

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataZoneMembersBuilder) = (*_BACnetConstructedDataZoneMembersBuilder)(nil)

func (b *_BACnetConstructedDataZoneMembersBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataZoneMembers
}

func (b *_BACnetConstructedDataZoneMembersBuilder) WithMandatoryFields(members []BACnetDeviceObjectReference) BACnetConstructedDataZoneMembersBuilder {
	return b.WithMembers(members...)
}

func (b *_BACnetConstructedDataZoneMembersBuilder) WithMembers(members ...BACnetDeviceObjectReference) BACnetConstructedDataZoneMembersBuilder {
	b.Members = members
	return b
}

func (b *_BACnetConstructedDataZoneMembersBuilder) Build() (BACnetConstructedDataZoneMembers, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataZoneMembers.deepCopy(), nil
}

func (b *_BACnetConstructedDataZoneMembersBuilder) MustBuild() BACnetConstructedDataZoneMembers {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataZoneMembersBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataZoneMembersBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataZoneMembersBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataZoneMembersBuilder().(*_BACnetConstructedDataZoneMembersBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataZoneMembersBuilder creates a BACnetConstructedDataZoneMembersBuilder
func (b *_BACnetConstructedDataZoneMembers) CreateBACnetConstructedDataZoneMembersBuilder() BACnetConstructedDataZoneMembersBuilder {
	if b == nil {
		return NewBACnetConstructedDataZoneMembersBuilder()
	}
	return &_BACnetConstructedDataZoneMembersBuilder{_BACnetConstructedDataZoneMembers: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataZoneMembers) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataZoneMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ZONE_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataZoneMembers) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataZoneMembers) GetMembers() []BACnetDeviceObjectReference {
	return m.Members
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataZoneMembers(structType any) BACnetConstructedDataZoneMembers {
	if casted, ok := structType.(BACnetConstructedDataZoneMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataZoneMembers) GetTypeName() string {
	return "BACnetConstructedDataZoneMembers"
}

func (m *_BACnetConstructedDataZoneMembers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.Members) > 0 {
		for _, element := range m.Members {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataZoneMembers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataZoneMembers) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataZoneMembers BACnetConstructedDataZoneMembers, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	members, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "members", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'members' field"))
	}
	m.Members = members

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneMembers")
	}

	return m, nil
}

func (m *_BACnetConstructedDataZoneMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataZoneMembers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneMembers")
		}

		if err := WriteComplexTypeArrayField(ctx, "members", m.GetMembers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'members' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneMembers")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataZoneMembers) IsBACnetConstructedDataZoneMembers() {}

func (m *_BACnetConstructedDataZoneMembers) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataZoneMembers) deepCopy() *_BACnetConstructedDataZoneMembers {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataZoneMembersCopy := &_BACnetConstructedDataZoneMembers{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetDeviceObjectReference, BACnetDeviceObjectReference](m.Members),
	}
	_BACnetConstructedDataZoneMembersCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataZoneMembersCopy
}

func (m *_BACnetConstructedDataZoneMembers) String() string {
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