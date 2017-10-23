import {grpc} from "grpc-web-client";
import {
    Courses,
    GetUserRequest,
    UpdateUserRequest,
    User,
    UsersResponse,
    Void
} from "../../_proto/aguis/library/aguis_service_pb"
import {AutograderService} from "../../_proto/aguis/library/aguis_service_pb_service";
import {IUser} from "../models";
import {Message} from "google-protobuf"

declare const USE_TLS: boolean;
const host = USE_TLS ? "https://localhost:8091" : "http://localhost:8090";

export interface IGrpcResult<T> {
    statusCode: number;
    data?: T;
}

export class GrpcHelper {

    public getUsers(): Promise<IGrpcResult<UsersResponse>> {
        const usersRequest = new Void();
        return this.grpcUnary<UsersResponse>(AutograderService.GetUsers, usersRequest);
    }

    public getUser(id: number): Promise<IGrpcResult<User>> {
        const userRequest = new GetUserRequest();
        userRequest.setId(id);
        return this.grpcUnary<User>(AutograderService.GetUser, userRequest);
    }

    public updateUser(user: IUser, isadmin?: boolean): Promise<IGrpcResult<User>> {
        console.log("update User obj = ", user);
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
        const userRequest = new UpdateUserRequest;
        userRequest.setUser(u);
        return this.grpcUnary<User>(AutograderService.UpdateUser, userRequest);
    }

    public getCourses(): Promise<IGrpcResult<Courses>> {
        const usersRequest = new Void();
        return this.grpcUnary<Courses>(AutograderService.GetCourses, usersRequest);
    }


    private grpcUnary<TReceive extends Message>(method: any, request: any): Promise<IGrpcResult<TReceive>> {
        const requestPromise = new Promise<IGrpcResult<TReceive>>((resolve, reject) => {
            grpc.unary(method, {
                request: request,
                host: host,
                onEnd: res => {
                    const {status, statusMessage, headers, message, trailers} = res;
                    const temp: IGrpcResult<TReceive> = {
                        data: message as TReceive,
                        statusCode: status,
                    };
                    resolve(temp);

                }
            });
        });
        return requestPromise
    }

}
