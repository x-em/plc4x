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

// BACnetConfirmedServiceRequestGetEventInformation is the corresponding interface of BACnetConfirmedServiceRequestGetEventInformation
type BACnetConfirmedServiceRequestGetEventInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetLastReceivedObjectIdentifier returns LastReceivedObjectIdentifier (property field)
	GetLastReceivedObjectIdentifier() BACnetContextTagObjectIdentifier
	// IsBACnetConfirmedServiceRequestGetEventInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestGetEventInformation()
	// CreateBuilder creates a BACnetConfirmedServiceRequestGetEventInformationBuilder
	CreateBACnetConfirmedServiceRequestGetEventInformationBuilder() BACnetConfirmedServiceRequestGetEventInformationBuilder
}

// _BACnetConfirmedServiceRequestGetEventInformation is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEventInformation struct {
	BACnetConfirmedServiceRequestContract
	LastReceivedObjectIdentifier BACnetContextTagObjectIdentifier
}

var _ BACnetConfirmedServiceRequestGetEventInformation = (*_BACnetConfirmedServiceRequestGetEventInformation)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestGetEventInformation)(nil)

// NewBACnetConfirmedServiceRequestGetEventInformation factory function for _BACnetConfirmedServiceRequestGetEventInformation
func NewBACnetConfirmedServiceRequestGetEventInformation(lastReceivedObjectIdentifier BACnetContextTagObjectIdentifier, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestGetEventInformation {
	_result := &_BACnetConfirmedServiceRequestGetEventInformation{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		LastReceivedObjectIdentifier:          lastReceivedObjectIdentifier,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestGetEventInformationBuilder is a builder for BACnetConfirmedServiceRequestGetEventInformation
type BACnetConfirmedServiceRequestGetEventInformationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConfirmedServiceRequestGetEventInformationBuilder
	// WithLastReceivedObjectIdentifier adds LastReceivedObjectIdentifier (property field)
	WithOptionalLastReceivedObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestGetEventInformationBuilder
	// WithOptionalLastReceivedObjectIdentifierBuilder adds LastReceivedObjectIdentifier (property field) which is build by the builder
	WithOptionalLastReceivedObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestGetEventInformationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestBuilder
	// Build builds the BACnetConfirmedServiceRequestGetEventInformation or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestGetEventInformation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestGetEventInformation
}

// NewBACnetConfirmedServiceRequestGetEventInformationBuilder() creates a BACnetConfirmedServiceRequestGetEventInformationBuilder
func NewBACnetConfirmedServiceRequestGetEventInformationBuilder() BACnetConfirmedServiceRequestGetEventInformationBuilder {
	return &_BACnetConfirmedServiceRequestGetEventInformationBuilder{_BACnetConfirmedServiceRequestGetEventInformation: new(_BACnetConfirmedServiceRequestGetEventInformation)}
}

type _BACnetConfirmedServiceRequestGetEventInformationBuilder struct {
	*_BACnetConfirmedServiceRequestGetEventInformation

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestGetEventInformationBuilder) = (*_BACnetConfirmedServiceRequestGetEventInformationBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
	contract.(*_BACnetConfirmedServiceRequest)._SubType = b._BACnetConfirmedServiceRequestGetEventInformation
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) WithMandatoryFields() BACnetConfirmedServiceRequestGetEventInformationBuilder {
	return b
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) WithOptionalLastReceivedObjectIdentifier(lastReceivedObjectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestGetEventInformationBuilder {
	b.LastReceivedObjectIdentifier = lastReceivedObjectIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) WithOptionalLastReceivedObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestGetEventInformationBuilder {
	builder := builderSupplier(b.LastReceivedObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.LastReceivedObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) Build() (BACnetConfirmedServiceRequestGetEventInformation, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestGetEventInformation.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) MustBuild() BACnetConfirmedServiceRequestGetEventInformation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestBuilder().(*_BACnetConfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestGetEventInformationBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestGetEventInformationBuilder().(*_BACnetConfirmedServiceRequestGetEventInformationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestGetEventInformationBuilder creates a BACnetConfirmedServiceRequestGetEventInformationBuilder
func (b *_BACnetConfirmedServiceRequestGetEventInformation) CreateBACnetConfirmedServiceRequestGetEventInformationBuilder() BACnetConfirmedServiceRequestGetEventInformationBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestGetEventInformationBuilder()
	}
	return &_BACnetConfirmedServiceRequestGetEventInformationBuilder{_BACnetConfirmedServiceRequestGetEventInformation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetLastReceivedObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.LastReceivedObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEventInformation(structType any) BACnetConfirmedServiceRequestGetEventInformation {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEventInformation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEventInformation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEventInformation"
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Optional Field (lastReceivedObjectIdentifier)
	if m.LastReceivedObjectIdentifier != nil {
		lengthInBits += m.LastReceivedObjectIdentifier.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestGetEventInformation BACnetConfirmedServiceRequestGetEventInformation, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEventInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEventInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var lastReceivedObjectIdentifier BACnetContextTagObjectIdentifier
	_lastReceivedObjectIdentifier, err := ReadOptionalField[BACnetContextTagObjectIdentifier](ctx, "lastReceivedObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastReceivedObjectIdentifier' field"))
	}
	if _lastReceivedObjectIdentifier != nil {
		lastReceivedObjectIdentifier = *_lastReceivedObjectIdentifier
		m.LastReceivedObjectIdentifier = lastReceivedObjectIdentifier
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEventInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEventInformation")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEventInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEventInformation")
		}

		if err := WriteOptionalField[BACnetContextTagObjectIdentifier](ctx, "lastReceivedObjectIdentifier", GetRef(m.GetLastReceivedObjectIdentifier()), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lastReceivedObjectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEventInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEventInformation")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) IsBACnetConfirmedServiceRequestGetEventInformation() {
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) deepCopy() *_BACnetConfirmedServiceRequestGetEventInformation {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestGetEventInformationCopy := &_BACnetConfirmedServiceRequestGetEventInformation{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.LastReceivedObjectIdentifier),
	}
	_BACnetConfirmedServiceRequestGetEventInformationCopy.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestGetEventInformationCopy
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) String() string {
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