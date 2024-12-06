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

public class QueryFirstResponse extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 618;
  }

  // Properties.
  protected final ResponseHeader responseHeader;
  protected final List<QueryDataSet> queryDataSets;
  protected final PascalByteString continuationPoint;
  protected final List<ParsingResult> parsingResults;
  protected final List<DiagnosticInfo> diagnosticInfos;
  protected final ContentFilterResult filterResult;

  public QueryFirstResponse(
      ResponseHeader responseHeader,
      List<QueryDataSet> queryDataSets,
      PascalByteString continuationPoint,
      List<ParsingResult> parsingResults,
      List<DiagnosticInfo> diagnosticInfos,
      ContentFilterResult filterResult) {
    super();
    this.responseHeader = responseHeader;
    this.queryDataSets = queryDataSets;
    this.continuationPoint = continuationPoint;
    this.parsingResults = parsingResults;
    this.diagnosticInfos = diagnosticInfos;
    this.filterResult = filterResult;
  }

  public ResponseHeader getResponseHeader() {
    return responseHeader;
  }

  public List<QueryDataSet> getQueryDataSets() {
    return queryDataSets;
  }

  public PascalByteString getContinuationPoint() {
    return continuationPoint;
  }

  public List<ParsingResult> getParsingResults() {
    return parsingResults;
  }

  public List<DiagnosticInfo> getDiagnosticInfos() {
    return diagnosticInfos;
  }

  public ContentFilterResult getFilterResult() {
    return filterResult;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("QueryFirstResponse");

    // Simple Field (responseHeader)
    writeSimpleField("responseHeader", responseHeader, writeComplex(writeBuffer));

    // Implicit Field (noOfQueryDataSets) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfQueryDataSets =
        (int) ((((getQueryDataSets()) == (null)) ? -(1) : COUNT(getQueryDataSets())));
    writeImplicitField("noOfQueryDataSets", noOfQueryDataSets, writeSignedInt(writeBuffer, 32));

    // Array Field (queryDataSets)
    writeComplexTypeArrayField("queryDataSets", queryDataSets, writeBuffer);

    // Simple Field (continuationPoint)
    writeSimpleField("continuationPoint", continuationPoint, writeComplex(writeBuffer));

    // Implicit Field (noOfParsingResults) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfParsingResults =
        (int) ((((getParsingResults()) == (null)) ? -(1) : COUNT(getParsingResults())));
    writeImplicitField("noOfParsingResults", noOfParsingResults, writeSignedInt(writeBuffer, 32));

    // Array Field (parsingResults)
    writeComplexTypeArrayField("parsingResults", parsingResults, writeBuffer);

    // Implicit Field (noOfDiagnosticInfos) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfDiagnosticInfos =
        (int) ((((getDiagnosticInfos()) == (null)) ? -(1) : COUNT(getDiagnosticInfos())));
    writeImplicitField("noOfDiagnosticInfos", noOfDiagnosticInfos, writeSignedInt(writeBuffer, 32));

    // Array Field (diagnosticInfos)
    writeComplexTypeArrayField("diagnosticInfos", diagnosticInfos, writeBuffer);

    // Simple Field (filterResult)
    writeSimpleField("filterResult", filterResult, writeComplex(writeBuffer));

    writeBuffer.popContext("QueryFirstResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    QueryFirstResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (responseHeader)
    lengthInBits += responseHeader.getLengthInBits();

    // Implicit Field (noOfQueryDataSets)
    lengthInBits += 32;

    // Array field
    if (queryDataSets != null) {
      int i = 0;
      for (QueryDataSet element : queryDataSets) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= queryDataSets.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (continuationPoint)
    lengthInBits += continuationPoint.getLengthInBits();

    // Implicit Field (noOfParsingResults)
    lengthInBits += 32;

    // Array field
    if (parsingResults != null) {
      int i = 0;
      for (ParsingResult element : parsingResults) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= parsingResults.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfDiagnosticInfos)
    lengthInBits += 32;

    // Array field
    if (diagnosticInfos != null) {
      int i = 0;
      for (DiagnosticInfo element : diagnosticInfos) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= diagnosticInfos.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (filterResult)
    lengthInBits += filterResult.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("QueryFirstResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ResponseHeader responseHeader =
        readSimpleField(
            "responseHeader",
            readComplex(
                () ->
                    (ResponseHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (394)),
                readBuffer));

    int noOfQueryDataSets = readImplicitField("noOfQueryDataSets", readSignedInt(readBuffer, 32));

    List<QueryDataSet> queryDataSets =
        readCountArrayField(
            "queryDataSets",
            readComplex(
                () -> (QueryDataSet) ExtensionObjectDefinition.staticParse(readBuffer, (int) (579)),
                readBuffer),
            noOfQueryDataSets);

    PascalByteString continuationPoint =
        readSimpleField(
            "continuationPoint",
            readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer));

    int noOfParsingResults = readImplicitField("noOfParsingResults", readSignedInt(readBuffer, 32));

    List<ParsingResult> parsingResults =
        readCountArrayField(
            "parsingResults",
            readComplex(
                () ->
                    (ParsingResult) ExtensionObjectDefinition.staticParse(readBuffer, (int) (612)),
                readBuffer),
            noOfParsingResults);

    int noOfDiagnosticInfos =
        readImplicitField("noOfDiagnosticInfos", readSignedInt(readBuffer, 32));

    List<DiagnosticInfo> diagnosticInfos =
        readCountArrayField(
            "diagnosticInfos",
            readComplex(() -> DiagnosticInfo.staticParse(readBuffer), readBuffer),
            noOfDiagnosticInfos);

    ContentFilterResult filterResult =
        readSimpleField(
            "filterResult",
            readComplex(
                () ->
                    (ContentFilterResult)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (609)),
                readBuffer));

    readBuffer.closeContext("QueryFirstResponse");
    // Create the instance
    return new QueryFirstResponseBuilderImpl(
        responseHeader,
        queryDataSets,
        continuationPoint,
        parsingResults,
        diagnosticInfos,
        filterResult);
  }

  public static class QueryFirstResponseBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ResponseHeader responseHeader;
    private final List<QueryDataSet> queryDataSets;
    private final PascalByteString continuationPoint;
    private final List<ParsingResult> parsingResults;
    private final List<DiagnosticInfo> diagnosticInfos;
    private final ContentFilterResult filterResult;

    public QueryFirstResponseBuilderImpl(
        ResponseHeader responseHeader,
        List<QueryDataSet> queryDataSets,
        PascalByteString continuationPoint,
        List<ParsingResult> parsingResults,
        List<DiagnosticInfo> diagnosticInfos,
        ContentFilterResult filterResult) {
      this.responseHeader = responseHeader;
      this.queryDataSets = queryDataSets;
      this.continuationPoint = continuationPoint;
      this.parsingResults = parsingResults;
      this.diagnosticInfos = diagnosticInfos;
      this.filterResult = filterResult;
    }

    public QueryFirstResponse build() {
      QueryFirstResponse queryFirstResponse =
          new QueryFirstResponse(
              responseHeader,
              queryDataSets,
              continuationPoint,
              parsingResults,
              diagnosticInfos,
              filterResult);
      return queryFirstResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof QueryFirstResponse)) {
      return false;
    }
    QueryFirstResponse that = (QueryFirstResponse) o;
    return (getResponseHeader() == that.getResponseHeader())
        && (getQueryDataSets() == that.getQueryDataSets())
        && (getContinuationPoint() == that.getContinuationPoint())
        && (getParsingResults() == that.getParsingResults())
        && (getDiagnosticInfos() == that.getDiagnosticInfos())
        && (getFilterResult() == that.getFilterResult())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getResponseHeader(),
        getQueryDataSets(),
        getContinuationPoint(),
        getParsingResults(),
        getDiagnosticInfos(),
        getFilterResult());
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