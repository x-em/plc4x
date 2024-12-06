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

// BACnetBinaryLightingPV is an enum
type BACnetBinaryLightingPV uint8

type IBACnetBinaryLightingPV interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetBinaryLightingPV_OFF                      BACnetBinaryLightingPV = 0
	BACnetBinaryLightingPV_ON                       BACnetBinaryLightingPV = 1
	BACnetBinaryLightingPV_WARN                     BACnetBinaryLightingPV = 2
	BACnetBinaryLightingPV_WARN_OFF                 BACnetBinaryLightingPV = 3
	BACnetBinaryLightingPV_WARN_RELINQUISH          BACnetBinaryLightingPV = 4
	BACnetBinaryLightingPV_STOP                     BACnetBinaryLightingPV = 5
	BACnetBinaryLightingPV_VENDOR_PROPRIETARY_VALUE BACnetBinaryLightingPV = 0xFF
)

var BACnetBinaryLightingPVValues []BACnetBinaryLightingPV

func init() {
	_ = errors.New
	BACnetBinaryLightingPVValues = []BACnetBinaryLightingPV{
		BACnetBinaryLightingPV_OFF,
		BACnetBinaryLightingPV_ON,
		BACnetBinaryLightingPV_WARN,
		BACnetBinaryLightingPV_WARN_OFF,
		BACnetBinaryLightingPV_WARN_RELINQUISH,
		BACnetBinaryLightingPV_STOP,
		BACnetBinaryLightingPV_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetBinaryLightingPVByValue(value uint8) (enum BACnetBinaryLightingPV, ok bool) {
	switch value {
	case 0:
		return BACnetBinaryLightingPV_OFF, true
	case 0xFF:
		return BACnetBinaryLightingPV_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetBinaryLightingPV_ON, true
	case 2:
		return BACnetBinaryLightingPV_WARN, true
	case 3:
		return BACnetBinaryLightingPV_WARN_OFF, true
	case 4:
		return BACnetBinaryLightingPV_WARN_RELINQUISH, true
	case 5:
		return BACnetBinaryLightingPV_STOP, true
	}
	return 0, false
}

func BACnetBinaryLightingPVByName(value string) (enum BACnetBinaryLightingPV, ok bool) {
	switch value {
	case "OFF":
		return BACnetBinaryLightingPV_OFF, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetBinaryLightingPV_VENDOR_PROPRIETARY_VALUE, true
	case "ON":
		return BACnetBinaryLightingPV_ON, true
	case "WARN":
		return BACnetBinaryLightingPV_WARN, true
	case "WARN_OFF":
		return BACnetBinaryLightingPV_WARN_OFF, true
	case "WARN_RELINQUISH":
		return BACnetBinaryLightingPV_WARN_RELINQUISH, true
	case "STOP":
		return BACnetBinaryLightingPV_STOP, true
	}
	return 0, false
}

func BACnetBinaryLightingPVKnows(value uint8) bool {
	for _, typeValue := range BACnetBinaryLightingPVValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetBinaryLightingPV(structType any) BACnetBinaryLightingPV {
	castFunc := func(typ any) BACnetBinaryLightingPV {
		if sBACnetBinaryLightingPV, ok := typ.(BACnetBinaryLightingPV); ok {
			return sBACnetBinaryLightingPV
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetBinaryLightingPV) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetBinaryLightingPV) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetBinaryLightingPVParse(ctx context.Context, theBytes []byte) (BACnetBinaryLightingPV, error) {
	return BACnetBinaryLightingPVParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetBinaryLightingPVParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetBinaryLightingPV, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("BACnetBinaryLightingPV", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetBinaryLightingPV")
	}
	if enum, ok := BACnetBinaryLightingPVByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetBinaryLightingPV")
		return BACnetBinaryLightingPV(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetBinaryLightingPV) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetBinaryLightingPV) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("BACnetBinaryLightingPV", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e BACnetBinaryLightingPV) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetBinaryLightingPV) PLC4XEnumName() string {
	switch e {
	case BACnetBinaryLightingPV_OFF:
		return "OFF"
	case BACnetBinaryLightingPV_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetBinaryLightingPV_ON:
		return "ON"
	case BACnetBinaryLightingPV_WARN:
		return "WARN"
	case BACnetBinaryLightingPV_WARN_OFF:
		return "WARN_OFF"
	case BACnetBinaryLightingPV_WARN_RELINQUISH:
		return "WARN_RELINQUISH"
	case BACnetBinaryLightingPV_STOP:
		return "STOP"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetBinaryLightingPV) String() string {
	return e.PLC4XEnumName()
}