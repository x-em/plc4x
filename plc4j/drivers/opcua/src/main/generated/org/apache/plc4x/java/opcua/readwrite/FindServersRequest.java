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

public class FindServersRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "422";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final PascalString endpointUrl;
  protected final int noOfLocaleIds;
  protected final List<PascalString> localeIds;
  protected final int noOfServerUris;
  protected final List<PascalString> serverUris;

  public FindServersRequest(
      ExtensionObjectDefinition requestHeader,
      PascalString endpointUrl,
      int noOfLocaleIds,
      List<PascalString> localeIds,
      int noOfServerUris,
      List<PascalString> serverUris) {
    super();
    this.requestHeader = requestHeader;
    this.endpointUrl = endpointUrl;
    this.noOfLocaleIds = noOfLocaleIds;
    this.localeIds = localeIds;
    this.noOfServerUris = noOfServerUris;
    this.serverUris = serverUris;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public PascalString getEndpointUrl() {
    return endpointUrl;
  }

  public int getNoOfLocaleIds() {
    return noOfLocaleIds;
  }

  public List<PascalString> getLocaleIds() {
    return localeIds;
  }

  public int getNoOfServerUris() {
    return noOfServerUris;
  }

  public List<PascalString> getServerUris() {
    return serverUris;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("FindServersRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (endpointUrl)
    writeSimpleField("endpointUrl", endpointUrl, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfLocaleIds)
    writeSimpleField("noOfLocaleIds", noOfLocaleIds, writeSignedInt(writeBuffer, 32));

    // Array Field (localeIds)
    writeComplexTypeArrayField("localeIds", localeIds, writeBuffer);

    // Simple Field (noOfServerUris)
    writeSimpleField("noOfServerUris", noOfServerUris, writeSignedInt(writeBuffer, 32));

    // Array Field (serverUris)
    writeComplexTypeArrayField("serverUris", serverUris, writeBuffer);

    writeBuffer.popContext("FindServersRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FindServersRequest _value = this;

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (endpointUrl)
    lengthInBits += endpointUrl.getLengthInBits();

    // Simple field (noOfLocaleIds)
    lengthInBits += 32;

    // Array field
    if (localeIds != null) {
      int i = 0;
      for (PascalString element : localeIds) {
        boolean last = ++i >= localeIds.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfServerUris)
    lengthInBits += 32;

    // Array field
    if (serverUris != null) {
      int i = 0;
      for (PascalString element : serverUris) {
        boolean last = ++i >= serverUris.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("FindServersRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    PascalString endpointUrl =
        readSimpleField(
            "endpointUrl",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfLocaleIds = readSimpleField("noOfLocaleIds", readSignedInt(readBuffer, 32));

    List<PascalString> localeIds =
        readCountArrayField(
            "localeIds",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfLocaleIds);

    int noOfServerUris = readSimpleField("noOfServerUris", readSignedInt(readBuffer, 32));

    List<PascalString> serverUris =
        readCountArrayField(
            "serverUris",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfServerUris);

    readBuffer.closeContext("FindServersRequest");
    // Create the instance
    return new FindServersRequestBuilderImpl(
        requestHeader, endpointUrl, noOfLocaleIds, localeIds, noOfServerUris, serverUris);
  }

  public static class FindServersRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final PascalString endpointUrl;
    private final int noOfLocaleIds;
    private final List<PascalString> localeIds;
    private final int noOfServerUris;
    private final List<PascalString> serverUris;

    public FindServersRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        PascalString endpointUrl,
        int noOfLocaleIds,
        List<PascalString> localeIds,
        int noOfServerUris,
        List<PascalString> serverUris) {
      this.requestHeader = requestHeader;
      this.endpointUrl = endpointUrl;
      this.noOfLocaleIds = noOfLocaleIds;
      this.localeIds = localeIds;
      this.noOfServerUris = noOfServerUris;
      this.serverUris = serverUris;
    }

    public FindServersRequest build() {
      FindServersRequest findServersRequest =
          new FindServersRequest(
              requestHeader, endpointUrl, noOfLocaleIds, localeIds, noOfServerUris, serverUris);
      return findServersRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FindServersRequest)) {
      return false;
    }
    FindServersRequest that = (FindServersRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getEndpointUrl() == that.getEndpointUrl())
        && (getNoOfLocaleIds() == that.getNoOfLocaleIds())
        && (getLocaleIds() == that.getLocaleIds())
        && (getNoOfServerUris() == that.getNoOfServerUris())
        && (getServerUris() == that.getServerUris())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getEndpointUrl(),
        getNoOfLocaleIds(),
        getLocaleIds(),
        getNoOfServerUris(),
        getServerUris());
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
