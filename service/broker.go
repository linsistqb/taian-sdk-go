// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package service

import (
	"fmt"
	"time"

	"github.com/hewenxiang/shanhe-sdk-go/config"
	"github.com/hewenxiang/shanhe-sdk-go/request"
	"github.com/hewenxiang/shanhe-sdk-go/request/data"
)

var _ fmt.State
var _ time.Time

type BrokerService struct {
	Config     *config.Config
	Properties *BrokerServiceProperties
}

type BrokerServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) Broker(zone string) (*BrokerService, error) {
	properties := &BrokerServiceProperties{
		Zone: &zone,
	}

	return &BrokerService{Config: s.Config, Properties: properties}, nil
}

func (s *BrokerService) DescribeBroker(i *DescribeBrokerInput) (*DescribeBrokerOutput, error) {
	if i == nil {
		i = &DescribeBrokerInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "CreateBrokers",
		RequestMethod: "GET",
	}

	x := &DescribeBrokerOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DescribeBrokerInput struct {
	Instances []*string `json:"instances" name:"instances" location:"params"`
}

func (v *DescribeBrokerInput) Validate() error {

	return nil
}

type DescribeBrokerOutput struct {
	Message *string `json:"message" name:"message"`
	Action  *string `json:"action" name:"action" location:"elements"`
	Brokers []*Brokers
	RetCode *int `json:"ret_code" name:"ret_code" location:"elements"`
}
