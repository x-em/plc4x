/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi.codegen.fields;

import org.apache.commons.lang3.ArrayUtils;
import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.spi.codegen.io.DataReader;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.WithReaderArgs;
import org.apache.plc4x.java.spi.generation.WithReaderWriterArgs;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.function.Supplier;

public class FieldReaderArray<T> implements FieldReader<T> {

    @Override
    public T readField(String logicalName, DataReader<T> dataReader, WithReaderArgs... readerArgs) throws ParseException {
        throw new NotImplementedException("use dedicated methods instead");
    }

    public List<T> readFieldCount(String logicalName, DataReader<T> dataReader, long count, WithReaderArgs... readerArgs) throws ParseException {
        if (count > Integer.MAX_VALUE) {
            throw new ParseException("Array count of " + count + " exceeds the maximum allowed count of " + Integer.MAX_VALUE);
        }
        // Ensure we have the render as list argument present
        readerArgs = ArrayUtils.add(readerArgs, WithReaderWriterArgs.WithRenderAsList(true));
        dataReader.pullContext(logicalName, readerArgs);
        int itemCount = Math.max(0, (int) count);
        List<T> result = new ArrayList<>(itemCount);
        for (int curItem = 0; curItem < itemCount; curItem++) {
            result.set(curItem, dataReader.read("", readerArgs));
        }
        dataReader.closeContext(logicalName, readerArgs);
        return result;
    }

    public List<T> readFieldLength(String logicalName, DataReader<T> dataReader, int length, WithReaderArgs... readerArgs) throws ParseException {
        // Ensure we have the render as list argument present
        readerArgs = ArrayUtils.add(readerArgs, WithReaderWriterArgs.WithRenderAsList(true));
        dataReader.pullContext(logicalName, readerArgs);
        List<T> result = new LinkedList<>();
        while (dataReader.getPos() < length) {
            result.add(dataReader.read("", readerArgs));
        }
        dataReader.closeContext(logicalName, readerArgs);
        return result;
    }

    public List<T> readFieldTerminated(String logicalName, DataReader<T> dataReader, Supplier<Boolean> termination, WithReaderArgs... readerArgs) throws ParseException {
        // Ensure we have the render as list argument present
        readerArgs = ArrayUtils.add(readerArgs, WithReaderWriterArgs.WithRenderAsList(true));
        dataReader.pullContext(logicalName, readerArgs);
        List<T> result = new LinkedList<>();
        while (!termination.get()) {
            result.add(dataReader.read("", readerArgs));
        }
        dataReader.closeContext(logicalName, readerArgs);
        return result;
    }

}