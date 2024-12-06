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

// BACnetNotificationParametersFloatingLimit is the corresponding interface of BACnetNotificationParametersFloatingLimit
type BACnetNotificationParametersFloatingLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetReferenceValue returns ReferenceValue (property field)
	GetReferenceValue() BACnetContextTagReal
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetSetPointValue returns SetPointValue (property field)
	GetSetPointValue() BACnetContextTagReal
	// GetErrorLimit returns ErrorLimit (property field)
	GetErrorLimit() BACnetContextTagReal
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersFloatingLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersFloatingLimit()
	// CreateBuilder creates a BACnetNotificationParametersFloatingLimitBuilder
	CreateBACnetNotificationParametersFloatingLimitBuilder() BACnetNotificationParametersFloatingLimitBuilder
}

// _BACnetNotificationParametersFloatingLimit is the data-structure of this message
type _BACnetNotificationParametersFloatingLimit struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	ReferenceValue  BACnetContextTagReal
	StatusFlags     BACnetStatusFlagsTagged
	SetPointValue   BACnetContextTagReal
	ErrorLimit      BACnetContextTagReal
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersFloatingLimit = (*_BACnetNotificationParametersFloatingLimit)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersFloatingLimit)(nil)

// NewBACnetNotificationParametersFloatingLimit factory function for _BACnetNotificationParametersFloatingLimit
func NewBACnetNotificationParametersFloatingLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, referenceValue BACnetContextTagReal, statusFlags BACnetStatusFlagsTagged, setPointValue BACnetContextTagReal, errorLimit BACnetContextTagReal, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersFloatingLimit {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersFloatingLimit must not be nil")
	}
	if referenceValue == nil {
		panic("referenceValue of type BACnetContextTagReal for BACnetNotificationParametersFloatingLimit must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersFloatingLimit must not be nil")
	}
	if setPointValue == nil {
		panic("setPointValue of type BACnetContextTagReal for BACnetNotificationParametersFloatingLimit must not be nil")
	}
	if errorLimit == nil {
		panic("errorLimit of type BACnetContextTagReal for BACnetNotificationParametersFloatingLimit must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersFloatingLimit must not be nil")
	}
	_result := &_BACnetNotificationParametersFloatingLimit{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		ReferenceValue:                       referenceValue,
		StatusFlags:                          statusFlags,
		SetPointValue:                        setPointValue,
		ErrorLimit:                           errorLimit,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersFloatingLimitBuilder is a builder for BACnetNotificationParametersFloatingLimit
type BACnetNotificationParametersFloatingLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, referenceValue BACnetContextTagReal, statusFlags BACnetStatusFlagsTagged, setPointValue BACnetContextTagReal, errorLimit BACnetContextTagReal, innerClosingTag BACnetClosingTag) BACnetNotificationParametersFloatingLimitBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersFloatingLimitBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersFloatingLimitBuilder
	// WithReferenceValue adds ReferenceValue (property field)
	WithReferenceValue(BACnetContextTagReal) BACnetNotificationParametersFloatingLimitBuilder
	// WithReferenceValueBuilder adds ReferenceValue (property field) which is build by the builder
	WithReferenceValueBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetNotificationParametersFloatingLimitBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersFloatingLimitBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersFloatingLimitBuilder
	// WithSetPointValue adds SetPointValue (property field)
	WithSetPointValue(BACnetContextTagReal) BACnetNotificationParametersFloatingLimitBuilder
	// WithSetPointValueBuilder adds SetPointValue (property field) which is build by the builder
	WithSetPointValueBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetNotificationParametersFloatingLimitBuilder
	// WithErrorLimit adds ErrorLimit (property field)
	WithErrorLimit(BACnetContextTagReal) BACnetNotificationParametersFloatingLimitBuilder
	// WithErrorLimitBuilder adds ErrorLimit (property field) which is build by the builder
	WithErrorLimitBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetNotificationParametersFloatingLimitBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersFloatingLimitBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersFloatingLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersBuilder
	// Build builds the BACnetNotificationParametersFloatingLimit or returns an error if something is wrong
	Build() (BACnetNotificationParametersFloatingLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersFloatingLimit
}

// NewBACnetNotificationParametersFloatingLimitBuilder() creates a BACnetNotificationParametersFloatingLimitBuilder
func NewBACnetNotificationParametersFloatingLimitBuilder() BACnetNotificationParametersFloatingLimitBuilder {
	return &_BACnetNotificationParametersFloatingLimitBuilder{_BACnetNotificationParametersFloatingLimit: new(_BACnetNotificationParametersFloatingLimit)}
}

type _BACnetNotificationParametersFloatingLimitBuilder struct {
	*_BACnetNotificationParametersFloatingLimit

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersFloatingLimitBuilder) = (*_BACnetNotificationParametersFloatingLimitBuilder)(nil)

func (b *_BACnetNotificationParametersFloatingLimitBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
	contract.(*_BACnetNotificationParameters)._SubType = b._BACnetNotificationParametersFloatingLimit
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, referenceValue BACnetContextTagReal, statusFlags BACnetStatusFlagsTagged, setPointValue BACnetContextTagReal, errorLimit BACnetContextTagReal, innerClosingTag BACnetClosingTag) BACnetNotificationParametersFloatingLimitBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithReferenceValue(referenceValue).WithStatusFlags(statusFlags).WithSetPointValue(setPointValue).WithErrorLimit(errorLimit).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersFloatingLimitBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersFloatingLimitBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithReferenceValue(referenceValue BACnetContextTagReal) BACnetNotificationParametersFloatingLimitBuilder {
	b.ReferenceValue = referenceValue
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithReferenceValueBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetNotificationParametersFloatingLimitBuilder {
	builder := builderSupplier(b.ReferenceValue.CreateBACnetContextTagRealBuilder())
	var err error
	b.ReferenceValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetNotificationParametersFloatingLimitBuilder {
	b.StatusFlags = statusFlags
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersFloatingLimitBuilder {
	builder := builderSupplier(b.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.StatusFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithSetPointValue(setPointValue BACnetContextTagReal) BACnetNotificationParametersFloatingLimitBuilder {
	b.SetPointValue = setPointValue
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithSetPointValueBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetNotificationParametersFloatingLimitBuilder {
	builder := builderSupplier(b.SetPointValue.CreateBACnetContextTagRealBuilder())
	var err error
	b.SetPointValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithErrorLimit(errorLimit BACnetContextTagReal) BACnetNotificationParametersFloatingLimitBuilder {
	b.ErrorLimit = errorLimit
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithErrorLimitBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetNotificationParametersFloatingLimitBuilder {
	builder := builderSupplier(b.ErrorLimit.CreateBACnetContextTagRealBuilder())
	var err error
	b.ErrorLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersFloatingLimitBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersFloatingLimitBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) Build() (BACnetNotificationParametersFloatingLimit, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.ReferenceValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referenceValue' not set"))
	}
	if b.StatusFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if b.SetPointValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'setPointValue' not set"))
	}
	if b.ErrorLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorLimit' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersFloatingLimit.deepCopy(), nil
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) MustBuild() BACnetNotificationParametersFloatingLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) Done() BACnetNotificationParametersBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersBuilder().(*_BACnetNotificationParametersBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersFloatingLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersFloatingLimitBuilder().(*_BACnetNotificationParametersFloatingLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersFloatingLimitBuilder creates a BACnetNotificationParametersFloatingLimitBuilder
func (b *_BACnetNotificationParametersFloatingLimit) CreateBACnetNotificationParametersFloatingLimitBuilder() BACnetNotificationParametersFloatingLimitBuilder {
	if b == nil {
		return NewBACnetNotificationParametersFloatingLimitBuilder()
	}
	return &_BACnetNotificationParametersFloatingLimitBuilder{_BACnetNotificationParametersFloatingLimit: b.deepCopy()}
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

func (m *_BACnetNotificationParametersFloatingLimit) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersFloatingLimit) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersFloatingLimit) GetReferenceValue() BACnetContextTagReal {
	return m.ReferenceValue
}

func (m *_BACnetNotificationParametersFloatingLimit) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersFloatingLimit) GetSetPointValue() BACnetContextTagReal {
	return m.SetPointValue
}

func (m *_BACnetNotificationParametersFloatingLimit) GetErrorLimit() BACnetContextTagReal {
	return m.ErrorLimit
}

func (m *_BACnetNotificationParametersFloatingLimit) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersFloatingLimit(structType any) BACnetNotificationParametersFloatingLimit {
	if casted, ok := structType.(BACnetNotificationParametersFloatingLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersFloatingLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersFloatingLimit) GetTypeName() string {
	return "BACnetNotificationParametersFloatingLimit"
}

func (m *_BACnetNotificationParametersFloatingLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (referenceValue)
	lengthInBits += m.ReferenceValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (setPointValue)
	lengthInBits += m.SetPointValue.GetLengthInBits(ctx)

	// Simple field (errorLimit)
	lengthInBits += m.ErrorLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersFloatingLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersFloatingLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersFloatingLimit BACnetNotificationParametersFloatingLimit, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersFloatingLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersFloatingLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	referenceValue, err := ReadSimpleField[BACnetContextTagReal](ctx, "referenceValue", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceValue' field"))
	}
	m.ReferenceValue = referenceValue

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	setPointValue, err := ReadSimpleField[BACnetContextTagReal](ctx, "setPointValue", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'setPointValue' field"))
	}
	m.SetPointValue = setPointValue

	errorLimit, err := ReadSimpleField[BACnetContextTagReal](ctx, "errorLimit", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorLimit' field"))
	}
	m.ErrorLimit = errorLimit

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersFloatingLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersFloatingLimit")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersFloatingLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersFloatingLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersFloatingLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersFloatingLimit")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "referenceValue", m.GetReferenceValue(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "setPointValue", m.GetSetPointValue(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'setPointValue' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "errorLimit", m.GetErrorLimit(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersFloatingLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersFloatingLimit")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersFloatingLimit) IsBACnetNotificationParametersFloatingLimit() {}

func (m *_BACnetNotificationParametersFloatingLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersFloatingLimit) deepCopy() *_BACnetNotificationParametersFloatingLimit {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersFloatingLimitCopy := &_BACnetNotificationParametersFloatingLimit{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.InnerOpeningTag),
		utils.DeepCopy[BACnetContextTagReal](m.ReferenceValue),
		utils.DeepCopy[BACnetStatusFlagsTagged](m.StatusFlags),
		utils.DeepCopy[BACnetContextTagReal](m.SetPointValue),
		utils.DeepCopy[BACnetContextTagReal](m.ErrorLimit),
		utils.DeepCopy[BACnetClosingTag](m.InnerClosingTag),
	}
	_BACnetNotificationParametersFloatingLimitCopy.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersFloatingLimitCopy
}

func (m *_BACnetNotificationParametersFloatingLimit) String() string {
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
