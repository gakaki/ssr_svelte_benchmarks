package lib

type BodyData[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type InformationFLow struct {
	NextCursor string `json:"next_cursor"`
	IsInhouse  bool   `json:"is_inhouse"`
	ItemCount  int    `json:"item_count"`
	Items      []struct {
		ResourceType  string `json:"resource_type"`
		ResourceOwner string `json:"resource_owner"`
		Resource      struct {
			Author struct {
				Avatar      string `json:"avatar"`
				DisplayName string `json:"display_name"`
				ID          int    `json:"id"`
				URI         string `json:"uri"`
			} `json:"author"`
			Categories []struct {
				PropertyKey  string `json:"property_key"`
				PropertyName string `json:"property_name"`
			} `json:"categories"`
			ContentArgs  []interface{} `json:"content_args"`
			ContentShort string        `json:"content_short"`
			DisplayTime  int           `json:"display_time"`
			HangDown     struct {
			} `json:"hang_down"`
			ID    int `json:"id"`
			Image struct {
				URI    string `json:"uri"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
				Size   int    `json:"size"`
			} `json:"image"`
			IsInVipPrivilege bool   `json:"is_in_vip_privilege"`
			IsIstioAPI       bool   `json:"is_istio_api"`
			IsPaid           bool   `json:"is_paid"`
			IsPriced         bool   `json:"is_priced"`
			IsTrial          bool   `json:"is_trial"`
			Layout           string `json:"layout"`
			LimitedTime      int    `json:"limited_time"`
			RelatedThemes    []struct {
				ID       int    `json:"id"`
				Title    string `json:"title"`
				ImageURI string `json:"image_uri"`
				URI      string `json:"uri"`
				Key      string `json:"key"`
			} `json:"related_themes"`
			SourceName string `json:"source_name"`
			SourceURI  string `json:"source_uri"`
			Subtitle   string `json:"subtitle"`
			Title      string `json:"title"`
			URI        string `json:"uri"`
			VipType    string `json:"vip_type"`
		} `json:"resource"`
	} `json:"items"`
	NextPageToken string `json:"next_page_token"`
}

type LiveFLow struct {
	Items []struct {
		Article interface{} `json:"article"`
		Author  struct {
			Avatar      string `json:"avatar"`
			DisplayName string `json:"display_name"`
			ID          int64  `json:"id"`
			IsFollowed  bool   `json:"is_followed"`
			URI         string `json:"uri"`
		} `json:"author"`
		BanComment        bool          `json:"ban_comment"`
		Channels          []string      `json:"channels"`
		CommentCount      int           `json:"comment_count"`
		Content           string        `json:"content"`
		ContentMore       string        `json:"content_more"`
		ContentText       string        `json:"content_text"`
		CoverImages       []interface{} `json:"cover_images"`
		DisplayTime       int           `json:"display_time"`
		FundCodes         []interface{} `json:"fund_codes"`
		GlobalChannelName string        `json:"global_channel_name"`
		GlobalMoreURI     string        `json:"global_more_uri"`
		HasLiveReading    bool          `json:"has_live_reading"`
		HighlightTitle    string        `json:"highlight_title"`
		ID                int           `json:"id"`
		Images            []interface{} `json:"images"`
		IsCalendar        bool          `json:"is_calendar"`
		IsFavourite       bool          `json:"is_favourite"`
		IsPriced          bool          `json:"is_priced"`
		IsScaling         bool          `json:"is_scaling"`
		Reference         string        `json:"reference"`
		RelatedThemes     []interface{} `json:"related_themes"`
		Score             int           `json:"score"`
		Symbols           []interface{} `json:"symbols"`
		Tags              []interface{} `json:"tags"`
		Title             string        `json:"title"`
		Type              string        `json:"type"`
		URI               string        `json:"uri"`
		CalendarKey       string        `json:"calendar_key,omitempty"`
		WscnTicker        string        `json:"wscn_ticker,omitempty"`
	} `json:"items"`
	NextCursor    string `json:"next_cursor"`
	OpCursor      string `json:"op_cursor"`
	PollingCursor string `json:"polling_cursor"`
}

type LessonMaster struct {
	Items []struct {
		AssetTags []struct {
			BannerKey    string `json:"banner_key"`
			IsUnderlying bool   `json:"is_underlying"`
			Key          string `json:"key"`
			Name         string `json:"name"`
			Path         string `json:"path"`
		} `json:"asset_tags"`
		CommentCount int         `json:"comment_count"`
		ContentShort string      `json:"content_short"`
		CustomTag    string      `json:"custom_tag"`
		DisplayTime  int         `json:"display_time"`
		HangDown     interface{} `json:"hang_down"`
		ID           int         `json:"id"`
		Image        struct {
			URI    string `json:"uri"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
			Size   int    `json:"size"`
		} `json:"image"`
		ImageScore int `json:"image_score"`
		Images     []struct {
			URI    string `json:"uri"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
			Size   int    `json:"size"`
		} `json:"images"`
		InfluenceScore   int           `json:"influence_score"`
		IsInVipPrivilege bool          `json:"is_in_vip_privilege"`
		IsPriced         bool          `json:"is_priced"`
		IsTrial          bool          `json:"is_trial"`
		LikeCount        int           `json:"like_count"`
		LikeImage        string        `json:"like_image"`
		LikeType         string        `json:"like_type"`
		LimitedTime      int           `json:"limited_time"`
		PrivilegeConfig  []interface{} `json:"privilege_config"`
		RelatedTopics    []struct {
			GroupID  int    `json:"group_id"`
			ID       int    `json:"id"`
			ImageURI string `json:"image_uri"`
			Title    string `json:"title"`
			URI      string `json:"uri"`
		} `json:"related_topics"`
		SourceName         string `json:"source_name"`
		SourceURI          string `json:"source_uri"`
		Subtitle           string `json:"subtitle"`
		Title              string `json:"title"`
		UnshowContentShort bool   `json:"unshow_content_short"`
		URI                string `json:"uri"`
		VipType            string `json:"vip_type"`
		VipURI             string `json:"vip_uri"`
	} `json:"items"`
	NextCursor string `json:"next_cursor"`
}

type MarketReal struct {
	Fields   []string `json:"fields"`
	Snapshot struct {
		Zero00001SS  []interface{} `json:"000001.SS"`
		Three99001SZ []interface{} `json:"399001.SZ"`
		Three99006SZ []interface{} `json:"399006.SZ"`
		EURUSDOTC    []interface{} `json:"EURUSD.OTC"`
		US10YROTC    []interface{} `json:"US10YR.OTC"`
		USDJPYOTC    []interface{} `json:"USDJPY.OTC"`
	} `json:"snapshot"`
}

type Article struct {
	DayItems []struct {
		CommentCount  int    `json:"comment_count"`
		DisplayTime   int    `json:"display_time"`
		DisplayUserID int64  `json:"display_user_id"`
		ID            int    `json:"id"`
		Pageviews     int    `json:"pageviews"`
		SourceURI     string `json:"source_uri"`
		Title         string `json:"title"`
		URI           string `json:"uri"`
	} `json:"day_items"`
	OfflineIds []int `json:"offline_ids"`
	WeekItems  []struct {
		CommentCount  int    `json:"comment_count"`
		DisplayTime   int    `json:"display_time"`
		DisplayUserID int64  `json:"display_user_id"`
		ID            int    `json:"id"`
		Pageviews     int    `json:"pageviews"`
		SourceURI     string `json:"source_uri"`
		Title         string `json:"title"`
		URI           string `json:"uri"`
	} `json:"week_items"`
}

type PageMain struct {
	MarketReal      *MarketReal
	InformationFlow *InformationFLow
	LiveFLow        *LiveFLow
	LessonMaster    *LessonMaster
	Article         *Article
}
