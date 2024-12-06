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

// AdsStampHeader is the corresponding interface of AdsStampHeader
type AdsStampHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() uint64
	// GetSamples returns Samples (property field)
	GetSamples() uint32
	// GetAdsNotificationSamples returns AdsNotificationSamples (property field)
	GetAdsNotificationSamples() []AdsNotificationSample
	// IsAdsStampHeader is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsStampHeader()
	// CreateBuilder creates a AdsStampHeaderBuilder
	CreateAdsStampHeaderBuilder() AdsStampHeaderBuilder
}

// _AdsStampHeader is the data-structure of this message
type _AdsStampHeader struct {
	Timestamp              uint64
	Samples                uint32
	AdsNotificationSamples []AdsNotificationSample
}

var _ AdsStampHeader = (*_AdsStampHeader)(nil)

// NewAdsStampHeader factory function for _AdsStampHeader
func NewAdsStampHeader(timestamp uint64, samples uint32, adsNotificationSamples []AdsNotificationSample) *_AdsStampHeader {
	return &_AdsStampHeader{Timestamp: timestamp, Samples: samples, AdsNotificationSamples: adsNotificationSamples}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsStampHeaderBuilder is a builder for AdsStampHeader
type AdsStampHeaderBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timestamp uint64, samples uint32, adsNotificationSamples []AdsNotificationSample) AdsStampHeaderBuilder
	// WithTimestamp adds Timestamp (property field)
	WithTimestamp(uint64) AdsStampHeaderBuilder
	// WithSamples adds Samples (property field)
	WithSamples(uint32) AdsStampHeaderBuilder
	// WithAdsNotificationSamples adds AdsNotificationSamples (property field)
	WithAdsNotificationSamples(...AdsNotificationSample) AdsStampHeaderBuilder
	// Build builds the AdsStampHeader or returns an error if something is wrong
	Build() (AdsStampHeader, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsStampHeader
}

// NewAdsStampHeaderBuilder() creates a AdsStampHeaderBuilder
func NewAdsStampHeaderBuilder() AdsStampHeaderBuilder {
	return &_AdsStampHeaderBuilder{_AdsStampHeader: new(_AdsStampHeader)}
}

type _AdsStampHeaderBuilder struct {
	*_AdsStampHeader

	err *utils.MultiError
}

var _ (AdsStampHeaderBuilder) = (*_AdsStampHeaderBuilder)(nil)

func (b *_AdsStampHeaderBuilder) WithMandatoryFields(timestamp uint64, samples uint32, adsNotificationSamples []AdsNotificationSample) AdsStampHeaderBuilder {
	return b.WithTimestamp(timestamp).WithSamples(samples).WithAdsNotificationSamples(adsNotificationSamples...)
}

func (b *_AdsStampHeaderBuilder) WithTimestamp(timestamp uint64) AdsStampHeaderBuilder {
	b.Timestamp = timestamp
	return b
}

func (b *_AdsStampHeaderBuilder) WithSamples(samples uint32) AdsStampHeaderBuilder {
	b.Samples = samples
	return b
}

func (b *_AdsStampHeaderBuilder) WithAdsNotificationSamples(adsNotificationSamples ...AdsNotificationSample) AdsStampHeaderBuilder {
	b.AdsNotificationSamples = adsNotificationSamples
	return b
}

func (b *_AdsStampHeaderBuilder) Build() (AdsStampHeader, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsStampHeader.deepCopy(), nil
}

func (b *_AdsStampHeaderBuilder) MustBuild() AdsStampHeader {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsStampHeaderBuilder) DeepCopy() any {
	_copy := b.CreateAdsStampHeaderBuilder().(*_AdsStampHeaderBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsStampHeaderBuilder creates a AdsStampHeaderBuilder
func (b *_AdsStampHeader) CreateAdsStampHeaderBuilder() AdsStampHeaderBuilder {
	if b == nil {
		return NewAdsStampHeaderBuilder()
	}
	return &_AdsStampHeaderBuilder{_AdsStampHeader: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsStampHeader) GetTimestamp() uint64 {
	return m.Timestamp
}

func (m *_AdsStampHeader) GetSamples() uint32 {
	return m.Samples
}

func (m *_AdsStampHeader) GetAdsNotificationSamples() []AdsNotificationSample {
	return m.AdsNotificationSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsStampHeader(structType any) AdsStampHeader {
	if casted, ok := structType.(AdsStampHeader); ok {
		return casted
	}
	if casted, ok := structType.(*AdsStampHeader); ok {
		return *casted
	}
	return nil
}

func (m *_AdsStampHeader) GetTypeName() string {
	return "AdsStampHeader"
}

func (m *_AdsStampHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (samples)
	lengthInBits += 32

	// Array field
	if len(m.AdsNotificationSamples) > 0 {
		for _curItem, element := range m.AdsNotificationSamples {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AdsNotificationSamples), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AdsStampHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsStampHeaderParse(ctx context.Context, theBytes []byte) (AdsStampHeader, error) {
	return AdsStampHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsStampHeaderParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsStampHeader, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsStampHeader, error) {
		return AdsStampHeaderParseWithBuffer(ctx, readBuffer)
	}
}

func AdsStampHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsStampHeader, error) {
	v, err := (&_AdsStampHeader{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AdsStampHeader) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__adsStampHeader AdsStampHeader, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsStampHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsStampHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField(ctx, "timestamp", ReadUnsignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	samples, err := ReadSimpleField(ctx, "samples", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'samples' field"))
	}
	m.Samples = samples

	adsNotificationSamples, err := ReadCountArrayField[AdsNotificationSample](ctx, "adsNotificationSamples", ReadComplex[AdsNotificationSample](AdsNotificationSampleParseWithBuffer, readBuffer), uint64(samples))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adsNotificationSamples' field"))
	}
	m.AdsNotificationSamples = adsNotificationSamples

	if closeErr := readBuffer.CloseContext("AdsStampHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsStampHeader")
	}

	return m, nil
}

func (m *_AdsStampHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsStampHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsStampHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsStampHeader")
	}

	if err := WriteSimpleField[uint64](ctx, "timestamp", m.GetTimestamp(), WriteUnsignedLong(writeBuffer, 64)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if err := WriteSimpleField[uint32](ctx, "samples", m.GetSamples(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'samples' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "adsNotificationSamples", m.GetAdsNotificationSamples(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'adsNotificationSamples' field")
	}

	if popErr := writeBuffer.PopContext("AdsStampHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsStampHeader")
	}
	return nil
}

func (m *_AdsStampHeader) IsAdsStampHeader() {}

func (m *_AdsStampHeader) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsStampHeader) deepCopy() *_AdsStampHeader {
	if m == nil {
		return nil
	}
	_AdsStampHeaderCopy := &_AdsStampHeader{
		m.Timestamp,
		m.Samples,
		utils.DeepCopySlice[AdsNotificationSample, AdsNotificationSample](m.AdsNotificationSamples),
	}
	return _AdsStampHeaderCopy
}

func (m *_AdsStampHeader) String() string {
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