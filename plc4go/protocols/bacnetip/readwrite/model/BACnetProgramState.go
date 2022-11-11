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

// BACnetProgramState is an enum
type BACnetProgramState uint8

type IBACnetProgramState interface {
	utils.Serializable
}

const (
	BACnetProgramState_IDLE      BACnetProgramState = 0
	BACnetProgramState_LOADING   BACnetProgramState = 1
	BACnetProgramState_RUNNING   BACnetProgramState = 2
	BACnetProgramState_WAITING   BACnetProgramState = 3
	BACnetProgramState_HALTED    BACnetProgramState = 4
	BACnetProgramState_UNLOADING BACnetProgramState = 5
)

var BACnetProgramStateValues []BACnetProgramState

func init() {
	_ = errors.New
	BACnetProgramStateValues = []BACnetProgramState{
		BACnetProgramState_IDLE,
		BACnetProgramState_LOADING,
		BACnetProgramState_RUNNING,
		BACnetProgramState_WAITING,
		BACnetProgramState_HALTED,
		BACnetProgramState_UNLOADING,
	}
}

func BACnetProgramStateByValue(value uint8) (enum BACnetProgramState, ok bool) {
	switch value {
	case 0:
		return BACnetProgramState_IDLE, true
	case 1:
		return BACnetProgramState_LOADING, true
	case 2:
		return BACnetProgramState_RUNNING, true
	case 3:
		return BACnetProgramState_WAITING, true
	case 4:
		return BACnetProgramState_HALTED, true
	case 5:
		return BACnetProgramState_UNLOADING, true
	}
	return 0, false
}

func BACnetProgramStateByName(value string) (enum BACnetProgramState, ok bool) {
	switch value {
	case "IDLE":
		return BACnetProgramState_IDLE, true
	case "LOADING":
		return BACnetProgramState_LOADING, true
	case "RUNNING":
		return BACnetProgramState_RUNNING, true
	case "WAITING":
		return BACnetProgramState_WAITING, true
	case "HALTED":
		return BACnetProgramState_HALTED, true
	case "UNLOADING":
		return BACnetProgramState_UNLOADING, true
	}
	return 0, false
}

func BACnetProgramStateKnows(value uint8) bool {
	for _, typeValue := range BACnetProgramStateValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetProgramState(structType interface{}) BACnetProgramState {
	castFunc := func(typ interface{}) BACnetProgramState {
		if sBACnetProgramState, ok := typ.(BACnetProgramState); ok {
			return sBACnetProgramState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetProgramState) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetProgramState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProgramStateParse(theBytes []byte) (BACnetProgramState, error) {
	return BACnetProgramStateParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetProgramStateParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetProgramState, error) {
	val, err := readBuffer.ReadUint8("BACnetProgramState", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetProgramState")
	}
	if enum, ok := BACnetProgramStateByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetProgramState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetProgramState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetProgramState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetProgramState", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetProgramState) PLC4XEnumName() string {
	switch e {
	case BACnetProgramState_IDLE:
		return "IDLE"
	case BACnetProgramState_LOADING:
		return "LOADING"
	case BACnetProgramState_RUNNING:
		return "RUNNING"
	case BACnetProgramState_WAITING:
		return "WAITING"
	case BACnetProgramState_HALTED:
		return "HALTED"
	case BACnetProgramState_UNLOADING:
		return "UNLOADING"
	}
	return ""
}

func (e BACnetProgramState) String() string {
	return e.PLC4XEnumName()
}
