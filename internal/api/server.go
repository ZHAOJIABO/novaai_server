package api

import (
	_ "context"
	_ "errors"
	_ "regexp"
	"time"

	_ "na_novaai_server/internal/na_interface"
)

const (
	NovaServerVersion    = "__NOVA_SERVER_VERSION__"
	lastNoLimitVersion   = "1.5.0"
	defaultStreamTimeout = 30 * time.Second
)

//var (
//	ErrNotImplemented    = errors.New("not implemented")
//	ErrInvalidPhone      = errors.New("invalid phone number")
//	ErrInvalidVerifyCode = errors.New("invalid verification code")
//	ErrInvalidChatID     = errors.New("chat id is required")
//
//	phoneNumberRegex = regexp.MustCompile(`^\+\d{2}1\d{10}$`)
//)
//
//type VisionAiServer struct {
//	UserService      *service.UserService
//	ChatService      *service.ChatService
//	MsgService       *service.MessageService
//	ConfigService    *service.ConfigService
//	UploadService    *service.UploadService
//	SubscribeService *service.SubscribeService
//	PromptService    *service.PromptService
//	UserServer       *UserServer
//	ChatServer       *ChatServer
//	MediaServer      *MediaServer
//	SystemServer     *SystemServer
//
//	vai.UnimplementedVisionAiServiceServer
//}
//
//func NewVisionAiServer(
//	mediaServer *MediaServer,
//	systemServer *SystemServer,
//	userService *service.UserService,
//	userPersonalInfoService *service.UserPersonalInfoService,
//	chatService *service.ChatService,
//	msgService *service.MessageService,
//	configService *service.ConfigService,
//	uploadService *service.UploadService,
//	subscribeService *service.SubscribeService,
//	promptService *service.PromptService,
//	userServer *UserServer,
//	chatServer *ChatServer,
//) *VisionAiServer {
//
//	server := &VisionAiServer{
//		UserService:      userService,
//		ChatService:      chatService,
//		MsgService:       msgService,
//		ConfigService:    configService,
//		UploadService:    uploadService,
//		SubscribeService: subscribeService,
//		PromptService:    promptService,
//	}
//
//	// Initialize specialized servers
//	server.UserServer = userServer
//	server.ChatServer = NewChatServer(chatService, userService, msgService, subscribeService, promptService)
//	server.MediaServer = mediaServer
//	server.SystemServer = systemServer
//	return server
//}
//
//// Ping tests if the service is healthy
//func (s *VisionAiServer) Ping(ctx context.Context, req *vai.Empty) (*vai.Empty, error) {
//	return &vai.Empty{}, nil
//}
//
//// GetAppConfiguration delegates to SystemServer
//func (s *VisionAiServer) GetAppConfiguration(ctx context.Context, req *vai.ConfigurationRequest) (*vai.ConfigurationResponse, error) {
//	return s.SystemServer.GetAppConfiguration(ctx, req)
//}
//
//// SendVerifyCode delegates to UserServer
//func (s *VisionAiServer) SendVerifyCode(ctx context.Context, req *vai.VerifyRequest) (*vai.VerifyResponse, error) {
//	return s.UserServer.SendVerifyCode(ctx, req)
//}
//
//// Login delegates to UserServer
//func (s *VisionAiServer) Login(ctx context.Context, req *vai.LoginRequest) (*vai.LoginResponse, error) {
//	return s.UserServer.Login(ctx, req)
//}
//
//// RefreshToken delegates to UserServer
//func (s *VisionAiServer) RefreshToken(ctx context.Context, req *vai.RefreshTokenRequest) (*vai.RefreshTokenResponse, error) {
//	return s.UserServer.RefreshToken(ctx, req)
//}
//
//// Logout delegates to UserServer
//func (s *VisionAiServer) Logout(ctx context.Context, req *vai.LogoutRequest) (*vai.LogoutResponse, error) {
//	return s.UserServer.Logout(ctx, req)
//}
//
//// GetUserInfo delegates to UserServer
//func (s *VisionAiServer) GetUserInfo(ctx context.Context, req *vai.UserInfoRequest) (*vai.UserInfoResponse, error) {
//	return s.UserServer.GetUserInfo(ctx, req)
//}
//
//// GetChatList delegates to ChatServer
//func (s *VisionAiServer) GetChatList(ctx context.Context, req *vai.ChatListRequest) (*vai.ChatListResponse, error) {
//	return s.ChatServer.GetChatList(ctx, req)
//}
//
//// GetChatMessage delegates to ChatServer
//func (s *VisionAiServer) GetChatMessage(ctx context.Context, req *vai.ChatMessageListRequest) (*vai.ChatMessageListResponse, error) {
//	return s.ChatServer.GetChatMessage(ctx, req)
//}
//
//// UploadImage delegates to MediaServer
//func (s *VisionAiServer) UploadImage(ctx context.Context, req *vai.UploadRequest) (*vai.UploadResponse, error) {
//	return s.MediaServer.UploadImage(ctx, req)
//}
//
//// ChatArchive delegates to ChatServer
//func (s *VisionAiServer) ChatArchive(ctx context.Context, req *vai.ChatArchiveRequest) (*vai.ChatArchiveResponse, error) {
//	return s.ChatServer.ChatArchive(ctx, req)
//}
//
//// ClearChatHistory delegates to ChatServer
//func (s *VisionAiServer) ClearChatHistory(ctx context.Context, req *vai.ClearChatHistoryRequest) (*vai.ClearChatHistoryResponse, error) {
//	return s.ChatServer.ClearChatHistory(ctx, req)
//}
//
//// GetFeedBackConfig delegates to SystemServer
//// It retrieves the feedback configuration, including the popup count and interval,
//// from the SystemServer. If an error occurs during retrieval, it returns an error response.
//// Otherwise, it returns a successful response containing the PopupCount and PopupInterval.
//// This function is used to determine how often and when to prompt the user for feedback.
//func (s *VisionAiServer) GetFeedBackConfig(ctx context.Context, req *vai.FeedbackConfigRequest) (*vai.FeedbackConfigResponse, error) {
//	return s.SystemServer.GetFeedBackConfig(ctx, req)
//}
//
//func (s *VisionAiServer) buildUserInfo(ctx context.Context, lang, os string, userInfo *vai.UserInfo) error {
//	var popAdsConfig struct {
//		ShowOpenAds  bool `json:"show_open_ads"`
//		ShowInnerAds bool `json:"show_inner_ads"`
//	}
//	if err := s.ConfigService.GetPopAdsConfig(&popAdsConfig); err != nil {
//		zlog.LogWithContext(ctx).Error("GetPopAdsConfig Error", zap.Error(err))
//	}
//
//	var popSubscribeConfig struct {
//		PopSubscribePageCount int `json:"pop_subscribe_page_count"`
//		PopSubscribeInterval  int `json:"pop_subscribe_interval"`
//	}
//	if err := s.ConfigService.GetPopSubscribeConfig(&popSubscribeConfig); err != nil {
//		zlog.LogWithContext(ctx).Error("GetPopSubscribeConfig Error", zap.Error(err))
//	}
//
//	shareURL, err := s.ConfigService.GetConfValue("android_share_url")
//	if err != nil {
//		zlog.LogWithContext(ctx).Error("GetConfValue Error", zap.Error(err))
//	}
//
//	subscribeInfo, err := s.SubscribeService.GetSubscribeInfo(ctx, userInfo.GetUserId(), os, lang, "")
//	if err != nil {
//		zlog.LogWithContext(ctx).Error("GetSubscribeInfo Error", zap.Error(err))
//	}
//
//	shouldPopSubscribe, err := s.ConfigService.ShouldPopSubscribe(ctx, userInfo.GetUserId())
//	if err != nil {
//		zlog.LogWithContext(ctx).Error("GetShouldPopSubscribe Error", zap.Error(err))
//	}
//
//	userInfo.UserConfig = &vai.UserConfig{
//		ShareUrl:     shareURL,
//		PopSubscribe: shouldPopSubscribe,
//	}
//	userInfo.SubscribeInfo = subscribeInfo
//
//	// Disable pop subscribe for active subscribers
//	if subscribeInfo != nil && subscribeInfo.GetStatus() == vai.SubscribeStatus_SubscribeActivate {
//		userInfo.UserConfig.PopSubscribe = false
//	}
//
//	return nil
//}
//
//// SendChatMessageStream delegates to ChatServer
//func (s *VisionAiServer) SendChatMessageStream(req *vai.ChatMessageSendRequest, streamServer vai.VisionAiService_SendChatMessageStreamServer) error {
//	return s.ChatServer.SendChatMessageStream(req, streamServer)
//}
