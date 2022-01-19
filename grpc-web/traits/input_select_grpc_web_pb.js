/**
 * @fileoverview gRPC-Web generated client stub for smartcore.traits
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var types_info_pb = require('../types/info_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./input_select_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.InputSelectApiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.InputSelectApiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateInputRequest,
 *   !proto.smartcore.traits.Input>}
 */
const methodDescriptor_InputSelectApi_UpdateInput = new grpc.web.MethodDescriptor(
  '/smartcore.traits.InputSelectApi/UpdateInput',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateInputRequest,
  proto.smartcore.traits.Input,
  /**
   * @param {!proto.smartcore.traits.UpdateInputRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Input.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateInputRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Input)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Input>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.InputSelectApiClient.prototype.updateInput =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.InputSelectApi/UpdateInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectApi_UpdateInput,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateInputRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Input>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.InputSelectApiPromiseClient.prototype.updateInput =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.InputSelectApi/UpdateInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectApi_UpdateInput);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.GetInputRequest,
 *   !proto.smartcore.traits.Input>}
 */
const methodDescriptor_InputSelectApi_GetInput = new grpc.web.MethodDescriptor(
  '/smartcore.traits.InputSelectApi/GetInput',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetInputRequest,
  proto.smartcore.traits.Input,
  /**
   * @param {!proto.smartcore.traits.GetInputRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Input.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetInputRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Input)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Input>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.InputSelectApiClient.prototype.getInput =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.InputSelectApi/GetInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectApi_GetInput,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetInputRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Input>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.InputSelectApiPromiseClient.prototype.getInput =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.InputSelectApi/GetInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectApi_GetInput);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullInputRequest,
 *   !proto.smartcore.traits.PullInputResponse>}
 */
const methodDescriptor_InputSelectApi_PullInput = new grpc.web.MethodDescriptor(
  '/smartcore.traits.InputSelectApi/PullInput',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullInputRequest,
  proto.smartcore.traits.PullInputResponse,
  /**
   * @param {!proto.smartcore.traits.PullInputRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullInputResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullInputRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullInputResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.InputSelectApiClient.prototype.pullInput =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.InputSelectApi/PullInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectApi_PullInput);
};


/**
 * @param {!proto.smartcore.traits.PullInputRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullInputResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.InputSelectApiPromiseClient.prototype.pullInput =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.InputSelectApi/PullInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectApi_PullInput);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.InputSelectInfoClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.InputSelectInfoPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.DescribeInputRequest,
 *   !proto.smartcore.traits.InputSupport>}
 */
const methodDescriptor_InputSelectInfo_DescribeInput = new grpc.web.MethodDescriptor(
  '/smartcore.traits.InputSelectInfo/DescribeInput',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DescribeInputRequest,
  proto.smartcore.traits.InputSupport,
  /**
   * @param {!proto.smartcore.traits.DescribeInputRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.InputSupport.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DescribeInputRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.InputSupport)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.InputSupport>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.InputSelectInfoClient.prototype.describeInput =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.InputSelectInfo/DescribeInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectInfo_DescribeInput,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DescribeInputRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.InputSupport>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.InputSelectInfoPromiseClient.prototype.describeInput =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.InputSelectInfo/DescribeInput',
      request,
      metadata || {},
      methodDescriptor_InputSelectInfo_DescribeInput);
};


module.exports = proto.smartcore.traits;

