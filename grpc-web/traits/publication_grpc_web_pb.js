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

var types_change_pb = require('../types/change_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./publication_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.PublicationApiClient =
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
proto.smartcore.traits.PublicationApiPromiseClient =
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
 *   !proto.smartcore.traits.CreatePublicationRequest,
 *   !proto.smartcore.traits.Publication>}
 */
const methodDescriptor_PublicationApi_CreatePublication = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/CreatePublication',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.CreatePublicationRequest,
  proto.smartcore.traits.Publication,
  /**
   * @param {!proto.smartcore.traits.CreatePublicationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Publication.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.CreatePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Publication)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Publication>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.createPublication =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/CreatePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_CreatePublication,
      callback);
};


/**
 * @param {!proto.smartcore.traits.CreatePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Publication>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.createPublication =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/CreatePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_CreatePublication);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.GetPublicationRequest,
 *   !proto.smartcore.traits.Publication>}
 */
const methodDescriptor_PublicationApi_GetPublication = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/GetPublication',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetPublicationRequest,
  proto.smartcore.traits.Publication,
  /**
   * @param {!proto.smartcore.traits.GetPublicationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Publication.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetPublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Publication)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Publication>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.getPublication =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/GetPublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_GetPublication,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetPublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Publication>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.getPublication =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/GetPublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_GetPublication);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdatePublicationRequest,
 *   !proto.smartcore.traits.Publication>}
 */
const methodDescriptor_PublicationApi_UpdatePublication = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/UpdatePublication',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdatePublicationRequest,
  proto.smartcore.traits.Publication,
  /**
   * @param {!proto.smartcore.traits.UpdatePublicationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Publication.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdatePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Publication)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Publication>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.updatePublication =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/UpdatePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_UpdatePublication,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdatePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Publication>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.updatePublication =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/UpdatePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_UpdatePublication);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.DeletePublicationRequest,
 *   !proto.smartcore.traits.Publication>}
 */
const methodDescriptor_PublicationApi_DeletePublication = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/DeletePublication',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DeletePublicationRequest,
  proto.smartcore.traits.Publication,
  /**
   * @param {!proto.smartcore.traits.DeletePublicationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Publication.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DeletePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Publication)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Publication>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.deletePublication =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/DeletePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_DeletePublication,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DeletePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Publication>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.deletePublication =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/DeletePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_DeletePublication);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullPublicationRequest,
 *   !proto.smartcore.traits.PullPublicationResponse>}
 */
const methodDescriptor_PublicationApi_PullPublication = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/PullPublication',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullPublicationRequest,
  proto.smartcore.traits.PullPublicationResponse,
  /**
   * @param {!proto.smartcore.traits.PullPublicationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullPublicationResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullPublicationRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullPublicationResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.pullPublication =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.PublicationApi/PullPublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_PullPublication);
};


/**
 * @param {!proto.smartcore.traits.PullPublicationRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullPublicationResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.pullPublication =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.PublicationApi/PullPublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_PullPublication);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.ListPublicationsRequest,
 *   !proto.smartcore.traits.ListPublicationsResponse>}
 */
const methodDescriptor_PublicationApi_ListPublications = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/ListPublications',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ListPublicationsRequest,
  proto.smartcore.traits.ListPublicationsResponse,
  /**
   * @param {!proto.smartcore.traits.ListPublicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ListPublicationsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ListPublicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ListPublicationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ListPublicationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.listPublications =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/ListPublications',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_ListPublications,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ListPublicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ListPublicationsResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.listPublications =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/ListPublications',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_ListPublications);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullPublicationsRequest,
 *   !proto.smartcore.traits.PullPublicationsResponse>}
 */
const methodDescriptor_PublicationApi_PullPublications = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/PullPublications',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullPublicationsRequest,
  proto.smartcore.traits.PullPublicationsResponse,
  /**
   * @param {!proto.smartcore.traits.PullPublicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullPublicationsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullPublicationsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullPublicationsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.pullPublications =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.PublicationApi/PullPublications',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_PullPublications);
};


/**
 * @param {!proto.smartcore.traits.PullPublicationsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullPublicationsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.pullPublications =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.PublicationApi/PullPublications',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_PullPublications);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.AcknowledgePublicationRequest,
 *   !proto.smartcore.traits.Publication>}
 */
const methodDescriptor_PublicationApi_AcknowledgePublication = new grpc.web.MethodDescriptor(
  '/smartcore.traits.PublicationApi/AcknowledgePublication',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.AcknowledgePublicationRequest,
  proto.smartcore.traits.Publication,
  /**
   * @param {!proto.smartcore.traits.AcknowledgePublicationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Publication.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.AcknowledgePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Publication)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Publication>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.PublicationApiClient.prototype.acknowledgePublication =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/AcknowledgePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_AcknowledgePublication,
      callback);
};


/**
 * @param {!proto.smartcore.traits.AcknowledgePublicationRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Publication>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.PublicationApiPromiseClient.prototype.acknowledgePublication =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.PublicationApi/AcknowledgePublication',
      request,
      metadata || {},
      methodDescriptor_PublicationApi_AcknowledgePublication);
};


module.exports = proto.smartcore.traits;

