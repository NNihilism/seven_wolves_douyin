namespace go like

struct Like {
    1: required i64 id
    2: required User user  
    3: required Video video
    4: required i32 action_type
}

struct User {
    1: i64 id
    2: string username
    3: i64 follow_count
    4: i64 follower_count
    5: string password
}

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

struct LikeActionRequest {
    1: required string token    (api.query="token")
    2: required i64 video_id    (api.query="video_id")
    3: required i32 action_type // 1-点赞，2-取消点赞
}

struct LikeActionResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg
}

struct GetLikeListRequest {
    1: required i64 user_id     (api.query="user_id")   // 用户id
    2: required string token    (api.query="token")
}

struct GetLikeListResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg
    3: list<i64> video_ids  //用户点赞的视频id列表
}

service LikeService {
    LikeActionResponse LikeAction(1: LikeActionRequest req) (api.post="/douyin/favorite/action/")
    GetLikeListResponse GetLikeList(1: GetLikeListRequest req) (api.get="/douyin/favorite/list/")
}