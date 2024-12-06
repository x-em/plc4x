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

// BACnetConstructedDataDaylightSavingsStatus is the corresponding interface of BACnetConstructedDataDaylightSavingsStatus
type BACnetConstructedDataDaylightSavingsStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDaylightSavingsStatus returns DaylightSavingsStatus (property field)
	GetDaylightSavingsStatus() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataDaylightSavingsStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDaylightSavingsStatus()
	// CreateBuilder creates a BACnetConstructedDataDaylightSavingsStatusBuilder
	CreateBACnetConstructedDataDaylightSavingsStatusBuilder() BACnetConstructedDataDaylightSavingsStatusBuilder
}

// _BACnetConstructedDataDaylightSavingsStatus is the data-structure of this message
type _BACnetConstructedDataDaylightSavingsStatus struct {
	BACnetConstructedDataContract
	DaylightSavingsStatus BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataDaylightSavingsStatus = (*_BACnetConstructedDataDaylightSavingsStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDaylightSavingsStatus)(nil)

// NewBACnetConstructedDataDaylightSavingsStatus factory function for _BACnetConstructedDataDaylightSavingsStatus
func NewBACnetConstructedDataDaylightSavingsStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, daylightSavingsStatus BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDaylightSavingsStatus {
	if daylightSavingsStatus == nil {
		panic("daylightSavingsStatus of type BACnetApplicationTagBoolean for BACnetConstructedDataDaylightSavingsStatus must not be nil")
	}
	_result := &_BACnetConstructedDataDaylightSavingsStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DaylightSavingsStatus:         daylightSavingsStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDaylightSavingsStatusBuilder is a builder for BACnetConstructedDataDaylightSavingsStatus
type BACnetConstructedDataDaylightSavingsStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(daylightSavingsStatus BACnetApplicationTagBoolean) BACnetConstructedDataDaylightSavingsStatusBuilder
	// WithDaylightSavingsStatus adds DaylightSavingsStatus (property field)
	WithDaylightSavingsStatus(BACnetApplicationTagBoolean) BACnetConstructedDataDaylightSavingsStatusBuilder
	// WithDaylightSavingsStatusBuilder adds DaylightSavingsStatus (property field) which is build by the builder
	WithDaylightSavingsStatusBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataDaylightSavingsStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDaylightSavingsStatus or returns an error if something is wrong
	Build() (BACnetConstructedDataDaylightSavingsStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDaylightSavingsStatus
}

// NewBACnetConstructedDataDaylightSavingsStatusBuilder() creates a BACnetConstructedDataDaylightSavingsStatusBuilder
func NewBACnetConstructedDataDaylightSavingsStatusBuilder() BACnetConstructedDataDaylightSavingsStatusBuilder {
	return &_BACnetConstructedDataDaylightSavingsStatusBuilder{_BACnetConstructedDataDaylightSavingsStatus: new(_BACnetConstructedDataDaylightSavingsStatus)}
}

type _BACnetConstructedDataDaylightSavingsStatusBuilder struct {
	*_BACnetConstructedDataDaylightSavingsStatus

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDaylightSavingsStatusBuilder) = (*_BACnetConstructedDataDaylightSavingsStatusBuilder)(nil)

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDaylightSavingsStatus
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) WithMandatoryFields(daylightSavingsStatus BACnetApplicationTagBoolean) BACnetConstructedDataDaylightSavingsStatusBuilder {
	return b.WithDaylightSavingsStatus(daylightSavingsStatus)
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) WithDaylightSavingsStatus(daylightSavingsStatus BACnetApplicationTagBoolean) BACnetConstructedDataDaylightSavingsStatusBuilder {
	b.DaylightSavingsStatus = daylightSavingsStatus
	return b
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) WithDaylightSavingsStatusBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataDaylightSavingsStatusBuilder {
	builder := builderSupplier(b.DaylightSavingsStatus.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.DaylightSavingsStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) Build() (BACnetConstructedDataDaylightSavingsStatus, error) {
	if b.DaylightSavingsStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'daylightSavingsStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDaylightSavingsStatus.deepCopy(), nil
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) MustBuild() BACnetConstructedDataDaylightSavingsStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDaylightSavingsStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDaylightSavingsStatusBuilder().(*_BACnetConstructedDataDaylightSavingsStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDaylightSavingsStatusBuilder creates a BACnetConstructedDataDaylightSavingsStatusBuilder
func (b *_BACnetConstructedDataDaylightSavingsStatus) CreateBACnetConstructedDataDaylightSavingsStatusBuilder() BACnetConstructedDataDaylightSavingsStatusBuilder {
	if b == nil {
		return NewBACnetConstructedDataDaylightSavingsStatusBuilder()
	}
	return &_BACnetConstructedDataDaylightSavingsStatusBuilder{_BACnetConstructedDataDaylightSavingsStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DAYLIGHT_SAVINGS_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetDaylightSavingsStatus() BACnetApplicationTagBoolean {
	return m.DaylightSavingsStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetDaylightSavingsStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDaylightSavingsStatus(structType any) BACnetConstructedDataDaylightSavingsStatus {
	if casted, ok := structType.(BACnetConstructedDataDaylightSavingsStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDaylightSavingsStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetTypeName() string {
	return "BACnetConstructedDataDaylightSavingsStatus"
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (daylightSavingsStatus)
	lengthInBits += m.DaylightSavingsStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDaylightSavingsStatus BACnetConstructedDataDaylightSavingsStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDaylightSavingsStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDaylightSavingsStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	daylightSavingsStatus, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "daylightSavingsStatus", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daylightSavingsStatus' field"))
	}
	m.DaylightSavingsStatus = daylightSavingsStatus

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), daylightSavingsStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDaylightSavingsStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDaylightSavingsStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDaylightSavingsStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDaylightSavingsStatus")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "daylightSavingsStatus", m.GetDaylightSavingsStatus(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'daylightSavingsStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDaylightSavingsStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDaylightSavingsStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) IsBACnetConstructedDataDaylightSavingsStatus() {
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) deepCopy() *_BACnetConstructedDataDaylightSavingsStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDaylightSavingsStatusCopy := &_BACnetConstructedDataDaylightSavingsStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.DaylightSavingsStatus),
	}
	_BACnetConstructedDataDaylightSavingsStatusCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDaylightSavingsStatusCopy
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) String() string {
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