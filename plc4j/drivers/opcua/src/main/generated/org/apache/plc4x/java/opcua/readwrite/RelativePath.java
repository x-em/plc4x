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

public class RelativePath extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 542;
  }

  // Properties.
  protected final List<RelativePathElement> elements;

  public RelativePath(List<RelativePathElement> elements) {
    super();
    this.elements = elements;
  }

  public List<RelativePathElement> getElements() {
    return elements;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RelativePath");

    // Implicit Field (noOfElements) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfElements = (int) ((((getElements()) == (null)) ? -(1) : COUNT(getElements())));
    writeImplicitField("noOfElements", noOfElements, writeSignedInt(writeBuffer, 32));

    // Array Field (elements)
    writeComplexTypeArrayField("elements", elements, writeBuffer);

    writeBuffer.popContext("RelativePath");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RelativePath _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (noOfElements)
    lengthInBits += 32;

    // Array field
    if (elements != null) {
      int i = 0;
      for (RelativePathElement element : elements) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= elements.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("RelativePath");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfElements = readImplicitField("noOfElements", readSignedInt(readBuffer, 32));

    List<RelativePathElement> elements =
        readCountArrayField(
            "elements",
            readComplex(
                () ->
                    (RelativePathElement)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (539)),
                readBuffer),
            noOfElements);

    readBuffer.closeContext("RelativePath");
    // Create the instance
    return new RelativePathBuilderImpl(elements);
  }

  public static class RelativePathBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final List<RelativePathElement> elements;

    public RelativePathBuilderImpl(List<RelativePathElement> elements) {
      this.elements = elements;
    }

    public RelativePath build() {
      RelativePath relativePath = new RelativePath(elements);
      return relativePath;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RelativePath)) {
      return false;
    }
    RelativePath that = (RelativePath) o;
    return (getElements() == that.getElements()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getElements());
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