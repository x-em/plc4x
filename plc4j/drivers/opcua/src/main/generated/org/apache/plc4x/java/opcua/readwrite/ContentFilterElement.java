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

public class ContentFilterElement extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 585;
  }

  // Properties.
  protected final FilterOperator filterOperator;
  protected final List<ExtensionObject> filterOperands;

  public ContentFilterElement(FilterOperator filterOperator, List<ExtensionObject> filterOperands) {
    super();
    this.filterOperator = filterOperator;
    this.filterOperands = filterOperands;
  }

  public FilterOperator getFilterOperator() {
    return filterOperator;
  }

  public List<ExtensionObject> getFilterOperands() {
    return filterOperands;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ContentFilterElement");

    // Simple Field (filterOperator)
    writeSimpleEnumField(
        "filterOperator",
        "FilterOperator",
        filterOperator,
        writeEnum(
            FilterOperator::getValue, FilterOperator::name, writeUnsignedLong(writeBuffer, 32)));

    // Implicit Field (noOfFilterOperands) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfFilterOperands =
        (int) ((((getFilterOperands()) == (null)) ? -(1) : COUNT(getFilterOperands())));
    writeImplicitField("noOfFilterOperands", noOfFilterOperands, writeSignedInt(writeBuffer, 32));

    // Array Field (filterOperands)
    writeComplexTypeArrayField("filterOperands", filterOperands, writeBuffer);

    writeBuffer.popContext("ContentFilterElement");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ContentFilterElement _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (filterOperator)
    lengthInBits += 32;

    // Implicit Field (noOfFilterOperands)
    lengthInBits += 32;

    // Array field
    if (filterOperands != null) {
      int i = 0;
      for (ExtensionObject element : filterOperands) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= filterOperands.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("ContentFilterElement");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    FilterOperator filterOperator =
        readEnumField(
            "filterOperator",
            "FilterOperator",
            readEnum(FilterOperator::enumForValue, readUnsignedLong(readBuffer, 32)));

    int noOfFilterOperands = readImplicitField("noOfFilterOperands", readSignedInt(readBuffer, 32));

    List<ExtensionObject> filterOperands =
        readCountArrayField(
            "filterOperands",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer),
            noOfFilterOperands);

    readBuffer.closeContext("ContentFilterElement");
    // Create the instance
    return new ContentFilterElementBuilderImpl(filterOperator, filterOperands);
  }

  public static class ContentFilterElementBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final FilterOperator filterOperator;
    private final List<ExtensionObject> filterOperands;

    public ContentFilterElementBuilderImpl(
        FilterOperator filterOperator, List<ExtensionObject> filterOperands) {
      this.filterOperator = filterOperator;
      this.filterOperands = filterOperands;
    }

    public ContentFilterElement build() {
      ContentFilterElement contentFilterElement =
          new ContentFilterElement(filterOperator, filterOperands);
      return contentFilterElement;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ContentFilterElement)) {
      return false;
    }
    ContentFilterElement that = (ContentFilterElement) o;
    return (getFilterOperator() == that.getFilterOperator())
        && (getFilterOperands() == that.getFilterOperands())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFilterOperator(), getFilterOperands());
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