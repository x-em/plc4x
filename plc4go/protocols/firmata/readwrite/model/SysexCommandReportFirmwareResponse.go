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

// SysexCommandReportFirmwareResponse is the corresponding interface of SysexCommandReportFirmwareResponse
type SysexCommandReportFirmwareResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint8
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint8
	// GetFileName returns FileName (property field)
	GetFileName() []byte
	// IsSysexCommandReportFirmwareResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandReportFirmwareResponse()
	// CreateBuilder creates a SysexCommandReportFirmwareResponseBuilder
	CreateSysexCommandReportFirmwareResponseBuilder() SysexCommandReportFirmwareResponseBuilder
}

// _SysexCommandReportFirmwareResponse is the data-structure of this message
type _SysexCommandReportFirmwareResponse struct {
	SysexCommandContract
	MajorVersion uint8
	MinorVersion uint8
	FileName     []byte
}

var _ SysexCommandReportFirmwareResponse = (*_SysexCommandReportFirmwareResponse)(nil)
var _ SysexCommandRequirements = (*_SysexCommandReportFirmwareResponse)(nil)

// NewSysexCommandReportFirmwareResponse factory function for _SysexCommandReportFirmwareResponse
func NewSysexCommandReportFirmwareResponse(majorVersion uint8, minorVersion uint8, fileName []byte) *_SysexCommandReportFirmwareResponse {
	_result := &_SysexCommandReportFirmwareResponse{
		SysexCommandContract: NewSysexCommand(),
		MajorVersion:         majorVersion,
		MinorVersion:         minorVersion,
		FileName:             fileName,
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandReportFirmwareResponseBuilder is a builder for SysexCommandReportFirmwareResponse
type SysexCommandReportFirmwareResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(majorVersion uint8, minorVersion uint8, fileName []byte) SysexCommandReportFirmwareResponseBuilder
	// WithMajorVersion adds MajorVersion (property field)
	WithMajorVersion(uint8) SysexCommandReportFirmwareResponseBuilder
	// WithMinorVersion adds MinorVersion (property field)
	WithMinorVersion(uint8) SysexCommandReportFirmwareResponseBuilder
	// WithFileName adds FileName (property field)
	WithFileName(...byte) SysexCommandReportFirmwareResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SysexCommandBuilder
	// Build builds the SysexCommandReportFirmwareResponse or returns an error if something is wrong
	Build() (SysexCommandReportFirmwareResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommandReportFirmwareResponse
}

// NewSysexCommandReportFirmwareResponseBuilder() creates a SysexCommandReportFirmwareResponseBuilder
func NewSysexCommandReportFirmwareResponseBuilder() SysexCommandReportFirmwareResponseBuilder {
	return &_SysexCommandReportFirmwareResponseBuilder{_SysexCommandReportFirmwareResponse: new(_SysexCommandReportFirmwareResponse)}
}

type _SysexCommandReportFirmwareResponseBuilder struct {
	*_SysexCommandReportFirmwareResponse

	parentBuilder *_SysexCommandBuilder

	err *utils.MultiError
}

var _ (SysexCommandReportFirmwareResponseBuilder) = (*_SysexCommandReportFirmwareResponseBuilder)(nil)

func (b *_SysexCommandReportFirmwareResponseBuilder) setParent(contract SysexCommandContract) {
	b.SysexCommandContract = contract
	contract.(*_SysexCommand)._SubType = b._SysexCommandReportFirmwareResponse
}

func (b *_SysexCommandReportFirmwareResponseBuilder) WithMandatoryFields(majorVersion uint8, minorVersion uint8, fileName []byte) SysexCommandReportFirmwareResponseBuilder {
	return b.WithMajorVersion(majorVersion).WithMinorVersion(minorVersion).WithFileName(fileName...)
}

func (b *_SysexCommandReportFirmwareResponseBuilder) WithMajorVersion(majorVersion uint8) SysexCommandReportFirmwareResponseBuilder {
	b.MajorVersion = majorVersion
	return b
}

func (b *_SysexCommandReportFirmwareResponseBuilder) WithMinorVersion(minorVersion uint8) SysexCommandReportFirmwareResponseBuilder {
	b.MinorVersion = minorVersion
	return b
}

func (b *_SysexCommandReportFirmwareResponseBuilder) WithFileName(fileName ...byte) SysexCommandReportFirmwareResponseBuilder {
	b.FileName = fileName
	return b
}

func (b *_SysexCommandReportFirmwareResponseBuilder) Build() (SysexCommandReportFirmwareResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SysexCommandReportFirmwareResponse.deepCopy(), nil
}

func (b *_SysexCommandReportFirmwareResponseBuilder) MustBuild() SysexCommandReportFirmwareResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SysexCommandReportFirmwareResponseBuilder) Done() SysexCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSysexCommandBuilder().(*_SysexCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_SysexCommandReportFirmwareResponseBuilder) buildForSysexCommand() (SysexCommand, error) {
	return b.Build()
}

func (b *_SysexCommandReportFirmwareResponseBuilder) DeepCopy() any {
	_copy := b.CreateSysexCommandReportFirmwareResponseBuilder().(*_SysexCommandReportFirmwareResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSysexCommandReportFirmwareResponseBuilder creates a SysexCommandReportFirmwareResponseBuilder
func (b *_SysexCommandReportFirmwareResponse) CreateSysexCommandReportFirmwareResponseBuilder() SysexCommandReportFirmwareResponseBuilder {
	if b == nil {
		return NewSysexCommandReportFirmwareResponseBuilder()
	}
	return &_SysexCommandReportFirmwareResponseBuilder{_SysexCommandReportFirmwareResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandReportFirmwareResponse) GetCommandType() uint8 {
	return 0x79
}

func (m *_SysexCommandReportFirmwareResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandReportFirmwareResponse) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandReportFirmwareResponse) GetMajorVersion() uint8 {
	return m.MajorVersion
}

func (m *_SysexCommandReportFirmwareResponse) GetMinorVersion() uint8 {
	return m.MinorVersion
}

func (m *_SysexCommandReportFirmwareResponse) GetFileName() []byte {
	return m.FileName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSysexCommandReportFirmwareResponse(structType any) SysexCommandReportFirmwareResponse {
	if casted, ok := structType.(SysexCommandReportFirmwareResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandReportFirmwareResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandReportFirmwareResponse) GetTypeName() string {
	return "SysexCommandReportFirmwareResponse"
}

func (m *_SysexCommandReportFirmwareResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Manual Array Field (fileName)
	lengthInBits += uint16(LengthSysexString(ctx, m.GetFileName()))

	return lengthInBits
}

func (m *_SysexCommandReportFirmwareResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandReportFirmwareResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandReportFirmwareResponse SysexCommandReportFirmwareResponse, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandReportFirmwareResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandReportFirmwareResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	majorVersion, err := ReadSimpleField(ctx, "majorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'majorVersion' field"))
	}
	m.MajorVersion = majorVersion

	minorVersion, err := ReadSimpleField(ctx, "minorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minorVersion' field"))
	}
	m.MinorVersion = minorVersion

	fileName, err := ReadManualByteArrayField(ctx, "fileName", readBuffer, IsSysexEnd(ctx, readBuffer), ParseSysexString(ctx, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileName' field"))
	}
	m.FileName = fileName

	if closeErr := readBuffer.CloseContext("SysexCommandReportFirmwareResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandReportFirmwareResponse")
	}

	return m, nil
}

func (m *_SysexCommandReportFirmwareResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandReportFirmwareResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandReportFirmwareResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandReportFirmwareResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "majorVersion", m.GetMajorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'majorVersion' field")
		}

		if err := WriteSimpleField[uint8](ctx, "minorVersion", m.GetMinorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'minorVersion' field")
		}

		if err := WriteManualArrayField[byte](ctx, "fileName", m.GetFileName(), func(ctx context.Context, writeBuffer utils.WriteBuffer, m byte) error {
			return SerializeSysexString(ctx, writeBuffer, m)
		}, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'fileName' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandReportFirmwareResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandReportFirmwareResponse")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandReportFirmwareResponse) IsSysexCommandReportFirmwareResponse() {}

func (m *_SysexCommandReportFirmwareResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandReportFirmwareResponse) deepCopy() *_SysexCommandReportFirmwareResponse {
	if m == nil {
		return nil
	}
	_SysexCommandReportFirmwareResponseCopy := &_SysexCommandReportFirmwareResponse{
		m.SysexCommandContract.(*_SysexCommand).deepCopy(),
		m.MajorVersion,
		m.MinorVersion,
		utils.DeepCopySlice[byte, byte](m.FileName),
	}
	_SysexCommandReportFirmwareResponseCopy.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandReportFirmwareResponseCopy
}

func (m *_SysexCommandReportFirmwareResponse) String() string {
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
