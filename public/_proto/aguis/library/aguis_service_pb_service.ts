// package: library
// file: aguis/library/aguis_service.proto

import * as aguis_library_aguis_service_pb from "../../aguis/library/aguis_service_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
export class AutograderService {
  static serviceName = "library.AutograderService";
}
export namespace AutograderService {
  export class GetUser {
    static readonly methodName = "GetUser";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.RecordRequest;
    static readonly responseType = aguis_library_aguis_service_pb.User;
  }
  export class GetUsers {
    static readonly methodName = "GetUsers";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.Void;
    static readonly responseType = aguis_library_aguis_service_pb.Users;
  }
  export class UpdateUser {
    static readonly methodName = "UpdateUser";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.User;
    static readonly responseType = aguis_library_aguis_service_pb.User;
  }
  export class CreateCourse {
    static readonly methodName = "CreateCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.Course;
    static readonly responseType = aguis_library_aguis_service_pb.Course;
  }
  export class GetCourse {
    static readonly methodName = "GetCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.RecordRequest;
    static readonly responseType = aguis_library_aguis_service_pb.Course;
  }
  export class UpdateCourse {
    static readonly methodName = "UpdateCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.Course;
    static readonly responseType = aguis_library_aguis_service_pb.Course;
  }
  export class GetCourses {
    static readonly methodName = "GetCourses";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.Void;
    static readonly responseType = aguis_library_aguis_service_pb.Courses;
  }
  export class GetCoursesWithEnrollment {
    static readonly methodName = "GetCoursesWithEnrollment";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.RecordWithStatusRequest;
    static readonly responseType = aguis_library_aguis_service_pb.Courses;
  }
  export class GetAssignments {
    static readonly methodName = "GetAssignments";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.RecordRequest;
    static readonly responseType = aguis_library_aguis_service_pb.Assignments;
  }
  export class GetEnrollmentsByCourse {
    static readonly methodName = "GetEnrollmentsByCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.RecordWithStatusRequest;
    static readonly responseType = aguis_library_aguis_service_pb.EnrollmentResponse;
  }
  export class CreateEnrollment {
    static readonly methodName = "CreateEnrollment";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.EnrollmentRequest;
    static readonly responseType = aguis_library_aguis_service_pb.StatusCode;
  }
  export class UpdateEnrollment {
    static readonly methodName = "UpdateEnrollment";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = aguis_library_aguis_service_pb.EnrollmentRequest;
    static readonly responseType = aguis_library_aguis_service_pb.StatusCode;
  }
}
