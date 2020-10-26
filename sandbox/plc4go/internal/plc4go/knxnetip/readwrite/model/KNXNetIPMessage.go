//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
    "strconv"
)

// Constant values.
const KNXNetIPMessage_PROTOCOLVERSION uint8 = 0x10

// The data-structure of this message
type KNXNetIPMessage struct {

}

// The corresponding interface
type IKNXNetIPMessage interface {
    spi.Message
    MsgType() uint16
    Serialize(io utils.WriteBuffer) error
}

type KNXNetIPMessageInitializer interface {
    initialize() spi.Message
}

func KNXNetIPMessageMsgType(m IKNXNetIPMessage) uint16 {
    return m.MsgType()
}


func CastIKNXNetIPMessage(structType interface{}) IKNXNetIPMessage {
    castFunc := func(typ interface{}) IKNXNetIPMessage {
        if iKNXNetIPMessage, ok := typ.(IKNXNetIPMessage); ok {
            return iKNXNetIPMessage
        }
        return nil
    }
    return castFunc(structType)
}

func CastKNXNetIPMessage(structType interface{}) KNXNetIPMessage {
    castFunc := func(typ interface{}) KNXNetIPMessage {
        if sKNXNetIPMessage, ok := typ.(KNXNetIPMessage); ok {
            return sKNXNetIPMessage
        }
        if sKNXNetIPMessage, ok := typ.(*KNXNetIPMessage); ok {
            return *sKNXNetIPMessage
        }
        return KNXNetIPMessage{}
    }
    return castFunc(structType)
}

func (m KNXNetIPMessage) LengthInBits() uint16 {
    var lengthInBits uint16 = 0

    // Implicit Field (headerLength)
    lengthInBits += 8

    // Const Field (protocolVersion)
    lengthInBits += 8

    // Discriminator Field (msgType)
    lengthInBits += 16

    // Implicit Field (totalLength)
    lengthInBits += 16

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits
}

func (m KNXNetIPMessage) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func KNXNetIPMessageParse(io *utils.ReadBuffer) (spi.Message, error) {

    // Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    _, _headerLengthErr := io.ReadUint8(8)
    if _headerLengthErr != nil {
        return nil, errors.New("Error parsing 'headerLength' field " + _headerLengthErr.Error())
    }

    // Const Field (protocolVersion)
    protocolVersion, _protocolVersionErr := io.ReadUint8(8)
    if _protocolVersionErr != nil {
        return nil, errors.New("Error parsing 'protocolVersion' field " + _protocolVersionErr.Error())
    }
    if protocolVersion != KNXNetIPMessage_PROTOCOLVERSION {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(KNXNetIPMessage_PROTOCOLVERSION)) + " but got " + strconv.Itoa(int(protocolVersion)))
    }

    // Discriminator Field (msgType) (Used as input to a switch field)
    msgType, _msgTypeErr := io.ReadUint16(16)
    if _msgTypeErr != nil {
        return nil, errors.New("Error parsing 'msgType' field " + _msgTypeErr.Error())
    }

    // Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    totalLength, _totalLengthErr := io.ReadUint16(16)
    if _totalLengthErr != nil {
        return nil, errors.New("Error parsing 'totalLength' field " + _totalLengthErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var initializer KNXNetIPMessageInitializer
    var typeSwitchError error
    switch {
    case msgType == 0x0201:
        initializer, typeSwitchError = SearchRequestParse(io)
    case msgType == 0x0202:
        initializer, typeSwitchError = SearchResponseParse(io)
    case msgType == 0x0203:
        initializer, typeSwitchError = DescriptionRequestParse(io)
    case msgType == 0x0204:
        initializer, typeSwitchError = DescriptionResponseParse(io)
    case msgType == 0x0205:
        initializer, typeSwitchError = ConnectionRequestParse(io)
    case msgType == 0x0206:
        initializer, typeSwitchError = ConnectionResponseParse(io)
    case msgType == 0x0207:
        initializer, typeSwitchError = ConnectionStateRequestParse(io)
    case msgType == 0x0208:
        initializer, typeSwitchError = ConnectionStateResponseParse(io)
    case msgType == 0x0209:
        initializer, typeSwitchError = DisconnectRequestParse(io)
    case msgType == 0x020A:
        initializer, typeSwitchError = DisconnectResponseParse(io)
    case msgType == 0x020B:
        initializer, typeSwitchError = UnknownMessageParse(io, totalLength)
    case msgType == 0x0310:
        initializer, typeSwitchError = DeviceConfigurationRequestParse(io, totalLength)
    case msgType == 0x0311:
        initializer, typeSwitchError = DeviceConfigurationAckParse(io)
    case msgType == 0x0420:
        initializer, typeSwitchError = TunnelingRequestParse(io, totalLength)
    case msgType == 0x0421:
        initializer, typeSwitchError = TunnelingResponseParse(io)
    case msgType == 0x0530:
        initializer, typeSwitchError = RoutingIndicationParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Create the instance
    return initializer.initialize(), nil
}

func KNXNetIPMessageSerialize(io utils.WriteBuffer, m KNXNetIPMessage, i IKNXNetIPMessage, childSerialize func() error) error {

    // Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    headerLength := uint8(uint8(6))
    _headerLengthErr := io.WriteUint8(8, (headerLength))
    if _headerLengthErr != nil {
        return errors.New("Error serializing 'headerLength' field " + _headerLengthErr.Error())
    }

    // Const Field (protocolVersion)
    _protocolVersionErr := io.WriteUint8(8, 0x10)
    if _protocolVersionErr != nil {
        return errors.New("Error serializing 'protocolVersion' field " + _protocolVersionErr.Error())
    }

    // Discriminator Field (msgType) (Used as input to a switch field)
    msgType := uint16(i.MsgType())
    _msgTypeErr := io.WriteUint16(16, (msgType))
    if _msgTypeErr != nil {
        return errors.New("Error serializing 'msgType' field " + _msgTypeErr.Error())
    }

    // Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    totalLength := uint16(uint16(m.LengthInBytes()))
    _totalLengthErr := io.WriteUint16(16, (totalLength))
    if _totalLengthErr != nil {
        return errors.New("Error serializing 'totalLength' field " + _totalLengthErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := childSerialize()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *KNXNetIPMessage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
    }
}

func (m KNXNetIPMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.KNXNetIPMessage"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
