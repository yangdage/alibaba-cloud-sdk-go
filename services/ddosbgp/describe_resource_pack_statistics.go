package ddosbgp

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeResourcePackStatistics invokes the ddosbgp.DescribeResourcePackStatistics API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeresourcepackstatistics.html
func (client *Client) DescribeResourcePackStatistics(request *DescribeResourcePackStatisticsRequest) (response *DescribeResourcePackStatisticsResponse, err error) {
	response = CreateDescribeResourcePackStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourcePackStatisticsWithChan invokes the ddosbgp.DescribeResourcePackStatistics API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeresourcepackstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourcePackStatisticsWithChan(request *DescribeResourcePackStatisticsRequest) (<-chan *DescribeResourcePackStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeResourcePackStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourcePackStatistics(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeResourcePackStatisticsWithCallback invokes the ddosbgp.DescribeResourcePackStatistics API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeresourcepackstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourcePackStatisticsWithCallback(request *DescribeResourcePackStatisticsRequest, callback func(response *DescribeResourcePackStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourcePackStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourcePackStatistics(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeResourcePackStatisticsRequest is the request struct for api DescribeResourcePackStatistics
type DescribeResourcePackStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp        string `position:"Query" name:"SourceIp"`
	DdosRegionId    string `position:"Query" name:"DdosRegionId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
}

// DescribeResourcePackStatisticsResponse is the response struct for api DescribeResourcePackStatistics
type DescribeResourcePackStatisticsResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	AvailablePackNum  int    `json:"AvailablePackNum" xml:"AvailablePackNum"`
	TotalCurrCapacity int64  `json:"TotalCurrCapacity" xml:"TotalCurrCapacity"`
	TotalUsedCapacity int64  `json:"TotalUsedCapacity" xml:"TotalUsedCapacity"`
	TotalInitCapacity int64  `json:"TotalInitCapacity" xml:"TotalInitCapacity"`
}

// CreateDescribeResourcePackStatisticsRequest creates a request to invoke DescribeResourcePackStatistics API
func CreateDescribeResourcePackStatisticsRequest() (request *DescribeResourcePackStatisticsRequest) {
	request = &DescribeResourcePackStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeResourcePackStatistics", "ddosbgp", "openAPI")
	return
}

// CreateDescribeResourcePackStatisticsResponse creates a response to parse from DescribeResourcePackStatistics response
func CreateDescribeResourcePackStatisticsResponse() (response *DescribeResourcePackStatisticsResponse) {
	response = &DescribeResourcePackStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}