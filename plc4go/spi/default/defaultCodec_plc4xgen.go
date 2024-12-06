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

// Code generated by "plc4xGenerator -type=defaultCodec"; DO NOT EDIT.

package _default

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *defaultCodec) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *defaultCodec) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("defaultCodec"); err != nil {
		return err
	}

	if d.transportInstance != nil {
		if serializableField, ok := any(d.transportInstance).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("transportInstance"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("transportInstance"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.transportInstance)
			if err := writeBuffer.WriteString("transportInstance", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PushContext("expectations", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.expectations {
		var elem any = elem

		if elem != nil {
			if serializableField, ok := any(elem).(utils.Serializable); ok {
				if err := writeBuffer.PushContext("value"); err != nil {
					return err
				}
				if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
					return err
				}
				if err := writeBuffer.PopContext("value"); err != nil {
					return err
				}
			} else {
				stringValue := fmt.Sprintf("%v", elem)
				if err := writeBuffer.WriteString("value", uint32(len(stringValue)*8), stringValue); err != nil {
					return err
				}
			}
		}
	}
	if err := writeBuffer.PopContext("expectations", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	_defaultIncomingMessageChannel_plx4gen_description := fmt.Sprintf("%d element(s)", len(d.defaultIncomingMessageChannel))
	if err := writeBuffer.WriteString("defaultIncomingMessageChannel", uint32(len(_defaultIncomingMessageChannel_plx4gen_description)*8), _defaultIncomingMessageChannel_plx4gen_description); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("customMessageHandling", d.customMessageHandling != nil); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("running", d.running.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("receiveTimeout", uint32(len(fmt.Sprintf("%s", d.receiveTimeout))*8), fmt.Sprintf("%s", d.receiveTimeout)); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("traceDefaultMessageCodecWorker", d.traceDefaultMessageCodecWorker); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("defaultCodec"); err != nil {
		return err
	}
	return nil
}

func (d *defaultCodec) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	wb := utils.NewWriteBufferBoxBased(utils.WithWriteBufferBoxBasedMergeSingleBoxes(), utils.WithWriteBufferBoxBasedOmitEmptyBoxes())
	if err := wb.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
