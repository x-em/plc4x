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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AdsSymbolTableEntry_NAMETERMINATOR uint8 = 0x00
const AdsSymbolTableEntry_DATATYPENAMETERMINATOR uint8 = 0x00
const AdsSymbolTableEntry_COMMENTTERMINATOR uint8 = 0x00

// AdsSymbolTableEntry is the corresponding interface of AdsSymbolTableEntry
type AdsSymbolTableEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetEntryLength returns EntryLength (property field)
	GetEntryLength() uint32
	// GetGroup returns Group (property field)
	GetGroup() uint32
	// GetOffset returns Offset (property field)
	GetOffset() uint32
	// GetSize returns Size (property field)
	GetSize() uint32
	// GetDataType returns DataType (property field)
	GetDataType() uint32
	// GetFlagMethodDeref returns FlagMethodDeref (property field)
	GetFlagMethodDeref() bool
	// GetFlagItfMethodAccess returns FlagItfMethodAccess (property field)
	GetFlagItfMethodAccess() bool
	// GetFlagReadOnly returns FlagReadOnly (property field)
	GetFlagReadOnly() bool
	// GetFlagTComInterfacePointer returns FlagTComInterfacePointer (property field)
	GetFlagTComInterfacePointer() bool
	// GetFlagTypeGuid returns FlagTypeGuid (property field)
	GetFlagTypeGuid() bool
	// GetFlagReferenceTo returns FlagReferenceTo (property field)
	GetFlagReferenceTo() bool
	// GetFlagBitValue returns FlagBitValue (property field)
	GetFlagBitValue() bool
	// GetFlagPersistent returns FlagPersistent (property field)
	GetFlagPersistent() bool
	// GetFlagExtendedFlags returns FlagExtendedFlags (property field)
	GetFlagExtendedFlags() bool
	// GetFlagInitOnReset returns FlagInitOnReset (property field)
	GetFlagInitOnReset() bool
	// GetFlagStatic returns FlagStatic (property field)
	GetFlagStatic() bool
	// GetFlagAttributes returns FlagAttributes (property field)
	GetFlagAttributes() bool
	// GetFlagContextMask returns FlagContextMask (property field)
	GetFlagContextMask() bool
	// GetName returns Name (property field)
	GetName() string
	// GetDataTypeName returns DataTypeName (property field)
	GetDataTypeName() string
	// GetComment returns Comment (property field)
	GetComment() string
	// GetRest returns Rest (property field)
	GetRest() []byte
	// IsAdsSymbolTableEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsSymbolTableEntry()
}

// _AdsSymbolTableEntry is the data-structure of this message
type _AdsSymbolTableEntry struct {
	EntryLength              uint32
	Group                    uint32
	Offset                   uint32
	Size                     uint32
	DataType                 uint32
	FlagMethodDeref          bool
	FlagItfMethodAccess      bool
	FlagReadOnly             bool
	FlagTComInterfacePointer bool
	FlagTypeGuid             bool
	FlagReferenceTo          bool
	FlagBitValue             bool
	FlagPersistent           bool
	FlagExtendedFlags        bool
	FlagInitOnReset          bool
	FlagStatic               bool
	FlagAttributes           bool
	FlagContextMask          bool
	Name                     string
	DataTypeName             string
	Comment                  string
	Rest                     []byte
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint16
}

var _ AdsSymbolTableEntry = (*_AdsSymbolTableEntry)(nil)

// NewAdsSymbolTableEntry factory function for _AdsSymbolTableEntry
func NewAdsSymbolTableEntry(entryLength uint32, group uint32, offset uint32, size uint32, dataType uint32, flagMethodDeref bool, flagItfMethodAccess bool, flagReadOnly bool, flagTComInterfacePointer bool, flagTypeGuid bool, flagReferenceTo bool, flagBitValue bool, flagPersistent bool, flagExtendedFlags bool, flagInitOnReset bool, flagStatic bool, flagAttributes bool, flagContextMask bool, name string, dataTypeName string, comment string, rest []byte) *_AdsSymbolTableEntry {
	return &_AdsSymbolTableEntry{EntryLength: entryLength, Group: group, Offset: offset, Size: size, DataType: dataType, FlagMethodDeref: flagMethodDeref, FlagItfMethodAccess: flagItfMethodAccess, FlagReadOnly: flagReadOnly, FlagTComInterfacePointer: flagTComInterfacePointer, FlagTypeGuid: flagTypeGuid, FlagReferenceTo: flagReferenceTo, FlagBitValue: flagBitValue, FlagPersistent: flagPersistent, FlagExtendedFlags: flagExtendedFlags, FlagInitOnReset: flagInitOnReset, FlagStatic: flagStatic, FlagAttributes: flagAttributes, FlagContextMask: flagContextMask, Name: name, DataTypeName: dataTypeName, Comment: comment, Rest: rest}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsSymbolTableEntry) GetEntryLength() uint32 {
	return m.EntryLength
}

func (m *_AdsSymbolTableEntry) GetGroup() uint32 {
	return m.Group
}

func (m *_AdsSymbolTableEntry) GetOffset() uint32 {
	return m.Offset
}

func (m *_AdsSymbolTableEntry) GetSize() uint32 {
	return m.Size
}

func (m *_AdsSymbolTableEntry) GetDataType() uint32 {
	return m.DataType
}

func (m *_AdsSymbolTableEntry) GetFlagMethodDeref() bool {
	return m.FlagMethodDeref
}

func (m *_AdsSymbolTableEntry) GetFlagItfMethodAccess() bool {
	return m.FlagItfMethodAccess
}

func (m *_AdsSymbolTableEntry) GetFlagReadOnly() bool {
	return m.FlagReadOnly
}

func (m *_AdsSymbolTableEntry) GetFlagTComInterfacePointer() bool {
	return m.FlagTComInterfacePointer
}

func (m *_AdsSymbolTableEntry) GetFlagTypeGuid() bool {
	return m.FlagTypeGuid
}

func (m *_AdsSymbolTableEntry) GetFlagReferenceTo() bool {
	return m.FlagReferenceTo
}

func (m *_AdsSymbolTableEntry) GetFlagBitValue() bool {
	return m.FlagBitValue
}

func (m *_AdsSymbolTableEntry) GetFlagPersistent() bool {
	return m.FlagPersistent
}

func (m *_AdsSymbolTableEntry) GetFlagExtendedFlags() bool {
	return m.FlagExtendedFlags
}

func (m *_AdsSymbolTableEntry) GetFlagInitOnReset() bool {
	return m.FlagInitOnReset
}

func (m *_AdsSymbolTableEntry) GetFlagStatic() bool {
	return m.FlagStatic
}

func (m *_AdsSymbolTableEntry) GetFlagAttributes() bool {
	return m.FlagAttributes
}

func (m *_AdsSymbolTableEntry) GetFlagContextMask() bool {
	return m.FlagContextMask
}

func (m *_AdsSymbolTableEntry) GetName() string {
	return m.Name
}

func (m *_AdsSymbolTableEntry) GetDataTypeName() string {
	return m.DataTypeName
}

func (m *_AdsSymbolTableEntry) GetComment() string {
	return m.Comment
}

func (m *_AdsSymbolTableEntry) GetRest() []byte {
	return m.Rest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsSymbolTableEntry) GetNameTerminator() uint8 {
	return AdsSymbolTableEntry_NAMETERMINATOR
}

func (m *_AdsSymbolTableEntry) GetDataTypeNameTerminator() uint8 {
	return AdsSymbolTableEntry_DATATYPENAMETERMINATOR
}

func (m *_AdsSymbolTableEntry) GetCommentTerminator() uint8 {
	return AdsSymbolTableEntry_COMMENTTERMINATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsSymbolTableEntry(structType any) AdsSymbolTableEntry {
	if casted, ok := structType.(AdsSymbolTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AdsSymbolTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AdsSymbolTableEntry) GetTypeName() string {
	return "AdsSymbolTableEntry"
}

func (m *_AdsSymbolTableEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (entryLength)
	lengthInBits += 32

	// Simple field (group)
	lengthInBits += 32

	// Simple field (offset)
	lengthInBits += 32

	// Simple field (size)
	lengthInBits += 32

	// Simple field (dataType)
	lengthInBits += 32

	// Simple field (flagMethodDeref)
	lengthInBits += 1

	// Simple field (flagItfMethodAccess)
	lengthInBits += 1

	// Simple field (flagReadOnly)
	lengthInBits += 1

	// Simple field (flagTComInterfacePointer)
	lengthInBits += 1

	// Simple field (flagTypeGuid)
	lengthInBits += 1

	// Simple field (flagReferenceTo)
	lengthInBits += 1

	// Simple field (flagBitValue)
	lengthInBits += 1

	// Simple field (flagPersistent)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 3

	// Simple field (flagExtendedFlags)
	lengthInBits += 1

	// Simple field (flagInitOnReset)
	lengthInBits += 1

	// Simple field (flagStatic)
	lengthInBits += 1

	// Simple field (flagAttributes)
	lengthInBits += 1

	// Simple field (flagContextMask)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (nameLength)
	lengthInBits += 16

	// Implicit Field (dataTypeNameLength)
	lengthInBits += 16

	// Implicit Field (commentLength)
	lengthInBits += 16

	// Simple field (name)
	lengthInBits += uint16(int32(uint16(len(m.GetName()))) * int32(int32(8)))

	// Const Field (nameTerminator)
	lengthInBits += 8

	// Simple field (dataTypeName)
	lengthInBits += uint16(int32(uint16(len(m.GetDataTypeName()))) * int32(int32(8)))

	// Const Field (dataTypeNameTerminator)
	lengthInBits += 8

	// Simple field (comment)
	lengthInBits += uint16(int32(uint16(len(m.GetComment()))) * int32(int32(8)))

	// Const Field (commentTerminator)
	lengthInBits += 8

	// Array field
	if len(m.Rest) > 0 {
		lengthInBits += 8 * uint16(len(m.Rest))
	}

	return lengthInBits
}

func (m *_AdsSymbolTableEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsSymbolTableEntryParse(ctx context.Context, theBytes []byte) (AdsSymbolTableEntry, error) {
	return AdsSymbolTableEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsSymbolTableEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsSymbolTableEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsSymbolTableEntry, error) {
		return AdsSymbolTableEntryParseWithBuffer(ctx, readBuffer)
	}
}

func AdsSymbolTableEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsSymbolTableEntry, error) {
	v, err := (&_AdsSymbolTableEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AdsSymbolTableEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__adsSymbolTableEntry AdsSymbolTableEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsSymbolTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsSymbolTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	var startPos = positionAware.GetPos()
	_ = startPos

	entryLength, err := ReadSimpleField(ctx, "entryLength", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entryLength' field"))
	}
	m.EntryLength = entryLength

	group, err := ReadSimpleField(ctx, "group", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	offset, err := ReadSimpleField(ctx, "offset", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'offset' field"))
	}
	m.Offset = offset

	size, err := ReadSimpleField(ctx, "size", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'size' field"))
	}
	m.Size = size

	dataType, err := ReadSimpleField(ctx, "dataType", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	flagMethodDeref, err := ReadSimpleField(ctx, "flagMethodDeref", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagMethodDeref' field"))
	}
	m.FlagMethodDeref = flagMethodDeref

	flagItfMethodAccess, err := ReadSimpleField(ctx, "flagItfMethodAccess", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagItfMethodAccess' field"))
	}
	m.FlagItfMethodAccess = flagItfMethodAccess

	flagReadOnly, err := ReadSimpleField(ctx, "flagReadOnly", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagReadOnly' field"))
	}
	m.FlagReadOnly = flagReadOnly

	flagTComInterfacePointer, err := ReadSimpleField(ctx, "flagTComInterfacePointer", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagTComInterfacePointer' field"))
	}
	m.FlagTComInterfacePointer = flagTComInterfacePointer

	flagTypeGuid, err := ReadSimpleField(ctx, "flagTypeGuid", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagTypeGuid' field"))
	}
	m.FlagTypeGuid = flagTypeGuid

	flagReferenceTo, err := ReadSimpleField(ctx, "flagReferenceTo", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagReferenceTo' field"))
	}
	m.FlagReferenceTo = flagReferenceTo

	flagBitValue, err := ReadSimpleField(ctx, "flagBitValue", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagBitValue' field"))
	}
	m.FlagBitValue = flagBitValue

	flagPersistent, err := ReadSimpleField(ctx, "flagPersistent", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagPersistent' field"))
	}
	m.FlagPersistent = flagPersistent

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(3)), uint8(0x00), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	flagExtendedFlags, err := ReadSimpleField(ctx, "flagExtendedFlags", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagExtendedFlags' field"))
	}
	m.FlagExtendedFlags = flagExtendedFlags

	flagInitOnReset, err := ReadSimpleField(ctx, "flagInitOnReset", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagInitOnReset' field"))
	}
	m.FlagInitOnReset = flagInitOnReset

	flagStatic, err := ReadSimpleField(ctx, "flagStatic", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagStatic' field"))
	}
	m.FlagStatic = flagStatic

	flagAttributes, err := ReadSimpleField(ctx, "flagAttributes", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagAttributes' field"))
	}
	m.FlagAttributes = flagAttributes

	flagContextMask, err := ReadSimpleField(ctx, "flagContextMask", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flagContextMask' field"))
	}
	m.FlagContextMask = flagContextMask

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	nameLength, err := ReadImplicitField[uint16](ctx, "nameLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nameLength' field"))
	}
	_ = nameLength

	dataTypeNameLength, err := ReadImplicitField[uint16](ctx, "dataTypeNameLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeNameLength' field"))
	}
	_ = dataTypeNameLength

	commentLength, err := ReadImplicitField[uint16](ctx, "commentLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commentLength' field"))
	}
	_ = commentLength

	name, err := ReadSimpleField(ctx, "name", ReadString(readBuffer, uint32(int32(nameLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	nameTerminator, err := ReadConstField[uint8](ctx, "nameTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsSymbolTableEntry_NAMETERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nameTerminator' field"))
	}
	_ = nameTerminator

	dataTypeName, err := ReadSimpleField(ctx, "dataTypeName", ReadString(readBuffer, uint32(int32(dataTypeNameLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeName' field"))
	}
	m.DataTypeName = dataTypeName

	dataTypeNameTerminator, err := ReadConstField[uint8](ctx, "dataTypeNameTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsSymbolTableEntry_DATATYPENAMETERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeNameTerminator' field"))
	}
	_ = dataTypeNameTerminator

	comment, err := ReadSimpleField(ctx, "comment", ReadString(readBuffer, uint32(int32(commentLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comment' field"))
	}
	m.Comment = comment

	commentTerminator, err := ReadConstField[uint8](ctx, "commentTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsSymbolTableEntry_COMMENTTERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commentTerminator' field"))
	}
	_ = commentTerminator

	rest, err := readBuffer.ReadByteArray("rest", int(int32(entryLength)-int32((positionAware.GetPos()-startPos))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rest' field"))
	}
	m.Rest = rest

	if closeErr := readBuffer.CloseContext("AdsSymbolTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsSymbolTableEntry")
	}

	return m, nil
}

func (m *_AdsSymbolTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsSymbolTableEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsSymbolTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsSymbolTableEntry")
	}

	if err := WriteSimpleField[uint32](ctx, "entryLength", m.GetEntryLength(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'entryLength' field")
	}

	if err := WriteSimpleField[uint32](ctx, "group", m.GetGroup(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'group' field")
	}

	if err := WriteSimpleField[uint32](ctx, "offset", m.GetOffset(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'offset' field")
	}

	if err := WriteSimpleField[uint32](ctx, "size", m.GetSize(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'size' field")
	}

	if err := WriteSimpleField[uint32](ctx, "dataType", m.GetDataType(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataType' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagMethodDeref", m.GetFlagMethodDeref(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagMethodDeref' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagItfMethodAccess", m.GetFlagItfMethodAccess(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagItfMethodAccess' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagReadOnly", m.GetFlagReadOnly(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagReadOnly' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagTComInterfacePointer", m.GetFlagTComInterfacePointer(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagTComInterfacePointer' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagTypeGuid", m.GetFlagTypeGuid(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagTypeGuid' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagReferenceTo", m.GetFlagReferenceTo(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagReferenceTo' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagBitValue", m.GetFlagBitValue(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagBitValue' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagPersistent", m.GetFlagPersistent(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagPersistent' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 3), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "flagExtendedFlags", m.GetFlagExtendedFlags(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagExtendedFlags' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagInitOnReset", m.GetFlagInitOnReset(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagInitOnReset' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagStatic", m.GetFlagStatic(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagStatic' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagAttributes", m.GetFlagAttributes(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagAttributes' field")
	}

	if err := WriteSimpleField[bool](ctx, "flagContextMask", m.GetFlagContextMask(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flagContextMask' field")
	}

	if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}
	nameLength := uint16(uint16(len(m.GetName())))
	if err := WriteImplicitField(ctx, "nameLength", nameLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'nameLength' field")
	}
	dataTypeNameLength := uint16(uint16(len(m.GetDataTypeName())))
	if err := WriteImplicitField(ctx, "dataTypeNameLength", dataTypeNameLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeNameLength' field")
	}
	commentLength := uint16(uint16(len(m.GetComment())))
	if err := WriteImplicitField(ctx, "commentLength", commentLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'commentLength' field")
	}

	if err := WriteSimpleField[string](ctx, "name", m.GetName(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetName())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'name' field")
	}

	if err := WriteConstField(ctx, "nameTerminator", AdsSymbolTableEntry_NAMETERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'nameTerminator' field")
	}

	if err := WriteSimpleField[string](ctx, "dataTypeName", m.GetDataTypeName(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetDataTypeName())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeName' field")
	}

	if err := WriteConstField(ctx, "dataTypeNameTerminator", AdsSymbolTableEntry_DATATYPENAMETERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeNameTerminator' field")
	}

	if err := WriteSimpleField[string](ctx, "comment", m.GetComment(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetComment())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'comment' field")
	}

	if err := WriteConstField(ctx, "commentTerminator", AdsSymbolTableEntry_COMMENTTERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'commentTerminator' field")
	}

	if err := WriteByteArrayField(ctx, "rest", m.GetRest(), WriteByteArray(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'rest' field")
	}

	if popErr := writeBuffer.PopContext("AdsSymbolTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsSymbolTableEntry")
	}
	return nil
}

func (m *_AdsSymbolTableEntry) IsAdsSymbolTableEntry() {}

func (m *_AdsSymbolTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
