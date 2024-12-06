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

public class SimpleAttributeOperand extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 603;
  }

  // Properties.
  protected final NodeId typeDefinitionId;
  protected final List<QualifiedName> browsePath;
  protected final long attributeId;
  protected final PascalString indexRange;

  public SimpleAttributeOperand(
      NodeId typeDefinitionId,
      List<QualifiedName> browsePath,
      long attributeId,
      PascalString indexRange) {
    super();
    this.typeDefinitionId = typeDefinitionId;
    this.browsePath = browsePath;
    this.attributeId = attributeId;
    this.indexRange = indexRange;
  }

  public NodeId getTypeDefinitionId() {
    return typeDefinitionId;
  }

  public List<QualifiedName> getBrowsePath() {
    return browsePath;
  }

  public long getAttributeId() {
    return attributeId;
  }

  public PascalString getIndexRange() {
    return indexRange;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SimpleAttributeOperand");

    // Simple Field (typeDefinitionId)
    writeSimpleField("typeDefinitionId", typeDefinitionId, writeComplex(writeBuffer));

    // Implicit Field (noOfBrowsePath) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfBrowsePath = (int) ((((getBrowsePath()) == (null)) ? -(1) : COUNT(getBrowsePath())));
    writeImplicitField("noOfBrowsePath", noOfBrowsePath, writeSignedInt(writeBuffer, 32));

    // Array Field (browsePath)
    writeComplexTypeArrayField("browsePath", browsePath, writeBuffer);

    // Simple Field (attributeId)
    writeSimpleField("attributeId", attributeId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (indexRange)
    writeSimpleField("indexRange", indexRange, writeComplex(writeBuffer));

    writeBuffer.popContext("SimpleAttributeOperand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SimpleAttributeOperand _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (typeDefinitionId)
    lengthInBits += typeDefinitionId.getLengthInBits();

    // Implicit Field (noOfBrowsePath)
    lengthInBits += 32;

    // Array field
    if (browsePath != null) {
      int i = 0;
      for (QualifiedName element : browsePath) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= browsePath.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (attributeId)
    lengthInBits += 32;

    // Simple field (indexRange)
    lengthInBits += indexRange.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("SimpleAttributeOperand");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId typeDefinitionId =
        readSimpleField(
            "typeDefinitionId", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    int noOfBrowsePath = readImplicitField("noOfBrowsePath", readSignedInt(readBuffer, 32));

    List<QualifiedName> browsePath =
        readCountArrayField(
            "browsePath",
            readComplex(() -> QualifiedName.staticParse(readBuffer), readBuffer),
            noOfBrowsePath);

    long attributeId = readSimpleField("attributeId", readUnsignedLong(readBuffer, 32));

    PascalString indexRange =
        readSimpleField(
            "indexRange", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SimpleAttributeOperand");
    // Create the instance
    return new SimpleAttributeOperandBuilderImpl(
        typeDefinitionId, browsePath, attributeId, indexRange);
  }

  public static class SimpleAttributeOperandBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId typeDefinitionId;
    private final List<QualifiedName> browsePath;
    private final long attributeId;
    private final PascalString indexRange;

    public SimpleAttributeOperandBuilderImpl(
        NodeId typeDefinitionId,
        List<QualifiedName> browsePath,
        long attributeId,
        PascalString indexRange) {
      this.typeDefinitionId = typeDefinitionId;
      this.browsePath = browsePath;
      this.attributeId = attributeId;
      this.indexRange = indexRange;
    }

    public SimpleAttributeOperand build() {
      SimpleAttributeOperand simpleAttributeOperand =
          new SimpleAttributeOperand(typeDefinitionId, browsePath, attributeId, indexRange);
      return simpleAttributeOperand;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SimpleAttributeOperand)) {
      return false;
    }
    SimpleAttributeOperand that = (SimpleAttributeOperand) o;
    return (getTypeDefinitionId() == that.getTypeDefinitionId())
        && (getBrowsePath() == that.getBrowsePath())
        && (getAttributeId() == that.getAttributeId())
        && (getIndexRange() == that.getIndexRange())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getTypeDefinitionId(),
        getBrowsePath(),
        getAttributeId(),
        getIndexRange());
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