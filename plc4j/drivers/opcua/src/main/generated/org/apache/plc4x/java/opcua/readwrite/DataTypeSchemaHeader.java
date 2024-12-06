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

public class DataTypeSchemaHeader extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 15536;
  }

  // Properties.
  protected final List<PascalString> namespaces;
  protected final List<StructureDescription> structureDataTypes;
  protected final List<EnumDescription> enumDataTypes;
  protected final List<SimpleTypeDescription> simpleDataTypes;

  public DataTypeSchemaHeader(
      List<PascalString> namespaces,
      List<StructureDescription> structureDataTypes,
      List<EnumDescription> enumDataTypes,
      List<SimpleTypeDescription> simpleDataTypes) {
    super();
    this.namespaces = namespaces;
    this.structureDataTypes = structureDataTypes;
    this.enumDataTypes = enumDataTypes;
    this.simpleDataTypes = simpleDataTypes;
  }

  public List<PascalString> getNamespaces() {
    return namespaces;
  }

  public List<StructureDescription> getStructureDataTypes() {
    return structureDataTypes;
  }

  public List<EnumDescription> getEnumDataTypes() {
    return enumDataTypes;
  }

  public List<SimpleTypeDescription> getSimpleDataTypes() {
    return simpleDataTypes;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DataTypeSchemaHeader");

    // Implicit Field (noOfNamespaces) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfNamespaces = (int) ((((getNamespaces()) == (null)) ? -(1) : COUNT(getNamespaces())));
    writeImplicitField("noOfNamespaces", noOfNamespaces, writeSignedInt(writeBuffer, 32));

    // Array Field (namespaces)
    writeComplexTypeArrayField("namespaces", namespaces, writeBuffer);

    // Implicit Field (noOfStructureDataTypes) (Used for parsing, but its value is not stored as
    // it's implicitly given by the objects content)
    int noOfStructureDataTypes =
        (int) ((((getStructureDataTypes()) == (null)) ? -(1) : COUNT(getStructureDataTypes())));
    writeImplicitField(
        "noOfStructureDataTypes", noOfStructureDataTypes, writeSignedInt(writeBuffer, 32));

    // Array Field (structureDataTypes)
    writeComplexTypeArrayField("structureDataTypes", structureDataTypes, writeBuffer);

    // Implicit Field (noOfEnumDataTypes) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfEnumDataTypes =
        (int) ((((getEnumDataTypes()) == (null)) ? -(1) : COUNT(getEnumDataTypes())));
    writeImplicitField("noOfEnumDataTypes", noOfEnumDataTypes, writeSignedInt(writeBuffer, 32));

    // Array Field (enumDataTypes)
    writeComplexTypeArrayField("enumDataTypes", enumDataTypes, writeBuffer);

    // Implicit Field (noOfSimpleDataTypes) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfSimpleDataTypes =
        (int) ((((getSimpleDataTypes()) == (null)) ? -(1) : COUNT(getSimpleDataTypes())));
    writeImplicitField("noOfSimpleDataTypes", noOfSimpleDataTypes, writeSignedInt(writeBuffer, 32));

    // Array Field (simpleDataTypes)
    writeComplexTypeArrayField("simpleDataTypes", simpleDataTypes, writeBuffer);

    writeBuffer.popContext("DataTypeSchemaHeader");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DataTypeSchemaHeader _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (noOfNamespaces)
    lengthInBits += 32;

    // Array field
    if (namespaces != null) {
      int i = 0;
      for (PascalString element : namespaces) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= namespaces.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfStructureDataTypes)
    lengthInBits += 32;

    // Array field
    if (structureDataTypes != null) {
      int i = 0;
      for (StructureDescription element : structureDataTypes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= structureDataTypes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfEnumDataTypes)
    lengthInBits += 32;

    // Array field
    if (enumDataTypes != null) {
      int i = 0;
      for (EnumDescription element : enumDataTypes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= enumDataTypes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfSimpleDataTypes)
    lengthInBits += 32;

    // Array field
    if (simpleDataTypes != null) {
      int i = 0;
      for (SimpleTypeDescription element : simpleDataTypes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= simpleDataTypes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("DataTypeSchemaHeader");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfNamespaces = readImplicitField("noOfNamespaces", readSignedInt(readBuffer, 32));

    List<PascalString> namespaces =
        readCountArrayField(
            "namespaces",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfNamespaces);

    int noOfStructureDataTypes =
        readImplicitField("noOfStructureDataTypes", readSignedInt(readBuffer, 32));

    List<StructureDescription> structureDataTypes =
        readCountArrayField(
            "structureDataTypes",
            readComplex(
                () ->
                    (StructureDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (15489)),
                readBuffer),
            noOfStructureDataTypes);

    int noOfEnumDataTypes = readImplicitField("noOfEnumDataTypes", readSignedInt(readBuffer, 32));

    List<EnumDescription> enumDataTypes =
        readCountArrayField(
            "enumDataTypes",
            readComplex(
                () ->
                    (EnumDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (15490)),
                readBuffer),
            noOfEnumDataTypes);

    int noOfSimpleDataTypes =
        readImplicitField("noOfSimpleDataTypes", readSignedInt(readBuffer, 32));

    List<SimpleTypeDescription> simpleDataTypes =
        readCountArrayField(
            "simpleDataTypes",
            readComplex(
                () ->
                    (SimpleTypeDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (15007)),
                readBuffer),
            noOfSimpleDataTypes);

    readBuffer.closeContext("DataTypeSchemaHeader");
    // Create the instance
    return new DataTypeSchemaHeaderBuilderImpl(
        namespaces, structureDataTypes, enumDataTypes, simpleDataTypes);
  }

  public static class DataTypeSchemaHeaderBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final List<PascalString> namespaces;
    private final List<StructureDescription> structureDataTypes;
    private final List<EnumDescription> enumDataTypes;
    private final List<SimpleTypeDescription> simpleDataTypes;

    public DataTypeSchemaHeaderBuilderImpl(
        List<PascalString> namespaces,
        List<StructureDescription> structureDataTypes,
        List<EnumDescription> enumDataTypes,
        List<SimpleTypeDescription> simpleDataTypes) {
      this.namespaces = namespaces;
      this.structureDataTypes = structureDataTypes;
      this.enumDataTypes = enumDataTypes;
      this.simpleDataTypes = simpleDataTypes;
    }

    public DataTypeSchemaHeader build() {
      DataTypeSchemaHeader dataTypeSchemaHeader =
          new DataTypeSchemaHeader(namespaces, structureDataTypes, enumDataTypes, simpleDataTypes);
      return dataTypeSchemaHeader;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DataTypeSchemaHeader)) {
      return false;
    }
    DataTypeSchemaHeader that = (DataTypeSchemaHeader) o;
    return (getNamespaces() == that.getNamespaces())
        && (getStructureDataTypes() == that.getStructureDataTypes())
        && (getEnumDataTypes() == that.getEnumDataTypes())
        && (getSimpleDataTypes() == that.getSimpleDataTypes())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNamespaces(),
        getStructureDataTypes(),
        getEnumDataTypes(),
        getSimpleDataTypes());
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