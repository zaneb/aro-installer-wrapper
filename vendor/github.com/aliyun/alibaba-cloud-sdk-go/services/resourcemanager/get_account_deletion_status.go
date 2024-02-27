package resourcemanager

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

// GetAccountDeletionStatus invokes the resourcemanager.GetAccountDeletionStatus API synchronously
func (client *Client) GetAccountDeletionStatus(request *GetAccountDeletionStatusRequest) (response *GetAccountDeletionStatusResponse, err error) {
	response = CreateGetAccountDeletionStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccountDeletionStatusWithChan invokes the resourcemanager.GetAccountDeletionStatus API asynchronously
func (client *Client) GetAccountDeletionStatusWithChan(request *GetAccountDeletionStatusRequest) (<-chan *GetAccountDeletionStatusResponse, <-chan error) {
	responseChan := make(chan *GetAccountDeletionStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountDeletionStatus(request)
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

// GetAccountDeletionStatusWithCallback invokes the resourcemanager.GetAccountDeletionStatus API asynchronously
func (client *Client) GetAccountDeletionStatusWithCallback(request *GetAccountDeletionStatusRequest, callback func(response *GetAccountDeletionStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccountDeletionStatusResponse
		var err error
		defer close(result)
		response, err = client.GetAccountDeletionStatus(request)
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

// GetAccountDeletionStatusRequest is the request struct for api GetAccountDeletionStatus
type GetAccountDeletionStatusRequest struct {
	*requests.RpcRequest
	AccountId string `position:"Query" name:"AccountId"`
}

// GetAccountDeletionStatusResponse is the response struct for api GetAccountDeletionStatus
type GetAccountDeletionStatusResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	RdAccountDeletionStatus RdAccountDeletionStatus `json:"RdAccountDeletionStatus" xml:"RdAccountDeletionStatus"`
}

// CreateGetAccountDeletionStatusRequest creates a request to invoke GetAccountDeletionStatus API
func CreateGetAccountDeletionStatusRequest() (request *GetAccountDeletionStatusRequest) {
	request = &GetAccountDeletionStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "GetAccountDeletionStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAccountDeletionStatusResponse creates a response to parse from GetAccountDeletionStatus response
func CreateGetAccountDeletionStatusResponse() (response *GetAccountDeletionStatusResponse) {
	response = &GetAccountDeletionStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}