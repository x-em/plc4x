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
package org.apache.plc4x.java.openprotocol.readwrite;

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

public class OpenProtocolMessageParameterSetSelectedUnsubscribeRev1
    extends OpenProtocolMessageParameterSetSelectedUnsubscribe implements Message {

  // Accessors for discriminator values.
  public Long getRevision() {
    return (long) 1;
  }

  public OpenProtocolMessageParameterSetSelectedUnsubscribeRev1(
      Long midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
  }

  @Override
  protected void serializeOpenProtocolMessageParameterSetSelectedUnsubscribeChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessageParameterSetSelectedUnsubscribeRev1");

    writeBuffer.popContext("OpenProtocolMessageParameterSetSelectedUnsubscribeRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageParameterSetSelectedUnsubscribeRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    return lengthInBits;
  }

  public static OpenProtocolMessageParameterSetSelectedUnsubscribeBuilder
      staticParseOpenProtocolMessageParameterSetSelectedUnsubscribeBuilder(
          ReadBuffer readBuffer, Long revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageParameterSetSelectedUnsubscribeRev1");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    readBuffer.closeContext("OpenProtocolMessageParameterSetSelectedUnsubscribeRev1");
    // Create the instance
    return new OpenProtocolMessageParameterSetSelectedUnsubscribeRev1BuilderImpl();
  }

  public static class OpenProtocolMessageParameterSetSelectedUnsubscribeRev1BuilderImpl
      implements OpenProtocolMessageParameterSetSelectedUnsubscribe
          .OpenProtocolMessageParameterSetSelectedUnsubscribeBuilder {

    public OpenProtocolMessageParameterSetSelectedUnsubscribeRev1BuilderImpl() {}

    public OpenProtocolMessageParameterSetSelectedUnsubscribeRev1 build(
        Long midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageParameterSetSelectedUnsubscribeRev1
          openProtocolMessageParameterSetSelectedUnsubscribeRev1 =
              new OpenProtocolMessageParameterSetSelectedUnsubscribeRev1(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber);
      return openProtocolMessageParameterSetSelectedUnsubscribeRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageParameterSetSelectedUnsubscribeRev1)) {
      return false;
    }
    OpenProtocolMessageParameterSetSelectedUnsubscribeRev1 that =
        (OpenProtocolMessageParameterSetSelectedUnsubscribeRev1) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
