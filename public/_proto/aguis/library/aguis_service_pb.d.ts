// package: library
// file: aguis/library/aguis_service.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Enrollment extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getCourseid(): number;
  setCourseid(value: number): void;

  getUserid(): number;
  setUserid(value: number): void;

  getGroupid(): number;
  setGroupid(value: number): void;

  hasUser(): boolean;
  clearUser(): void;
  getUser(): User | undefined;
  setUser(value?: User): void;

  hasCourse(): boolean;
  clearCourse(): void;
  getCourse(): Course | undefined;
  setCourse(value?: Course): void;

  hasGroup(): boolean;
  clearGroup(): void;
  getGroup(): Group | undefined;
  setGroup(value?: Group): void;

  getStatus(): Enrollment.Status;
  setStatus(value: Enrollment.Status): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Enrollment.AsObject;
  static toObject(includeInstance: boolean, msg: Enrollment): Enrollment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Enrollment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Enrollment;
  static deserializeBinaryFromReader(message: Enrollment, reader: jspb.BinaryReader): Enrollment;
}

export namespace Enrollment {
  export type AsObject = {
    id: number,
    courseid: number,
    userid: number,
    groupid: number,
    user?: User.AsObject,
    course?: Course.AsObject,
    group?: Group.AsObject,
    status: Enrollment.Status,
  }

  export enum Status {
    None = 0,
    Pending = 1,
    Rejected = 2,
    Student = 3,
    Teacher = 4,
  }
}

export class EnrollmentResponse extends jspb.Message {
  clearEnrollmentsList(): void;
  getEnrollmentsList(): Array<Enrollment>;
  setEnrollmentsList(value: Array<Enrollment>): void;
  addEnrollments(value?: Enrollment, index?: number): Enrollment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EnrollmentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EnrollmentResponse): EnrollmentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EnrollmentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EnrollmentResponse;
  static deserializeBinaryFromReader(message: EnrollmentResponse, reader: jspb.BinaryReader): EnrollmentResponse;
}

export namespace EnrollmentResponse {
  export type AsObject = {
    enrollmentsList: Array<Enrollment.AsObject>,
  }
}

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getIsadmin(): boolean;
  setIsadmin(value: boolean): void;

  getName(): string;
  setName(value: string): void;

  getStudentid(): string;
  setStudentid(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getAvatarurl(): string;
  setAvatarurl(value: string): void;

  clearRemoteidentitiesList(): void;
  getRemoteidentitiesList(): Array<RemoteIdentity>;
  setRemoteidentitiesList(value: Array<RemoteIdentity>): void;
  addRemoteidentities(value?: RemoteIdentity, index?: number): RemoteIdentity;

  clearEnrollmentsList(): void;
  getEnrollmentsList(): Array<Enrollment>;
  setEnrollmentsList(value: Array<Enrollment>): void;
  addEnrollments(value?: Enrollment, index?: number): Enrollment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    isadmin: boolean,
    name: string,
    studentid: string,
    email: string,
    avatarurl: string,
    remoteidentitiesList: Array<RemoteIdentity.AsObject>,
    enrollmentsList: Array<Enrollment.AsObject>,
  }
}

export class RemoteIdentity extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getProvider(): string;
  setProvider(value: string): void;

  getRemoteid(): number;
  setRemoteid(value: number): void;

  getAccesstoken(): string;
  setAccesstoken(value: string): void;

  getUserid(): number;
  setUserid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoteIdentity.AsObject;
  static toObject(includeInstance: boolean, msg: RemoteIdentity): RemoteIdentity.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RemoteIdentity, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoteIdentity;
  static deserializeBinaryFromReader(message: RemoteIdentity, reader: jspb.BinaryReader): RemoteIdentity;
}

export namespace RemoteIdentity {
  export type AsObject = {
    id: number,
    provider: string,
    remoteid: number,
    accesstoken: string,
    userid: number,
  }
}

export class Users extends jspb.Message {
  clearUsersList(): void;
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): void;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Users.AsObject;
  static toObject(includeInstance: boolean, msg: Users): Users.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Users, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Users;
  static deserializeBinaryFromReader(message: Users, reader: jspb.BinaryReader): Users;
}

export namespace Users {
  export type AsObject = {
    usersList: Array<User.AsObject>,
  }
}

export class Void extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Void.AsObject;
  static toObject(includeInstance: boolean, msg: Void): Void.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Void, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Void;
  static deserializeBinaryFromReader(message: Void, reader: jspb.BinaryReader): Void;
}

export namespace Void {
  export type AsObject = {
  }
}

export class RecordRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RecordRequest): RecordRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RecordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecordRequest;
  static deserializeBinaryFromReader(message: RecordRequest, reader: jspb.BinaryReader): RecordRequest;
}

export namespace RecordRequest {
  export type AsObject = {
    id: number,
  }
}

export class Assignment extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getCourseid(): number;
  setCourseid(value: number): void;

  getName(): string;
  setName(value: string): void;

  getLanguage(): string;
  setLanguage(value: string): void;

  hasDeadline(): boolean;
  clearDeadline(): void;
  getDeadline(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeadline(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getAutoapprove(): boolean;
  setAutoapprove(value: boolean): void;

  getOrder(): number;
  setOrder(value: number): void;

  hasSubmission(): boolean;
  clearSubmission(): void;
  getSubmission(): Submission | undefined;
  setSubmission(value?: Submission): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Assignment.AsObject;
  static toObject(includeInstance: boolean, msg: Assignment): Assignment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Assignment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Assignment;
  static deserializeBinaryFromReader(message: Assignment, reader: jspb.BinaryReader): Assignment;
}

export namespace Assignment {
  export type AsObject = {
    id: number,
    courseid: number,
    name: string,
    language: string,
    deadline?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    autoapprove: boolean,
    order: number,
    submission?: Submission.AsObject,
  }
}

export class Submission extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getAssignmentid(): number;
  setAssignmentid(value: number): void;

  getUserid(): number;
  setUserid(value: number): void;

  getGroupid(): number;
  setGroupid(value: number): void;

  getScore(): number;
  setScore(value: number): void;

  getScoreobjects(): string;
  setScoreobjects(value: string): void;

  getBuildinfo(): string;
  setBuildinfo(value: string): void;

  getCommithash(): string;
  setCommithash(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Submission.AsObject;
  static toObject(includeInstance: boolean, msg: Submission): Submission.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Submission, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Submission;
  static deserializeBinaryFromReader(message: Submission, reader: jspb.BinaryReader): Submission;
}

export namespace Submission {
  export type AsObject = {
    id: number,
    assignmentid: number,
    userid: number,
    groupid: number,
    score: number,
    scoreobjects: string,
    buildinfo: string,
    commithash: string,
  }
}

export class Assignments extends jspb.Message {
  clearAssignmentsList(): void;
  getAssignmentsList(): Array<Assignment>;
  setAssignmentsList(value: Array<Assignment>): void;
  addAssignments(value?: Assignment, index?: number): Assignment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Assignments.AsObject;
  static toObject(includeInstance: boolean, msg: Assignments): Assignments.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Assignments, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Assignments;
  static deserializeBinaryFromReader(message: Assignments, reader: jspb.BinaryReader): Assignments;
}

export namespace Assignments {
  export type AsObject = {
    assignmentsList: Array<Assignment.AsObject>,
  }
}

export class Group extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  getStatus(): number;
  setStatus(value: number): void;

  getCourseid(): number;
  setCourseid(value: number): void;

  clearUsersList(): void;
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): void;
  addUsers(value?: User, index?: number): User;

  clearEnrollmentsList(): void;
  getEnrollmentsList(): Array<Enrollment>;
  setEnrollmentsList(value: Array<Enrollment>): void;
  addEnrollments(value?: Enrollment, index?: number): Enrollment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Group.AsObject;
  static toObject(includeInstance: boolean, msg: Group): Group.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Group, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Group;
  static deserializeBinaryFromReader(message: Group, reader: jspb.BinaryReader): Group;
}

export namespace Group {
  export type AsObject = {
    id: number,
    name: string,
    status: number,
    courseid: number,
    usersList: Array<User.AsObject>,
    enrollmentsList: Array<Enrollment.AsObject>,
  }
}

export class Course extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  getYear(): number;
  setYear(value: number): void;

  getTag(): string;
  setTag(value: string): void;

  getProvider(): string;
  setProvider(value: string): void;

  getDirectoryid(): number;
  setDirectoryid(value: number): void;

  getEnrolled(): Enrollment.Status;
  setEnrolled(value: Enrollment.Status): void;

  clearEnrollmentsList(): void;
  getEnrollmentsList(): Array<Enrollment>;
  setEnrollmentsList(value: Array<Enrollment>): void;
  addEnrollments(value?: Enrollment, index?: number): Enrollment;

  clearAssignmentsList(): void;
  getAssignmentsList(): Array<Assignment>;
  setAssignmentsList(value: Array<Assignment>): void;
  addAssignments(value?: Assignment, index?: number): Assignment;

  clearGroupsList(): void;
  getGroupsList(): Array<Group>;
  setGroupsList(value: Array<Group>): void;
  addGroups(value?: Group, index?: number): Group;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Course.AsObject;
  static toObject(includeInstance: boolean, msg: Course): Course.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Course, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Course;
  static deserializeBinaryFromReader(message: Course, reader: jspb.BinaryReader): Course;
}

export namespace Course {
  export type AsObject = {
    id: number,
    name: string,
    code: string,
    year: number,
    tag: string,
    provider: string,
    directoryid: number,
    enrolled: Enrollment.Status,
    enrollmentsList: Array<Enrollment.AsObject>,
    assignmentsList: Array<Assignment.AsObject>,
    groupsList: Array<Group.AsObject>,
  }
}

export class Courses extends jspb.Message {
  clearCoursesList(): void;
  getCoursesList(): Array<Course>;
  setCoursesList(value: Array<Course>): void;
  addCourses(value?: Course, index?: number): Course;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Courses.AsObject;
  static toObject(includeInstance: boolean, msg: Courses): Courses.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Courses, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Courses;
  static deserializeBinaryFromReader(message: Courses, reader: jspb.BinaryReader): Courses;
}

export namespace Courses {
  export type AsObject = {
    coursesList: Array<Course.AsObject>,
  }
}

export class RecordWithStatusRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  clearStatusesList(): void;
  getStatusesList(): Array<Enrollment.Status>;
  setStatusesList(value: Array<Enrollment.Status>): void;
  addStatuses(value: Enrollment.Status, index?: number): Enrollment.Status;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecordWithStatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RecordWithStatusRequest): RecordWithStatusRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RecordWithStatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecordWithStatusRequest;
  static deserializeBinaryFromReader(message: RecordWithStatusRequest, reader: jspb.BinaryReader): RecordWithStatusRequest;
}

export namespace RecordWithStatusRequest {
  export type AsObject = {
    id: number,
    statusesList: Array<Enrollment.Status>,
  }
}

export class UserIDCourseID extends jspb.Message {
  getUserid(): number;
  setUserid(value: number): void;

  getCourseid(): number;
  setCourseid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserIDCourseID.AsObject;
  static toObject(includeInstance: boolean, msg: UserIDCourseID): UserIDCourseID.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserIDCourseID, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserIDCourseID;
  static deserializeBinaryFromReader(message: UserIDCourseID, reader: jspb.BinaryReader): UserIDCourseID;
}

export namespace UserIDCourseID {
  export type AsObject = {
    userid: number,
    courseid: number,
  }
}

export class StatusCode extends jspb.Message {
  getStatuscode(): number;
  setStatuscode(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatusCode.AsObject;
  static toObject(includeInstance: boolean, msg: StatusCode): StatusCode.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StatusCode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatusCode;
  static deserializeBinaryFromReader(message: StatusCode, reader: jspb.BinaryReader): StatusCode;
}

export namespace StatusCode {
  export type AsObject = {
    statuscode: number,
  }
}

export class EnrollmentRequest extends jspb.Message {
  getUserid(): number;
  setUserid(value: number): void;

  getCourseid(): number;
  setCourseid(value: number): void;

  getEnrolled(): Enrollment.Status;
  setEnrolled(value: Enrollment.Status): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EnrollmentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EnrollmentRequest): EnrollmentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EnrollmentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EnrollmentRequest;
  static deserializeBinaryFromReader(message: EnrollmentRequest, reader: jspb.BinaryReader): EnrollmentRequest;
}

export namespace EnrollmentRequest {
  export type AsObject = {
    userid: number,
    courseid: number,
    enrolled: Enrollment.Status,
  }
}

