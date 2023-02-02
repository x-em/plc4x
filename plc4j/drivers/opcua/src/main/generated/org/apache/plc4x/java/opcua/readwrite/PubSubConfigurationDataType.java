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

public class PubSubConfigurationDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "15532";
  }

  // Properties.
  protected final int noOfPublishedDataSets;
  protected final List<ExtensionObjectDefinition> publishedDataSets;
  protected final int noOfConnections;
  protected final List<ExtensionObjectDefinition> connections;
  protected final boolean enabled;

  public PubSubConfigurationDataType(
      int noOfPublishedDataSets,
      List<ExtensionObjectDefinition> publishedDataSets,
      int noOfConnections,
      List<ExtensionObjectDefinition> connections,
      boolean enabled) {
    super();
    this.noOfPublishedDataSets = noOfPublishedDataSets;
    this.publishedDataSets = publishedDataSets;
    this.noOfConnections = noOfConnections;
    this.connections = connections;
    this.enabled = enabled;
  }

  public int getNoOfPublishedDataSets() {
    return noOfPublishedDataSets;
  }

  public List<ExtensionObjectDefinition> getPublishedDataSets() {
    return publishedDataSets;
  }

  public int getNoOfConnections() {
    return noOfConnections;
  }

  public List<ExtensionObjectDefinition> getConnections() {
    return connections;
  }

  public boolean getEnabled() {
    return enabled;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PubSubConfigurationDataType");

    // Simple Field (noOfPublishedDataSets)
    writeSimpleField(
        "noOfPublishedDataSets", noOfPublishedDataSets, writeSignedInt(writeBuffer, 32));

    // Array Field (publishedDataSets)
    writeComplexTypeArrayField("publishedDataSets", publishedDataSets, writeBuffer);

    // Simple Field (noOfConnections)
    writeSimpleField("noOfConnections", noOfConnections, writeSignedInt(writeBuffer, 32));

    // Array Field (connections)
    writeComplexTypeArrayField("connections", connections, writeBuffer);

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 7));

    // Simple Field (enabled)
    writeSimpleField("enabled", enabled, writeBoolean(writeBuffer));

    writeBuffer.popContext("PubSubConfigurationDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PubSubConfigurationDataType _value = this;

    // Simple field (noOfPublishedDataSets)
    lengthInBits += 32;

    // Array field
    if (publishedDataSets != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : publishedDataSets) {
        boolean last = ++i >= publishedDataSets.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfConnections)
    lengthInBits += 32;

    // Array field
    if (connections != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : connections) {
        boolean last = ++i >= connections.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (enabled)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("PubSubConfigurationDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int noOfPublishedDataSets =
        readSimpleField("noOfPublishedDataSets", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> publishedDataSets =
        readCountArrayField(
            "publishedDataSets",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("15580")),
                readBuffer),
            noOfPublishedDataSets);

    int noOfConnections = readSimpleField("noOfConnections", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> connections =
        readCountArrayField(
            "connections",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("15619")),
                readBuffer),
            noOfConnections);

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 7), (short) 0x00);

    boolean enabled = readSimpleField("enabled", readBoolean(readBuffer));

    readBuffer.closeContext("PubSubConfigurationDataType");
    // Create the instance
    return new PubSubConfigurationDataTypeBuilderImpl(
        noOfPublishedDataSets, publishedDataSets, noOfConnections, connections, enabled);
  }

  public static class PubSubConfigurationDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int noOfPublishedDataSets;
    private final List<ExtensionObjectDefinition> publishedDataSets;
    private final int noOfConnections;
    private final List<ExtensionObjectDefinition> connections;
    private final boolean enabled;

    public PubSubConfigurationDataTypeBuilderImpl(
        int noOfPublishedDataSets,
        List<ExtensionObjectDefinition> publishedDataSets,
        int noOfConnections,
        List<ExtensionObjectDefinition> connections,
        boolean enabled) {
      this.noOfPublishedDataSets = noOfPublishedDataSets;
      this.publishedDataSets = publishedDataSets;
      this.noOfConnections = noOfConnections;
      this.connections = connections;
      this.enabled = enabled;
    }

    public PubSubConfigurationDataType build() {
      PubSubConfigurationDataType pubSubConfigurationDataType =
          new PubSubConfigurationDataType(
              noOfPublishedDataSets, publishedDataSets, noOfConnections, connections, enabled);
      return pubSubConfigurationDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PubSubConfigurationDataType)) {
      return false;
    }
    PubSubConfigurationDataType that = (PubSubConfigurationDataType) o;
    return (getNoOfPublishedDataSets() == that.getNoOfPublishedDataSets())
        && (getPublishedDataSets() == that.getPublishedDataSets())
        && (getNoOfConnections() == that.getNoOfConnections())
        && (getConnections() == that.getConnections())
        && (getEnabled() == that.getEnabled())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNoOfPublishedDataSets(),
        getPublishedDataSets(),
        getNoOfConnections(),
        getConnections(),
        getEnabled());
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
