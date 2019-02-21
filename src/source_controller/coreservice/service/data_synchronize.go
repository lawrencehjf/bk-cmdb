/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/coreservice/core"
)

func (s *coreService) SynchronizeInstance(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {
	inputData := &metadata.SynchronizeParameter{}
	if err := data.MarshalJSONInto(inputData); nil != err {
		blog.Errorf("SynchronizeInstance MarshalJSONInto error, err:%s,input:%v,rid:%s", err.Error(), data, params.ReqID)
		return nil, err
	}
	inputData.OperateDataType = metadata.SynchronizeOperateDataTypeInstance
	return s.core.DataSynchronizeOperation().SynchronizeInstanceAdapter(params, inputData)
}

func (s *coreService) SynchronizeModel(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := &metadata.SynchronizeParameter{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		blog.Errorf("SynchronizeModel MarshalJSONInto error, err:%s,input:%v,rid:%s", err.Error(), data, params.ReqID)
		return nil, err
	}
	inputData.OperateDataType = metadata.SynchronizeOperateDataTypeModel
	return s.core.DataSynchronizeOperation().SynchronizeModelAdapter(params, inputData)
}

func (s *coreService) SynchronizeAssociation(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := &metadata.SynchronizeParameter{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		blog.Errorf("SynchronizeAssociation MarshalJSONInto error, err:%s,input:%v,rid:%s", err.Error(), data, params.ReqID)
		return nil, err
	}
	inputData.OperateDataType = metadata.SynchronizeOperateDataTypeAssociation
	return s.core.DataSynchronizeOperation().SynchronizeAssociationAdapter(params, inputData)
}

func (s *coreService) SynchronizeFetch(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := &metadata.SynchronizeFetchInfoParameter{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		blog.Errorf("SynchronizeAssociation MarshalJSONInto error, err:%s,input:%v,rid:%s", err.Error(), data, params.ReqID)
		return nil, err
	}
	info, cnt, err := s.core.DataSynchronizeOperation().GetAssociationInfo(params, inputData)
	if err != nil {
		blog.Errorf("SynchronizeAssociation GetAssociationInfo error, err:%s,input:%v,rid:%s", err.Error(), data, params.ReqID)
		return nil, err
	}
	return mapstr.MapStr{"info": info, "count": cnt}, nil
}
