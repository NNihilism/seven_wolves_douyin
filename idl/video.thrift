namespace go video

struct Video {
    1: required i64 id               // 视频唯一标识
    2: required User author          // 视频作者信息
    3: required string play_url      // 视频播放地址
    4: required string cover_url     // 视频封面地址
    5: required i64 favorite_count   // 视频的点赞总数
    6: required i64 comment_count     // 视频的评论总数
    7: required bool is_favorite      // true-已点赞，false-未点赞
    8: required string title          // 视频标题
}

//struct User {
//    1: required i64 id               // 用户id
//    2: required string name          // 用户名称
//    3: optional i64 follow_count     // 关注总数
//    4: optional i64 follower_count   // 粉丝总数
//    5: required bool is_follow       // true-已关注，false-未关注
//}
struct User {
    1: i64 id
    2: string username
    3: i64 follow_count
    4: i64 follower_count
    5: string password
}

struct FeedRequest {
    1: optional i64 latest_time (api.query="latest_time")  // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token    (api.query="token")        // 可选参数，登录用户设置
}

struct FeedResponse {
    1: required i32 status_code     // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 返回状态描述
    3: list<Video> video_list       // 视频列表
    4: optional i64 next_time       // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct PublishActionRequest {
    1: required string token    (api.query="token")            // 用户鉴权token
    2: required binary data     (api.query="data")             // 视频数据
    3: required string title    (api.query="title")            // 视频标题
}

struct PublishActionResponse {
    1: required i32 status_code         // 状态码，0-成功，其他值-失败
    2: optional string status_msg       // 返回状态描述
}

struct PublishListRequest {
    1: required i64 user_id     (api.query="user_id")   // 用户id
    2: required string token    (api.query="token")     // 用户鉴权token
}

struct PublishListResponse {
    1: required i32 status_code      // 状态码，0-成功，其他值-失败
    2: optional string status_msg    // 返回状态描述
    3: list<Video> video_list       // 用户发布的视频列表
}

service VideoService {
    FeedResponse GetFeed(1: FeedRequest req)    (api.get="/douyin/feed")
    PublishActionResponse PublishVideo(1: PublishActionRequest req)     (api.post="/douyin/publish/action")
    PublishListResponse GetPublishVideoList(1: PublishListRequest req)  (api.get="/douyin/publish/list")
}