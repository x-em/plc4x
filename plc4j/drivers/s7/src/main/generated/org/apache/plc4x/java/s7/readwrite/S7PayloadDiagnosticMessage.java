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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7PayloadDiagnosticMessage extends S7PayloadUserDataItem implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionGroup() {
    return (byte) 0x04;
  }

  public Byte getCpuFunctionType() {
    return (byte) 0x00;
  }

  public Short getCpuSubfunction() {
    return (short) 0x03;
  }

  // Properties.
  protected final int eventId;
  protected final short priorityClass;
  protected final short obNumber;
  protected final int datId;
  protected final int info1;
  protected final long info2;
  protected final DateAndTime timeStamp;

  public S7PayloadDiagnosticMessage(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      int dataLength,
      int eventId,
      short priorityClass,
      short obNumber,
      int datId,
      int info1,
      long info2,
      DateAndTime timeStamp) {
    super(returnCode, transportSize, dataLength);
    this.eventId = eventId;
    this.priorityClass = priorityClass;
    this.obNumber = obNumber;
    this.datId = datId;
    this.info1 = info1;
    this.info2 = info2;
    this.timeStamp = timeStamp;
  }

  public int getEventId() {
    return eventId;
  }

  public short getPriorityClass() {
    return priorityClass;
  }

  public short getObNumber() {
    return obNumber;
  }

  public int getDatId() {
    return datId;
  }

  public int getInfo1() {
    return info1;
  }

  public long getInfo2() {
    return info2;
  }

  public DateAndTime getTimeStamp() {
    return timeStamp;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7PayloadDiagnosticMessage");

    // Simple Field (eventId)
    writeSimpleField("eventId", eventId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (priorityClass)
    writeSimpleField("priorityClass", priorityClass, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (obNumber)
    writeSimpleField("obNumber", obNumber, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (datId)
    writeSimpleField("datId", datId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (info1)
    writeSimpleField("info1", info1, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (info2)
    writeSimpleField("info2", info2, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (timeStamp)
    writeSimpleField("timeStamp", timeStamp, writeComplex(writeBuffer));

    writeBuffer.popContext("S7PayloadDiagnosticMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadDiagnosticMessage _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (eventId)
    lengthInBits += 16;

    // Simple field (priorityClass)
    lengthInBits += 8;

    // Simple field (obNumber)
    lengthInBits += 8;

    // Simple field (datId)
    lengthInBits += 16;

    // Simple field (info1)
    lengthInBits += 16;

    // Simple field (info2)
    lengthInBits += 32;

    // Simple field (timeStamp)
    lengthInBits += timeStamp.getLengthInBits();

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionGroup, Byte cpuFunctionType, Short cpuSubfunction)
      throws ParseException {
    readBuffer.pullContext("S7PayloadDiagnosticMessage");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int eventId = readSimpleField("eventId", readUnsignedInt(readBuffer, 16));

    short priorityClass = readSimpleField("priorityClass", readUnsignedShort(readBuffer, 8));

    short obNumber = readSimpleField("obNumber", readUnsignedShort(readBuffer, 8));

    int datId = readSimpleField("datId", readUnsignedInt(readBuffer, 16));

    int info1 = readSimpleField("info1", readUnsignedInt(readBuffer, 16));

    long info2 = readSimpleField("info2", readUnsignedLong(readBuffer, 32));

    DateAndTime timeStamp =
        readSimpleField(
            "timeStamp", readComplex(() -> DateAndTime.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("S7PayloadDiagnosticMessage");
    // Create the instance
    return new S7PayloadDiagnosticMessageBuilderImpl(
        eventId, priorityClass, obNumber, datId, info1, info2, timeStamp);
  }

  public static class S7PayloadDiagnosticMessageBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final int eventId;
    private final short priorityClass;
    private final short obNumber;
    private final int datId;
    private final int info1;
    private final long info2;
    private final DateAndTime timeStamp;

    public S7PayloadDiagnosticMessageBuilderImpl(
        int eventId,
        short priorityClass,
        short obNumber,
        int datId,
        int info1,
        long info2,
        DateAndTime timeStamp) {
      this.eventId = eventId;
      this.priorityClass = priorityClass;
      this.obNumber = obNumber;
      this.datId = datId;
      this.info1 = info1;
      this.info2 = info2;
      this.timeStamp = timeStamp;
    }

    public S7PayloadDiagnosticMessage build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize, int dataLength) {
      S7PayloadDiagnosticMessage s7PayloadDiagnosticMessage =
          new S7PayloadDiagnosticMessage(
              returnCode,
              transportSize,
              dataLength,
              eventId,
              priorityClass,
              obNumber,
              datId,
              info1,
              info2,
              timeStamp);
      return s7PayloadDiagnosticMessage;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadDiagnosticMessage)) {
      return false;
    }
    S7PayloadDiagnosticMessage that = (S7PayloadDiagnosticMessage) o;
    return (getEventId() == that.getEventId())
        && (getPriorityClass() == that.getPriorityClass())
        && (getObNumber() == that.getObNumber())
        && (getDatId() == that.getDatId())
        && (getInfo1() == that.getInfo1())
        && (getInfo2() == that.getInfo2())
        && (getTimeStamp() == that.getTimeStamp())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getEventId(),
        getPriorityClass(),
        getObNumber(),
        getDatId(),
        getInfo1(),
        getInfo2(),
        getTimeStamp());
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