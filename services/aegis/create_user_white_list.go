package aegis

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

// invoke CreateUserWhiteList api with *CreateUserWhiteListRequest synchronously
// api document: https://help.aliyun.com/api/aegis/createuserwhitelist.html
func (client *Client) CreateUserWhiteList(request *CreateUserWhiteListRequest) (response *CreateUserWhiteListResponse, err error) {
	response = CreateCreateUserWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// invoke CreateUserWhiteList api with *CreateUserWhiteListRequest asynchronously
// api document: https://help.aliyun.com/api/aegis/createuserwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserWhiteListWithChan(request *CreateUserWhiteListRequest) (<-chan *CreateUserWhiteListResponse, <-chan error) {
	responseChan := make(chan *CreateUserWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserWhiteList(request)
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

// invoke CreateUserWhiteList api with *CreateUserWhiteListRequest asynchronously
// api document: https://help.aliyun.com/api/aegis/createuserwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserWhiteListWithCallback(request *CreateUserWhiteListRequest, callback func(response *CreateUserWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserWhiteListResponse
		var err error
		defer close(result)
		response, err = client.CreateUserWhiteList(request)
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

type CreateUserWhiteListRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RiskIdList      string           `position:"Query" name:"RiskIdList"`
}

type CreateUserWhiteListResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// create a request to invoke CreateUserWhiteList API
func CreateCreateUserWhiteListRequest() (request *CreateUserWhiteListRequest) {
	request = &CreateUserWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "CreateUserWhiteList", "vipaegis", "openAPI")
	return
}

// create a response to parse from CreateUserWhiteList response
func CreateCreateUserWhiteListResponse() (response *CreateUserWhiteListResponse) {
	response = &CreateUserWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}