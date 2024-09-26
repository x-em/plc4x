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

// ErrorEnclosed is the corresponding interface of ErrorEnclosed
type ErrorEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetError returns Error (property field)
	GetError() Error
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsErrorEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorEnclosed()
}

// _ErrorEnclosed is the data-structure of this message
type _ErrorEnclosed struct {
	OpeningTag BACnetOpeningTag
	Error      Error
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ ErrorEnclosed = (*_ErrorEnclosed)(nil)

// NewErrorEnclosed factory function for _ErrorEnclosed
func NewErrorEnclosed(openingTag BACnetOpeningTag, error Error, closingTag BACnetClosingTag, tagNumber uint8) *_ErrorEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for ErrorEnclosed must not be nil")
	}
	if error == nil {
		panic("error of type Error for ErrorEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for ErrorEnclosed must not be nil")
	}
	return &_ErrorEnclosed{OpeningTag: openingTag, Error: error, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_ErrorEnclosed) GetError() Error {
	return m.Error
}

func (m *_ErrorEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorEnclosed(structType any) ErrorEnclosed {
	if casted, ok := structType.(ErrorEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorEnclosed) GetTypeName() string {
	return "ErrorEnclosed"
}

func (m *_ErrorEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ErrorEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (ErrorEnclosed, error) {
	return ErrorEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func ErrorEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorEnclosed, error) {
		return ErrorEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func ErrorEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (ErrorEnclosed, error) {
	v, err := (&_ErrorEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ErrorEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__errorEnclosed ErrorEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	error, err := ReadSimpleField[Error](ctx, "error", ReadComplex[Error](ErrorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("ErrorEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorEnclosed")
	}

	return m, nil
}

func (m *_ErrorEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[Error](ctx, "error", m.GetError(), WriteComplex[Error](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'error' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("ErrorEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_ErrorEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_ErrorEnclosed) IsErrorEnclosed() {}

func (m *_ErrorEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
