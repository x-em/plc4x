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

public class PubSubConfigurationValueDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 25522;
  }

  // Properties.
  protected final PubSubConfigurationRefDataType configurationElement;
  protected final PascalString name;
  protected final Variant identifier;

  public PubSubConfigurationValueDataType(
      PubSubConfigurationRefDataType configurationElement, PascalString name, Variant identifier) {
    super();
    this.configurationElement = configurationElement;
    this.name = name;
    this.identifier = identifier;
  }

  public PubSubConfigurationRefDataType getConfigurationElement() {
    return configurationElement;
  }

  public PascalString getName() {
    return name;
  }

  public Variant getIdentifier() {
    return identifier;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PubSubConfigurationValueDataType");

    // Simple Field (configurationElement)
    writeSimpleField("configurationElement", configurationElement, writeComplex(writeBuffer));

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    // Simple Field (identifier)
    writeSimpleField("identifier", identifier, writeComplex(writeBuffer));

    writeBuffer.popContext("PubSubConfigurationValueDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PubSubConfigurationValueDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (configurationElement)
    lengthInBits += configurationElement.getLengthInBits();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Simple field (identifier)
    lengthInBits += identifier.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("PubSubConfigurationValueDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PubSubConfigurationRefDataType configurationElement =
        readSimpleField(
            "configurationElement",
            readComplex(
                () ->
                    (PubSubConfigurationRefDataType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (25521)),
                readBuffer));

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    Variant identifier =
        readSimpleField(
            "identifier", readComplex(() -> Variant.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("PubSubConfigurationValueDataType");
    // Create the instance
    return new PubSubConfigurationValueDataTypeBuilderImpl(configurationElement, name, identifier);
  }

  public static class PubSubConfigurationValueDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PubSubConfigurationRefDataType configurationElement;
    private final PascalString name;
    private final Variant identifier;

    public PubSubConfigurationValueDataTypeBuilderImpl(
        PubSubConfigurationRefDataType configurationElement,
        PascalString name,
        Variant identifier) {
      this.configurationElement = configurationElement;
      this.name = name;
      this.identifier = identifier;
    }

    public PubSubConfigurationValueDataType build() {
      PubSubConfigurationValueDataType pubSubConfigurationValueDataType =
          new PubSubConfigurationValueDataType(configurationElement, name, identifier);
      return pubSubConfigurationValueDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PubSubConfigurationValueDataType)) {
      return false;
    }
    PubSubConfigurationValueDataType that = (PubSubConfigurationValueDataType) o;
    return (getConfigurationElement() == that.getConfigurationElement())
        && (getName() == that.getName())
        && (getIdentifier() == that.getIdentifier())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getConfigurationElement(), getName(), getIdentifier());
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