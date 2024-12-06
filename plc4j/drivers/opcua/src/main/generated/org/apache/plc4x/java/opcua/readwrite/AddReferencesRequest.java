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

public class AddReferencesRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 494;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final List<AddReferencesItem> referencesToAdd;

  public AddReferencesRequest(
      RequestHeader requestHeader, List<AddReferencesItem> referencesToAdd) {
    super();
    this.requestHeader = requestHeader;
    this.referencesToAdd = referencesToAdd;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public List<AddReferencesItem> getReferencesToAdd() {
    return referencesToAdd;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AddReferencesRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Implicit Field (noOfReferencesToAdd) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfReferencesToAdd =
        (int) ((((getReferencesToAdd()) == (null)) ? -(1) : COUNT(getReferencesToAdd())));
    writeImplicitField("noOfReferencesToAdd", noOfReferencesToAdd, writeSignedInt(writeBuffer, 32));

    // Array Field (referencesToAdd)
    writeComplexTypeArrayField("referencesToAdd", referencesToAdd, writeBuffer);

    writeBuffer.popContext("AddReferencesRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AddReferencesRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Implicit Field (noOfReferencesToAdd)
    lengthInBits += 32;

    // Array field
    if (referencesToAdd != null) {
      int i = 0;
      for (AddReferencesItem element : referencesToAdd) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= referencesToAdd.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("AddReferencesRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    int noOfReferencesToAdd =
        readImplicitField("noOfReferencesToAdd", readSignedInt(readBuffer, 32));

    List<AddReferencesItem> referencesToAdd =
        readCountArrayField(
            "referencesToAdd",
            readComplex(
                () ->
                    (AddReferencesItem)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (381)),
                readBuffer),
            noOfReferencesToAdd);

    readBuffer.closeContext("AddReferencesRequest");
    // Create the instance
    return new AddReferencesRequestBuilderImpl(requestHeader, referencesToAdd);
  }

  public static class AddReferencesRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final List<AddReferencesItem> referencesToAdd;

    public AddReferencesRequestBuilderImpl(
        RequestHeader requestHeader, List<AddReferencesItem> referencesToAdd) {
      this.requestHeader = requestHeader;
      this.referencesToAdd = referencesToAdd;
    }

    public AddReferencesRequest build() {
      AddReferencesRequest addReferencesRequest =
          new AddReferencesRequest(requestHeader, referencesToAdd);
      return addReferencesRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AddReferencesRequest)) {
      return false;
    }
    AddReferencesRequest that = (AddReferencesRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getReferencesToAdd() == that.getReferencesToAdd())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequestHeader(), getReferencesToAdd());
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