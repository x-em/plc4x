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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PcDcp_Pdu_DelayReq extends PnDcp_Pdu implements Message {

  // Accessors for discriminator values.

  // Constant values.
  public static final Short PARAMETERTYPE = 6;
  public static final Integer PARAMETERLENGTH = 6;
  public static final Short ENDTYPE = 0;
  public static final Integer ENDLENGTH = 0;

  // Properties.
  protected final int sequenceId;
  protected final long delayInNs;
  protected final MacAddress portMacAddress;

  // Reserved Fields
  private Long reservedField0;
  private Long reservedField1;
  private Long reservedField2;
  private Integer reservedField3;

  public PcDcp_Pdu_DelayReq(
      int frameIdValue, int sequenceId, long delayInNs, MacAddress portMacAddress) {
    super(frameIdValue);
    this.sequenceId = sequenceId;
    this.delayInNs = delayInNs;
    this.portMacAddress = portMacAddress;
  }

  public int getSequenceId() {
    return sequenceId;
  }

  public long getDelayInNs() {
    return delayInNs;
  }

  public MacAddress getPortMacAddress() {
    return portMacAddress;
  }

  public short getParameterType() {
    return PARAMETERTYPE;
  }

  public int getParameterLength() {
    return PARAMETERLENGTH;
  }

  public short getEndType() {
    return ENDTYPE;
  }

  public int getEndLength() {
    return ENDLENGTH;
  }

  @Override
  protected void serializePnDcp_PduChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PcDcp_Pdu_DelayReq");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (long) 0x00000000,
        writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (long) 0x00000000,
        writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField2 != null ? reservedField2 : (long) 0x00000000,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (sequenceId)
    writeSimpleField("sequenceId", sequenceId, writeUnsignedInt(writeBuffer, 16));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField3 != null ? reservedField3 : (int) 0x0000,
        writeUnsignedInt(writeBuffer, 16));

    // Simple Field (delayInNs)
    writeSimpleField("delayInNs", delayInNs, writeUnsignedLong(writeBuffer, 32));

    // Const Field (parameterType)
    writeConstField("parameterType", PARAMETERTYPE, writeUnsignedShort(writeBuffer, 7));

    // Const Field (parameterLength)
    writeConstField("parameterLength", PARAMETERLENGTH, writeUnsignedInt(writeBuffer, 9));

    // Simple Field (portMacAddress)
    writeSimpleField("portMacAddress", portMacAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Const Field (endType)
    writeConstField("endType", ENDTYPE, writeUnsignedShort(writeBuffer, 7));

    // Const Field (endLength)
    writeConstField("endLength", ENDLENGTH, writeUnsignedInt(writeBuffer, 9));

    writeBuffer.popContext("PcDcp_Pdu_DelayReq");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PcDcp_Pdu_DelayReq _value = this;

    // Reserved Field (reserved)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 32;

    // Simple field (sequenceId)
    lengthInBits += 16;

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Simple field (delayInNs)
    lengthInBits += 32;

    // Const Field (parameterType)
    lengthInBits += 7;

    // Const Field (parameterLength)
    lengthInBits += 9;

    // Simple field (portMacAddress)
    lengthInBits += portMacAddress.getLengthInBits();

    // Const Field (endType)
    lengthInBits += 7;

    // Const Field (endLength)
    lengthInBits += 9;

    return lengthInBits;
  }

  public static PnDcp_PduBuilder staticParsePnDcp_PduBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PcDcp_Pdu_DelayReq");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Long reservedField0 =
        readReservedField("reserved", readUnsignedLong(readBuffer, 32), (long) 0x00000000);

    Long reservedField1 =
        readReservedField("reserved", readUnsignedLong(readBuffer, 32), (long) 0x00000000);

    Long reservedField2 =
        readReservedField("reserved", readUnsignedLong(readBuffer, 32), (long) 0x00000000);

    int sequenceId = readSimpleField("sequenceId", readUnsignedInt(readBuffer, 16));

    Integer reservedField3 =
        readReservedField("reserved", readUnsignedInt(readBuffer, 16), (int) 0x0000);

    long delayInNs = readSimpleField("delayInNs", readUnsignedLong(readBuffer, 32));

    short parameterType =
        readConstField(
            "parameterType", readUnsignedShort(readBuffer, 7), PcDcp_Pdu_DelayReq.PARAMETERTYPE);

    int parameterLength =
        readConstField(
            "parameterLength", readUnsignedInt(readBuffer, 9), PcDcp_Pdu_DelayReq.PARAMETERLENGTH);

    MacAddress portMacAddress =
        readSimpleField(
            "portMacAddress",
            new DataReaderComplexDefault<>(() -> MacAddress.staticParse(readBuffer), readBuffer));

    short endType =
        readConstField("endType", readUnsignedShort(readBuffer, 7), PcDcp_Pdu_DelayReq.ENDTYPE);

    int endLength =
        readConstField("endLength", readUnsignedInt(readBuffer, 9), PcDcp_Pdu_DelayReq.ENDLENGTH);

    readBuffer.closeContext("PcDcp_Pdu_DelayReq");
    // Create the instance
    return new PcDcp_Pdu_DelayReqBuilderImpl(
        sequenceId,
        delayInNs,
        portMacAddress,
        reservedField0,
        reservedField1,
        reservedField2,
        reservedField3);
  }

  public static class PcDcp_Pdu_DelayReqBuilderImpl implements PnDcp_Pdu.PnDcp_PduBuilder {
    private final int sequenceId;
    private final long delayInNs;
    private final MacAddress portMacAddress;
    private final Long reservedField0;
    private final Long reservedField1;
    private final Long reservedField2;
    private final Integer reservedField3;

    public PcDcp_Pdu_DelayReqBuilderImpl(
        int sequenceId,
        long delayInNs,
        MacAddress portMacAddress,
        Long reservedField0,
        Long reservedField1,
        Long reservedField2,
        Integer reservedField3) {
      this.sequenceId = sequenceId;
      this.delayInNs = delayInNs;
      this.portMacAddress = portMacAddress;
      this.reservedField0 = reservedField0;
      this.reservedField1 = reservedField1;
      this.reservedField2 = reservedField2;
      this.reservedField3 = reservedField3;
    }

    public PcDcp_Pdu_DelayReq build(int frameIdValue) {
      PcDcp_Pdu_DelayReq pcDcp_Pdu_DelayReq =
          new PcDcp_Pdu_DelayReq(frameIdValue, sequenceId, delayInNs, portMacAddress);
      pcDcp_Pdu_DelayReq.reservedField0 = reservedField0;
      pcDcp_Pdu_DelayReq.reservedField1 = reservedField1;
      pcDcp_Pdu_DelayReq.reservedField2 = reservedField2;
      pcDcp_Pdu_DelayReq.reservedField3 = reservedField3;
      return pcDcp_Pdu_DelayReq;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PcDcp_Pdu_DelayReq)) {
      return false;
    }
    PcDcp_Pdu_DelayReq that = (PcDcp_Pdu_DelayReq) o;
    return (getSequenceId() == that.getSequenceId())
        && (getDelayInNs() == that.getDelayInNs())
        && (getPortMacAddress() == that.getPortMacAddress())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSequenceId(), getDelayInNs(), getPortMacAddress());
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
