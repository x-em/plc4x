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

public class CreateSessionRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 461;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final ApplicationDescription clientDescription;
  protected final PascalString serverUri;
  protected final PascalString endpointUrl;
  protected final PascalString sessionName;
  protected final PascalByteString clientNonce;
  protected final PascalByteString clientCertificate;
  protected final double requestedSessionTimeout;
  protected final long maxResponseMessageSize;

  public CreateSessionRequest(
      RequestHeader requestHeader,
      ApplicationDescription clientDescription,
      PascalString serverUri,
      PascalString endpointUrl,
      PascalString sessionName,
      PascalByteString clientNonce,
      PascalByteString clientCertificate,
      double requestedSessionTimeout,
      long maxResponseMessageSize) {
    super();
    this.requestHeader = requestHeader;
    this.clientDescription = clientDescription;
    this.serverUri = serverUri;
    this.endpointUrl = endpointUrl;
    this.sessionName = sessionName;
    this.clientNonce = clientNonce;
    this.clientCertificate = clientCertificate;
    this.requestedSessionTimeout = requestedSessionTimeout;
    this.maxResponseMessageSize = maxResponseMessageSize;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public ApplicationDescription getClientDescription() {
    return clientDescription;
  }

  public PascalString getServerUri() {
    return serverUri;
  }

  public PascalString getEndpointUrl() {
    return endpointUrl;
  }

  public PascalString getSessionName() {
    return sessionName;
  }

  public PascalByteString getClientNonce() {
    return clientNonce;
  }

  public PascalByteString getClientCertificate() {
    return clientCertificate;
  }

  public double getRequestedSessionTimeout() {
    return requestedSessionTimeout;
  }

  public long getMaxResponseMessageSize() {
    return maxResponseMessageSize;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CreateSessionRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Simple Field (clientDescription)
    writeSimpleField("clientDescription", clientDescription, writeComplex(writeBuffer));

    // Simple Field (serverUri)
    writeSimpleField("serverUri", serverUri, writeComplex(writeBuffer));

    // Simple Field (endpointUrl)
    writeSimpleField("endpointUrl", endpointUrl, writeComplex(writeBuffer));

    // Simple Field (sessionName)
    writeSimpleField("sessionName", sessionName, writeComplex(writeBuffer));

    // Simple Field (clientNonce)
    writeSimpleField("clientNonce", clientNonce, writeComplex(writeBuffer));

    // Simple Field (clientCertificate)
    writeSimpleField("clientCertificate", clientCertificate, writeComplex(writeBuffer));

    // Simple Field (requestedSessionTimeout)
    writeSimpleField(
        "requestedSessionTimeout", requestedSessionTimeout, writeDouble(writeBuffer, 64));

    // Simple Field (maxResponseMessageSize)
    writeSimpleField(
        "maxResponseMessageSize", maxResponseMessageSize, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("CreateSessionRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CreateSessionRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (clientDescription)
    lengthInBits += clientDescription.getLengthInBits();

    // Simple field (serverUri)
    lengthInBits += serverUri.getLengthInBits();

    // Simple field (endpointUrl)
    lengthInBits += endpointUrl.getLengthInBits();

    // Simple field (sessionName)
    lengthInBits += sessionName.getLengthInBits();

    // Simple field (clientNonce)
    lengthInBits += clientNonce.getLengthInBits();

    // Simple field (clientCertificate)
    lengthInBits += clientCertificate.getLengthInBits();

    // Simple field (requestedSessionTimeout)
    lengthInBits += 64;

    // Simple field (maxResponseMessageSize)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("CreateSessionRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    ApplicationDescription clientDescription =
        readSimpleField(
            "clientDescription",
            readComplex(
                () ->
                    (ApplicationDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (310)),
                readBuffer));

    PascalString serverUri =
        readSimpleField(
            "serverUri", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString endpointUrl =
        readSimpleField(
            "endpointUrl", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString sessionName =
        readSimpleField(
            "sessionName", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalByteString clientNonce =
        readSimpleField(
            "clientNonce", readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer));

    PascalByteString clientCertificate =
        readSimpleField(
            "clientCertificate",
            readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer));

    double requestedSessionTimeout =
        readSimpleField("requestedSessionTimeout", readDouble(readBuffer, 64));

    long maxResponseMessageSize =
        readSimpleField("maxResponseMessageSize", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("CreateSessionRequest");
    // Create the instance
    return new CreateSessionRequestBuilderImpl(
        requestHeader,
        clientDescription,
        serverUri,
        endpointUrl,
        sessionName,
        clientNonce,
        clientCertificate,
        requestedSessionTimeout,
        maxResponseMessageSize);
  }

  public static class CreateSessionRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final ApplicationDescription clientDescription;
    private final PascalString serverUri;
    private final PascalString endpointUrl;
    private final PascalString sessionName;
    private final PascalByteString clientNonce;
    private final PascalByteString clientCertificate;
    private final double requestedSessionTimeout;
    private final long maxResponseMessageSize;

    public CreateSessionRequestBuilderImpl(
        RequestHeader requestHeader,
        ApplicationDescription clientDescription,
        PascalString serverUri,
        PascalString endpointUrl,
        PascalString sessionName,
        PascalByteString clientNonce,
        PascalByteString clientCertificate,
        double requestedSessionTimeout,
        long maxResponseMessageSize) {
      this.requestHeader = requestHeader;
      this.clientDescription = clientDescription;
      this.serverUri = serverUri;
      this.endpointUrl = endpointUrl;
      this.sessionName = sessionName;
      this.clientNonce = clientNonce;
      this.clientCertificate = clientCertificate;
      this.requestedSessionTimeout = requestedSessionTimeout;
      this.maxResponseMessageSize = maxResponseMessageSize;
    }

    public CreateSessionRequest build() {
      CreateSessionRequest createSessionRequest =
          new CreateSessionRequest(
              requestHeader,
              clientDescription,
              serverUri,
              endpointUrl,
              sessionName,
              clientNonce,
              clientCertificate,
              requestedSessionTimeout,
              maxResponseMessageSize);
      return createSessionRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CreateSessionRequest)) {
      return false;
    }
    CreateSessionRequest that = (CreateSessionRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getClientDescription() == that.getClientDescription())
        && (getServerUri() == that.getServerUri())
        && (getEndpointUrl() == that.getEndpointUrl())
        && (getSessionName() == that.getSessionName())
        && (getClientNonce() == that.getClientNonce())
        && (getClientCertificate() == that.getClientCertificate())
        && (getRequestedSessionTimeout() == that.getRequestedSessionTimeout())
        && (getMaxResponseMessageSize() == that.getMaxResponseMessageSize())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getClientDescription(),
        getServerUri(),
        getEndpointUrl(),
        getSessionName(),
        getClientNonce(),
        getClientCertificate(),
        getRequestedSessionTimeout(),
        getMaxResponseMessageSize());
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
