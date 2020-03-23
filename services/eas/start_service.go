package eas

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

// StartService invokes the eas.StartService API synchronously
// api document: https://help.aliyun.com/api/eas/startservice.html
func (client *Client) StartService(request *StartServiceRequest) (response *StartServiceResponse, err error) {
	response = CreateStartServiceResponse()
	err = client.DoAction(request, response)
	return
}

// StartServiceWithChan invokes the eas.StartService API asynchronously
// api document: https://help.aliyun.com/api/eas/startservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartServiceWithChan(request *StartServiceRequest) (<-chan *StartServiceResponse, <-chan error) {
	responseChan := make(chan *StartServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartService(request)
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

// StartServiceWithCallback invokes the eas.StartService API asynchronously
// api document: https://help.aliyun.com/api/eas/startservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartServiceWithCallback(request *StartServiceRequest, callback func(response *StartServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartServiceResponse
		var err error
		defer close(result)
		response, err = client.StartService(request)
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

// StartServiceRequest is the request struct for api StartService
type StartServiceRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"service_name"`
	Region      string `position:"Path" name:"region"`
}

// StartServiceResponse is the response struct for api StartService
type StartServiceResponse struct {
	*responses.BaseResponse
}

// CreateStartServiceRequest creates a request to invoke StartService API
func CreateStartServiceRequest() (request *StartServiceRequest) {
	request = &StartServiceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "StartService", "/api/services/[region]/[service_name]/start", "", "")
	request.Method = requests.PUT
	return
}

// CreateStartServiceResponse creates a response to parse from StartService response
func CreateStartServiceResponse() (response *StartServiceResponse) {
	response = &StartServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}