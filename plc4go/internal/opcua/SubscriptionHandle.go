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

package opcua

import (
	"time"

	spiModel "github.com/apache/plc4x/plc4go/spi/model"
)

type SubscriptionHandle struct {
	*spiModel.DefaultPlcSubscriptionHandle
	tagName          string
	tag              any
	subscriptionType spiModel.SubscriptionType
	interval         time.Duration
}

func NewSubscriptionHandle(subscriber *Subscriber, tagName string, tag any, subscriptionType spiModel.SubscriptionType, interval time.Duration) *SubscriptionHandle {
	s := &SubscriptionHandle{
		tagName:          tagName,
		tag:              tag,
		subscriptionType: subscriptionType,
		interval:         interval,
	}
	s.DefaultPlcSubscriptionHandle = spiModel.NewDefaultPlcSubscriptionHandleWithHandleToRegister(subscriber, s)
	return s
}
