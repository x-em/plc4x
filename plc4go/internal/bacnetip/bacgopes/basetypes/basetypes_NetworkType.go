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

package basetypes

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type NetworkType struct {
	*Enumerated
	vendorRange  vendorRange
	enumerations map[string]uint64
}

func NewNetworkType(arg Arg) (*NetworkType, error) {
	s := &NetworkType{
		enumerations: map[string]uint64{"ethernet": 0,
			"arcnet":  1,
			"mstp":    2,
			"ptp":     3,
			"lontalk": 4,
			"ipv4":    5,
			"zigbee":  6,
			"virtual": 7,
			// "non-bacnet":  8  Removed in Version 1, Revision 18,
			"ipv6":   9,
			"serial": 10,
		},
	}
	var err error
	s.Enumerated, err = NewEnumerated(NoArgs)
	if err != nil {
		return nil, errors.Wrap(err, "error creating enumerated")
	}
	return s, nil
}