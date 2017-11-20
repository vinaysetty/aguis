import {Assignment, Course, Enrollment, User} from "../../_proto/aguis/library/aguis_service_pb";
import {
    CourseGroupStatus,
    CourseUserState,
    courseUserStateToString,
    IAssignment,
    IBuildInfo,
    ICourse,
    ICourseGroup,
    ICourseUserLink,
    ICourseWithEnrollStatus,
    IError, INewCourse,
    INewGroup,
    IOrganization, IStatusCode,
    ISubmission,
    ITestCases,
    IUser,
} from "../models";

import { HttpHelper, IHTTPResult } from "../HttpHelper";
import { ICourseProvider } from "./CourseManager";

import {
    ICourseEnrollemtnt,
    IEnrollment,
    isCourseEnrollment,
    isUserEnrollment,
    IUserEnrollment,
    IUserProvider,
} from "../managers";

import { IMap, mapify } from "../map";
import { combinePath } from "../NavigationHelper";
import {GrpcHelper} from "./GrpcHelper";
import { ILogger } from "./LogManager";

export class ServerProvider implements IUserProvider, ICourseProvider {
    private helper: HttpHelper;
    private logger: ILogger;
    private grpcHelper: GrpcHelper;

    constructor(helper: HttpHelper,grpc: GrpcHelper, logger: ILogger) {
        this.helper = helper;
        this.logger = logger;
        this.grpcHelper = grpc;
    }

    public async getCourses(): Promise<ICourse[]> {
        // const result = await this.helper.get<any>("courses");
        const result = await this.grpcHelper.getCourses();
        if (result.statusCode !== 0 || !result.data) {
            this.handleError(result, "getCourses");
            return [];
        }
        const data = result.data.getCoursesList().map((course) => {
            return this.toICourse(course);
        });
        return data;
        // const data = JSON.parse(JSON.stringify(result.data).toLowerCase()) as ICourse[];
        // return mapify(result.data, (ele) => ele.id);
        // throw new Error("Method not implemented.");
    }

    public async getCoursesFor(user: IUser, state?: CourseUserState[]): Promise<ICourseEnrollemtnt[]> {
        // TODO: Fix to use correct url request
        const status = state ? courseUserStateToString(state) : "";
        // const result = await this.helper.get<ICourseWithEnrollStatus[]>("/users/" + user.id + "/courses" + status);
        const result = await this.grpcHelper.getCoursesWithEnrollment(user.id, status);
        if (result.statusCode !== 0 || !result.data) {
            this.handleError(result, "getCoursesFor");
            return [];
        }

        const arr: ICourseEnrollemtnt[] = [];
        result.data.getCoursesList().forEach((ele) => {
            const course = this.toICourse(ele);
            const enroll = ele.getEnrolled() as number >= 0 ? ele.getEnrolled() : undefined;
            arr.push({
                course,
                status: enroll,
                courseid: course.id,
                userid: user.id,
                user,
            });
        });
        return arr;
    }

    public async getUsersForCourse(course: ICourse, state?: CourseUserState[]): Promise<IUserEnrollment[]> {
        const status = state ? courseUserStateToString(state) : "";
        const result = await this.grpcHelper.getEnrollmentsByCourse(course.id, status);
        if (result.statusCode !== 0 || !result.data) {
            this.handleError(result, "getUserForCourse");
            return [];
        }

        const arr: IUserEnrollment[] = [];
        result.data.getEnrollmentsList().forEach((ele) => {
            const enroll: IEnrollment = this.toIEnrollment(ele);
            if (isCourseEnrollment(enroll)) {
                enroll.user = this.makeUserInfo(enroll.user);
                arr.push(enroll);
            }
        });
        return arr;
    }

    public async getAssignments(courseId: number): Promise<IMap<IAssignment>> {
        // const result = await this.helper.get<any>("courses/" + courseId.toString() + "/assignments");
        const result = await this.grpcHelper.getAssignments(courseId);
        if (result.statusCode !== 0 || !result.data) {
            this.handleError(result, "getAssignments");
            throw new Error("Problem with the request");
        }
        const assignments: IAssignment[] = [];
        result.data.getAssignmentsList().forEach((ele) => {
            const asg = this.toIAssignment(ele);
            assignments.push(asg);
        });

        return mapify(assignments, (ele) => {
            ele.deadline = new Date(2017, 7, 18);
            return ele.id;
        });
    }

    public async addUserToCourse(user: IUser, course: ICourse): Promise<boolean> {
        const result = await this.helper.post<{ status: CourseUserState }, undefined>
            ("/courses/" + course.id + "/users/" + user.id, {
                status: CourseUserState.pending,
            });
        if (result.statusCode === 201) {
            return true;
        } else {
            this.handleError(result, "addUserToCourse");
        }
        return false;
    }

    public async changeUserState(link: ICourseUserLink, state: CourseUserState): Promise<boolean> {
        const result = await this.helper.patch<{ courseid: number, userid: number, status: CourseUserState }, undefined>
            ("/courses/" + link.courseId + "/users/" + link.userid, {
                courseid: link.courseId,
                userid: link.userid,
                status: state,
            });
        if (result.statusCode <= 202) {
            return true;
        } else {
            this.handleError(result, "changeUserState");
        }
        return false;
    }

    public async createNewCourse(courseData: INewCourse): Promise<ICourse | IError> {
        // const uri: string = "courses";
        // const result = await this.helper.post<INewCourse, ICourse>(uri, courseData);
        const result = await this.grpcHelper.createCourse(courseData);
        if (result.statusCode !== 0 || !result.data) {
            this.handleError(result, "createNewCourse");
            return this.parseError(result);
        }
        return this.toICourse(result.data);
    }

    public async getCourse(id: number): Promise<ICourse | null> {
        // const result = await this.helper.get<any>("courses/" + id);
        const result = await this.grpcHelper.getCourse(id);
        if (result.statusCode !== 0 || !result.data) {
            this.handleError(result, "getCourse");
            return null;
        }

        return this.toICourse(result.data);
    }

    public async updateCourse(courseId: number, courseData: ICourse): Promise<IStatusCode | IError> {
        const uri: string = "courses/" + courseId;
        const result = await this.helper.put<ICourse, ICourse>(uri, courseData);
        if (result.statusCode !== 200) {
            this.handleError(result, "updateCourse");
            return this.parseError(result);
        }
        return JSON.parse(JSON.stringify(result.statusCode)) as IStatusCode;
    }

    public async createGroup(groupData: INewGroup, courseId: number): Promise<ICourseGroup | IError> {
        const uri: string = "courses/" + courseId + "/groups";
        const result = await this.helper.post<INewGroup, ICourseGroup>(uri, groupData);
        if (result.statusCode !== 201 || !result.data) {
            this.handleError(result, "createGroup");
            return this.parseError(result);
        }
        return JSON.parse(JSON.stringify(result.data)) as ICourseGroup;
    }

    public async getCourseGroups(courseId: number): Promise<ICourseGroup[]> {
        const uri: string = "courses/" + courseId + "/groups";
        const result = await this.helper.get<ICourseGroup>(uri);
        if (result.statusCode !== 200 || !result.data) {
            this.handleError(result, "getCourseGroups");
            return [];
        }
        return JSON.parse(JSON.stringify(result.data)) as ICourseGroup[];
    }

    public async getGroupByUserAndCourse(userid: number, courseid: number): Promise<ICourseGroup | null> {
        const uri: string = "users/" + userid + "/courses/" + courseid + "/group";
        const result = await this.helper.get<ICourseGroup>(uri);
        if (result.statusCode !== 302 || !result.data) {
            return null;
        }
        return JSON.parse(JSON.stringify(result.data)) as ICourseGroup;
    }

    public async updateGroupStatus(groupId: number, st: CourseGroupStatus): Promise<boolean> {
        const uri: string = "groups/" + groupId;
        const data = { status: st };

        const result = await this.helper.patch<{ status: CourseGroupStatus }, undefined>(uri, data);
        if (result.statusCode === 200) {
            return true;
        } else {
            this.handleError(result, "updateGroupStatus");
        }
        return false;
    }

    public async getGroup(gid: number): Promise<ICourseGroup | null> {
        const uri: string = "groups/" + gid;
        const result = await this.helper.get<ICourseGroup>(uri);
        if (result.statusCode !== 200 || !result.data) {
            this.handleError(result, "getGroup");
            return null;
        }
        return JSON.parse(JSON.stringify(result.data)) as ICourseGroup;
    }

    public async deleteGroup(groupId: number): Promise<boolean> {
        const uri: string = "groups/" + groupId;
        const result = await this.helper.delete(uri);
        if (result.statusCode === 200) {
            return true;
        } else {
            this.handleError(result, "deleteGroup");
        }
        return false;
    }

    public async updateGroup(groupData: INewGroup, groupId: number, courseId: number): Promise<IStatusCode | IError> {
        const uri: string = "courses/" + courseId + "/groups/" + groupId;
        const result = await this.helper.put<INewGroup, any>(uri, groupData);
        if (result.statusCode !== 200) {
            this.handleError(result, "updateGroup");
            return this.parseError(result);
        }
        return JSON.parse(JSON.stringify(result.statusCode)) as IStatusCode;
    }

    // TODO change to use course id instead of getting all of them
    public async getAllLabInfos(courseId: number, userId: number): Promise<IMap<ISubmission>> {
        const result = await this.helper.get<ISubmission[]>(
            ("courses/" + courseId.toString() + "/users/" + userId + "/submissions"),
        );
        if (result.statusCode === 200 && result.data === undefined) {
            return {};
        }
        if (!result.data) {
            this.handleError(result, "getAllLabInfos");
            return {};
        }
        return mapify(result.data, (e) => {
            let a = "{\"builddate\": \"2017-07-28\", \"buildid\": 1, \"buildlog\": \"This is cool\", \"execTime\": 1}";
            let b = "[{\"name\": \"Test 1\", \"score\": 3, \"points\": 4, \"weight\": 100}]";
            if ((e as any).buildinfo && ((e as any).buildinfo as string).trim().length > 2) {
                a = (e as any).buildinfo as string;
            }
            if ((e as any).scoreobjects && ((e as any).scoreobjects as string).trim().length > 2) {
                b = (e as any).scoreobjects;
            }
            let tempInfo: IBuildInfo;
            let scoreObj: ITestCases[];
            try {
                tempInfo = JSON.parse(a);
            } catch (e) {
                tempInfo = JSON.parse(
                    "{\"builddate\": \"2017-07-28\", \"buildid\": 1, \"buildlog\": \"This is cool\", \"execTime\": 1}",
                );
            }
            try {
                scoreObj = JSON.parse(b);
            } catch (e) {
                scoreObj = JSON.parse(
                    "[{\"name\": \"Test 1\", \"score\": 3, \"points\": 4, \"weight\": 100}]",
                );
            }

            e.buildDate = tempInfo.builddate;
            e.buildId = tempInfo.buildid;
            e.buildLog = tempInfo.buildlog;
            e.executetionTime = tempInfo.exectime;
            e.failedTests = 0;
            e.passedTests = 1;
            e.testCases = scoreObj;
            return e.id;
        });
    }

    public async tryLogin(username: string, password: string): Promise<IUser | null> {
        throw new Error("tryLogin This could be removed since there is no normal login.");
    }

    public async grpcLogin(id: number): Promise<IUser | null> {
        const result = await this.grpcHelper.getUser(id);
        if (result.statusCode === 0 && result.data) {
            return this.toIUser(result.data);
        }
        return null;
    }

    public async logout(user: IUser): Promise<boolean> {
        window.location.assign("/logout");
        return true;
    }

    public async getAllUser(): Promise<IUser[]> {
        const result = await this.grpcHelper.getUsers();
        if (result.statusCode !== 0 || !result.data) {
            this.handleError(result, "getAllUser");
            return [];
        }

        const data = result.data.getUsersList().map((user) => {
            return this.toIUser(user);
        });
        const newArray = data.map<IUser>((ele) => this.makeUserInfo(ele));
        return newArray;
    }

    public async tryRemoteLogin(provider: string): Promise<IUser | null> {
        // TODO this needs to be fixed to return user data from provider
        if (provider.length > 0) {
            const requestString = "/auth/" + provider;
            window.location.assign(requestString);
        }
        return null;
    }

    public async changeAdminRole(user: IUser): Promise<boolean> {
        const result = await this.grpcHelper.updateUser(user, true);
        if (result.statusCode !== 0) {
            return false;
        } else {
            this.handleError(result, "changeAdminRole");
        }
        return true;
    }

    public async updateUser(user: IUser): Promise<boolean> {
        const result = await this.grpcHelper.updateUser(user);
        if (result.statusCode !== 0) {
            this.handleError(result, "updateUser");
        }
        return true;
    }

    // TODO: check if resp.status contain correct status
    public async getDirectories(provider: string): Promise<IOrganization[]> {
        const uri: string = "directories";
        const data: { provider: string } = { provider };
        const result = await this.helper.post<{ provider: string }, IOrganization[]>(uri, data);
        if (result.statusCode === 200 && result.data) {
            return result.data;
        } else {
            this.handleError(result, "getDirectories");
        }
        return [];
    }

    public async getProviders(): Promise<string[]> {
        // https://nicolasf.itest.run/api/v1/providers
        const result = await this.helper.get<string[]>("providers");
        if (result.data) {
            return result.data;
        } else {
            this.handleError(result, "getProviders");
        }
        return [];
    }

    public async getLoggedInUser(): Promise<IUser | null> {
        const result = await this.helper.get<IUser>("user");
        if (result.statusCode !== 302 || !result.data) {
            this.handleError(result, "getLoggedInUser");
            return null;
        }
        return this.makeUserInfo(result.data);
    }

    public async refreshCoursesFor(courseid: number): Promise<any> {
        const result = await this.helper.post<any, null>("courses/" + courseid + "/refresh", null);
        if (result.statusCode !== 200 || !result.data) {
            this.handleError(result, "getLoggedInUser");
            return null;
        }
        return this.makeUserInfo(result.data);
    }

    private makeUserInfo(data: IUser): IUser {
        return data;
    }

    private handleError(result: IHTTPResult<any>, sender?: string): void {
        this.logger.warn("Request to server failed with status code: " + result.statusCode, true);
        if (sender) {
            this.logger.warn("Failed request from: " + sender);
        }
    }

    private parseError(result: IHTTPResult<any>): IError {
        const error: IError = {
            statusCode: result.statusCode,
        };
        if (result.data) {
            error.data = JSON.parse(JSON.stringify(result.data)) as any;
        }
        return error;
    }
    // this method convert a grpc User to IUSer
    private toIUser(user: User): IUser {
        return {
            id: user.getId(),
            name: user.getName(),
            avatarurl: user.getAvatarurl(),
            email: user.getEmail(),
            studentid: user.getStudentid(),
            isadmin: user.getIsadmin(),
        } as IUser
    }

    // this method convert a grpc Course to ICourse
    private toICourse(course: Course): ICourse {
        return {
            id: course.getId(),
            name: course.getName(),
            code: course.getCode(),
            tag: course.getTag(),
            year: course.getYear(),
            provider: course.getProvider(),
            directoryid: course.getDirectoryid(),
        } as ICourse
    }

    // this method convert a grpc Assignment to IAssignment
    private toIAssignment(assg: Assignment): IAssignment {
        return {
            id: assg.getId(),
            name: assg.getName(),
            courseid: assg.getCourseid(),
        } as IAssignment
    }

    // this method convert a grpc Enrollment to IEnrollment
    private toIEnrollment(enrollment: Enrollment): IEnrollment {
        const ienroll: IEnrollment =  {
            userid: enrollment.getUserid(),
            courseid: enrollment.getCourseid(),
        };
        if (enrollment.getStatus() !== undefined) {
            ienroll.status = enrollment.getStatus();
        }

        const user: User | undefined = enrollment.getUser();
        if (user !== undefined) {
            ienroll.user = this.toIUser(user);
        }
        const course: Course | undefined = enrollment.getCourse();
        if (course !== undefined) {
            ienroll.course = this.toICourse(course);
        }

        return ienroll as IEnrollment;
    }

}
