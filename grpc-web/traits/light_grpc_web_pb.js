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


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var types_info_pb = require('../types/info_pb.js')

var types_number_pb = require('../types/number_pb.js')

var types_tween_pb = require('../types/tween_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./light_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.LightApiClient =
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
proto.smartcore.traits.LightApiPromiseClient =
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
 *   !proto.smartcore.traits.UpdateBrightnessRequest,
 *   !proto.smartcore.traits.Brightness>}
 */
const methodDescriptor_LightApi_UpdateBrightness = new grpc.web.MethodDescriptor(
  '/smartcore.traits.LightApi/UpdateBrightness',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateBrightnessRequest,
  proto.smartcore.traits.Brightness,
  /**
   * @param {!proto.smartcore.traits.UpdateBrightnessRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Brightness.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateBrightnessRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Brightness)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Brightness>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LightApiClient.prototype.updateBrightness =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.LightApi/UpdateBrightness',
      request,
      metadata || {},
      methodDescriptor_LightApi_UpdateBrightness,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateBrightnessRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Brightness>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.LightApiPromiseClient.prototype.updateBrightness =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.LightApi/UpdateBrightness',
      request,
      metadata || {},
      methodDescriptor_LightApi_UpdateBrightness);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.GetBrightnessRequest,
 *   !proto.smartcore.traits.Brightness>}
 */
const methodDescriptor_LightApi_GetBrightness = new grpc.web.MethodDescriptor(
  '/smartcore.traits.LightApi/GetBrightness',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetBrightnessRequest,
  proto.smartcore.traits.Brightness,
  /**
   * @param {!proto.smartcore.traits.GetBrightnessRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Brightness.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetBrightnessRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Brightness)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Brightness>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LightApiClient.prototype.getBrightness =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.LightApi/GetBrightness',
      request,
      metadata || {},
      methodDescriptor_LightApi_GetBrightness,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetBrightnessRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Brightness>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.LightApiPromiseClient.prototype.getBrightness =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.LightApi/GetBrightness',
      request,
      metadata || {},
      methodDescriptor_LightApi_GetBrightness);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullBrightnessRequest,
 *   !proto.smartcore.traits.PullBrightnessResponse>}
 */
const methodDescriptor_LightApi_PullBrightness = new grpc.web.MethodDescriptor(
  '/smartcore.traits.LightApi/PullBrightness',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullBrightnessRequest,
  proto.smartcore.traits.PullBrightnessResponse,
  /**
   * @param {!proto.smartcore.traits.PullBrightnessRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullBrightnessResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullBrightnessRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullBrightnessResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LightApiClient.prototype.pullBrightness =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.LightApi/PullBrightness',
      request,
      metadata || {},
      methodDescriptor_LightApi_PullBrightness);
};


/**
 * @param {!proto.smartcore.traits.PullBrightnessRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullBrightnessResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LightApiPromiseClient.prototype.pullBrightness =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.LightApi/PullBrightness',
      request,
      metadata || {},
      methodDescriptor_LightApi_PullBrightness);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.LightInfoClient =
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
proto.smartcore.traits.LightInfoPromiseClient =
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
 *   !proto.smartcore.traits.DescribeBrightnessRequest,
 *   !proto.smartcore.traits.BrightnessSupport>}
 */
const methodDescriptor_LightInfo_DescribeBrightness = new grpc.web.MethodDescriptor(
  '/smartcore.traits.LightInfo/DescribeBrightness',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DescribeBrightnessRequest,
  proto.smartcore.traits.BrightnessSupport,
  /**
   * @param {!proto.smartcore.traits.DescribeBrightnessRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.BrightnessSupport.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DescribeBrightnessRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.BrightnessSupport)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.BrightnessSupport>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LightInfoClient.prototype.describeBrightness =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.LightInfo/DescribeBrightness',
      request,
      metadata || {},
      methodDescriptor_LightInfo_DescribeBrightness,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DescribeBrightnessRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.BrightnessSupport>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.LightInfoPromiseClient.prototype.describeBrightness =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.LightInfo/DescribeBrightness',
      request,
      metadata || {},
      methodDescriptor_LightInfo_DescribeBrightness);
};


module.exports = proto.smartcore.traits;

