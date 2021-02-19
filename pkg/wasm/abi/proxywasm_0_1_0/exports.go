/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proxywasm_0_1_0

import "mosn.io/mosn/pkg/log"

func (a *abiContext) waitAsyncHttpCallout() {
	if a.httpCallout != nil && a.httpCallout.asyncRetChan != nil {
		<-a.httpCallout.asyncRetChan
	}
}

func (a *abiContext) ProxyOnContextCreate(contextId int32, parentContextId int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] ProxyOnContextCreate contextID: %v, parentContextId: %v", contextId, parentContextId)
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_context_create")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_context_create, err: %v", err)
		return err
	}

	_, err = ff.Call(contextId, parentContextId)
	if err != nil {
		return err
	}

	a.waitAsyncHttpCallout()

	return nil
}

func (a *abiContext) ProxyOnDone(contextId int32) (int32, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Infof("[proxywasm_0_1_0][export] ProxyOnDone contextID: %v", contextId)
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_done")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_done, err: %v", err)
		return 0, err
	}

	res, err := ff.Call(contextId)
	if err != nil {
		return 0, err
	}

	a.waitAsyncHttpCallout()

	return res.(int32), nil
}

func (a *abiContext) ProxyOnLog(contextId int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[wasmer][instance] ProxyOnLog")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_log")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_log, err: %v", err)
		return err
	}

	_, err = ff.Call(contextId)

	return err
}

func (a *abiContext) ProxyOnVmStart(rootContextId int32, vmConfigurationSize int32) (int32, error) {
	log.DefaultLogger.Infof("[proxywasm_0_1_0][export] ProxyOnVmStart rootContextId: %v, vmConfigurationSize: %v", rootContextId, vmConfigurationSize)

	ff, err := a.instance.GetExportsFunc("proxy_on_vm_start")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_vm_start, err: %v", err)
		return 0, err
	}

	res, err := ff.Call(rootContextId, vmConfigurationSize)
	if err != nil {
		return 0, err
	}

	a.waitAsyncHttpCallout()

	return res.(int32), nil
}

func (a *abiContext) ProxyOnDelete(contextId int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnDelete")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_delete")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_delete, err: %v", err)
		return err
	}

	_, err = ff.Call(contextId)

	a.waitAsyncHttpCallout()

	return err
}

func (a *abiContext) ProxyOnConfigure(rootContextId int32, configurationSize int32) (int32, error) {
	log.DefaultLogger.Infof("[proxywasm_0_1_0][export] ProxyOnConfigure rootContextId: %v, configurationSize: %v", rootContextId, configurationSize)

	ff, err := a.instance.GetExportsFunc("proxy_on_configure")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_configure, err: %v", err)
		return 0, err
	}

	res, err := ff.Call(rootContextId, configurationSize)
	if err != nil {
		return 0, err
	}

	a.waitAsyncHttpCallout()

	return res.(int32), nil
}

func (a *abiContext) ProxyOnTick(rootContextId int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnTick")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_tick")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_tick, err: %v", err)
		return err
	}

	_, err = ff.Call(rootContextId)

	a.waitAsyncHttpCallout()

	return err
}

func (a *abiContext) ProxyOnNewConnection(contextId int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnNewConnection")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_new_connection")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_new_connection, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId)

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), err
}

func (a *abiContext) ProxyOnDownstreamData(contextId int32, dataLength int32, endOfStream int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnDownstreamData")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_downstream_data")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_downstream_data, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, dataLength, endOfStream)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnDownstreamConnectionClose(contextId int32, closeType int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnDownstreamConnectionClose")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_downstream_connection_close")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_downstream_connection_close, err: %v", err)
		return err
	}

	_, err = ff.Call(contextId, closeType)

	a.waitAsyncHttpCallout()

	return err
}

func (a *abiContext) ProxyOnUpstreamData(contextId int32, dataLength int32, endOfStream int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnUpstreamData")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_upstream_data")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_upstream_data, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, dataLength, endOfStream)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnUpstreamConnectionClose(contextId int32, closeType int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnUpstreamConnectionClose")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_upstream_connection_close")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_upstream_connection_close, err: %v", err)
		return err
	}

	_, err = ff.Call(contextId, closeType)

	a.waitAsyncHttpCallout()

	return err
}

func (a *abiContext) ProxyOnRequestHeaders(contextID int32, numHeaders int32, endOfStream int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnRequestHeaders")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_request_headers")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_request_headers, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextID, numHeaders, endOfStream)
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to call proxy_on_request_headers func, err: %v", err)
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnRequestBody(contextId int32, bodyBufferLength int32, endOfStream int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnRequestBody")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_request_body")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_request_body, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, bodyBufferLength, endOfStream)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnRequestTrailers(contextId int32, trailers int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnRequestTrailers")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_request_trailers")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_request_trailers, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, trailers)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnRequestMetadata(contextId int32, nElements int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnRequestMetadata")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_request_metadata")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_request_metadata, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, nElements)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnResponseHeaders(contextId int32, headers int32, endOfStream int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnResponseHeaders")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_response_headers")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_response_headers, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, headers, endOfStream)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnResponseBody(contextId int32, bodyBufferLength int32, endOfStream int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnResponseBody")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_response_body")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_response_body, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, bodyBufferLength, endOfStream)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnResponseTrailers(contextId int32, trailers int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnResponseTrailers")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_response_trailers")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_response_trailers, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, trailers)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnResponseMetadata(contextId int32, nElements int32) (Action, error) {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnResponseMetadata")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_response_metadata")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_response_metadata, err: %v", err)
		return ActionPause, err
	}

	res, err := ff.Call(contextId, nElements)
	if err != nil {
		return ActionPause, err
	}

	a.waitAsyncHttpCallout()

	return Action(res.(int32)), nil
}

func (a *abiContext) ProxyOnHttpCallResponse(contextId int32, token int32, headers int32, bodySize int32, trailers int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnHttpCallResponse")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_http_call_response")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_http_call_response, err: %v", err)
		return err
	}

	_, err = ff.Call(contextId, token, headers, bodySize, trailers)

	return err
}

func (a *abiContext) ProxyOnQueueReady(rootContextId int32, token int32) error {
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[proxywasm_0_1_0][export] WasmerInstance ProxyOnQueueReady")
	}

	ff, err := a.instance.GetExportsFunc("proxy_on_queue_ready")
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm_0_1_0][export] WasmerInstance fail to get export func: proxy_on_queue_ready, err: %v", err)
		return err
	}

	_, err = ff.Call(rootContextId, token)

	return err
}
