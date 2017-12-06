import {Message} from "google-protobuf";
import {grpc} from "grpc-web-client";
import {
    Assignments,
    Course,
    Courses,
    EnrollmentRequest,
    EnrollmentResponse,
    RecordRequest,
    RecordWithStatusRequest,
    StatusCode,
    User,
    Users,
    Void,
} from "../../_proto/ag_service_pb";
import { AutograderService } from "../../_proto/ag_service_pb_service";
import { ICourse, INewCourse, IUser } from "../models";

declare const USE_TLS: boolean;
const host = USE_TLS ? "https://localhost:8091" : "http://localhost:8090";

export interface IGrpcResult<T> {
    statusCode: number;
    data?: T;
}

export class GrpcHelper {

    public getUsers(): Promise<IGrpcResult<Users>> {
        const usersRequest = new Void();
        return this.grpcUnary<Users>(AutograderService.GetUsers, usersRequest);
    }

    public getUser(id: number): Promise<IGrpcResult<User>> {
        const userRequest = new RecordRequest();
        userRequest.setId(id);
        return this.grpcUnary<User>(AutograderService.GetUser, userRequest);
    }

    public updateUser(user: IUser, isadmin?: boolean): Promise<IGrpcResult<User>> {
        const u = new User();
        u.setId(user.id);
        u.setAvatarurl(user.avatarurl);
        u.setEmail(user.email);
        u.setName(user.name);
        u.setStudentid(user.studentid);
        if (isadmin) {
            u.setIsadmin(isadmin);
        } else {
            u.setIsadmin(user.isadmin);
        }
        // const userRequest = new UpdateUserRequest();
        // userRequest.setUser(u);
        return this.grpcUnary<User>(AutograderService.UpdateUser, u);
    }

    public createCourse(course: INewCourse): Promise<IGrpcResult<Course>> {
        const nc = new Course();
        nc.setName(course.name);
        nc.setCode(course.code);
        nc.setProvider(course.provider);
        nc.setDirectoryid(course.directoryid);
        nc.setTag(course.tag);
        nc.setYear(course.year);
        return this.grpcUnary<Course>(AutograderService.CreateCourse, nc);
    }

    public updateCourse(course: ICourse): Promise<IGrpcResult<Course>> {
        const crs = new Course();
        crs.setId(course.id);
        crs.setName(course.name);
        crs.setCode(course.code);
        crs.setDirectoryid(course.directoryid);
        crs.setProvider(course.provider);
        crs.setYear(course.year);
        crs.setTag(course.tag);
        return this.grpcUnary<Course>(AutograderService.UpdateCourse, crs);
    }

    public getCourse(id: number): Promise<IGrpcResult<Course>> {
        const query = new RecordRequest();
        query.setId(id);
        return this.grpcUnary(AutograderService.GetCourse, query);
    }

    public getCourses(): Promise<IGrpcResult<Courses>> {
        const usersRequest = new Void();
        return this.grpcUnary<Courses>(AutograderService.GetCourses, usersRequest);
    }

    public getCoursesWithEnrollment(userid: number, state: any): Promise<IGrpcResult<Courses>> {
        const courseReq = new RecordWithStatusRequest();
        courseReq.setId(userid);
        courseReq.setStatusesList(state);
        return this.grpcUnary<Courses>(AutograderService.GetCoursesWithEnrollment, courseReq);
    }

    public getAssignments(courseId: number): Promise<IGrpcResult<Assignments>> {
        const req = new RecordRequest();
        req.setId(courseId);
        return this.grpcUnary<Assignments>(AutograderService.GetAssignments, req);
    }

    public getEnrollmentsByCourse(courseid: number, state: any): Promise<IGrpcResult<EnrollmentResponse>> {
        const req = new RecordWithStatusRequest();
        req.setId(courseid);
        req.setStatusesList(state);
        return this.grpcUnary<EnrollmentResponse>(AutograderService.GetEnrollmentsByCourse, req);
    }

    public createEnrollment(userid: number, courseid: number): Promise<IGrpcResult<StatusCode>> {
        const req = new EnrollmentRequest();
        req.setUserid(userid);
        req.setCourseid(courseid);
        return this.grpcUnary<StatusCode>(AutograderService.CreateEnrollment, req);
    }

    public updateEnrollment(userid: number,
                            courseid: number,
                            state: any): Promise<IGrpcResult<StatusCode>> {
        const req = new EnrollmentRequest();
        req.setUserid(userid);
        req.setCourseid(courseid);
        req.setEnrolled(state);
        return this.grpcUnary<StatusCode>(AutograderService.UpdateEnrollment, req);
    }

    private grpcUnary<TReceive extends Message>(method: any, request: any): Promise<IGrpcResult<TReceive>> {
        const requestPromise = new Promise<IGrpcResult<TReceive>>((resolve, reject) => {
            grpc.unary(method, {
                request,
                host,
                onEnd: (res) => {
                    const {status, statusMessage, headers, message, trailers} = res;
                    const temp: IGrpcResult<TReceive> = {
                        data: message as TReceive,
                        statusCode: status,
                    };
                    resolve(temp);

                },
            });
        });
        return requestPromise;
    }

}
