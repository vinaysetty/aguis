// package: library
// file: aguis/library/aguis_service.proto

import * as aguis_library_aguis_service_pb from "../../aguis/library/aguis_service_pb";
export class AutograderService {
  static serviceName = "library.AutograderService";
}
export namespace AutograderService {
  export class GetUser {
    static readonly methodName = "GetUser";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.GetUserRequest;
    static readonly responseType = aguis_library_aguis_service_pb.User;
  }
  export class GetUsers {
    static readonly methodName = "GetUsers";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.Void;
    static readonly responseType = aguis_library_aguis_service_pb.UsersResponse;
  }
  export class UpdateUser {
    static readonly methodName = "UpdateUser";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.UpdateUserRequest;
    static readonly responseType = aguis_library_aguis_service_pb.User;
  }
  export class GetCourses {
    static readonly methodName = "GetCourses";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.Void;
    static readonly responseType = aguis_library_aguis_service_pb.Courses;
  }
}
