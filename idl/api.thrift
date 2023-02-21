namespace go follows

struct UserReq {
    1: string name (api.query="username");
    2: string pwd (api.query="password");
}

struct UserResp {
    1: i64 status_code;
    2: string status_msg;
    3: i64 user_id;
    4: string token;
}

service UserService {
    UserResp Login(1: UserReq request) (api.get="/login");
    UserResp Register(1: UserReq request) (api.get="/register");
}
# struct CommentReq {
#     1: string Name (api.query="name");
# }

# struct CommentResp {
#     1: string RespBody;
# }

# service CommentService {
#     CommentResp AddComment(1: CommentReq request) (api.get="/addcomment");
#     CommentResp DelComment(1: CommentReq request) (api.get="/delcomment");
# }

# struct VideoReq {
#     1: string Name (api.query="name");
# }

# struct VideoResp {
#     1: string RespBody;
# }

# service VideoService {
#     VideoResp AddVideo(1: VideoReq request) (api.get="/addVideo");
#     VideoResp DelVideo(1: VideoReq request) (api.get="/delVideo");
# }

# service AnotherVideoService {
#     VideoResp AddVideo(1: VideoReq request) (api.get="/addVideo");
#     VideoResp DelVideo(1: VideoReq request) (api.get="/delVideo");
# }