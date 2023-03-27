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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ErrorReportingSystemCategoryTypeForOutputUnits is an enum
type ErrorReportingSystemCategoryTypeForOutputUnits uint8

type IErrorReportingSystemCategoryTypeForOutputUnits interface {
	utils.Serializable
}

const (
	ErrorReportingSystemCategoryTypeForOutputUnits_LE_MONOBLOCK_DIMMERS                        ErrorReportingSystemCategoryTypeForOutputUnits = 0x0
	ErrorReportingSystemCategoryTypeForOutputUnits_TE_MONOBLOCK_DIMMERS                        ErrorReportingSystemCategoryTypeForOutputUnits = 0x1
	ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_2                                  ErrorReportingSystemCategoryTypeForOutputUnits = 0x2
	ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_3                                  ErrorReportingSystemCategoryTypeForOutputUnits = 0x3
	ErrorReportingSystemCategoryTypeForOutputUnits_RELAYS_AND_OTHER_ON_OFF_SWITCHING_DEVICES   ErrorReportingSystemCategoryTypeForOutputUnits = 0x4
	ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_5                                  ErrorReportingSystemCategoryTypeForOutputUnits = 0x5
	ErrorReportingSystemCategoryTypeForOutputUnits_PWM_DIMMERS_INCLUDES_LED_CONTROL            ErrorReportingSystemCategoryTypeForOutputUnits = 0x6
	ErrorReportingSystemCategoryTypeForOutputUnits_SINEWAVE_MONOBLOCK_DIMMERS                  ErrorReportingSystemCategoryTypeForOutputUnits = 0x7
	ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_8                                  ErrorReportingSystemCategoryTypeForOutputUnits = 0x8
	ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_9                                  ErrorReportingSystemCategoryTypeForOutputUnits = 0x9
	ErrorReportingSystemCategoryTypeForOutputUnits_DALI_DSI_AND_OTHER_BALLAST_CONTROL_GATEWAYS ErrorReportingSystemCategoryTypeForOutputUnits = 0xA
	ErrorReportingSystemCategoryTypeForOutputUnits_MODULAR_DIMMERS                             ErrorReportingSystemCategoryTypeForOutputUnits = 0xB
	ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_12                                 ErrorReportingSystemCategoryTypeForOutputUnits = 0xC
	ErrorReportingSystemCategoryTypeForOutputUnits_UNIVERSAL_MONOBLOCK_DIMMERS                 ErrorReportingSystemCategoryTypeForOutputUnits = 0xD
	ErrorReportingSystemCategoryTypeForOutputUnits_DEVICE_CONTROLLERS_IR_RS_232_etc            ErrorReportingSystemCategoryTypeForOutputUnits = 0xE
	ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_15                                 ErrorReportingSystemCategoryTypeForOutputUnits = 0xF
)

var ErrorReportingSystemCategoryTypeForOutputUnitsValues []ErrorReportingSystemCategoryTypeForOutputUnits

func init() {
	_ = errors.New
	ErrorReportingSystemCategoryTypeForOutputUnitsValues = []ErrorReportingSystemCategoryTypeForOutputUnits{
		ErrorReportingSystemCategoryTypeForOutputUnits_LE_MONOBLOCK_DIMMERS,
		ErrorReportingSystemCategoryTypeForOutputUnits_TE_MONOBLOCK_DIMMERS,
		ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_2,
		ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_3,
		ErrorReportingSystemCategoryTypeForOutputUnits_RELAYS_AND_OTHER_ON_OFF_SWITCHING_DEVICES,
		ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_5,
		ErrorReportingSystemCategoryTypeForOutputUnits_PWM_DIMMERS_INCLUDES_LED_CONTROL,
		ErrorReportingSystemCategoryTypeForOutputUnits_SINEWAVE_MONOBLOCK_DIMMERS,
		ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_8,
		ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_9,
		ErrorReportingSystemCategoryTypeForOutputUnits_DALI_DSI_AND_OTHER_BALLAST_CONTROL_GATEWAYS,
		ErrorReportingSystemCategoryTypeForOutputUnits_MODULAR_DIMMERS,
		ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_12,
		ErrorReportingSystemCategoryTypeForOutputUnits_UNIVERSAL_MONOBLOCK_DIMMERS,
		ErrorReportingSystemCategoryTypeForOutputUnits_DEVICE_CONTROLLERS_IR_RS_232_etc,
		ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_15,
	}
}

func ErrorReportingSystemCategoryTypeForOutputUnitsByValue(value uint8) (enum ErrorReportingSystemCategoryTypeForOutputUnits, ok bool) {
	switch value {
	case 0x0:
		return ErrorReportingSystemCategoryTypeForOutputUnits_LE_MONOBLOCK_DIMMERS, true
	case 0x1:
		return ErrorReportingSystemCategoryTypeForOutputUnits_TE_MONOBLOCK_DIMMERS, true
	case 0x2:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_2, true
	case 0x3:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_3, true
	case 0x4:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RELAYS_AND_OTHER_ON_OFF_SWITCHING_DEVICES, true
	case 0x5:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_5, true
	case 0x6:
		return ErrorReportingSystemCategoryTypeForOutputUnits_PWM_DIMMERS_INCLUDES_LED_CONTROL, true
	case 0x7:
		return ErrorReportingSystemCategoryTypeForOutputUnits_SINEWAVE_MONOBLOCK_DIMMERS, true
	case 0x8:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_8, true
	case 0x9:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_9, true
	case 0xA:
		return ErrorReportingSystemCategoryTypeForOutputUnits_DALI_DSI_AND_OTHER_BALLAST_CONTROL_GATEWAYS, true
	case 0xB:
		return ErrorReportingSystemCategoryTypeForOutputUnits_MODULAR_DIMMERS, true
	case 0xC:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_12, true
	case 0xD:
		return ErrorReportingSystemCategoryTypeForOutputUnits_UNIVERSAL_MONOBLOCK_DIMMERS, true
	case 0xE:
		return ErrorReportingSystemCategoryTypeForOutputUnits_DEVICE_CONTROLLERS_IR_RS_232_etc, true
	case 0xF:
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForOutputUnitsByName(value string) (enum ErrorReportingSystemCategoryTypeForOutputUnits, ok bool) {
	switch value {
	case "LE_MONOBLOCK_DIMMERS":
		return ErrorReportingSystemCategoryTypeForOutputUnits_LE_MONOBLOCK_DIMMERS, true
	case "TE_MONOBLOCK_DIMMERS":
		return ErrorReportingSystemCategoryTypeForOutputUnits_TE_MONOBLOCK_DIMMERS, true
	case "RESERVED_2":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_3, true
	case "RELAYS_AND_OTHER_ON_OFF_SWITCHING_DEVICES":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RELAYS_AND_OTHER_ON_OFF_SWITCHING_DEVICES, true
	case "RESERVED_5":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_5, true
	case "PWM_DIMMERS_INCLUDES_LED_CONTROL":
		return ErrorReportingSystemCategoryTypeForOutputUnits_PWM_DIMMERS_INCLUDES_LED_CONTROL, true
	case "SINEWAVE_MONOBLOCK_DIMMERS":
		return ErrorReportingSystemCategoryTypeForOutputUnits_SINEWAVE_MONOBLOCK_DIMMERS, true
	case "RESERVED_8":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_8, true
	case "RESERVED_9":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_9, true
	case "DALI_DSI_AND_OTHER_BALLAST_CONTROL_GATEWAYS":
		return ErrorReportingSystemCategoryTypeForOutputUnits_DALI_DSI_AND_OTHER_BALLAST_CONTROL_GATEWAYS, true
	case "MODULAR_DIMMERS":
		return ErrorReportingSystemCategoryTypeForOutputUnits_MODULAR_DIMMERS, true
	case "RESERVED_12":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_12, true
	case "UNIVERSAL_MONOBLOCK_DIMMERS":
		return ErrorReportingSystemCategoryTypeForOutputUnits_UNIVERSAL_MONOBLOCK_DIMMERS, true
	case "DEVICE_CONTROLLERS_IR_RS_232_etc":
		return ErrorReportingSystemCategoryTypeForOutputUnits_DEVICE_CONTROLLERS_IR_RS_232_etc, true
	case "RESERVED_15":
		return ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForOutputUnitsKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingSystemCategoryTypeForOutputUnitsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingSystemCategoryTypeForOutputUnits(structType interface{}) ErrorReportingSystemCategoryTypeForOutputUnits {
	castFunc := func(typ interface{}) ErrorReportingSystemCategoryTypeForOutputUnits {
		if sErrorReportingSystemCategoryTypeForOutputUnits, ok := typ.(ErrorReportingSystemCategoryTypeForOutputUnits); ok {
			return sErrorReportingSystemCategoryTypeForOutputUnits
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSystemCategoryTypeForOutputUnits) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m ErrorReportingSystemCategoryTypeForOutputUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeForOutputUnitsParse(ctx context.Context, theBytes []byte) (ErrorReportingSystemCategoryTypeForOutputUnits, error) {
	return ErrorReportingSystemCategoryTypeForOutputUnitsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryTypeForOutputUnitsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategoryTypeForOutputUnits, error) {
	val, err := readBuffer.ReadUint8("ErrorReportingSystemCategoryTypeForOutputUnits", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSystemCategoryTypeForOutputUnits")
	}
	if enum, ok := ErrorReportingSystemCategoryTypeForOutputUnitsByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorReportingSystemCategoryTypeForOutputUnits(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSystemCategoryTypeForOutputUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingSystemCategoryTypeForOutputUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ErrorReportingSystemCategoryTypeForOutputUnits", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSystemCategoryTypeForOutputUnits) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSystemCategoryTypeForOutputUnits_LE_MONOBLOCK_DIMMERS:
		return "LE_MONOBLOCK_DIMMERS"
	case ErrorReportingSystemCategoryTypeForOutputUnits_TE_MONOBLOCK_DIMMERS:
		return "TE_MONOBLOCK_DIMMERS"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_3:
		return "RESERVED_3"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RELAYS_AND_OTHER_ON_OFF_SWITCHING_DEVICES:
		return "RELAYS_AND_OTHER_ON_OFF_SWITCHING_DEVICES"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_5:
		return "RESERVED_5"
	case ErrorReportingSystemCategoryTypeForOutputUnits_PWM_DIMMERS_INCLUDES_LED_CONTROL:
		return "PWM_DIMMERS_INCLUDES_LED_CONTROL"
	case ErrorReportingSystemCategoryTypeForOutputUnits_SINEWAVE_MONOBLOCK_DIMMERS:
		return "SINEWAVE_MONOBLOCK_DIMMERS"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_8:
		return "RESERVED_8"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_9:
		return "RESERVED_9"
	case ErrorReportingSystemCategoryTypeForOutputUnits_DALI_DSI_AND_OTHER_BALLAST_CONTROL_GATEWAYS:
		return "DALI_DSI_AND_OTHER_BALLAST_CONTROL_GATEWAYS"
	case ErrorReportingSystemCategoryTypeForOutputUnits_MODULAR_DIMMERS:
		return "MODULAR_DIMMERS"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_12:
		return "RESERVED_12"
	case ErrorReportingSystemCategoryTypeForOutputUnits_UNIVERSAL_MONOBLOCK_DIMMERS:
		return "UNIVERSAL_MONOBLOCK_DIMMERS"
	case ErrorReportingSystemCategoryTypeForOutputUnits_DEVICE_CONTROLLERS_IR_RS_232_etc:
		return "DEVICE_CONTROLLERS_IR_RS_232_etc"
	case ErrorReportingSystemCategoryTypeForOutputUnits_RESERVED_15:
		return "RESERVED_15"
	}
	return ""
}

func (e ErrorReportingSystemCategoryTypeForOutputUnits) String() string {
	return e.PLC4XEnumName()
}
