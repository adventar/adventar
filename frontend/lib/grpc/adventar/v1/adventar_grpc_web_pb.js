/**
 * @fileoverview gRPC-Web generated client stub for adventar.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var adventar_v1_resources_pb = require('../../adventar/v1/resources_pb.js')

var adventar_v1_rpc_messages_pb = require('../../adventar/v1/rpc_messages_pb.js')
const proto = {};
proto.adventar = {};
proto.adventar.v1 = require('./adventar_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.adventar.v1.AdventarClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.adventar.v1.AdventarPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.adventar.v1.AdventarClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.adventar.v1.AdventarClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.ListCalendarsRequest,
 *   !proto.adventar.v1.ListCalendarsResponse>}
 */
const methodInfo_Adventar_ListCalendars = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_rpc_messages_pb.ListCalendarsResponse,
  /** @param {!proto.adventar.v1.ListCalendarsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_rpc_messages_pb.ListCalendarsResponse.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.ListCalendarsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.ListCalendarsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.ListCalendarsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.listCalendars =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/ListCalendars',
      request,
      metadata || {},
      methodInfo_Adventar_ListCalendars,
      callback);
};


/**
 * @param {!proto.adventar.v1.ListCalendarsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.ListCalendarsResponse>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.listCalendars =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.listCalendars(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.GetCalendarRequest,
 *   !proto.adventar.v1.GetCalendarResponse>}
 */
const methodInfo_Adventar_GetCalendar = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_rpc_messages_pb.GetCalendarResponse,
  /** @param {!proto.adventar.v1.GetCalendarRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_rpc_messages_pb.GetCalendarResponse.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.GetCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.GetCalendarResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.GetCalendarResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.getCalendar =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/GetCalendar',
      request,
      metadata || {},
      methodInfo_Adventar_GetCalendar,
      callback);
};


/**
 * @param {!proto.adventar.v1.GetCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.GetCalendarResponse>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.getCalendar =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.getCalendar(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.CreateCalendarRequest,
 *   !proto.adventar.v1.Calendar>}
 */
const methodInfo_Adventar_CreateCalendar = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_resources_pb.Calendar,
  /** @param {!proto.adventar.v1.CreateCalendarRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_resources_pb.Calendar.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.CreateCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.Calendar)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.Calendar>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.createCalendar =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/CreateCalendar',
      request,
      metadata || {},
      methodInfo_Adventar_CreateCalendar,
      callback);
};


/**
 * @param {!proto.adventar.v1.CreateCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.Calendar>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.createCalendar =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.createCalendar(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.UpdateCalendarRequest,
 *   !proto.adventar.v1.Calendar>}
 */
const methodInfo_Adventar_UpdateCalendar = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_resources_pb.Calendar,
  /** @param {!proto.adventar.v1.UpdateCalendarRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_resources_pb.Calendar.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.UpdateCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.Calendar)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.Calendar>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.updateCalendar =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/UpdateCalendar',
      request,
      metadata || {},
      methodInfo_Adventar_UpdateCalendar,
      callback);
};


/**
 * @param {!proto.adventar.v1.UpdateCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.Calendar>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.updateCalendar =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.updateCalendar(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.DeleteCalendarRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Adventar_DeleteCalendar = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.adventar.v1.DeleteCalendarRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.DeleteCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.deleteCalendar =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/DeleteCalendar',
      request,
      metadata || {},
      methodInfo_Adventar_DeleteCalendar,
      callback);
};


/**
 * @param {!proto.adventar.v1.DeleteCalendarRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.deleteCalendar =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.deleteCalendar(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.ListEntriesRequest,
 *   !proto.adventar.v1.ListEntriesResponse>}
 */
const methodInfo_Adventar_ListEntries = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_rpc_messages_pb.ListEntriesResponse,
  /** @param {!proto.adventar.v1.ListEntriesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_rpc_messages_pb.ListEntriesResponse.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.ListEntriesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.ListEntriesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.ListEntriesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.listEntries =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/ListEntries',
      request,
      metadata || {},
      methodInfo_Adventar_ListEntries,
      callback);
};


/**
 * @param {!proto.adventar.v1.ListEntriesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.ListEntriesResponse>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.listEntries =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.listEntries(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.CreateEntryRequest,
 *   !proto.adventar.v1.Entry>}
 */
const methodInfo_Adventar_CreateEntry = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_resources_pb.Entry,
  /** @param {!proto.adventar.v1.CreateEntryRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_resources_pb.Entry.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.CreateEntryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.Entry)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.Entry>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.createEntry =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/CreateEntry',
      request,
      metadata || {},
      methodInfo_Adventar_CreateEntry,
      callback);
};


/**
 * @param {!proto.adventar.v1.CreateEntryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.Entry>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.createEntry =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.createEntry(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.UpdateEntryRequest,
 *   !proto.adventar.v1.Entry>}
 */
const methodInfo_Adventar_UpdateEntry = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_resources_pb.Entry,
  /** @param {!proto.adventar.v1.UpdateEntryRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_resources_pb.Entry.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.UpdateEntryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.Entry)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.Entry>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.updateEntry =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/UpdateEntry',
      request,
      metadata || {},
      methodInfo_Adventar_UpdateEntry,
      callback);
};


/**
 * @param {!proto.adventar.v1.UpdateEntryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.Entry>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.updateEntry =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.updateEntry(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.DeleteEntryRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Adventar_DeleteEntry = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /** @param {!proto.adventar.v1.DeleteEntryRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.DeleteEntryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.deleteEntry =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/DeleteEntry',
      request,
      metadata || {},
      methodInfo_Adventar_DeleteEntry,
      callback);
};


/**
 * @param {!proto.adventar.v1.DeleteEntryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.deleteEntry =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.deleteEntry(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.SignInRequest,
 *   !proto.adventar.v1.User>}
 */
const methodInfo_Adventar_SignIn = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_resources_pb.User,
  /** @param {!proto.adventar.v1.SignInRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_resources_pb.User.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.SignInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.signIn =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/SignIn',
      request,
      metadata || {},
      methodInfo_Adventar_SignIn,
      callback);
};


/**
 * @param {!proto.adventar.v1.SignInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.User>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.signIn =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.signIn(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.GetUserRequest,
 *   !proto.adventar.v1.User>}
 */
const methodInfo_Adventar_GetUser = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_resources_pb.User,
  /** @param {!proto.adventar.v1.GetUserRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_resources_pb.User.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/GetUser',
      request,
      metadata || {},
      methodInfo_Adventar_GetUser,
      callback);
};


/**
 * @param {!proto.adventar.v1.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.User>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.getUser =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.getUser(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.adventar.v1.UpdateUserRequest,
 *   !proto.adventar.v1.User>}
 */
const methodInfo_Adventar_UpdateUser = new grpc.web.AbstractClientBase.MethodInfo(
  adventar_v1_resources_pb.User,
  /** @param {!proto.adventar.v1.UpdateUserRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  adventar_v1_resources_pb.User.deserializeBinary
);


/**
 * @param {!proto.adventar.v1.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.adventar.v1.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.adventar.v1.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarClient.prototype.updateUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/adventar.v1.Adventar/UpdateUser',
      request,
      metadata || {},
      methodInfo_Adventar_UpdateUser,
      callback);
};


/**
 * @param {!proto.adventar.v1.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.adventar.v1.User>}
 *     The XHR Node Readable Stream
 */
proto.adventar.v1.AdventarPromiseClient.prototype.updateUser =
    function(request, metadata) {
  var _this = this;
  return new Promise(function (resolve, reject) {
    _this.delegateClient_.updateUser(
      request, metadata, function (error, response) {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.adventar.v1;

