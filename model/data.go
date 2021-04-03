package model

type Insta struct {
	Test                  func(int) string
	LoggingPageID         string `json:"logging_page_id"`
	ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
	ShowFollowDialog      bool   `json:"show_follow_dialog"`
	LenghtEdge            []int
	Graphql               struct {
		User struct {
			Biography              string `json:"biography"`
			BlockedByViewer        bool   `json:"blocked_by_viewer"`
			RestrictedByViewer     bool   `json:"restricted_by_viewer"`
			CountryBlock           bool   `json:"country_block"`
			ExternalURL            string `json:"external_url"`
			ExternalURLLinkshimmed string `json:"external_url_linkshimmed"`
			EdgeFollowedBy         struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
			Fbid             string `json:"fbid"`
			FollowedByViewer bool   `json:"followed_by_viewer"`
			EdgeFollow       struct {
				Count int `json:"count"`
			} `json:"edge_follow"`
			FollowsViewer        bool        `json:"follows_viewer"`
			FullName             string      `json:"full_name"`
			HasArEffects         bool        `json:"has_ar_effects"`
			HasClips             bool        `json:"has_clips"`
			HasGuides            bool        `json:"has_guides"`
			HasChannel           bool        `json:"has_channel"`
			HasBlockedViewer     bool        `json:"has_blocked_viewer"`
			HighlightReelCount   int         `json:"highlight_reel_count"`
			HasRequestedViewer   bool        `json:"has_requested_viewer"`
			ID                   string      `json:"id"`
			IsBusinessAccount    bool        `json:"is_business_account"`
			IsJoinedRecently     bool        `json:"is_joined_recently"`
			BusinessCategoryName interface{} `json:"business_category_name"`
			OverallCategoryName  interface{} `json:"overall_category_name"`
			CategoryEnum         interface{} `json:"category_enum"`
			CategoryName         interface{} `json:"category_name"`
			IsPrivate            bool        `json:"is_private"`
			IsVerified           bool        `json:"is_verified"`
			EdgeMutualFollowedBy struct {
				Count int           `json:"count"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_mutual_followed_by"`
			ProfilePicURL          string      `json:"profile_pic_url"`
			ProfilePicURLHd        string      `json:"profile_pic_url_hd"`
			RequestedByViewer      bool        `json:"requested_by_viewer"`
			ShouldShowCategory     bool        `json:"should_show_category"`
			Username               string      `json:"username"`
			ConnectedFbPage        interface{} `json:"connected_fb_page"`
			EdgeFelixVideoTimeline struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_felix_video_timeline"`
			EdgeOwnerToTimelineMedia struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						Typename   string `json:"__typename"`
						ID         string `json:"id"`
						Shortcode  string `json:"shortcode"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayURL            string `json:"display_url"`
						EdgeMediaToTaggedUser struct {
							Edges []interface{} `json:"edges"`
						} `json:"edge_media_to_tagged_user"`
						FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
						FactCheckInformation   interface{} `json:"fact_check_information"`
						GatingInfo             interface{} `json:"gating_info"`
						SharingFrictionInfo    struct {
							ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
							BloksAppURL               interface{} `json:"bloks_app_url"`
						} `json:"sharing_friction_info"`
						MediaOverlayInfo interface{} `json:"media_overlay_info"`
						MediaPreview     interface{} `json:"media_preview"`
						Owner            struct {
							ID       string `json:"id"`
							Username string `json:"username"`
						} `json:"owner"`
						IsVideo              bool   `json:"is_video"`
						AccessibilityCaption string `json:"accessibility_caption"`
						EdgeMediaToCaption   struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						EdgeMediaToComment struct {
							Count int `json:"count"`
						} `json:"edge_media_to_comment"`
						CommentsDisabled bool `json:"comments_disabled"`
						TakenAtTimestamp int  `json:"taken_at_timestamp"`
						TakenAt          string
						EdgeLikedBy      struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						EdgeMediaPreviewLike struct {
							Count int `json:"count"`
						} `json:"edge_media_preview_like"`
						Location           interface{} `json:"location"`
						ThumbnailSrc       string      `json:"thumbnail_src"`
						ThumbnailResources []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"thumbnail_resources"`
						EdgeSidecarToChildren struct {
							Edges []struct {
								Node struct {
									Typename   string `json:"__typename"`
									ID         string `json:"id"`
									Shortcode  string `json:"shortcode"`
									Dimensions struct {
										Height int `json:"height"`
										Width  int `json:"width"`
									} `json:"dimensions"`
									DisplayURL            string `json:"display_url"`
									EdgeMediaToTaggedUser struct {
										Edges []interface{} `json:"edges"`
									} `json:"edge_media_to_tagged_user"`
									FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
									FactCheckInformation   interface{} `json:"fact_check_information"`
									GatingInfo             interface{} `json:"gating_info"`
									SharingFrictionInfo    struct {
										ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
										BloksAppURL               interface{} `json:"bloks_app_url"`
									} `json:"sharing_friction_info"`
									MediaOverlayInfo interface{} `json:"media_overlay_info"`
									MediaPreview     string      `json:"media_preview"`
									Owner            struct {
										ID       string `json:"id"`
										Username string `json:"username"`
									} `json:"owner"`
									IsVideo              bool   `json:"is_video"`
									AccessibilityCaption string `json:"accessibility_caption"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_sidecar_to_children"`
					} `json:"node,omitempty"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
			EdgeSavedMedia struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_saved_media"`
			EdgeMediaCollections struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_collections"`
		} `json:"user"`
	} `json:"graphql"`
	ToastContentOnLoad      interface{} `json:"toast_content_on_load"`
	ShowViewShop            bool        `json:"show_view_shop"`
	ProfilePicEditSyncProps interface{} `json:"profile_pic_edit_sync_props"`
}
