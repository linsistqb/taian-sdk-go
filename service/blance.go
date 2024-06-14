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

type BalanceService struct {
	Config     *config.Config
	Properties *BalanceServiceProperties
}

type BalanceServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) Balance(zone string) (*BalanceService, error) {
	properties := &BalanceServiceProperties{
		Zone: &zone,
	}

	return &BalanceService{Config: s.Config, Properties: properties}, nil
}

func (s *BalanceService) DescribeBalance(i *DescribeBalanceInput) (*DescribeBalanceOutput, error) {
	if i == nil {
		i = &DescribeBalanceInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "GetBalance",
		RequestMethod: "GET",
	}

	x := &DescribeBalanceOutput{}
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

type DescribeBalanceInput struct {
	Balance    *string   `json:"access_keys" name:"access_keys" location:"params"`
	Limit      *int      `json:"limit" name:"limit" default:"20" location:"params"`
	Offset     *int      `json:"offset" name:"offset" default:"0" location:"params"`
	Owner      *string   `json:"owner" name:"owner" location:"params"`
	SearchWord *string   `json:"search_word" name:"search_word" location:"params"`
	Status     []*string `json:"status" name:"status" location:"params"`
	Verbose    *int      `json:"verbose" name:"verbose" default:"0" location:"params"`
	Zone       *string   `json:"zone" name:"zone"` // Required
}

func (v *DescribeBalanceInput) Validate() error {

	return nil
}

type DescribeBalanceOutput struct {
	Action  *string `json:"action" name:"action" location:"elements"`
	Balance *string
	Bonus   *string
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
	UserId  *string `json:"user_id" name:"user_id" location:"elements"`

	PaidMode *string `json:"paid_mode" name:"paid_mode" location:"elements"`
}
