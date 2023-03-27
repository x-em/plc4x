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

public class PnIoCm_Block_ArRes extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.AR_BLOCK_RES;
  }

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final PnIoCm_ArType arType;
  protected final Uuid arUuid;
  protected final int sessionKey;
  protected final MacAddress cmResponderMacAddr;
  protected final int responderUDPRTPort;

  public PnIoCm_Block_ArRes(
      short blockVersionHigh,
      short blockVersionLow,
      PnIoCm_ArType arType,
      Uuid arUuid,
      int sessionKey,
      MacAddress cmResponderMacAddr,
      int responderUDPRTPort) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.arType = arType;
    this.arUuid = arUuid;
    this.sessionKey = sessionKey;
    this.cmResponderMacAddr = cmResponderMacAddr;
    this.responderUDPRTPort = responderUDPRTPort;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public PnIoCm_ArType getArType() {
    return arType;
  }

  public Uuid getArUuid() {
    return arUuid;
  }

  public int getSessionKey() {
    return sessionKey;
  }

  public MacAddress getCmResponderMacAddr() {
    return cmResponderMacAddr;
  }

  public int getResponderUDPRTPort() {
    return responderUDPRTPort;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Block_ArRes");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (arType)
    writeSimpleEnumField(
        "arType",
        "PnIoCm_ArType",
        arType,
        new DataWriterEnumDefault<>(
            PnIoCm_ArType::getValue, PnIoCm_ArType::name, writeUnsignedInt(writeBuffer, 16)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (arUuid)
    writeSimpleField(
        "arUuid",
        arUuid,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (sessionKey)
    writeSimpleField(
        "sessionKey",
        sessionKey,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (cmResponderMacAddr)
    writeSimpleField(
        "cmResponderMacAddr",
        cmResponderMacAddr,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (responderUDPRTPort)
    writeSimpleField(
        "responderUDPRTPort",
        responderUDPRTPort,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_Block_ArRes");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Block_ArRes _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Simple field (arType)
    lengthInBits += 16;

    // Simple field (arUuid)
    lengthInBits += arUuid.getLengthInBits();

    // Simple field (sessionKey)
    lengthInBits += 16;

    // Simple field (cmResponderMacAddr)
    lengthInBits += cmResponderMacAddr.getLengthInBits();

    // Simple field (responderUDPRTPort)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Block_ArRes");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    PnIoCm_ArType arType =
        readEnumField(
            "arType",
            "PnIoCm_ArType",
            new DataReaderEnumDefault<>(
                PnIoCm_ArType::enumForValue, readUnsignedInt(readBuffer, 16)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Uuid arUuid =
        readSimpleField(
            "arUuid",
            new DataReaderComplexDefault<>(() -> Uuid.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int sessionKey =
        readSimpleField(
            "sessionKey",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    MacAddress cmResponderMacAddr =
        readSimpleField(
            "cmResponderMacAddr",
            new DataReaderComplexDefault<>(() -> MacAddress.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int responderUDPRTPort =
        readSimpleField(
            "responderUDPRTPort",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_Block_ArRes");
    // Create the instance
    return new PnIoCm_Block_ArResBuilderImpl(
        blockVersionHigh,
        blockVersionLow,
        arType,
        arUuid,
        sessionKey,
        cmResponderMacAddr,
        responderUDPRTPort);
  }

  public static class PnIoCm_Block_ArResBuilderImpl implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final PnIoCm_ArType arType;
    private final Uuid arUuid;
    private final int sessionKey;
    private final MacAddress cmResponderMacAddr;
    private final int responderUDPRTPort;

    public PnIoCm_Block_ArResBuilderImpl(
        short blockVersionHigh,
        short blockVersionLow,
        PnIoCm_ArType arType,
        Uuid arUuid,
        int sessionKey,
        MacAddress cmResponderMacAddr,
        int responderUDPRTPort) {
      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.arType = arType;
      this.arUuid = arUuid;
      this.sessionKey = sessionKey;
      this.cmResponderMacAddr = cmResponderMacAddr;
      this.responderUDPRTPort = responderUDPRTPort;
    }

    public PnIoCm_Block_ArRes build() {
      PnIoCm_Block_ArRes pnIoCm_Block_ArRes =
          new PnIoCm_Block_ArRes(
              blockVersionHigh,
              blockVersionLow,
              arType,
              arUuid,
              sessionKey,
              cmResponderMacAddr,
              responderUDPRTPort);
      return pnIoCm_Block_ArRes;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block_ArRes)) {
      return false;
    }
    PnIoCm_Block_ArRes that = (PnIoCm_Block_ArRes) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getArType() == that.getArType())
        && (getArUuid() == that.getArUuid())
        && (getSessionKey() == that.getSessionKey())
        && (getCmResponderMacAddr() == that.getCmResponderMacAddr())
        && (getResponderUDPRTPort() == that.getResponderUDPRTPort())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getBlockVersionHigh(),
        getBlockVersionLow(),
        getArType(),
        getArUuid(),
        getSessionKey(),
        getCmResponderMacAddr(),
        getResponderUDPRTPort());
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
