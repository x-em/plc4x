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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetSecurityLevel is an enum
type BACnetSecurityLevel uint8

type IBACnetSecurityLevel interface {
	utils.Serializable
}

const (
	BACnetSecurityLevel_INCAPABLE            BACnetSecurityLevel = 0
	BACnetSecurityLevel_PLAIN                BACnetSecurityLevel = 1
	BACnetSecurityLevel_SIGNED               BACnetSecurityLevel = 2
	BACnetSecurityLevel_ENCRYPTED            BACnetSecurityLevel = 3
	BACnetSecurityLevel_SIGNED_END_TO_END    BACnetSecurityLevel = 4
	BACnetSecurityLevel_ENCRYPTED_END_TO_END BACnetSecurityLevel = 5
)

var BACnetSecurityLevelValues []BACnetSecurityLevel

func init() {
	_ = errors.New
	BACnetSecurityLevelValues = []BACnetSecurityLevel{
		BACnetSecurityLevel_INCAPABLE,
		BACnetSecurityLevel_PLAIN,
		BACnetSecurityLevel_SIGNED,
		BACnetSecurityLevel_ENCRYPTED,
		BACnetSecurityLevel_SIGNED_END_TO_END,
		BACnetSecurityLevel_ENCRYPTED_END_TO_END,
	}
}

func BACnetSecurityLevelByValue(value uint8) (enum BACnetSecurityLevel, ok bool) {
	switch value {
	case 0:
		return BACnetSecurityLevel_INCAPABLE, true
	case 1:
		return BACnetSecurityLevel_PLAIN, true
	case 2:
		return BACnetSecurityLevel_SIGNED, true
	case 3:
		return BACnetSecurityLevel_ENCRYPTED, true
	case 4:
		return BACnetSecurityLevel_SIGNED_END_TO_END, true
	case 5:
		return BACnetSecurityLevel_ENCRYPTED_END_TO_END, true
	}
	return 0, false
}

func BACnetSecurityLevelByName(value string) (enum BACnetSecurityLevel, ok bool) {
	switch value {
	case "INCAPABLE":
		return BACnetSecurityLevel_INCAPABLE, true
	case "PLAIN":
		return BACnetSecurityLevel_PLAIN, true
	case "SIGNED":
		return BACnetSecurityLevel_SIGNED, true
	case "ENCRYPTED":
		return BACnetSecurityLevel_ENCRYPTED, true
	case "SIGNED_END_TO_END":
		return BACnetSecurityLevel_SIGNED_END_TO_END, true
	case "ENCRYPTED_END_TO_END":
		return BACnetSecurityLevel_ENCRYPTED_END_TO_END, true
	}
	return 0, false
}

func BACnetSecurityLevelKnows(value uint8) bool {
	for _, typeValue := range BACnetSecurityLevelValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetSecurityLevel(structType interface{}) BACnetSecurityLevel {
	castFunc := func(typ interface{}) BACnetSecurityLevel {
		if sBACnetSecurityLevel, ok := typ.(BACnetSecurityLevel); ok {
			return sBACnetSecurityLevel
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetSecurityLevel) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetSecurityLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSecurityLevelParse(theBytes []byte) (BACnetSecurityLevel, error) {
	return BACnetSecurityLevelParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetSecurityLevelParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetSecurityLevel, error) {
	val, err := readBuffer.ReadUint8("BACnetSecurityLevel", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetSecurityLevel")
	}
	if enum, ok := BACnetSecurityLevelByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetSecurityLevel(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetSecurityLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetSecurityLevel) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetSecurityLevel", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetSecurityLevel) PLC4XEnumName() string {
	switch e {
	case BACnetSecurityLevel_INCAPABLE:
		return "INCAPABLE"
	case BACnetSecurityLevel_PLAIN:
		return "PLAIN"
	case BACnetSecurityLevel_SIGNED:
		return "SIGNED"
	case BACnetSecurityLevel_ENCRYPTED:
		return "ENCRYPTED"
	case BACnetSecurityLevel_SIGNED_END_TO_END:
		return "SIGNED_END_TO_END"
	case BACnetSecurityLevel_ENCRYPTED_END_TO_END:
		return "ENCRYPTED_END_TO_END"
	}
	return ""
}

func (e BACnetSecurityLevel) String() string {
	return e.PLC4XEnumName()
}
