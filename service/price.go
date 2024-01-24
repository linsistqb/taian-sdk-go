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

type PriceService struct {
	Config     *config.Config
	Properties *PriceServiceProperties
}

type PriceServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) Price(zone string) (*PriceService, error) {
	properties := &PriceServiceProperties{
		Zone: &zone,
	}

	return &PriceService{Config: s.Config, Properties: properties}, nil
}

func (s *PriceService) DescribePrice(i *DescribePriceInput) (*DescribePriceOutput, error) {
	if i == nil {
		i = &DescribePriceInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "GetPrice",
		RequestMethod: "GET",
	}

	x := &DescribePriceOutput{}
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

type Resources struct {
	Type         *string `json:"duration" name:"duration" location:"params"`
	InstanceType *string `json:"instance_type" name:"instance_type" location:"params"`
	ImageId    *string `json:"image_id" name:"image_id" default:"img-7vnii9pb" location:"params"`
	Sequence   *int    `json:"sequence" name:"sequence" default:"0" location:"params"`
	VolumeType *string `json:"volume_type" name:"volume_type" default:"0" location:"params"`
	Size *int `json:"size" name:"size" location:"params"`
}
type DescribePriceInput struct {
	Resources []*Resources `json:"access_keys" name:"access_keys" location:"params"`
	Duration  *int         `json:"duration" name:"duration" location:"params"`
}

func (v *DescribePriceInput) Validate() error {

	return nil
}

type DescribePriceOutput struct {
	Message    *string  `json:"message" name:"message"`
	Action     *string  `json:"action" name:"action" location:"elements"`
	PriceSet   []*Price `json:"price_set" name:"price_set" location:"elements"`
	RetCode    *int     `json:"ret_code" name:"ret_code" location:"elements"`

}
