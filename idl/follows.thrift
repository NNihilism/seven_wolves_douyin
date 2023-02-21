namespace go follows

struct User {
    1: i64 id
    2: string username
    3: i64 follow_count
    4: i64 follower_count
    5: string password
}

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct RelationActionRequest {
    1: i64 user_id
    2: optional string token
    3: i64 to_user_id  // 对方id
    4: string action_type // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: BaseResp base_resp
}

// 关注的用户列表
struct GetFollowListRequest {
    1: i64 user_id
    2: optional string token
}

struct GetFollowListResponse {
    1: list<User> users
    2: BaseResp base_resp
}

// 粉丝列表
struct GetFollowerListRequest {
    1: i64 user_id
    2: optional string token
}

struct GetFollowerListResponse {
    1: list<User> users
    2: BaseResp base_resp
}

// 关注的用户数量
struct GetFollowCountRequest {
    1: i64 user_id
    2: optional string token
}

struct GetFollowCountResponse {
    1: i64 follow_count
    2: BaseResp base_resp
}

// 粉丝数量
struct GetFollowerCountRequest {
    1: i64 user_id
    2: optional string token
}

struct GetFollowerCountResponse {
    1: i64 follower_count
    2: BaseResp base_resp
}

service FollowService {
    // 关注或取关
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    // 获取关注的用户列表
    GetFollowListResponse GetFollowList(1: GetFollowListRequest req)
    // 获取粉丝列表
    GetFollowerListResponse GetFollowerList(1: GetFollowerListRequest req)
    // 获取关注的用户数量
    GetFollowCountResponse GetFollowCount(1: GetFollowCountRequest req)
    // 获取粉丝数量
    GetFollowerCountResponse GetFollowerCount(1: GetFollowerCountRequest req)
}