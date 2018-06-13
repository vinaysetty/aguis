// package: ag
// file: ag_service.proto

import * as ag_service_pb from "./ag_service_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
//import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "./github.com/gogo/protobuf/gogoproto/gogo_pb";
export class AutograderService {
  static serviceName = "ag.AutograderService";
}
export namespace AutograderService {
  export class GetUser {
    static readonly methodName = "GetUser";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.RecordRequest;
    static readonly responseType = ag_service_pb.User;
  }
  export class GetUsers {
    static readonly methodName = "GetUsers";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.Void;
    static readonly responseType = ag_service_pb.Users;
  }
  export class UpdateUser {
    static readonly methodName = "UpdateUser";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.User;
    static readonly responseType = ag_service_pb.User;
  }
  export class CreateCourse {
    static readonly methodName = "CreateCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.Course;
    static readonly responseType = ag_service_pb.Course;
  }
  export class GetCourse {
    static readonly methodName = "GetCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.RecordRequest;
    static readonly responseType = ag_service_pb.Course;
  }
  export class UpdateCourse {
    static readonly methodName = "UpdateCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.Course;
    static readonly responseType = ag_service_pb.Course;
  }
  export class GetCourses {
    static readonly methodName = "GetCourses";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.Void;
    static readonly responseType = ag_service_pb.Courses;
  }
  export class GetCoursesWithEnrollment {
    static readonly methodName = "GetCoursesWithEnrollment";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.RecordWithStatusRequest;
    static readonly responseType = ag_service_pb.Courses;
  }
  export class GetAssignments {
    static readonly methodName = "GetAssignments";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.RecordRequest;
    static readonly responseType = ag_service_pb.Assignments;
  }
  export class GetEnrollmentsByCourse {
    static readonly methodName = "GetEnrollmentsByCourse";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.RecordWithStatusRequest;
    static readonly responseType = ag_service_pb.EnrollmentResponse;
  }
  export class CreateEnrollment {
    static readonly methodName = "CreateEnrollment";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.EnrollmentRequest;
    static readonly responseType = ag_service_pb.StatusCode;
  }
  export class UpdateEnrollment {
    static readonly methodName = "UpdateEnrollment";
    static readonly service = AutograderService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = ag_service_pb.EnrollmentRequest;
    static readonly responseType = ag_service_pb.StatusCode;
  }
}
