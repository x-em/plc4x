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

public class AggregateFilter extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 730;
  }

  // Properties.
  protected final long startTime;
  protected final NodeId aggregateType;
  protected final double processingInterval;
  protected final AggregateConfiguration aggregateConfiguration;

  public AggregateFilter(
      long startTime,
      NodeId aggregateType,
      double processingInterval,
      AggregateConfiguration aggregateConfiguration) {
    super();
    this.startTime = startTime;
    this.aggregateType = aggregateType;
    this.processingInterval = processingInterval;
    this.aggregateConfiguration = aggregateConfiguration;
  }

  public long getStartTime() {
    return startTime;
  }

  public NodeId getAggregateType() {
    return aggregateType;
  }

  public double getProcessingInterval() {
    return processingInterval;
  }

  public AggregateConfiguration getAggregateConfiguration() {
    return aggregateConfiguration;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AggregateFilter");

    // Simple Field (startTime)
    writeSimpleField("startTime", startTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (aggregateType)
    writeSimpleField("aggregateType", aggregateType, writeComplex(writeBuffer));

    // Simple Field (processingInterval)
    writeSimpleField("processingInterval", processingInterval, writeDouble(writeBuffer, 64));

    // Simple Field (aggregateConfiguration)
    writeSimpleField("aggregateConfiguration", aggregateConfiguration, writeComplex(writeBuffer));

    writeBuffer.popContext("AggregateFilter");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AggregateFilter _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (startTime)
    lengthInBits += 64;

    // Simple field (aggregateType)
    lengthInBits += aggregateType.getLengthInBits();

    // Simple field (processingInterval)
    lengthInBits += 64;

    // Simple field (aggregateConfiguration)
    lengthInBits += aggregateConfiguration.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("AggregateFilter");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long startTime = readSimpleField("startTime", readSignedLong(readBuffer, 64));

    NodeId aggregateType =
        readSimpleField(
            "aggregateType", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    double processingInterval = readSimpleField("processingInterval", readDouble(readBuffer, 64));

    AggregateConfiguration aggregateConfiguration =
        readSimpleField(
            "aggregateConfiguration",
            readComplex(
                () ->
                    (AggregateConfiguration)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (950)),
                readBuffer));

    readBuffer.closeContext("AggregateFilter");
    // Create the instance
    return new AggregateFilterBuilderImpl(
        startTime, aggregateType, processingInterval, aggregateConfiguration);
  }

  public static class AggregateFilterBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long startTime;
    private final NodeId aggregateType;
    private final double processingInterval;
    private final AggregateConfiguration aggregateConfiguration;

    public AggregateFilterBuilderImpl(
        long startTime,
        NodeId aggregateType,
        double processingInterval,
        AggregateConfiguration aggregateConfiguration) {
      this.startTime = startTime;
      this.aggregateType = aggregateType;
      this.processingInterval = processingInterval;
      this.aggregateConfiguration = aggregateConfiguration;
    }

    public AggregateFilter build() {
      AggregateFilter aggregateFilter =
          new AggregateFilter(startTime, aggregateType, processingInterval, aggregateConfiguration);
      return aggregateFilter;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AggregateFilter)) {
      return false;
    }
    AggregateFilter that = (AggregateFilter) o;
    return (getStartTime() == that.getStartTime())
        && (getAggregateType() == that.getAggregateType())
        && (getProcessingInterval() == that.getProcessingInterval())
        && (getAggregateConfiguration() == that.getAggregateConfiguration())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getStartTime(),
        getAggregateType(),
        getProcessingInterval(),
        getAggregateConfiguration());
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
