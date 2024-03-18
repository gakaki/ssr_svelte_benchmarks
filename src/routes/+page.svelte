<script>
  import { error } from "@sveltejs/kit";

  export let data;
  // console.log(data.Article.offline_ids[0]);
  // console.log(data);

  // console.log(data);

  function toFixed2(num) {
    return parseFloat(num).toFixed(2) + "";
  }
  function formatNumberAndPercentage(item) {
    const percentage = item[4];
    const formattedNumber = toFixed2(item[3]);
    // 格式化百分比部分，将百分比转换为字符串，并保留两位小数
    const formattedPercentage = toFixed2(percentage) + "%";
    return `+${formattedNumber}(+${formattedPercentage})`;
  }

  let data_real_obj = {};
  for (let [key, value] of Object.entries(data.MarketReal.snapshot)) {
    if (value) {
      data_real_obj[key] = {
        name: value[0],
        num: toFixed2(value[2]),
        description: formatNumberAndPercentage(value),
      };
    }
  }
  // console.log(data_real_obj);

  let data_infomation_flows5 = [];
  for (const [key, value] of Object.entries(
    data.InformationFlow.items.slice(0, 5),
  )) {
    if (value) {
      // console.log(value);
      data_infomation_flows5.push({
        image: value.resource.image.uri,
        title: value.resource.title,
        url: value.resource.uri,
      });
    }
  }
  // console.log("dataflow5", data_infomation_flows5);

  let data_infomation_flows = [];
  for (const [key, value] of Object.entries(
    data.InformationFlow.items.slice(5, 25),
  )) {
    if (value) {
      // console.log(value);
      data_infomation_flows.push({
        image: value.resource.image.uri,
        title: value.resource.title,
        short: value.resource.content_short,
        author: value.resource.author.display_name,
        display_time: timeSince(value.resource.display_time),
        url: value.resource.uri,
      });
    }
  }
  // console.log("dataflow", data.InformationFlow.items, data_infomation_flows);

  function timeSince(timestamp) {
    const now = Date.now(); // 获取当前时间的时间戳（毫秒）
    const then = new Date(timestamp).getTime(); // 将给定的时间戳转换为时间戳（毫秒）

    const diff = now - then; // 计算时间差（毫秒）

    const seconds = Math.floor(diff / 1000); // 秒
    const minutes = Math.floor(seconds / 60); // 分钟
    const hours = Math.floor(minutes / 60); // 小时
    const days = Math.floor(hours / 24); // 天

    // 计算剩余的秒数、分钟数、小时数和天数
    const remainingSeconds = seconds % 60;
    const remainingMinutes = minutes % 60;
    const remainingHours = hours % 24;
    const remainingDays = days % 365; // 简化为不考虑闰年的每年365天

    // 根据时间差返回相应的描述
    if (remainingSeconds < 10) {
      return "刚刚";
    } else if (remainingSeconds < 60) {
      return `${remainingSeconds}秒前`;
    } else if (remainingMinutes < 60) {
      return `${remainingMinutes}分钟前`;
    } else if (remainingHours < 24) {
      return `${remainingHours}小时前`;
    } else if (remainingDays < 365) {
      return `${remainingDays}天前`;
    } else {
      return `${Math.floor(remainingDays / 365)}年前`;
    }
  }

  let data_articles = [];
  for (const [key, value] of Object.entries(data.Article.day_items)) {
    if (value) {
      // console.log(value);
      data_articles.push({
        num: parseInt(key) + 1,
        title: value.title,
        url: value.uri,
      });
    }
  }
  // console.log("data_article", data_articles);

  function formatTimestamp(timestamp) {
    // 创建一个新的Date对象，传入的时间戳需要乘以1000
    var date = new Date(timestamp * 1000);

    // 获取小时和分钟
    var hours = date.getHours();
    var minutes = date.getMinutes();

    // 将小时和分钟转换为字符串，并确保都是两位数
    var hoursStr = hours.toString().padStart(2, "0");
    var minutesStr = minutes.toString().padStart(2, "0");

    // 拼接成 HH:MM 格式
    var timeString = hoursStr + ":" + minutesStr;

    return timeString;
  }
  let data_live_flow_3 = [];
  for (const [key, value] of Object.entries(data.LiveFLow.items.slice(0, 3))) {
    if (value) {
      // console.log(value);
      data_live_flow_3.push({
        content: value.content_text,
        url: value.uri,
        display_time: formatTimestamp(value.display_time),
      });
    }
  }
  // console.log("data_live_flow_3", data_live_flow_3);

  function formatDateString(timestamp) {
    // 创建一个新的Date对象，传入的时间戳需要乘以1000
    var date = new Date(timestamp * 1000);

    // 获取日期的各个部分
    var month = ("0" + (date.getMonth() + 1)).slice(-2); // 月份是从0开始的，所以+1
    var day = ("0" + date.getDate()).slice(-2);
    var hours = ("0" + date.getHours()).slice(-2);
    var minutes = ("0" + date.getMinutes()).slice(-2);

    // 拼接成 "MM月DD日 HH:MM 更新：" 格式
    var formattedDate = month + "月" + day + "日 " + " 更新";
    // + hours + ":" + minutes
    return formattedDate;
  }
  let data_lession_master_4 = [];
  for (const [key, value] of Object.entries(
    data.LessonMaster.items.slice(0, 4),
  )) {
    if (value) {
      // console.log(value);
      data_lession_master_4.push({
        title: value.related_topics[0].title,
        content: value.content_short,
        url: value.uri,
        image: value.image.uri,
        display_time: formatDateString(value.display_time),
      });
    }
  }
  // console.log("data_lession_master_4", data_lession_master_4);
</script>

<svelte:head>
  <title>华尔街见闻</title>
  <meta name="description" content="华尔街见闻" />
</svelte:head>

<!-- 各种指数 -->
<div class="bk_real">
  {#each Object.entries(data_real_obj) as [k, v]}
    {#if v}
      <div class="flex">
        <span>{v.name}</span>
        <strong class="">{v.num}</strong>
        <span class="text-red-400 red">{v.description}</span>
      </div>
    {/if}
  {/each}
</div>

<!-- 五个news -->
<div class="bk_news5">
  {#each data_infomation_flows5 as v, index}
    {#if v}
      <a
        href={v.url}
        target="_blank"
        class={index === 0 ? "new5_item a_first" : "new5_item"}
        ><img src={v.image} class="img" alt="" />
        <div class="text">
          <div class="title">
            {v.title}
          </div>
        </div>
      </a>
    {/if}
  {/each}
</div>

<!-- 三个快讯 -->
<div class="bk_right_panel">
  <div class="title_bar">
    <div>快讯</div>
    <div>查看更多></div>
  </div>
  <div class="bk_right_panel_container">
    {#each data_live_flow_3 as v, index}
      {#if v}
        <div class="bk_right_panel_container_item">
          <div class="time">{v.display_time}</div>
          <a href={v.url} class="content" target="_blank">{v.content}</a>
        </div>
        {#if index != data_live_flow_3.length - 1}
          <div class="seprate"></div>
        {/if}
      {/if}
    {/each}
  </div>
</div>

<!-- 实时行情 财经日历跳过 -->

<!-- 最新资讯 -->
<div class="bk_information_flow">
  <div class="title_bar">最新资讯</div>
  <div class="bk_information_flow_items">
    {#each data_infomation_flows as v, index}
      {#if v}
        <div class="bk_information_flow_item">
          <img src={v.image} alt="" />
          <div class="title">{v.title}</div>
          <div class="short">{v.short}</div>
          <div class="author_time">
            {v.author}
            <div class="time_ago">{v.display_time}</div>
          </div>
        </div>
      {/if}
    {/each}
  </div>
</div>

<!-- <div class="bk_liveflow">top news</div> -->

<!-- 特辑  这里用大师课凑活一下吧  -->

<div class="bk_masterlessons">
  <div class="title_bar">特辑</div>
  <div class="bk_masterlessons_container">
    {#each data_lession_master_4 as v, index}
      {#if v}
        <div class="bk_masterlessons_container_item">
          <a href="/premium/topics/1003713" target="_blank" class="premium-img"
            ><img
              src={v.image}
              alt=""
              width="90"
              height="120"
              class="qn-img lazy qn-img-lazyloaded"
            /></a
          >
          <div class="text">
            <a href="/premium/" target="_blank" class="title">
              {v.title}
            </a>
            <div class="content">
              <time class="last-updated"> {v.display_time}： </time>
              <a href="/premium/articles/3708258" target="_blank">
                {v.content}
              </a>
            </div>
          </div>
        </div>
      {/if}
    {/each}
  </div>
</div>

<!-- 最热文章10条 -->
<div class="bk_hot_article">
  <div class="title_bar">最热文章</div>
  <div class="bk_hot_article_container">
    {#each data_articles as v, index}
      {#if v}
        <div
          class={index == 0
            ? "bk_hot_article_item empty_border"
            : "bk_hot_article_item"}
        >
          <span>{v.num}.</span>
          <div>{v.title}</div>
        </div>
      {/if}
    {/each}
  </div>
</div>

<style>
  .bk_real {
    grid-column: 1/3;
    display: grid;
    grid-auto-flow: column dense;
    grid-auto-columns: max-content;
    font-size: 14px;
    line-height: 20px;
    padding: 20px 0;
    /* font-size: 18px; */
    /* color: #000; */
  }
  .bk_real > div {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    height: 100%;
    padding-left: 16px;
    padding-right: 16px;
    border-left: 1px solid #e6e6e6;
  }
  .bk_real > div .red {
    color: red;
    font-size: 12px;
  }
  .bk_news5 {
    grid-column: 1/3;
    display: grid;
    grid-template-columns: 2fr 1fr 1fr;
    grid-template-rows: 1fr 1fr;
    font-size: 14px;
    gap: 15px;
    /* max-width: 65vw; */
    /* height: 279px; */
  }
  .bk_news5 > a {
    line-height: 20px;
    position: relative;
    height: 100%;
    width: 100%;
    overflow: hidden;
    display: block;
  }
  .bk_news5 .a_first {
    grid-row: 1 / 3;
  }
  .bk_news5 > a img {
    width: 100%;
    aspect-ratio: auto 180 / 135;
    height: 100%;
    object-fit: cover;
  }

  /* .bk_news5 .a_first img {

  } */
  .bk_news5 > a .text {
    position: absolute;
    bottom: 0;
    /* opacity: 44%; */
    left: 0;
    /* background-color: #17181a; */
    padding: 15px;
    overflow: hidden;
    line-height: 26px;
    height: 44px;
    width: 100%;
    color: #fff;
    text-shadow: 1px 1px 0 #000;
    background-image: linear-gradient(to top, #000000, rgba(0, 0, 0, 0));
  }
  .bk_news5 > a .text .title {
    font-size: 14px;
    font-weight: 400;
    width: 91%;
  }
  .bk_information_flow {
    grid-row: 3;
  }
  .bk_information_flow_item {
    background-color: white;
    display: grid;
    grid-template-columns: 180px 1fr;
    grid-template-rows: repeat(3, min-content);
    height: 135px;
    line-height: 25px;
    margin-bottom: 15px;
    column-gap: 15px;
    padding: 20px;
  }
  .bk_information_flow_item img {
    grid-row: 1 / 4;
    width: 180px;
    aspect-ratio: auto 180 / 135;
    height: 135px;
    object-fit: cover;
  }

  .bk_information_flow_item .short {
    color: #666;
    line-height: 21px;
    font-size: 13px;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2; /* 限制在一个块元素显示的文本的行数 */
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: normal; /* 或者使用 break-spaces 根据需要 */
    max-height: 42px; /* 根据字体大小调整最大高度，以便显示两行文本 */
  }
  .bk_information_flow_item .author_time {
    font-size: 13px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
  }
  .bk_information_flow_item .time_ago {
    color: gray;
    border-left: 1px solid gray;
    padding-left: 10px;
    margin: 10px;
  }
  .bk_right_panel {
    grid-row: 1/2;
    grid-column: 3/4;
  }
  .bk_right_panel_container {
    background-color: #fff;
    padding: 10px;
  }

  .bk_right_panel_container_item {
    padding: 6px 10px;
  }
  .bk_right_panel_container_item:hover {
    background-color: #f2f4f5;
  }
  .bk_right_panel_container_item .time {
    border-radius: 2px;
    display: inline-block;
    border-radius: 2px;
    background-color: #191919;
    padding: 2px 4px 1px 5px;
    color: #fff;
    font-size: 12px;
    position: relative;
  }
  .bk_right_panel_container_item .content {
    font-size: 13px;
    line-height: 21px;
    max-height: 63px;
    color: #222;
    margin-top: 10px;
    padding-bottom: 1px;
    display: block;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 3;
    overflow-wrap: break-word;
    overflow: hidden;
  }
  .seprate {
    width: 260px;
    height: 0px;
    border-top: 1px #e6e6e6 dashed;
    margin: 10px auto;
  }
  .title_bar {
    height: 48px;
    display: flex;
    padding: 0px 20px;
    justify-content: space-between;
    align-items: center;
    background-color: #f5f6f7;
  }
  .title_bar div:nth-child(2) {
    font-size: 12px;
    line-height: 1.08;
    color: #666;
  }
  .empty_border {
    border-top: 0px !important;
  }
  .bk_hot_article {
    grid-row: 3;
    grid-column: 2/3;
  }
  .bk_hot_article_container {
    padding: 14px;
    background-color: white;
    padding-top: 0px;
  }
  .bk_hot_article_item {
    display: flex;
    gap: 12px;
    font-size: 14px;
    line-height: 24px;
    align-items: center;
  }
  .bk_hot_article_item span {
    font-weight: 600;
    font-size: 25px;
    font-family: monospace;
  }
  .bk_hot_article_container h4 {
  }
  .bk_hot_article_container > div {
    display: flex;
    align-items: flex-start;
    padding: 10px 0;
    border-top: 1px dashed #e6e6e6;
  }

  .bk_masterlessons {
    display: none;

    grid-column: 2/3;
    grid-row: 3/4;
  }
  .bk_masterlessons_container {
    border-radius: 4px;
    padding: 10px 10px;
    background-color: #fff;
  }
  .bk_masterlessons_container_item {
    display: grid;
    grid-template-rows: 1fr 1fr 1fr;
    grid-template-columns: 80px 1fr;
  }
  .bk_masterlessons_container_item a {
    display: block;
  }
  .bk_masterlessons_container_item a img {
    width: 100%;
    height: 100%;
  }

  .bk_masterlessons_container_item time {
    color: #1677d9;
    font-size: 13px;
  }
  .bk_masterlessons_container_item .content a {
    line-height: 21px;
    color: #666;
    font-size: 13px;
  }
</style>
