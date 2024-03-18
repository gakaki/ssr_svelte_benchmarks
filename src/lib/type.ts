export interface PageMain {
    MarketReal: MarketReal
    InformationFlow: InformationFlow
    LiveFLow: LiveFlow
    LessonMaster: LessonMaster
    Article: Article
  }
  
  export interface MarketReal {
    fields: string[]
    snapshot: Snapshot
  }
  
  export interface Snapshot {
    "000001.SS": [string, number, number, number, number, string, number]
    "399001.SZ": [string, number, number, number, number, string, number]
    "399006.SZ": [string, number, number, number, number, string, number]
    "EURUSD.OTC": any
    "US10YR.OTC": [string, number, number, number, number, string, number]
    "USDJPY.OTC": any
  }
  
  export interface InformationFlow {
    next_cursor: string
    is_inhouse: boolean
    item_count: number
    items: Item[]
    next_page_token: string
  }
  
  export interface Item {
    resource_type: string
    resource_owner: string
    resource: Resource
  }
  
  export interface Resource {
    author: Author
    categories: Category[]
    content_args?: ContentArg[]
    content_short: string
    display_time: number
    hang_down: HangDown
    id: number
    image: Image
    is_in_vip_privilege: boolean
    is_istio_api: boolean
    is_paid: boolean
    is_priced: boolean
    is_trial: boolean
    layout: string
    limited_time: number
    related_themes?: RelatedTheme[]
    source_name: string
    source_uri: string
    subtitle: string
    title: string
    uri: string
    vip_type: string
  }
  
  export interface Author {
    avatar: string
    display_name: string
    id: number
    uri: string
  }
  
  export interface Category {
    property_key: string
    property_name: string
  }
  
  export interface ContentArg {
    class: string
    height: string
    placeholder: string
    src: string
    type: string
    width: string
    size?: string
  }
  
  export interface HangDown {}
  
  export interface Image {
    uri: string
    height: number
    width: number
    size: number
  }
  
  export interface RelatedTheme {
    id: number
    title: string
    image_uri: string
    uri: string
    key: string
  }
  
  export interface LiveFlow {
    items: LiveFlowItem[]
    next_cursor: string
    op_cursor: string
    polling_cursor: string
  }
  
  export interface LiveFlowItem {
    article?: Article
    author: LiveFlowItemAuthor
    ban_comment: boolean
    channels: string[]
    comment_count: number
    content: string
    content_more: string
    content_text: string
    cover_images: CoverImage[]
    display_time: number
    fund_codes: any[]
    global_channel_name: string
    global_more_uri: string
    has_live_reading: boolean
    highlight_title: string
    id: number
    images: LiveFlowItemItem[]
    is_calendar: boolean
    is_favourite: boolean
    is_priced: boolean
    is_scaling: boolean
    reference: string
    related_themes: any[]
    score: number
    symbols: any[]
    tags: any[]
    title: string
    type: string
    uri: string
  }
  
  export interface Article {
    id: number
    image: ArticleImage
    platforms: string[]
    title: string
    uri: string
  }
  
  export interface ArticleImage {
    height: number
    size: number
    uri: string
    width: number
  }
  
  export interface LiveFlowItemAuthor {
    avatar: string
    display_name: string
    id: number
    is_followed: boolean
    uri: string
  }
  
  export interface CoverImage {
    height: number
    size: number
    uri: string
    width: number
  }
  
  export interface LiveFlowItemItem {
    height: number
    size: number
    uri: string
    width: number
  }
  
  export interface LessonMaster {
    items: LessonMasterItem[]
    next_cursor: string
  }
  
  export interface LessonMasterItem {
    asset_tags: AssetTag[]
    comment_count: number
    content_short: string
    custom_tag: string
    display_time: number
    hang_down: any
    id: number
    image: LessonMasterItemImage
    image_score: number
    images: LessonMasterItemImages[]
    influence_score: number
    is_in_vip_privilege: boolean
    is_priced: boolean
    is_trial: boolean
    like_count: number
    like_image: string
    like_type: string
    limited_time: number
    privilege_config: any[]
    related_topics: RelatedTopic[]
    source_name: string
    source_uri: string
    subtitle: string
    title: string
    unshow_content_short: boolean
    uri: string
    vip_type: string
    vip_uri: string
  }
  
  export interface AssetTag {
    banner_key: string
    is_underlying: boolean
    key: string
    name: string
    path: string
  }
  
  export interface LessonMasterItemImage {
    uri: string
    height: number
    width: number
    size: number
  }
  
  export interface LessonMasterItemImages {
    uri: string
    height: number
    width: number
    size: number
  }
  
  export interface RelatedTopic {
    group_id: number
    id: number
    image_uri: string
    title: string
    uri: string
  }
  
  export interface Article {
    day_items: DayItem[]
    offline_ids: number[]
    week_items: WeekItem[]
  }
  
  export interface DayItem {
    comment_count: number
    display_time: number
    display_user_id: number
    id: number
    pageviews: number
    source_uri: string
    title: string
    uri: string
  }
  
  export interface WeekItem {
    comment_count: number
    display_time: number
    display_user_id: number
    id: number
    pageviews: number
    source_uri: string
    title: string
    uri: string
  }
  