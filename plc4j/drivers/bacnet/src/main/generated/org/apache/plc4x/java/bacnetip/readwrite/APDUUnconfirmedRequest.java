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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class APDUUnconfirmedRequest extends APDU implements Message {

  // Accessors for discriminator values.
  public ApduType getApduType() {
    return ApduType.UNCONFIRMED_REQUEST_PDU;
  }

  // Properties.
  protected final BACnetUnconfirmedServiceRequest serviceRequest;

  // Arguments.
  protected final Integer apduLength;
  // Reserved Fields
  private Byte reservedField0;

  public APDUUnconfirmedRequest(
      BACnetUnconfirmedServiceRequest serviceRequest, Integer apduLength) {
    super(apduLength);
    this.serviceRequest = serviceRequest;
    this.apduLength = apduLength;
  }

  public BACnetUnconfirmedServiceRequest getServiceRequest() {
    return serviceRequest;
  }

  @Override
  protected void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("APDUUnconfirmedRequest");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0,
        writeUnsignedByte(writeBuffer, 4));

    // Simple Field (serviceRequest)
    writeSimpleField("serviceRequest", serviceRequest, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("APDUUnconfirmedRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    APDUUnconfirmedRequest _value = this;

    // Reserved Field (reserved)
    lengthInBits += 4;

    // Simple field (serviceRequest)
    lengthInBits += serviceRequest.getLengthInBits();

    return lengthInBits;
  }

  public static APDUBuilder staticParseAPDUBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("APDUUnconfirmedRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Byte reservedField0 = readReservedField("reserved", readUnsignedByte(readBuffer, 4), (byte) 0);

    BACnetUnconfirmedServiceRequest serviceRequest =
        readSimpleField(
            "serviceRequest",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetUnconfirmedServiceRequest.staticParse(
                        readBuffer, (int) ((apduLength) - (1))),
                readBuffer));

    readBuffer.closeContext("APDUUnconfirmedRequest");
    // Create the instance
    return new APDUUnconfirmedRequestBuilderImpl(serviceRequest, apduLength, reservedField0);
  }

  public static class APDUUnconfirmedRequestBuilderImpl implements APDU.APDUBuilder {
    private final BACnetUnconfirmedServiceRequest serviceRequest;
    private final Integer apduLength;
    private final Byte reservedField0;

    public APDUUnconfirmedRequestBuilderImpl(
        BACnetUnconfirmedServiceRequest serviceRequest, Integer apduLength, Byte reservedField0) {
      this.serviceRequest = serviceRequest;
      this.apduLength = apduLength;
      this.reservedField0 = reservedField0;
    }

    public APDUUnconfirmedRequest build(Integer apduLength) {

      APDUUnconfirmedRequest aPDUUnconfirmedRequest =
          new APDUUnconfirmedRequest(serviceRequest, apduLength);
      aPDUUnconfirmedRequest.reservedField0 = reservedField0;
      return aPDUUnconfirmedRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDUUnconfirmedRequest)) {
      return false;
    }
    APDUUnconfirmedRequest that = (APDUUnconfirmedRequest) o;
    return (getServiceRequest() == that.getServiceRequest()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getServiceRequest());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
