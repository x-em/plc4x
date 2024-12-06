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

public class BACnetConfirmedServiceRequestVTClose extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.VT_CLOSE;
  }

  // Properties.
  protected final List<BACnetApplicationTagUnsignedInteger> listOfRemoteVtSessionIdentifiers;

  // Arguments.
  protected final Long serviceRequestPayloadLength;
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestVTClose(
      List<BACnetApplicationTagUnsignedInteger> listOfRemoteVtSessionIdentifiers,
      Long serviceRequestPayloadLength,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.listOfRemoteVtSessionIdentifiers = listOfRemoteVtSessionIdentifiers;
    this.serviceRequestPayloadLength = serviceRequestPayloadLength;
    this.serviceRequestLength = serviceRequestLength;
  }

  public List<BACnetApplicationTagUnsignedInteger> getListOfRemoteVtSessionIdentifiers() {
    return listOfRemoteVtSessionIdentifiers;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestVTClose");

    // Array Field (listOfRemoteVtSessionIdentifiers)
    writeComplexTypeArrayField(
        "listOfRemoteVtSessionIdentifiers", listOfRemoteVtSessionIdentifiers, writeBuffer);

    writeBuffer.popContext("BACnetConfirmedServiceRequestVTClose");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestVTClose _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (listOfRemoteVtSessionIdentifiers != null) {
      for (Message element : listOfRemoteVtSessionIdentifiers) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestPayloadLength, Long serviceRequestLength)
          throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestVTClose");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    List<BACnetApplicationTagUnsignedInteger> listOfRemoteVtSessionIdentifiers =
        readLengthArrayField(
            "listOfRemoteVtSessionIdentifiers",
            readComplex(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer),
            serviceRequestPayloadLength);

    readBuffer.closeContext("BACnetConfirmedServiceRequestVTClose");
    // Create the instance
    return new BACnetConfirmedServiceRequestVTCloseBuilderImpl(
        listOfRemoteVtSessionIdentifiers, serviceRequestPayloadLength, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestVTCloseBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final List<BACnetApplicationTagUnsignedInteger> listOfRemoteVtSessionIdentifiers;
    private final Long serviceRequestPayloadLength;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestVTCloseBuilderImpl(
        List<BACnetApplicationTagUnsignedInteger> listOfRemoteVtSessionIdentifiers,
        Long serviceRequestPayloadLength,
        Long serviceRequestLength) {
      this.listOfRemoteVtSessionIdentifiers = listOfRemoteVtSessionIdentifiers;
      this.serviceRequestPayloadLength = serviceRequestPayloadLength;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestVTClose build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestVTClose bACnetConfirmedServiceRequestVTClose =
          new BACnetConfirmedServiceRequestVTClose(
              listOfRemoteVtSessionIdentifiers, serviceRequestPayloadLength, serviceRequestLength);
      return bACnetConfirmedServiceRequestVTClose;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestVTClose)) {
      return false;
    }
    BACnetConfirmedServiceRequestVTClose that = (BACnetConfirmedServiceRequestVTClose) o;
    return (getListOfRemoteVtSessionIdentifiers() == that.getListOfRemoteVtSessionIdentifiers())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getListOfRemoteVtSessionIdentifiers());
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