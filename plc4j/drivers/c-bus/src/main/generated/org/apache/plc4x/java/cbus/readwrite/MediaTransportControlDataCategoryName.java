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
package org.apache.plc4x.java.cbus.readwrite;

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

public class MediaTransportControlDataCategoryName extends MediaTransportControlData
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final String categoryName;

  public MediaTransportControlDataCategoryName(
      MediaTransportControlCommandTypeContainer commandTypeContainer,
      byte mediaLinkGroup,
      String categoryName) {
    super(commandTypeContainer, mediaLinkGroup);
    this.categoryName = categoryName;
  }

  public String getCategoryName() {
    return categoryName;
  }

  @Override
  protected void serializeMediaTransportControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MediaTransportControlDataCategoryName");

    // Simple Field (categoryName)
    writeSimpleField(
        "categoryName",
        categoryName,
        writeString(writeBuffer, (((commandTypeContainer.getNumBytes()) - (1))) * (8)));

    writeBuffer.popContext("MediaTransportControlDataCategoryName");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MediaTransportControlDataCategoryName _value = this;

    // Simple field (categoryName)
    lengthInBits += (((commandTypeContainer.getNumBytes()) - (1))) * (8);

    return lengthInBits;
  }

  public static MediaTransportControlDataBuilder staticParseMediaTransportControlDataBuilder(
      ReadBuffer readBuffer, MediaTransportControlCommandTypeContainer commandTypeContainer)
      throws ParseException {
    readBuffer.pullContext("MediaTransportControlDataCategoryName");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    String categoryName =
        readSimpleField(
            "categoryName",
            readString(readBuffer, (((commandTypeContainer.getNumBytes()) - (1))) * (8)));

    readBuffer.closeContext("MediaTransportControlDataCategoryName");
    // Create the instance
    return new MediaTransportControlDataCategoryNameBuilderImpl(categoryName);
  }

  public static class MediaTransportControlDataCategoryNameBuilderImpl
      implements MediaTransportControlData.MediaTransportControlDataBuilder {
    private final String categoryName;

    public MediaTransportControlDataCategoryNameBuilderImpl(String categoryName) {
      this.categoryName = categoryName;
    }

    public MediaTransportControlDataCategoryName build(
        MediaTransportControlCommandTypeContainer commandTypeContainer, byte mediaLinkGroup) {
      MediaTransportControlDataCategoryName mediaTransportControlDataCategoryName =
          new MediaTransportControlDataCategoryName(
              commandTypeContainer, mediaLinkGroup, categoryName);
      return mediaTransportControlDataCategoryName;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MediaTransportControlDataCategoryName)) {
      return false;
    }
    MediaTransportControlDataCategoryName that = (MediaTransportControlDataCategoryName) o;
    return (getCategoryName() == that.getCategoryName()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCategoryName());
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
