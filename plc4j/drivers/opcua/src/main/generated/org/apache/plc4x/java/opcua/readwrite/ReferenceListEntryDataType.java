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
package org.apache.plc4x.java.opcua.readwrite;

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

public class ReferenceListEntryDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 32662;
  }

  // Properties.
  protected final NodeId referenceType;
  protected final boolean isForward;
  protected final ExpandedNodeId targetNode;

  public ReferenceListEntryDataType(
      NodeId referenceType, boolean isForward, ExpandedNodeId targetNode) {
    super();
    this.referenceType = referenceType;
    this.isForward = isForward;
    this.targetNode = targetNode;
  }

  public NodeId getReferenceType() {
    return referenceType;
  }

  public boolean getIsForward() {
    return isForward;
  }

  public ExpandedNodeId getTargetNode() {
    return targetNode;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ReferenceListEntryDataType");

    // Simple Field (referenceType)
    writeSimpleField("referenceType", referenceType, writeComplex(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (isForward)
    writeSimpleField("isForward", isForward, writeBoolean(writeBuffer));

    // Simple Field (targetNode)
    writeSimpleField("targetNode", targetNode, writeComplex(writeBuffer));

    writeBuffer.popContext("ReferenceListEntryDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ReferenceListEntryDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (referenceType)
    lengthInBits += referenceType.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (isForward)
    lengthInBits += 1;

    // Simple field (targetNode)
    lengthInBits += targetNode.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("ReferenceListEntryDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId referenceType =
        readSimpleField(
            "referenceType", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean isForward = readSimpleField("isForward", readBoolean(readBuffer));

    ExpandedNodeId targetNode =
        readSimpleField(
            "targetNode", readComplex(() -> ExpandedNodeId.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("ReferenceListEntryDataType");
    // Create the instance
    return new ReferenceListEntryDataTypeBuilderImpl(referenceType, isForward, targetNode);
  }

  public static class ReferenceListEntryDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId referenceType;
    private final boolean isForward;
    private final ExpandedNodeId targetNode;

    public ReferenceListEntryDataTypeBuilderImpl(
        NodeId referenceType, boolean isForward, ExpandedNodeId targetNode) {
      this.referenceType = referenceType;
      this.isForward = isForward;
      this.targetNode = targetNode;
    }

    public ReferenceListEntryDataType build() {
      ReferenceListEntryDataType referenceListEntryDataType =
          new ReferenceListEntryDataType(referenceType, isForward, targetNode);
      return referenceListEntryDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ReferenceListEntryDataType)) {
      return false;
    }
    ReferenceListEntryDataType that = (ReferenceListEntryDataType) o;
    return (getReferenceType() == that.getReferenceType())
        && (getIsForward() == that.getIsForward())
        && (getTargetNode() == that.getTargetNode())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReferenceType(), getIsForward(), getTargetNode());
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