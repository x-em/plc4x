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

// Code generated by "plc4xGenerator -type=DefaultPlcSubscriptionEventItem"; DO NOT EDIT.

package model

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *DefaultPlcSubscriptionEventItem) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DefaultPlcSubscriptionEventItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("PlcSubscriptionEventItem"); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("code", uint32(len(d.code.String())*8), d.code.String()); err != nil {
		return err
	}

	if d.tag != nil {
		if serializableField, ok := any(d.tag).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("tag"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("tag"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.tag)
			if err := writeBuffer.WriteString("tag", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("subscriptionType", uint32(len(d.subscriptionType.String())*8), d.subscriptionType.String()); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("interval", uint32(len(fmt.Sprintf("%s", d.interval))*8), fmt.Sprintf("%s", d.interval)); err != nil {
		return err
	}

	if d.value != nil {
		if serializableField, ok := any(d.value).(utils.Serializable); ok {
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
			stringValue := fmt.Sprintf("%v", d.value)
			if err := writeBuffer.WriteString("value", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("PlcSubscriptionEventItem"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcSubscriptionEventItem) String() string {
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