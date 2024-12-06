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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// LightingCommandType is an enum
type LightingCommandType uint8

type ILightingCommandType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	LightingCommandType_OFF            LightingCommandType = 0x00
	LightingCommandType_ON             LightingCommandType = 0x01
	LightingCommandType_RAMP_TO_LEVEL  LightingCommandType = 0x02
	LightingCommandType_TERMINATE_RAMP LightingCommandType = 0x03
	LightingCommandType_LABEL          LightingCommandType = 0x04
)

var LightingCommandTypeValues []LightingCommandType

func init() {
	_ = errors.New
	LightingCommandTypeValues = []LightingCommandType{
		LightingCommandType_OFF,
		LightingCommandType_ON,
		LightingCommandType_RAMP_TO_LEVEL,
		LightingCommandType_TERMINATE_RAMP,
		LightingCommandType_LABEL,
	}
}

func (e LightingCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 1
		}
	case 0x01:
		{ /* '0x01' */
			return 1
		}
	case 0x02:
		{ /* '0x02' */
			return 2
		}
	case 0x03:
		{ /* '0x03' */
			return 1
		}
	case 0x04:
		{ /* '0x04' */
			return 4
		}
	default:
		{
			return 0
		}
	}
}

func LightingCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (enum LightingCommandType, ok bool) {
	for _, sizeValue := range LightingCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, true
		}
	}
	return 0, false
}
func LightingCommandTypeByValue(value uint8) (enum LightingCommandType, ok bool) {
	switch value {
	case 0x00:
		return LightingCommandType_OFF, true
	case 0x01:
		return LightingCommandType_ON, true
	case 0x02:
		return LightingCommandType_RAMP_TO_LEVEL, true
	case 0x03:
		return LightingCommandType_TERMINATE_RAMP, true
	case 0x04:
		return LightingCommandType_LABEL, true
	}
	return 0, false
}

func LightingCommandTypeByName(value string) (enum LightingCommandType, ok bool) {
	switch value {
	case "OFF":
		return LightingCommandType_OFF, true
	case "ON":
		return LightingCommandType_ON, true
	case "RAMP_TO_LEVEL":
		return LightingCommandType_RAMP_TO_LEVEL, true
	case "TERMINATE_RAMP":
		return LightingCommandType_TERMINATE_RAMP, true
	case "LABEL":
		return LightingCommandType_LABEL, true
	}
	return 0, false
}

func LightingCommandTypeKnows(value uint8) bool {
	for _, typeValue := range LightingCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastLightingCommandType(structType any) LightingCommandType {
	castFunc := func(typ any) LightingCommandType {
		if sLightingCommandType, ok := typ.(LightingCommandType); ok {
			return sLightingCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m LightingCommandType) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m LightingCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LightingCommandTypeParse(ctx context.Context, theBytes []byte) (LightingCommandType, error) {
	return LightingCommandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LightingCommandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LightingCommandType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("LightingCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LightingCommandType")
	}
	if enum, ok := LightingCommandTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for LightingCommandType")
		return LightingCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e LightingCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e LightingCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("LightingCommandType", 4, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e LightingCommandType) GetValue() uint8 {
	return uint8(e)
}

func (e LightingCommandType) GetNumberOfArguments() uint8 {
	return e.NumberOfArguments()
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LightingCommandType) PLC4XEnumName() string {
	switch e {
	case LightingCommandType_OFF:
		return "OFF"
	case LightingCommandType_ON:
		return "ON"
	case LightingCommandType_RAMP_TO_LEVEL:
		return "RAMP_TO_LEVEL"
	case LightingCommandType_TERMINATE_RAMP:
		return "TERMINATE_RAMP"
	case LightingCommandType_LABEL:
		return "LABEL"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e LightingCommandType) String() string {
	return e.PLC4XEnumName()
}