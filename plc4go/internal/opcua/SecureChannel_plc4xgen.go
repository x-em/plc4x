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

// Code generated by "plc4xGenerator -type=SecureChannel"; DO NOT EDIT.

package opcua

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *SecureChannel) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *SecureChannel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("SecureChannel"); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("sessionName", uint32(len(d.sessionName)*8), d.sessionName); err != nil {
		return err
	}
	if err := writeBuffer.WriteByteArray("clientNonce", d.clientNonce); err != nil {
		return err
	}

	if err := writeBuffer.WriteUint32("requestHandleGenerator", 32, d.requestHandleGenerator.Load()); err != nil {
		return err
	}

	if d.policyId != nil {
		if serializableField, ok := any(d.policyId).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("policyId"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("policyId"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.policyId)
			if err := writeBuffer.WriteString("policyId", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("tokenType", uint32(len(d.tokenType.String())*8), d.tokenType.String()); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("discovery", d.discovery); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("certFile", uint32(len(d.certFile)*8), d.certFile); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("keyStoreFile", uint32(len(d.keyStoreFile)*8), d.keyStoreFile); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.ckp)

		if err := writeBuffer.WriteString("ckp", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}

	if d.endpoint != nil {
		if serializableField, ok := any(d.endpoint).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("endpoint"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("endpoint"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.endpoint)
			if err := writeBuffer.WriteString("endpoint", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("username", uint32(len(d.username)*8), d.username); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("password", uint32(len(d.password)*8), d.password); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("securityPolicy", uint32(len(d.securityPolicy)*8), d.securityPolicy); err != nil {
		return err
	}

	if d.publicCertificate != nil {
		if serializableField, ok := any(d.publicCertificate).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("publicCertificate"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("publicCertificate"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.publicCertificate)
			if err := writeBuffer.WriteString("publicCertificate", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if d.thumbprint != nil {
		if serializableField, ok := any(d.thumbprint).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("thumbprint"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("thumbprint"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.thumbprint)
			if err := writeBuffer.WriteString("thumbprint", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteBit("isEncrypted", d.isEncrypted); err != nil {
		return err
	}
	if err := writeBuffer.WriteByteArray("senderCertificate", d.senderCertificate); err != nil {
		return err
	}
	if err := writeBuffer.WriteByteArray("senderNonce", d.senderNonce); err != nil {
		return err
	}

	if d.certificateThumbprint != nil {
		if serializableField, ok := any(d.certificateThumbprint).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("certificateThumbprint"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("certificateThumbprint"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.certificateThumbprint)
			if err := writeBuffer.WriteString("certificateThumbprint", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteBit("checkedEndpoints", d.checkedEndpoints); err != nil {
		return err
	}
	if d.encryptionHandler != nil {
		{
			_value := fmt.Sprintf("%v", d.encryptionHandler)

			if err := writeBuffer.WriteString("encryptionHandler", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("configuration", uint32(len(d.configuration.String())*8), d.configuration.String()); err != nil {
		return err
	}

	if err := writeBuffer.WriteUint32("channelId", 32, d.channelId.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteUint32("tokenId", 32, d.tokenId.Load()); err != nil {
		return err
	}

	if d.authenticationToken != nil {
		if serializableField, ok := any(d.authenticationToken).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("authenticationToken"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("authenticationToken"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.authenticationToken)
			if err := writeBuffer.WriteString("authenticationToken", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
	if d.codec != nil {
		{
			_value := fmt.Sprintf("%v", d.codec)

			if err := writeBuffer.WriteString("codec", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.channelTransactionManager != nil {
		{
			_value := fmt.Sprintf("%v", d.channelTransactionManager)

			if err := writeBuffer.WriteString("channelTransactionManager", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteUint32("lifetime", 32, d.lifetime); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("keepAliveIndicator", d.keepAliveIndicator.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt64("sendBufferSize", 64, int64(d.sendBufferSize)); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt64("maxMessageSize", 64, int64(d.maxMessageSize)); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("endpoints", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.endpoints {
		if err := writeBuffer.WriteString("", uint32(len(elem)*8), elem); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("endpoints", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt32("senderSequenceNumber", 32, d.senderSequenceNumber.Load()); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("SecureChannel"); err != nil {
		return err
	}
	return nil
}

func (d *SecureChannel) String() string {
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